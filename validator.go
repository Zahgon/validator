package validator

import (
	"context"
	"reflect"
)

type validate struct {
	v              *Validate
	top            reflect.Value
	ns             []byte
	actualNs       []byte
	errs           ValidationErrors
	includeExclude map[string]struct{}
	ffn            FilterFunc
	slflParent     reflect.Value
	slCurrent      reflect.Value
	flField        reflect.Value
	cf             *cField
	ct             *cTag
	misc           []byte
	str1           string
	str2           string
	fldIsPointer   bool
	isPartial      bool
	hasExcludes    bool
}

func (v *validate) validateStruct(ctx context.Context, parent reflect.Value, current reflect.Value, typ reflect.Type, ns []byte, structNs []byte, ct *cTag) {
	_ = "STUB: not implemented"
	return
}

func (v *validate) traverseField(ctx context.Context, parent reflect.Value, current reflect.Value, ns []byte, structNs []byte, cf *cField, ct *cTag) {
	_ = "STUB: not implemented"
	return
}

func appendAltName(ns []byte, altName string) string { _ = "STUB: not implemented"; return "" }

func getValue(val reflect.Value) interface{} { _ = "STUB: not implemented"; return nil }
