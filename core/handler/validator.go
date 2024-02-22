package core_handler

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func IsBeforeNow(field validator.FieldLevel) bool {
	fmt.Println(field.Field().String())
	return true
}
