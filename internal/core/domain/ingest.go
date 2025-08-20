package domain

type Ingest struct {
	APIKey  string
	AppID   string
	Level   string
	Message string
	Context map[string]string
}

type IngestOutput struct {
	ID     string
	Status string
}
