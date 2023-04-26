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
		{Role: "user", Content: "解释负载均衡"},
		{Role: "assistant", Content: "🧰 毛毛正在网络上搜索...\n喵负载均衡（Load Balance）是一种在多个服务器、网络连接、CPU、磁盘驱动器或其他资源中分配负载的技术，目的是优化资源使用、最大化吞吐率、最小化响应时间，同时避免过载。负载均衡的主要作用是提高系统整体的响应速度和可用性，通过将网络流量平均分发到多个服务器上。喵"},
		{Role: "user", Content: "吞吐量是什么意思"},
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
		{Role: "user", Content: "解释负载均衡"},
		{Role: "assistant", Content: "🧰 毛毛正在网络上搜索...\n喵负载均衡（Load Balance）是一种在多个服务器、网络连接、CPU、磁盘驱动器或其他资源中分配负载的技术，目的是优化资源使用、最大化吞吐率、最小化响应时间，同时避免过载。负载均衡的主要作用是提高系统整体的响应速度和可用性，通过将网络流量平均分发到多个服务器上。喵"},
		{Role: "user", Content: "吞吐量是什么意思"},
	}

	gpt := NewMMGPT(*config)

	fullContent := ""
	f := func(resp openai.Messages, err error) {
		fmt.Println(resp.Content, resp.Role)
		fullContent += resp.Content
	}

	ff := func(error) {
		fmt.Printf("full content: %s", fullContent)
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
