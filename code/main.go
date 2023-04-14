package main

import (
	"context"
	"fmt"
	"log"
	"start-feishubot/handlers"
	"start-feishubot/initialization"
	"start-feishubot/services/mmai"
	"start-feishubot/services/openai"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"

	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"

	sdkginext "github.com/larksuite/oapi-sdk-gin"

	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
)

var (
	cfg = pflag.StringP("config", "c", "./config.yaml", "apiserver config file path.")
)

func test() {
	config := initialization.LoadConfig("../config.yaml")

	msgs := []mmai.Messages{
		// {Role: "system", Content: "你是一个专业的翻译官，负责中英文翻译。"},
		{Role: "user", Content: "帮我总结下这个文档的内容： https://rzvo5fieru.feishu.cn/wiki/wikcn9YeuITFAD4M9hwpgARH7Ke"},
	}

	gpt := mmai.NewMMGPT(*config)

	resp, err := gpt.Completions(msgs)
	if err != nil {
		fmt.Errorf("TestCompletions failed with error: %v", err)
	}

	fmt.Println(resp.Content, resp.Role)

}

func main() {
	test()
	return

	pflag.Parse()
	config := initialization.LoadConfig(*cfg)
	initialization.LoadLarkClient(*config)
	gpt := openai.NewChatGPT(*config)
	mm := mmai.NewMMGPT(*config)
	handlers.InitHandlers(gpt, mm, *config)

	eventHandler := dispatcher.NewEventDispatcher(
		config.FeishuAppVerificationToken, config.FeishuAppEncryptKey).
		OnP2MessageReceiveV1(handlers.Handler).
		OnP2MessageReadV1(func(ctx context.Context, event *larkim.P2MessageReadV1) error {
			return handlers.ReadHandler(ctx, event)
		})

	cardHandler := larkcard.NewCardActionHandler(
		config.FeishuAppVerificationToken, config.FeishuAppEncryptKey,
		handlers.CardHandler())

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/webhook/event",
		sdkginext.NewEventHandlerFunc(eventHandler))
	r.POST("/webhook/card",
		sdkginext.NewCardActionHandlerFunc(
			cardHandler))

	err := initialization.StartServer(*config, r)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
