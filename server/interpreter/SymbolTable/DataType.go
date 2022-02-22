package SymbolTable

type DataType int

const (
	INTEGER DataType = iota
	FLOAT
	STRING
	CHAR
	BOOLEAN
	NULL
	ERROR
)

type ReturnType struct {
	Type  DataType
	Value interface{}
}
