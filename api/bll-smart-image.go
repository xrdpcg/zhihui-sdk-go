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

func (s *BllSmartImage) RenderApollo() string {
	fmt.Println("renderApollo")
	return "renderApollo"
}
