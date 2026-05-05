package model

// Geotab API Models
type GeotabRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}