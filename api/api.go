package api

type Service struct {
	token string
}
type API struct {
	BllSmartImage  *BllSmartImage
	MiddleSmartApi *MiddleSmartApi
}

func NewAPI(token string) *API {
	return &API{
		BllSmartImage:  NewBllSmartImage(token),
		MiddleSmartApi: NewMiddleSmartApi(token),
	}
}
