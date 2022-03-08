package models

type ToDoRequest struct {
	Details string `json:"details"`
	Status  string `json:"status"`
}

func (model *ToDoRequest) Validate() bool {
	var isDetailNil bool = false
	var details *string = &model.Details 
	if details == nil || *details == "" {
		isDetailNil = true
	}
	return isDetailNil
}