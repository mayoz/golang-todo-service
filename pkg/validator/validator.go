package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(st interface{}) map[string]string {
	errors := make(map[string]string)
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	if err := validate.Struct(st); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}
	}
	return errors
}
