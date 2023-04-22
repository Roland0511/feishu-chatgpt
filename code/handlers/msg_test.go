package handlers

import (
	"context"
	"fmt"
	"log"
	"start-feishubot/initialization"
	"testing"
	"time"
)

func TestCardMsg(t *testing.T) {
	config := initialization.LoadConfig("../../config.yaml")
	initialization.LoadLarkClient(*config)
	println(config)
	ctx := context.Background()
	newCard, _ := newSendCard(
		nil,
		withMainText("test"),
		withNote("提醒：点击对话框参与回复，可保持话题连贯"))
	msgId := "om_68ff679096752db83c90c82cc91c7f91"
	replayMsgId, err := replyCard(ctx, &msgId, newCard)

	if err != nil {
		log.Fatal(err)
	}

	msg := ""
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		msg += "a"
		newCard, _ = newSendCard(
			nil,
			withMarkdownText(msg),
			withNote("提醒：点击对话框参与回复，可保持话题连贯"))
		fmt.Println(msg)
		// err = updateCard(ctx, &replayMsgId, newCard)
		go updateCard(ctx, &replayMsgId, newCard)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}

}
