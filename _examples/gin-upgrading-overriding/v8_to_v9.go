package main

import (
	"reflect"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &defaultValidator{}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *defaultValidator) Engine() interface{} { _ = "STUB: not implemented"; return nil }

func (v *defaultValidator) lazyinit() { _ = "STUB: not implemented"; return }

func kindOfData(data interface{}) reflect.Kind {
	_ = "STUB: not implemented"
	return *new(reflect.Kind)
}
