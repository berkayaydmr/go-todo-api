package models

type ToDoRequest struct {
	Details string `json:"details"`
	Status string `json:"status"`
}
