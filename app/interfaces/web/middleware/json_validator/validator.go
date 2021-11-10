package json_validator

import "gopkg.in/go-playground/validator.v9"

type customValidator struct {
	validator *validator.Validate
}

func NewJsonValidator() *customValidator {
	return &customValidator{
		validator: validator.New(),
	}
}
func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
