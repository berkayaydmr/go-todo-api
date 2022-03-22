package models

type UserLogRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
