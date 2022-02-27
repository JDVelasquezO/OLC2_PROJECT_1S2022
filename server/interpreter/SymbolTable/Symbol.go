package SymbolTable

type Symbol struct {
	Id       string
	DataType DataType
	Value    interface{}
	Row      int
	Col      int
	IsConst  bool
	IsFunc   bool
}

func NewSymbolId(id string, row int, col int, dataType DataType, value interface{}, isConst bool) Symbol {
	in := Symbol{Id: id, Row: row, Col: col, IsConst: isConst,
		IsFunc: false, Value: value, DataType: dataType}
	return in
}
