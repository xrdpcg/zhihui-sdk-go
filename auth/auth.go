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
	"strings"
)

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

type Auth struct {
	id        string
	secretId  string
	secretKey string
	host      string
}

func NewAuth(id string, secretId string, secretKey string) *Auth {
	return &Auth{
		id:        id,
		secretId:  secretId,
		secretKey: secretKey,
		host:      "https://zhihui.qq.com",
	}
}
func (a *Auth) SetHost(host string) {
	a.host = strings.Replace(host, "api.", "", 1)
}

func (a *Auth) Init() string {
	params := map[string]interface{}{
		"id":       a.id,
		"secretId": a.secretId,
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
	// 生成sig
	paramsMap = a.sigCode(paramsMap, a.secretKey)
	// 将params转换为JSON格式的字符串
	body, err := json.Marshal(paramsMap)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return "error"
	}
	// 发起POST请求
	apiURL := a.host + "/account/api/auth/secret"
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

func (a *Auth) sigCode(params map[string]string, key string) map[string]string {
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
