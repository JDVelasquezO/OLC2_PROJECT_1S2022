package Abstract

import (
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
)

type Value struct {
	Value       interface{}
	Type        SymbolTable.DataType
	IsTemp      bool
	AuxType     string
	TrueLabel   string
	FalseLabel  string
	ValuesArray *arrayList.List
	TypeArray   interface{}
	Size        int
	IsNegative  bool
	IsLogical   bool
}

func NewValue(value interface{}, dataType SymbolTable.DataType, isTemp bool, auxType string) Value {
	return Value{
		Value:       value,
		Type:        dataType,
		IsTemp:      isTemp,
		AuxType:     auxType,
		TrueLabel:   "",
		FalseLabel:  "",
		ValuesArray: arrayList.New(),
		TypeArray:   nil,
		Size:        1,
		IsNegative:  false,
		IsLogical:   false,
	}
}

func (v Value) GetValuesArray() *arrayList.List {
	return v.ValuesArray
}

func (v Value) GetType() SymbolTable.DataType {
	return v.Type
}
