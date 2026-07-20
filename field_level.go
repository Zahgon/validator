package validator

import "reflect"

type FieldLevel interface {
	Top() reflect.Value

	Parent() reflect.Value

	Field() reflect.Value

	FieldName() string

	StructFieldName() string

	Param() string

	GetTag() string

	ExtractType(field reflect.Value) (value reflect.Value, kind reflect.Kind, nullable bool)

	GetStructFieldOK() (reflect.Value, reflect.Kind, bool)

	GetStructFieldOKAdvanced(val reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool)

	GetStructFieldOK2() (reflect.Value, reflect.Kind, bool, bool)

	GetStructFieldOKAdvanced2(val reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool, bool)
}

var _ FieldLevel = new(validate)

func (v *validate) Field() reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func (v *validate) FieldName() string { _ = "STUB: not implemented"; return "" }

func (v *validate) GetTag() string { _ = "STUB: not implemented"; return "" }

func (v *validate) StructFieldName() string { _ = "STUB: not implemented"; return "" }

func (v *validate) Param() string { _ = "STUB: not implemented"; return "" }

func (v *validate) GetStructFieldOK() (reflect.Value, reflect.Kind, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Kind), false
}

func (v *validate) GetStructFieldOKAdvanced(val reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Kind), false
}

func (v *validate) GetStructFieldOK2() (reflect.Value, reflect.Kind, bool, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Kind), false, false
}

func (v *validate) GetStructFieldOKAdvanced2(val reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Kind), false, false
}
