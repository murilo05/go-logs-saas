package domain

type Ingest struct {
	ID      string
	APIKey  string
	AppID   string
	Level   string
	Message string
	Context map[string]string
}

type IngestOutput struct {
	ID     string
	Status string
	Log    Ingest
}
