package main

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`
	Addresses      []*Address `validate:"required,dive,required"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func main() {

	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	translateAll(trans)
	translateIndividual(trans)
	translateOverride(trans)
}

func translateAll(trans ut.Translator) { _ = "STUB: not implemented"; return }

func translateIndividual(trans ut.Translator) { _ = "STUB: not implemented"; return }

func translateOverride(trans ut.Translator) { _ = "STUB: not implemented"; return }
