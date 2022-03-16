package models

type UserPatchRequest struct {
	OldPassword     string `json:"currentPassword"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"confirmPassword"`
}

func (model *UserPatchRequest) Validate() bool {
	var isPasswordConfirmed bool = false
	if model.Password == model.PasswordConfirm {
		isPasswordConfirmed = true
		return isPasswordConfirmed
	}
	return isPasswordConfirmed
}
