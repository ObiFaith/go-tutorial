package models

type ClassifyResponseData struct {
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
	SampleSize  int     `json:"sample_size"`
	IsConfident bool    `json:"is_confident"`
	ProcessedAt string  `json:"processed_at"`
}

type ClassifyResponse struct {
	Status string               `json:"status"`
	Data   ClassifyResponseData `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type HealthResponse struct {
    Status string `json:"status"`
}
