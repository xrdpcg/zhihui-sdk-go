package api

import (
	"fmt"
)

type BllSmartImage struct {
	token string
}

func NewBllSmartImage(token string) *BllSmartImage {
	return &BllSmartImage{token: token}
}

func (s *BllSmartImage) renderApollo() string {
	fmt.Println("renderApollo")
	return "renderApollo"
}
