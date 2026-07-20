package main

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type DbBackedUser struct {
	Name sql.NullString `validate:"required"`
	Age  sql.NullInt64  `validate:"required"`
}

var validate *validator.Validate

func main() {

	validate = validator.New()

	validate.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})

	x := DbBackedUser{Name: sql.NullString{String: "", Valid: true}, Age: sql.NullInt64{Int64: 0, Valid: false}}

	err := validate.Struct(x)

	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
}

func ValidateValuer(field reflect.Value) interface{} { _ = "STUB: not implemented"; return nil }
