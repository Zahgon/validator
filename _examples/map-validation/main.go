package main

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func main() {
	validate = validator.New()

	validateMap()
	validateNestedMap()
}

func validateMap() { _ = "STUB: not implemented"; return }

func validateNestedMap() { _ = "STUB: not implemented"; return }
