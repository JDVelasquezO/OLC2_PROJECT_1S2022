package Expression

import (
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type Cast struct {
	Expression Abstract.Expression
	Type       SymbolTable.DataType
}

func (c Cast) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	res := c.GetValue(symbolTable)

	return res
}

func NewCast(expr Abstract.Expression, dataType SymbolTable.DataType) Cast {
	return Cast{
		Expression: expr,
		Type:       dataType,
	}
}

func (c Cast) GetValue(table SymbolTable.SymbolTable) SymbolTable.ReturnType {
	var resRet SymbolTable.ReturnType
	if c.Type == SymbolTable.INTEGER {
		newVal := int(c.Expression.GetValue(table).Value.(float64))
		resRet.Value = newVal
		resRet.Type = SymbolTable.INTEGER

	} else if c.Type == SymbolTable.FLOAT {
		newVal := float64(c.Expression.GetValue(table).Value.(int))
		resRet.Value = newVal
		resRet.Type = SymbolTable.FLOAT
	}

	return resRet
}
