package SymbolTable

import arrayList "github.com/colegno/arraylist"

type Symbol struct {
	Id         string
	DataType   DataType
	Value      interface{}
	Row        int
	Col        int
	IsConst    bool
	IsFunc     bool
	ListParams *arrayList.List
}

func NewSymbolId(id string, row int, col int, dataType DataType, value interface{}, isConst bool) Symbol {
	in := Symbol{Id: id, Row: row, Col: col, IsConst: isConst,
		IsFunc: false, Value: value, DataType: dataType}
	return in
}

func NewSymbolFunction(line int, col int, id string, typeRet DataType, listParams *arrayList.List) Symbol {
	e := Symbol{
		Row:        line,
		Col:        col,
		Id:         id,
		IsConst:    false,
		IsFunc:     true,
		Value:      nil,
		DataType:   typeRet,
		ListParams: listParams,
	}

	return e
}
