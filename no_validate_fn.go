//go:build !validator_novalidatefn

package validator

import (
	"reflect"
)

func isValidateFn(fl FieldLevel) bool { _ = "STUB: not implemented"; return false }

func tryCallValidateFn(field reflect.Value, validateFn string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
