package api

import (
	"fmt"
)

type CreateTask struct {
	Custom []PlugData
}

type PlugData struct {
	Plug   string
	Params map[string]interface{}
}

type CreateBatchTask []CreateTask

type WorkflowBaseInput struct {
	Workflow string   `json:"workflow"`
	Name     string   `json:"name"`
	Params   []string `json:"params"`
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

type MiddleSmartService struct {
	*BaseService
}

var BASE_FILTER = []string{
	"id",
	"progress",
	"result",
	"creator",
	"createdAt",
	"updatedAt",
}

func NewMiddleSmartService(host string, httpHandler HttpHandler) *MiddleSmartService {
	baseRoute := "/middle-smart"
	baseService := NewBaseService(host, baseRoute, httpHandler)
	return &MiddleSmartService{
		BaseService: baseService,
	}
}

func (s *MiddleSmartService) CreateBatchTask(params map[string]interface{}) (SuccessResponse, error) {
	url := fmt.Sprintf("%s/workflow/task/batch", s.baseUrl)
	res, err := s.httpHandler.Post(url, params, nil)
	if err != nil {
		return SuccessResponse{}, err
	}
	return res.(SuccessResponse), nil
}

func (s *MiddleSmartService) GetBatchTaskStatus(id string) (SuccessResponse, error) {
	url := fmt.Sprintf("%s/workflow/task/collect/%s?field=name,total,fail,complete&nopop=true", s.baseUrl, id)
	res, err := s.httpHandler.Get(url, nil)
	if err != nil {
		return SuccessResponse{}, err
	}
	return res.(SuccessResponse), nil
}

func (s *MiddleSmartService) GetBatchTaskResList(id string, page int, limit int) (SuccessResponse, error) {
	url := fmt.Sprintf("%s/workflow/task/list", s.baseUrl)
	params := map[string]interface{}{
		"field":  BASE_FILTER,
		"filter": map[string]interface{}{"collect": id},
		"page":   page,
		"limit":  limit,
	}
	res, err := s.httpHandler.Post(url, params, nil)
	if err != nil {
		return SuccessResponse{}, err
	}
	return res.(SuccessResponse), nil
}
