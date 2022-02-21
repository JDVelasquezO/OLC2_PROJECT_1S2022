package SymbolTable

type Symbol struct {
	Id       string
	DataType DataType
	Value    interface{}
	Row      int
	Col      int
	isConst  bool
	isFunc   bool
}

func NewSymbolId(id string, row int, col int, dataType DataType, value interface{}) Symbol {
	in := Symbol{Id: id, Row: row, Col: col, isConst: false, isFunc: false, Value: value, DataType: dataType}
	return in
}
