package api

import (
	"testing"
)

func TestAuth(t *testing.T) {
	api := NewAPI("tokengsss")
	api.BllSmartImage.renderApollo()
	api.MiddleSmartApi.getPip()
	// expected := "error"
	// if result != expected {
	// 	t.Errorf("Expected %s, but got %s", expected, result)
	// }
	// result = Auth(appId, appKey)
	// if result == expected {
	// 	t.Errorf("Expected Right Access Token, but got error")
	// }
}
