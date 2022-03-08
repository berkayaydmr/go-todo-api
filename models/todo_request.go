package models

type ToDoRequest struct {
	Details *string `json:"details"`
	Status string `json:"status"`
}

type ToDoPatchRequest struct {
	Details *string `json:"details,omitempty"`
	Status *string  `json:"status"`
}