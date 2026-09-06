package validator

import (
	"reflect"

	ut "github.com/go-playground/universal-translator"
)

const (
	fieldErrMsg = "Key: '%s' Error:Field validation for '%s' failed on the '%s' tag"
)

type ValidationErrorsTranslations map[string]string

type InvalidValidationError struct {
	Type reflect.Type
}

func (e *InvalidValidationError) Error() string { _ = "STUB: not implemented"; return "" }

type ValidationErrors []FieldError

func (ve ValidationErrors) Error() string { _ = "STUB: not implemented"; return "" }

func (ve ValidationErrors) Translate(ut ut.Translator) ValidationErrorsTranslations {
	_ = "STUB: not implemented"
	return *new(ValidationErrorsTranslations)
}

type FieldError interface {
	Tag() string

	ActualTag() string

	Namespace() string

	StructNamespace() string

	Field() string

	StructField() string

	Value() interface{}

	Param() string

	Kind() reflect.Kind

	Type() reflect.Type

	Translate(ut ut.Translator) string

	Error() string
}

var _ FieldError = new(fieldError)
var _ error = new(fieldError)

type fieldError struct {
	v              *Validate
	tag            string
	actualTag      string
	ns             string
	structNs       string
	fieldLen       uint8
	structfieldLen uint8
	value          interface{}
	param          string
	kind           reflect.Kind
	typ            reflect.Type
}

func (fe *fieldError) Tag() string { _ = "STUB: not implemented"; return "" }

func (fe *fieldError) ActualTag() string { _ = "STUB: not implemented"; return "" }

func (fe *fieldError) Namespace() string { _ = "STUB: not implemented"; return "" }

func (fe *fieldError) StructNamespace() string { _ = "STUB: not implemented"; return "" }

func (fe *fieldError) Field() string { _ = "STUB: not implemented"; return "" }

func (fe *fieldError) StructField() string { _ = "STUB: not implemented"; return "" }

func (fe *fieldError) Value() interface{} { _ = "STUB: not implemented"; return nil }

func (fe *fieldError) Param() string { _ = "STUB: not implemented"; return "" }

func (fe *fieldError) Kind() reflect.Kind { _ = "STUB: not implemented"; return *new(reflect.Kind) }

func (fe *fieldError) Type() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (fe *fieldError) Error() string { _ = "STUB: not implemented"; return "" }

func (fe *fieldError) Translate(ut ut.Translator) string { _ = "STUB: not implemented"; return "" }
