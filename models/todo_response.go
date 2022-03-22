package models

type ToDo struct {
	ID        uint64 `json:"Id"`
	UserId    uint64 `json:"UserId"`
	Details   string `json:"Details"`
	Status    string `json:"Status"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}
