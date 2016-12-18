package validator

type ValidationResult struct {
	Message   string `json:"message"`
	IsSuccess bool   `json:"-"`
}
