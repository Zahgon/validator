package uk

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func RegisterDefaultTranslations(v *validator.Validate, trans ut.Translator) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	_ = "STUB: not implemented"
	return *new(validator.RegisterTranslationsFunc)
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	_ = "STUB: not implemented"
	return ""
}

func translateFuncWithParam(ut ut.Translator, fe validator.FieldError) string {
	_ = "STUB: not implemented"
	return ""
}

func registerCardinals(ut ut.Translator, prefix string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
