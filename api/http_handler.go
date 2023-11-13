package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IHttpHandler interface {
	SetToken(token string)
	Post(requestUrl string, params interface{}, headers map[string]interface{}) (interface{}, error)
	Patch(requestUrl string, params interface{}, headers map[string]interface{}) (interface{}, error)
	Get(requestUrl string, headers map[string]interface{}) (interface{}, error)
}

type HttpHandler struct {
	token string
}
type SuccessResponse struct {
	StatusCode int                    `json:"statusCode"`
	Data       map[string]interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func NewHttpHandler(token string) *HttpHandler {
	return &HttpHandler{token: token}
}

func (h *HttpHandler) SetToken(token string) {
	h.token = token
}

func (h *HttpHandler) sendRequest(config map[string]interface{}) (interface{}, error) {
	method := config["method"].(string)
	url := config["url"].(string)
	params := config["data"]

	var req *http.Request
	var err error

	if params != nil {
		jsonParams, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonParams))
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.token))

	// 设置 headers
	headers := config["headers"].(map[string]interface{})
	for key, value := range headers {
		req.Header.Set(key, fmt.Sprintf("%v", value))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 保存 resp.Body 的内容到变量 responseBody
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err, nil
	}
	// 解码成功响应
	var successResponse SuccessResponse
	err = json.Unmarshal(responseBody, &successResponse)
	if err != nil {
		fmt.Println("json.NewDecoder.Decode error:", err)
		return err, nil
	}
	if err == nil {
		switch successResponse.StatusCode {
		case 200, 201, 210:
			return successResponse, nil
		}
	}
	var errorResponse ErrorResponse
	err = json.Unmarshal(responseBody, &errorResponse)
	if err != nil {
		fmt.Println("json.NewDecoder.Decode ErrorResponse error:", err)
		return err, nil
	} else {
		fmt.Println("Request Error:", errorResponse.Message)
		return errorResponse, nil
	}
}
func (h *HttpHandler) Post(requestUrl string, params map[string]interface{}, headers map[string]interface{}) (interface{}, error) {
	config := map[string]interface{}{
		"method":  "POST",
		"url":     requestUrl,
		"data":    params,
		"headers": headers,
	}
	return h.sendRequest(config)
}

func (h *HttpHandler) Patch(requestUrl string, params map[string]interface{}, headers map[string]interface{}) (interface{}, error) {
	config := map[string]interface{}{
		"method":  "PATCH",
		"url":     requestUrl,
		"data":    params,
		"headers": headers,
	}
	return h.sendRequest(config)
}

func (h *HttpHandler) Get(requestUrl string, headers map[string]interface{}) (interface{}, error) {
	config := map[string]interface{}{
		"method":  "GET",
		"url":     requestUrl,
		"headers": headers,
	}
	return h.sendRequest(config)
}
