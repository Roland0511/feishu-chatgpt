package mmai

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"start-feishubot/initialization"
	"start-feishubot/services/openai"
	"strings"
	"time"

	goopenai "github.com/sashabaranov/go-openai"
)

type MMGPT struct {
	ApiKey    string
	ApiUrl    string
	HttpProxy string
}
type requestBodyType int

const (
	jsonBody requestBodyType = iota
	formVoiceDataBody
	formPictureDataBody
)

func (gpt MMGPT) doAPIRequestWithRetry(url, method string, bodyType requestBodyType,
	requestBody interface{}, responseBody interface{}, client *http.Client, maxRetries int) error {
	var requestBodyData []byte
	var err error
	var writer *multipart.Writer

	switch bodyType {
	case jsonBody:
		requestBodyData, err = json.Marshal(requestBody)
		if err != nil {
			return err
		}
	case formVoiceDataBody:
		return errors.New("unsupported request body type: formVoiceDataBody")
	case formPictureDataBody:
		return errors.New("unsupported request body type: formPictureDataBody")
	default:
		return errors.New("unknown request body type")
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(requestBodyData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	if bodyType == formVoiceDataBody || bodyType == formPictureDataBody {
		req.Header.Set("Content-Type", writer.FormDataContentType())
	}
	var response *http.Response
	var retry int
	for retry = 0; retry <= maxRetries; retry++ {
		response, err = client.Do(req)
		fmt.Println("--------------------")
		fmt.Println("req", req.Header)
		fmt.Printf("response: %v", response)
		// read body
		if err != nil || response.StatusCode < 200 || response.StatusCode >= 300 {

			if err != nil {
				fmt.Println("err:", err)
			}

			if response != nil {
				body, _ := ioutil.ReadAll(response.Body)
				fmt.Println("body", string(body))
			}

			if retry == maxRetries {
				break
			}
			time.Sleep(time.Duration(retry+1) * time.Second)
		} else {
			break
		}
	}
	if response != nil {
		defer response.Body.Close()
	}

	if response == nil || response.StatusCode < 200 || response.StatusCode >= 300 {
		return fmt.Errorf("%s api failed after %d retries", strings.ToUpper(method), retry)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, responseBody)
	if err != nil {
		return err
	}
	return nil
}

func (gpt MMGPT) sendRequestWithBodyType(link, method string, bodyType requestBodyType,
	requestBody interface{}, responseBody interface{}) error {
	var err error
	if gpt.HttpProxy == "" {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		client := &http.Client{
			Transport: transport,
			Timeout:   110 * time.Second,
		}
		err = gpt.doAPIRequestWithRetry(link, method, bodyType,
			requestBody, responseBody, client, 3)
	} else {
		proxyUrl, err := url.Parse(gpt.HttpProxy)
		if err != nil {
			return err
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		proxyClient := &http.Client{
			Transport: transport,
			Timeout:   110 * time.Second,
		}
		err = gpt.doAPIRequestWithRetry(link, method, bodyType,
			requestBody, responseBody, proxyClient, 3)
	}

	return err
}

const (
	maxTokens   = 2000
	temperature = 0.7
	engine      = "gpt-3.5-turbo"
)

// ChatGPTResponseBody 请求体
type ChatGPTResponseBody struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int                    `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChatGPTChoiceItem    `json:"choices"`
	Usage   map[string]interface{} `json:"usage"`
}
type ChatGPTChoiceItem struct {
	Message      openai.Messages `json:"message"`
	Index        int             `json:"index"`
	FinishReason string          `json:"finish_reason"`
}

// ChatGPTRequestBody 响应体
type ChatGPTRequestBody struct {
	Model            string            `json:"model"`
	Messages         []openai.Messages `json:"messages"`
	MaxTokens        int               `json:"max_tokens"`
	Temperature      float32           `json:"temperature"`
	TopP             int               `json:"top_p"`
	FrequencyPenalty int               `json:"frequency_penalty"`
	PresencePenalty  int               `json:"presence_penalty"`
	Stream           bool              `json:"stream"`
}

type CompletionsWithStreamRespCallback func(resp openai.Messages, err error)

func (gpt MMGPT) CompletionsWithStream(msg []openai.Messages, callback CompletionsWithStreamRespCallback) error {
	config := goopenai.DefaultConfig(gpt.ApiKey)
	config.BaseURL = gpt.ApiUrl + "/v1"

	c := goopenai.NewClientWithConfig(config)
	ctx := context.Background()

	req := goopenai.ChatCompletionRequest{
		Model:     goopenai.GPT3Dot5Turbo,
		MaxTokens: 20,
		Messages:  msg,
		Stream:    true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return err
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
	for {
		response, streamErr := stream.Recv()
		if errors.Is(streamErr, io.EOF) {
			fmt.Println("\nStream finished")
			return nil
		}

		if streamErr != nil {
			fmt.Printf("\nStream error: %v\n", streamErr)
			callback(openai.Messages{}, streamErr)
			return streamErr
		}

		resp := openai.Messages{
			Content: response.Choices[0].Delta.Content,
		}
		// fmt.Printf(resp.Content)
		callback(resp, nil)
	}
}

func (gpt MMGPT) Completions(msg []openai.Messages) (resp openai.Messages, err error) {
	requestBody := ChatGPTRequestBody{
		Model:            engine,
		Messages:         msg,
		MaxTokens:        maxTokens,
		Temperature:      temperature,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Stream:           false,
	}
	gptResponseBody := &ChatGPTResponseBody{}
	err = gpt.sendRequestWithBodyType(gpt.ApiUrl+"/v1/chat/completions", "POST",
		jsonBody,
		requestBody, gptResponseBody)

	if err == nil && len(gptResponseBody.Choices) > 0 {
		resp = gptResponseBody.Choices[0].Message
	} else {
		resp = openai.Messages{}
		err = errors.New("openai 请求失败")
	}
	return resp, err
}

func NewMMGPT(config initialization.Config) *MMGPT {
	apiKeys := config.OpenaiApiKeys
	apiUrl := config.MMApiUrl
	httpProxy := config.HttpProxy
	return &MMGPT{
		ApiKey:    apiKeys[0],
		ApiUrl:    apiUrl,
		HttpProxy: httpProxy,
	}
}
