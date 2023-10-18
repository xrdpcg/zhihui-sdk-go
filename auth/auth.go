package auth

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

func sigCode(params map[string]string, key string) map[string]string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sortedParams string
	for _, k := range keys {
		sortedParams += k + "=" + params[k] + "&"
	}
	sortedParams = sortedParams[:len(sortedParams)-1]
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(url.QueryEscape(sortedParams)))

	params["sig"] = hex.EncodeToString(h.Sum(nil))

	return params
}

type AuthOptions struct {
	channel string
	stamp   string
}
type SuccessResponse struct {
	StatusCode int  `json:"statusCode"`
	Data       Data `json:"data"`
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type Data struct {
	AccessToken string `json:"access_token"`
}
type Response interface {
	// 定义接口方法（如果有需要的话）
}

func Auth(appId string, appKey string, options ...AuthOptions) string {
	defaultOptions := AuthOptions{
		channel: "api",
		stamp:   "apiUser",
	}
	if len(options) > 0 {
		defaultOptions = options[0]
	}
	params := map[string]interface{}{
		"appid":      appId,
		"channel":    defaultOptions.channel,
		"stamp":      defaultOptions.stamp,
		"updateInfo": true,
		"timestamp":  time.Now().UnixNano() / int64(time.Millisecond),
	}
	paramsMap := make(map[string]string)
	for k, v := range params {
		switch v := v.(type) {
		case string:
			paramsMap[k] = v
		case []string:
			paramsMap[k] = v[0] // or join all elements if needed
		case bool:
			paramsMap[k] = strconv.FormatBool(v)
		case int64:
			paramsMap[k] = strconv.FormatInt(v, 10)
		}
	}
	key := appKey
	// 生成sig
	paramsMap = sigCode(paramsMap, key)
	// 将params转换为JSON格式的字符串
	body, err := json.Marshal(paramsMap)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return "error"
	}
	// 发起POST请求
	apiURL := "https://zhihui.qq.com/account/api/auth/token"
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("POST request error:", err)
		return "error"
	}
	defer resp.Body.Close()
	// 保存 resp.Body 的内容到变量 responseBody
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "error"
	}
	// 解码成功响应
	var successResponse SuccessResponse
	err = json.Unmarshal(responseBody, &successResponse)
	if err != nil {
		fmt.Println("json.NewDecoder.Decode error:", err)
		return "error"
	}
	if err == nil && successResponse.StatusCode == 200 {
		return successResponse.Data.AccessToken
	}
	var errorResponse ErrorResponse
	err = json.Unmarshal(responseBody, &errorResponse)
	if err != nil {
		fmt.Println("json.NewDecoder.Decode ErrorResponse error:", err)
		return "error"
	}
	if err == nil && errorResponse.StatusCode != 200 {
		fmt.Println("Request Auth API Error:", errorResponse.Message)
		return "error"
	}
	return "error"
}
