package Natives

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"strconv"
)

type Assign struct {
	Val     Abstract.Expression
	ListIds *arrayList.List
}

func NewAssign(listIds *arrayList.List, val Abstract.Expression) *Assign {
	return &Assign{
		ListIds: listIds,
		Val:     val,
	}
}

func (d *Assign) Execute(table SymbolTable.SymbolTable) interface{} {
	if d.ListIds.Len() > 1 {
		return nil
	}

	if d.Val == nil {
		row := d.ListIds.GetValue(0).(Expression.Identifier).Row
		col := d.ListIds.GetValue(0).(Expression.Identifier).Col
		errors.CounterError += 1
		err := errors.Error{Id: errors.CounterError, Line: row, Col: col, Msg: "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No se puede mezclar esos tipos de datos", Type: "Semantic", Ambit: "global"}
		errors.TypeError = append(errors.TypeError, err)
		interpreter.Console += fmt.Sprintf("%v", err.Msg)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
	}

	retExpr := d.Val.GetValue(table)
	typeExpr := retExpr.Type
	varDec := d.ListIds.GetValue(0).(Expression.Identifier).Id
	val := table.GetSymbol(varDec)

	typeDec := val.DataType
	typeRes := typeDef[typeDec][typeExpr]

	if typeRes == SymbolTable.NULL {
		row := d.Val.(Expression.Primitive).Row
		col := d.Val.(Expression.Primitive).Col
		errors.CounterError += 1
		err := errors.Error{Id: errors.CounterError, Line: row, Col: col, Msg: "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No se puede mezclar esos tipos de datos", Type: "Semantic", Ambit: "global"}
		errors.TypeError = append(errors.TypeError, err)
		interpreter.Console += fmt.Sprintf("%v", err.Msg)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
	}

	val.Value = d.Val.GetValue(table).Value
	table.ChangeVal(varDec, val)

	return nil
}
