package dto

type IngestInput struct {
	APIKey  string            `json:"api_key"`
	AppID   string            `json:"app_id"`
	Level   string            `json:"level"`
	Message string            `json:"message"`
	Context map[string]string `json:"context"`
}
