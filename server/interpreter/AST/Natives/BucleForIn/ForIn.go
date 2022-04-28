package BucleForIn

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/AST/ExpressionSpecial"
	"OLC2_Project1/server/interpreter/AST/Natives"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"reflect"
	"strconv"
)

type ForIn struct {
	Iterator       Abstract.Expression
	ValueToIterate Abstract.Expression
	Instructions   *arrayList.List
}

func (f ForIn) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	generator.AddComment("---- For Cycle ----")

	//var := Natives.NewDeclaration()

	listRank := f.ValueToIterate.Compile(symbolTable, generator).(Abstract.Value).Value
	list := listRank.(arrayList.List)
	expr1 := list.GetValue(0)
	expr2 := list.GetValue(1)

	idIterator := f.Iterator.(Expression.Identifier).Id
	symbol := SymbolTable.NewSymbolId(idIterator, 0, 0, SymbolTable.INTEGER, 0, false, false, "", "")
	symbolTable.AddNewSymbol(idIterator, &symbol)

	variable := symbolTable.GetSymbol(idIterator)
	tempPos := generator.AddTemp()
	generator.AddExpression(tempPos, "P", fmt.Sprintf("%v", variable.(*SymbolTable.Symbol).Pos), "+")

	generator.SetStack(tempPos, expr1.(Abstract.Value).Value, true)

	start := generator.NewLabel()
	labelTrue := generator.NewLabel()
	exit := generator.NewLabel()

	t1 := generator.AddTemp()
	t4 := generator.AddTemp()
	t5 := generator.AddTemp()

	generator.SetLabel(start)

	generator.GetStack(t1, tempPos)
	generator.AddIf(t1, fmt.Sprintf("%v", expr2.(Abstract.Value).Value), "<", labelTrue)
	generator.AddGoTo(exit)

	generator.SetLabel(labelTrue)

	for i := 0; i < f.Instructions.Len(); i++ {
		f.Instructions.GetValue(i).(Abstract.Instruction).Compile(symbolTable, generator)
	}

	generator.GetStack(t4, tempPos)
	generator.AddExpression(t5, t4, "1", "+")
	generator.SetStack(tempPos, t5, true)

	generator.AddGoTo(start)
	generator.SetLabel(exit)

	return nil
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
	var newTable SymbolTable.SymbolTable
	if typeof(f.ValueToIterate) == "Expression.Range" {

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

		for i := initVal.(int); i < finishVal.(int); i++ {

			if newTable.Table == nil {
				newTable = SymbolTable.NewSymbolTable("forIn", &symbolTable)
			} else {
				oldTable := newTable
				newTable = SymbolTable.NewSymbolTable("forIn", &oldTable)
			}

			newExpr := Expression.NewPrimitive(i, SymbolTable.INTEGER, 0, 0)
			newDeclaration := Natives.NewDeclarationInit(newArr, newExpr, true, "i64")
			newDeclaration.Execute(newTable)
			for j := 0; j < f.Instructions.Len(); j++ {
				instr := f.Instructions.GetValue(j).(Abstract.Instruction)
				instr.Execute(newTable)
			}
		}

	} else if typeof(f.ValueToIterate) == "ExpressionSpecial.ValueArray" {
		for i := 0; i < f.ValueToIterate.(ExpressionSpecial.ValueArray).Expressions.Len(); i++ {
			if newTable.Table == nil {
				newTable = SymbolTable.NewSymbolTable("forIn", &symbolTable)
			} else {
				oldTable := newTable
				newTable = SymbolTable.NewSymbolTable("forIn", &oldTable)
			}

			newExpr := f.ValueToIterate.(ExpressionSpecial.ValueArray).Expressions.GetValue(i)
			dataType := newExpr.(Abstract.Expression).GetValue(symbolTable).Type

			var data string
			switch dataType {
			case SymbolTable.INTEGER:
				data = "i64"
			case SymbolTable.FLOAT:
				data = "f64"
			case SymbolTable.STR:
				data = "&str"
			case SymbolTable.STRING:
				data = "String"
			case SymbolTable.BOOLEAN:
				data = "bool"
			}

			newDeclaration := Natives.NewDeclarationInit(newArr, newExpr.(Abstract.Expression), true, data)
			newDeclaration.Execute(newTable)
			for j := 0; j < f.Instructions.Len(); j++ {
				instr := f.Instructions.GetValue(j).(Abstract.Instruction)
				instr.Execute(newTable)
			}
		}
	}

	return nil
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
