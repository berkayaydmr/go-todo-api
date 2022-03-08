package models

type ToDoRequest struct {
	Details *string `json:"details,omitempty"`
	Status string `json:"status"`
}
