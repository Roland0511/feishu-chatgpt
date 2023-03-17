package initialization

import (
	"crypto/tls"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"net/http"
)

var larkClient *lark.Client

func LoadLarkClient(config Config) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: transport,
	}
	larkClient = lark.NewClient(config.FeishuAppId, config.FeishuAppSecret, lark.WithHttpClient(client))
}

func GetLarkClient() *lark.Client {
	return larkClient
}
