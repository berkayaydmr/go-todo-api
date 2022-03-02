package models

type ToDo struct {
	ID        int    `json:"ID"`
	Details   string `json:"Details"`
	Status    string `json:"Status"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}
