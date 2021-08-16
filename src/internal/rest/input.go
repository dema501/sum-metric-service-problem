package rest

// AddMetricRequest wrap POST requests to add metrics
type AddMetricRequest struct {
	Value int `json:"value" binding:"required,number"`
}
