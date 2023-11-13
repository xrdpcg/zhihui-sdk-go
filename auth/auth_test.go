package auth

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// 加载 .env 文件
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	exitCode := m.Run()
	os.Exit(exitCode)
}
func TestAuth_Init(t *testing.T) {
	auth := Auth{
		id:        os.Getenv("USERID"),
		secretId:  os.Getenv("SECRET_ID"),
		secretKey: os.Getenv("SECRET_KEY"),
	}
	auth.SetHost("https://api.zhdev.woa.com")
	result := auth.Init()
	expected := "error"
	if result == expected {
		t.Errorf("Expected Right Access Token, but got error")
	}
	authErro := Auth{
		id:        os.Getenv("USERID"),
		secretId:  os.Getenv("SECRET_ID"),
		secretKey: os.Getenv("SECRET_KEY_ERROR"),
	}
	authErro.SetHost("https://api.zhdev.woa.com")
	resultError := authErro.Init()
	if resultError != expected {
		t.Errorf("Expected %s, but got %s", expected, resultError)
	}
}
