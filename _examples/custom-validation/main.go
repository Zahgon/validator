package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type MyStruct struct {
	String string `validate:"is-awesome"`
}

var validate *validator.Validate

func main() {

	validate = validator.New()
	validate.RegisterValidation("is-awesome", ValidateMyVal)

	s := MyStruct{String: "awesome"}

	err := validate.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}

	s.String = "not awesome"
	err = validate.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
}

func ValidateMyVal(fl validator.FieldLevel) bool { _ = "STUB: not implemented"; return false }
