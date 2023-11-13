package client

import (
	"github.com/xrdpcg/zhihui-sdk-go/api"
	"github.com/xrdpcg/zhihui-sdk-go/auth"
)

type ZhihuiSDKClient struct {
	auth        *auth.Auth
	HttpHandler *api.HttpHandler
	Workflow    *api.MiddleSmartService
	Id          string
	SecretId    string
	SecretKey   string
	Host        string
}

func NewZhihuiSDKClient(options map[string]string) *ZhihuiSDKClient {
	host := options["host"]
	if host == "" {
		host = "https://api.zhihui.qq.com"
	}
	auth := auth.NewAuth(options["id"], options["secretId"], options["secretKey"])
	auth.SetHost(host)
	token := auth.Init()
	if token == "error" {
		return nil
	}
	httpHandler := api.NewHttpHandler(token)
	middleSmartService := api.NewMiddleSmartService(host, *httpHandler)
	return &ZhihuiSDKClient{
		auth:     auth,
		Workflow: middleSmartService,
	}
}
