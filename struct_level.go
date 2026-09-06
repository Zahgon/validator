package validator

import (
	"context"
	"reflect"
)

type StructLevelFunc func(sl StructLevel)

type StructLevelFuncCtx func(ctx context.Context, sl StructLevel)

func wrapStructLevelFunc(fn StructLevelFunc) StructLevelFuncCtx {
	_ = "STUB: not implemented"
	return *new(StructLevelFuncCtx)
}

type StructLevel interface {
	Validator() *Validate

	Top() reflect.Value

	Parent() reflect.Value

	Current() reflect.Value

	ExtractType(field reflect.Value) (value reflect.Value, kind reflect.Kind, nullable bool)

	ReportError(field interface{}, fieldName, structFieldName string, tag, param string)

	ReportValidationErrors(relativeNamespace, relativeActualNamespace string, errs ValidationErrors)
}

var _ StructLevel = new(validate)

func (v *validate) Top() reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func (v *validate) Parent() reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func (v *validate) Current() reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func (v *validate) Validator() *Validate { _ = "STUB: not implemented"; return nil }

func (v *validate) ExtractType(field reflect.Value) (reflect.Value, reflect.Kind, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Kind), false
}

func (v *validate) ReportError(field interface{}, fieldName, structFieldName, tag, param string) {
	_ = "STUB: not implemented"
	return
}

func (v *validate) ReportValidationErrors(relativeNamespace, relativeStructNamespace string, errs ValidationErrors) {
	_ = "STUB: not implemented"
	return
}
