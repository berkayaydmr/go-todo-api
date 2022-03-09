package models

type ToDoPatchRequest struct {
	Details *string `json:"details,omitempty"`
	Status *string  `json:"status"`
}