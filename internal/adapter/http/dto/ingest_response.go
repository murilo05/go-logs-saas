package dto

type IngestResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Log    Log    `json:"log"`
}

type Log struct {
	AppID   string
	Level   string
	Message string
	Context map[string]string
}
