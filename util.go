package validator

import (
	"reflect"
	"regexp"
)

type Valuer interface {
	ValidatorValue() any
}

func (v *validate) extractTypeInternal(current reflect.Value, nullable bool) (reflect.Value, reflect.Kind, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Kind), false
}

func (v *validate) getStructFieldOKInternal(val reflect.Value, namespace string) (current reflect.Value, kind reflect.Kind, nullable bool, found bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Kind), false, false
}

func asInt(param string) int64 { _ = "STUB: not implemented"; return 0 }

func asIntFromTimeDuration(param string) int64 { _ = "STUB: not implemented"; return 0 }

func asIntFromType(t reflect.Type, param string) int64 { _ = "STUB: not implemented"; return 0 }

func asUint(param string) uint64 { _ = "STUB: not implemented"; return 0 }

func asFloat64(param string) float64 { _ = "STUB: not implemented"; return 0 }

func asFloat32(param string) float64 { _ = "STUB: not implemented"; return 0 }

func asBool(param string) bool { _ = "STUB: not implemented"; return false }

func panicIf(err error) { _ = "STUB: not implemented"; return }

func fieldMatchesRegexByStringerValOrString(regexFn func() *regexp.Regexp, fl FieldLevel) bool {
	_ = "STUB: not implemented"
	return false
}
