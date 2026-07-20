package main

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	Gender         string     `validate:"oneof=male female prefer_not_to"`
	FavouriteColor string     `validate:"iscolor"`
	Addresses      []*Address `validate:"required,dive,required"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var validate *validator.Validate

func main() {

	validate = validator.New(validator.WithRequiredStructEnabled())

	validateStruct()
	validateVariable()
}

func validateStruct() { _ = "STUB: not implemented"; return }

func validateVariable() { _ = "STUB: not implemented"; return }
