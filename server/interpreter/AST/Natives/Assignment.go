package Natives

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Array"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Vector"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"reflect"
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
		msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No se puede mezclar esos tipos de datos \n"
		err := errors.NewError(errors.CounterError, row, col, msg, table.Name)
		errors.TypeError = append(errors.TypeError, err)
		interpreter.Console += fmt.Sprintf("%v", err.Msg)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
	}

	retExpr := d.Val.GetValue(table)
	typeExpr := retExpr.Type
	varDec := d.ListIds.GetValue(0).(Expression.Identifier).Id
	var val SymbolTable.Symbol
	if typeExpr == SymbolTable.ARRAY {
		valBefore := table.GetSymbolArray(varDec)
		if typeof(valBefore) == "Array.Array" {
			val = valBefore.(Array.Array).Symbol
		} else {
			val = valBefore.(Vector.Vector).Symbol
		}
	} else {
		val = table.GetSymbol(varDec).(SymbolTable.Symbol)
	}

	if val.IsConst {
		typeVal := typeof(d.Val)
		var row int
		var col int
		switch typeVal {
		case "Expression.Primitive":
			row = d.Val.(Expression.Primitive).Row
			col = d.Val.(Expression.Primitive).Col
		case "Expression.Operation":
			row = d.Val.(Expression.Operation).Row
			col = d.Val.(Expression.Operation).Col
		case "Expression.Identifier":
			row = d.Val.(Expression.Identifier).Row
			col = d.Val.(Expression.Identifier).Col
		}
		errors.CounterError += 1
		msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No se puede asignar valor a un no mut \n"
		err := errors.NewError(errors.CounterError, row, col, msg, table.Name)
		interpreter.Console += fmt.Sprintf("%v", err.Msg)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
	}

	typeDec := val.DataType
	typeRes := typeDef[typeDec][typeExpr]

	if typeRes == SymbolTable.NULL {
		typeVal := typeof(d.Val)
		var row int
		var col int
		switch typeVal {
		case "Expression.Primitive":
			row = d.Val.(Expression.Primitive).Row
			col = d.Val.(Expression.Primitive).Col
		case "Expression.Operation":
			row = d.Val.(Expression.Operation).Row
			col = d.Val.(Expression.Operation).Col
		case "Expression.Identifier":
			row = d.Val.(Expression.Identifier).Row
			col = d.Val.(Expression.Identifier).Col
		}

		errors.CounterError += 1
		msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No se puede mezclar esos tipos de datos \n"
		err := errors.NewError(errors.CounterError, row, col, msg, table.Name)
		interpreter.Console += fmt.Sprintf("%v", err.Msg)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
	}

	if retExpr.Type == SymbolTable.ARRAY {
		val.Value = retExpr.Value.(SymbolTable.ReturnType).Value
		table.ChangeValArray(varDec, val)
	} else {
		val.Value = retExpr.Value
		table.ChangeVal(varDec, val)
	}

	return val
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
