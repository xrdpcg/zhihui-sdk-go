package api

type BaseService struct {
	host        string
	baseUrl     string
	baseRoute   string
	httpHandler HttpHandler
}

func NewBaseService(host string, baseRoute string, httpHandler HttpHandler) *BaseService {
	baseUrl := host + baseRoute
	return &BaseService{
		host:        host,
		baseUrl:     baseUrl,
		baseRoute:   baseRoute,
		httpHandler: httpHandler,
	}
}

func (s *BaseService) SetHost(host string) {
	s.host = host
	s.baseUrl = host + s.baseRoute
}

func (s *BaseService) FilterResponse(obj map[string]interface{}, filter []string) map[string]interface{} {
	result := make(map[string]interface{})
	for _, key := range filter {
		if val, ok := obj[key]; ok {
			result[key] = val
		}
	}
	return result
}
