package models

type UserRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"confirmPassword"`
}

func (model *UserRequest) Validate() bool {
	var isPasswordConfirmed bool = false
	if model.Password == model.PasswordConfirm {
		isPasswordConfirmed = true
		return isPasswordConfirmed
	}
	return isPasswordConfirmed
}
