package client

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
)

type WorkflowBaseInput struct {
	Workflow string              `json:"workflow"`
	Name     string              `json:"name"`
	Params   []WorkflowTaskInput `json:"params"`
}
type WorkflowTaskInput struct {
	Custom []CustomInput `json:"custom"`
}
type CustomInput struct {
	Plug   string            `json:"plug"`
	Params CustomParamsInput `json:"params"`
}

type CustomParamsInput struct {
	JsonData string `json:"jsonData"`
}

func TestMain(m *testing.M) {
	// 加载 .env 文件
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}
func TestNewZhihuiSDKClient(t *testing.T) {
	options := map[string]string{
		"id":        os.Getenv("USERID"),
		"secretId":  os.Getenv("SECRET_ID"),
		"secretKey": os.Getenv("SECRET_KEY"),
		"host":      os.Getenv("HOST"),
	}
	client := NewZhihuiSDKClient(options)
	if client == nil {
		t.Fatal("Expected client to not be nil")
	}

	task1, err := os.ReadFile("../testdata/datasiceInputA.json")
	if err != nil {
		panic(err)
	}
	task2, err := os.ReadFile("../testdata/datasiceInputB.json")
	if err != nil {
		panic(err)
	}
	// Create a task
	workflowBaseInput := WorkflowBaseInput{
		Workflow: "6538da7ca4c67b45d712e4f1",
		Name:     "DS业务测试",
		Params:   []WorkflowTaskInput{},
	}
	taskFrame1 := WorkflowTaskInput{
		Custom: []CustomInput{
			{
				Plug: "6538d63ccc29928c335470df",
				Params: CustomParamsInput{
					JsonData: "",
				},
			},
		},
	}

	taskFrame2 := WorkflowTaskInput{
		Custom: []CustomInput{
			{
				Plug: "6538d63ccc29928c335470df",
				Params: CustomParamsInput{
					JsonData: "",
				},
			},
		},
	}
	taskFrame1.Custom[0].Params.JsonData = string(task1)
	taskFrame2.Custom[0].Params.JsonData = string(task2)

	workflowBaseInput.Params = append(workflowBaseInput.Params, taskFrame1)
	workflowBaseInput.Params = append(workflowBaseInput.Params, taskFrame2)

	workflowBaseInputJSON, _ := json.Marshal(workflowBaseInput)
	var workflowBaseInput2 WorkflowBaseInput
	json.Unmarshal(workflowBaseInputJSON, &workflowBaseInput2)
	res, _ := client.Workflow.CreateBatchTask(convertToMap(workflowBaseInput2))
	if res.StatusCode != 201 {
		t.Errorf("Expected Right Access Token, but got error")
	}
	taskID := res.Data["id"].(string)
	res, _ = client.Workflow.GetBatchTaskStatus(taskID)
	if res.StatusCode != 200 {
		t.Errorf("Expected Right Access Token, but got error")
	}
	// fmt.Println(res.Data)
}

func convertToMap(input WorkflowBaseInput) map[string]interface{} {
	data, _ := json.Marshal(input)
	var result map[string]interface{}
	json.Unmarshal(data, &result)
	return result
}
func getType(v interface{}) reflect.Kind {
	return reflect.TypeOf(v).Kind()
}
