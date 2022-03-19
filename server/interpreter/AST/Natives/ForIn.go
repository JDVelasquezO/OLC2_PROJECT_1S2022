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

type ForIn struct {
	Iterator       Abstract.Expression
	ValueToIterate Abstract.Expression
	Instructions   *arrayList.List
}

func NewForIn(iterator Abstract.Expression, valToIterate Abstract.Expression, instructions *arrayList.List) ForIn {
	return ForIn{
		Iterator:       iterator,
		ValueToIterate: valToIterate,
		Instructions:   instructions,
	}
}

func (f ForIn) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	newArr := arrayList.New()
	newArr.Add(f.Iterator)

	initVal := f.ValueToIterate.(Expression.Range).Expr1.GetValue(symbolTable).Value
	finishVal := f.ValueToIterate.(Expression.Range).Expr2.GetValue(symbolTable).Value

	if typeof(initVal) != "int" || typeof(finishVal) != "int" {
		errors.CounterError += 1
		row := f.ValueToIterate.(Expression.Range).Row
		col := f.ValueToIterate.(Expression.Range).Col
		msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") El valor de rango no es entero. \n"
		err := errors.NewError(errors.CounterError, row, col, msg, symbolTable.Name)
		errors.TypeError = append(errors.TypeError, err)
		interpreter.Console += fmt.Sprintf("%v", err.Msg)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
	}

	var newTable SymbolTable.SymbolTable
	for i := initVal.(int); i < finishVal.(int); i++ {

		if newTable.Table == nil {
			newTable = SymbolTable.NewSymbolTable("while", &symbolTable)
		} else {
			oldTable := newTable
			newTable = SymbolTable.NewSymbolTable("while", &oldTable)
		}

		newExpr := Expression.NewPrimitive(i, SymbolTable.INTEGER, 0, 0)
		newDeclaration := NewDeclarationInit(newArr, newExpr, true, "i64")
		newDeclaration.Execute(newTable)
		for j := 0; j < f.Instructions.Len(); j++ {
			instr := f.Instructions.GetValue(j).(Abstract.Instruction)
			instr.Execute(newTable)
		}
	}

	return nil
}
