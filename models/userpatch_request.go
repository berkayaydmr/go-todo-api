package models

type UserPatchRequest struct {
	OldPassword     string `json:"currentPassword"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"confirmPassword"`
}

func (model *UserPatchRequest) Validate() bool {
	var isValidationConfirmed bool = false
	if model.Password == model.PasswordConfirm {
		if model.OldPassword != "" || model.Password != "" || model.PasswordConfirm != "" {
			isValidationConfirmed = true
			return isValidationConfirmed
		}
	}
	return isValidationConfirmed
}
