package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func InitlizeValidator() *Validator {
	v := &Validator{}
	v.newValidator()

	return v
}

func (v *Validator) newValidator() {
	v.validate = validator.New(func(v *validator.Validate) {
		v.SetTagName("validate")
	})
	v.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return strings.Split(fld.Tag.Get("json"), ",")[0]
	})
}

func (v *Validator) Struct(objStruct any) Errors {
	if err := v.validate.Struct(objStruct); err != nil {

		return v.ParseValidationErrors(err)
	}
	return make(Errors)
}
