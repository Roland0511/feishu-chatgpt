package handlers

import (
	"context"
	"log"
	"start-feishubot/initialization"
	"testing"
	"time"

	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
)

func TestCardMsg(t *testing.T) {
	config := initialization.LoadConfig("../../config.yaml")
	initialization.LoadLarkClient(*config)
	println(config)
	ctx := context.Background()
	newCard, _ := newSendCard(
		withHeader("👻️ 已开启新的话题", larkcard.TemplateBlue),
		withMainText("test"),
		withNote("提醒：点击对话框参与回复，可保持话题连贯"))
	msgId := "om_f6f58404f1262cce59c1134defba9789"
	replayMsgId, err := replyCard(ctx, &msgId, newCard)

	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)

	newCard, _ = newSendCard(
		nil,
		withMarkdownText("test, \n**aaaaa**"),
		withNote("提醒：点击对话框参与回复，可保持话题连贯"))

	err = updateCard(ctx, &replayMsgId, newCard)
	if err != nil {
		log.Fatal(err)
	}
}
