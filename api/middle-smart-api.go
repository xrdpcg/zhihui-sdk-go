package api

import (
	"fmt"
)

type MiddleSmartApi struct {
	token string
}

func NewMiddleSmartApi(token string) *MiddleSmartApi {
	return &MiddleSmartApi{token: token}
}

func (s *MiddleSmartApi) getPip() string {
	fmt.Println("MiddleSmartApi getPip")
	return "pip"
}
