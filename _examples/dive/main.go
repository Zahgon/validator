package main

import (
	"github.com/go-playground/validator/v10"
)

type Test struct {
	Array []string          `validate:"required,gt=0,dive,required"`
	Map   map[string]string `validate:"required,gt=0,dive,keys,keymax,endkeys,required,max=1000"`
}

var validate *validator.Validate

func main() {

	validate = validator.New()

	validate.RegisterAlias("keymax", "max=10")

	var test Test

	val(test)

	test.Array = []string{""}
	test.Map = map[string]string{"test > than 10": ""}
	val(test)
}

func val(test Test) { _ = "STUB: not implemented"; return }
