package mmai

import (
	"fmt"
	"start-feishubot/initialization"
	"start-feishubot/services/openai"
	"testing"
)

func TestCompletions(t *testing.T) {
	config := initialization.LoadConfig("../../../config.yaml")

	msgs := []openai.Messages{
		// {Role: "system", Content: "你是一个专业的翻译官，负责中英文翻译。"},
		{Role: "user", Content: "宇宙探索编辑部 这部电影的主演是谁？"},
	}

	gpt := NewMMGPT(*config)

	resp, err := gpt.Completions(msgs)
	if err != nil {
		t.Errorf("TestCompletions failed with error: %v", err)
	}

	fmt.Println(resp.Content, resp.Role)
}

func TestStreamCompletions(t *testing.T) {
	config := initialization.LoadConfig("../../../config.yaml")

	msgs := []openai.Messages{
		// {Role: "system", Content: "你是一个专业的翻译官，负责中英文翻译。"},
		{Role: "user", Content: "宇宙探索编辑部 这部电影的主演是谁？"},
	}

	gpt := NewMMGPT(*config)

	f := func(resp openai.Messages, err error) {
		fmt.Println(resp.Content, resp.Role)
	}

	ff := func(error) {
		fmt.Println("stream closed")
	}

	err := gpt.CompletionsWithStream(msgs, f, ff)
	if err != nil {
		t.Errorf("TestCompletions failed with error: %v", err)
	}

}

// func TestGenerateOneImage(t *testing.T) {
// 	config := initialization.LoadConfig("../../config.yaml")
// 	gpt := NewMMGPT(*config)
// 	prompt := "a red apple"
// 	size := "256x256"
// 	imageURL, err := gpt.GenerateOneImage(prompt, size)
// 	if err != nil {
// 		t.Errorf("TestGenerateOneImage failed with error: %v", err)
// 	}
// 	if imageURL == "" {
// 		t.Errorf("TestGenerateOneImage returned empty imageURL")
// 	}
// }

// func TestAudioToText(t *testing.T) {
// 	config := initialization.LoadConfig("../../config.yaml")
// 	gpt := NewMMGPT(*config)
// 	audio := "./test_file/test.wav"
// 	text, err := gpt.AudioToText(audio)
// 	if err != nil {
// 		t.Errorf("TestAudioToText failed with error: %v", err)
// 	}
// 	fmt.Printf("TestAudioToText returned text: %s \n", text)
// 	if text == "" {
// 		t.Errorf("TestAudioToText returned empty text")
// 	}

// }

// func TestVariateOneImage(t *testing.T) {
// 	config := initialization.LoadConfig("../../config.yaml")
// 	gpt := NewMMGPT(*config)
// 	image := "./test_file/img.png"
// 	size := "256x256"
// 	//compressionType, err := GetImageCompressionType(image)
// 	//if err != nil {
// 	//	return
// 	//}
// 	//fmt.Println("compressionType: ", compressionType)
// 	ConvertToRGBA(image, image)
// 	err := VerifyPngs([]string{image})
// 	if err != nil {
// 		t.Errorf("TestVariateOneImage failed with error: %v", err)
// 		return
// 	}

// 	imageBs64, err := gpt.GenerateOneImageVariation(image, size)
// 	if err != nil {
// 		t.Errorf("TestVariateOneImage failed with error: %v", err)
// 	}
// 	//fmt.Printf("TestVariateOneImage returned imageBs64: %s \n", imageBs64)
// 	if imageBs64 == "" {
// 		t.Errorf("TestVariateOneImage returned empty imageURL")
// 	}
// }

// func TestVariateOneImageWithJpg(t *testing.T) {
// 	config := initialization.LoadConfig("../../config.yaml")
// 	gpt := NewChatGPT(*config)
// 	image := "./test_file/test.jpg"
// 	size := "256x256"
// 	compressionType, err := GetImageCompressionType(image)
// 	if err != nil {
// 		return
// 	}
// 	fmt.Println("compressionType: ", compressionType)
// 	//ConvertJPGtoPNG(image)
// 	ConvertToRGBA(image, image)
// 	err = VerifyPngs([]string{image})
// 	if err != nil {
// 		t.Errorf("TestVariateOneImage failed with error: %v", err)
// 		return
// 	}

// 	imageBs64, err := gpt.GenerateOneImageVariation(image, size)
// 	if err != nil {
// 		t.Errorf("TestVariateOneImage failed with error: %v", err)
// 	}
// 	fmt.Printf("TestVariateOneImage returned imageBs64: %s \n", imageBs64)
// 	if imageBs64 == "" {
// 		t.Errorf("TestVariateOneImage returned empty imageURL")
// 	}
// }
