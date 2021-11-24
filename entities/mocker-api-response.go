package entities

type MockerApiResponse struct {
	Query string                 `json:"query"`
	Value map[string]interface{} `json:"value"`
}
