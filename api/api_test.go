package api

import (
	"testing"
)

func TestAuth(t *testing.T) {
	api := NewAPI("tokengsss")
	b := api.BllSmartImage.renderApollo()
	c := api.MiddleSmartApi.getPip()
	expected := "renderApollo"
	e2 := "pip"

	if b != expected {
		t.Errorf("Expected %s, but got %s", expected, b)
	}
	// result = Auth(appId, appKey)
	if c != e2 {
		t.Errorf("Expected %s, but got %s", e2, c)
	}
}
