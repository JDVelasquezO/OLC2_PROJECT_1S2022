package Expression

import (
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type Cast struct {
	Expression Abstract.Expression
	Type       SymbolTable.DataType
	Row        int
	Col        int
}

func (c Cast) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	res := c.GetValue(symbolTable)

	return res
}

func NewCast(expr Abstract.Expression, dataType SymbolTable.DataType, row int, col int) Cast {
	return Cast{
		Expression: expr,
		Type:       dataType,
		Row:        row,
		Col:        col,
	}
}

func (c Cast) GetValue(table SymbolTable.SymbolTable) SymbolTable.ReturnType {
	var resRet SymbolTable.ReturnType
	if c.Type == SymbolTable.INTEGER {
		newVal := c.Expression.GetValue(table)
		if newVal.Type == SymbolTable.ERROR {
			return newVal
		}
		resRet.Value = int(newVal.Value.(float64))
		resRet.Type = SymbolTable.INTEGER

	} else if c.Type == SymbolTable.FLOAT {
		newVal := float64(c.Expression.GetValue(table).Value.(int))
		resRet.Value = newVal
		resRet.Type = SymbolTable.FLOAT
	}

	return resRet
}
