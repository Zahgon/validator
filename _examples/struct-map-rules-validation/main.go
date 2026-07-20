package main

import (
	"github.com/go-playground/validator/v10"
)

type Data struct {
	Name    string
	Email   string
	Details *Details
}

type Details struct {
	FamilyMembers *FamilyMembers
	Salary        string
}

type FamilyMembers struct {
	FatherName string
	MotherName string
}

type Data2 struct {
	Name string
	Age  uint32
}

var validate = validator.New()

func main() {
	validateStruct()

	validateStructNested()

}

func validateStruct() { _ = "STUB: not implemented"; return }

func validateStructNested() { _ = "STUB: not implemented"; return }
