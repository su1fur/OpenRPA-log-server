package validators

import (
	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

func Init() {
	validate.RegisterValidation("existTaskLine", ValidateExistTaskLine)
}

// Struct valid the structure
func Struct(s interface{}) error {

	return validate.Struct(s)
}
