package auth

import (
	"testing"
)

func TestAuth(t *testing.T) {
	appId := "appId"
	appKey := "appKey"
	appKey_Error := "appKey"

	// 测试输入为 123 和 456 的情况
	result := Auth(appId, appKey_Error)
	expected := "error"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
	result = Auth(appId, appKey)
	if result == expected {
		t.Errorf("Expected Right Access Token, but got error")
	}
}
