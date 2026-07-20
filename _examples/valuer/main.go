package main

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Nullable[T any] struct {
	Data T
}

func (n Nullable[T]) ValidatorValue() any { _ = "STUB: not implemented"; return *new(any) }

type Config struct {
	Name string `validate:"required"`
}

type Record struct {
	Config Nullable[Config] `validate:"required"`
}

var validate *validator.Validate

func main() {
	validate = validator.New()

	valid := Record{
		Config: Nullable[Config]{
			Data: Config{Name: "validator"},
		},
	}
	err := validate.Struct(valid)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}

	invalid := Record{
		Config: Nullable[Config]{},
	}
	err = validate.Struct(invalid)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) && len(validationErrs) > 0 {
			fmt.Printf("First error namespace: %s\n", validationErrs[0].Namespace())
		}
	}
}
