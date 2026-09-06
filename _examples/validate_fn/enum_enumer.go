package main

const _EnumName = "ZeroOneTwoThree"

var _EnumIndex = [...]uint8{0, 4, 7, 10, 15}

const _EnumLowerName = "zeroonetwothree"

func (i Enum) String() string { _ = "STUB: not implemented"; return "" }

func _EnumNoOp() { _ = "STUB: not implemented"; return }

var _EnumValues = []Enum{Zero, One, Two, Three}

var _EnumNameToValueMap = map[string]Enum{
	_EnumName[0:4]:        Zero,
	_EnumLowerName[0:4]:   Zero,
	_EnumName[4:7]:        One,
	_EnumLowerName[4:7]:   One,
	_EnumName[7:10]:       Two,
	_EnumLowerName[7:10]:  Two,
	_EnumName[10:15]:      Three,
	_EnumLowerName[10:15]: Three,
}

var _EnumNames = []string{
	_EnumName[0:4],
	_EnumName[4:7],
	_EnumName[7:10],
	_EnumName[10:15],
}

func EnumString(s string) (Enum, error) { _ = "STUB: not implemented"; return *new(Enum), nil }

func EnumValues() []Enum { _ = "STUB: not implemented"; return nil }

func EnumStrings() []string { _ = "STUB: not implemented"; return nil }

func (i Enum) IsAEnum() bool { _ = "STUB: not implemented"; return false }
