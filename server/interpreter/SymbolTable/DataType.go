package SymbolTable

type DataType int

const (
	INTEGER DataType = iota
	FLOAT
	STR
	STRING
	CHAR
	BOOLEAN
	NULL
	ERROR
	VOID
)

type ReturnType struct {
	Type  DataType
	Value interface{}
}
