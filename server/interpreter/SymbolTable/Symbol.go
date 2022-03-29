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
	InHeap     bool
	Pos        int
	LabelTrue  string
	LabelFalse string
	Size       int
}

func NewSymbolId(id string, row int, col int, dataType DataType, value interface{}, isConst bool,
	inHeap bool, labelTrue string, labelFalse string) Symbol {
	in := Symbol{
		Id:         id,
		Row:        row,
		Col:        col,
		IsConst:    isConst,
		IsFunc:     false,
		Value:      value,
		DataType:   dataType,
		InHeap:     inHeap,
		LabelTrue:  labelTrue,
		LabelFalse: labelFalse,
		Size:       1,
	}
	return in
}

func (s Symbol) SetPosition(pos int) {
	s.Pos = pos
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

func NewSymbolArray(line int, col int, id string, dataType DataType) Symbol {
	e := Symbol{
		Row:        line,
		Col:        col,
		Id:         id,
		IsConst:    false,
		IsFunc:     true,
		Value:      nil,
		DataType:   dataType,
		ListParams: nil,
	}

	return e
}

func NewSymbolVector(line int, col int, id string, dataType DataType) Symbol {
	e := Symbol{
		Row:        line,
		Col:        col,
		Id:         id,
		IsConst:    false,
		IsFunc:     true,
		Value:      nil,
		DataType:   dataType,
		ListParams: nil,
	}

	return e
}
