package Environment

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/AST/Natives"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"reflect"
	"strconv"
)

var typeDef = [7][7]SymbolTable.DataType{
	{SymbolTable.INTEGER, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.FLOAT, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.STR, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.STRING, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.CHAR, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.BOOLEAN, SymbolTable.NULL},
	{SymbolTable.INTEGER, SymbolTable.FLOAT, SymbolTable.STR, SymbolTable.STRING, SymbolTable.CHAR, SymbolTable.BOOLEAN, SymbolTable.NULL},
}

type Function struct {
	SymbolTable.Symbol
	ListParams    *arrayList.List
	ListInstructs *arrayList.List
	Line          int
	Col           int
}

func NewFunction(line int, col int, name string, listParams *arrayList.List, listInstructs *arrayList.List,
	dataType SymbolTable.DataType) Function {
	funcSymbol := SymbolTable.NewSymbolFunction(line, col, name, dataType, listParams)

	return Function{
		ListInstructs: listInstructs,
		ListParams:    listParams,
		Symbol:        funcSymbol,
	}
}

func (f Function) ExecuteParams(table SymbolTable.SymbolTable, expression *arrayList.List) bool {
	declarations := f.ListParams.Clone()

	if declarations.Len() != expression.Len() {
		errors.CounterError += 1
		msg := "Se esperaban " + strconv.Itoa(expression.Len()) + " parametros. Llegaron " + strconv.Itoa(declarations.Len()) + " \n"
		err := errors.NewError(errors.CounterError, 0, 0, msg, table.Name)
		errors.TypeError = append(errors.TypeError, err)
		interpreter.Console += fmt.Sprintf("%v", err.Msg)
		return false
	}

	for i := 0; i < declarations.Len(); i++ {
		pivot := declarations.GetValue(i).(*Natives.Declaration)
		pivot.InitVal = expression.GetValue(i).(Abstract.Expression)
		res := pivot.Execute(table)

		if res.(SymbolTable.ReturnType).Type == SymbolTable.ERROR {
			return false
		}
	}

	return true
}

func (f Function) Execute(table SymbolTable.SymbolTable) interface{} {
	for i := 0; i < f.ListInstructs.Len(); i++ {
		instrPivot := f.ListInstructs.GetValue(i).(Abstract.Instruction)
		valInstr := instrPivot.Execute(table)
		typeRet := f.DataType

		if valInstr != nil {
			if reflect.TypeOf(valInstr) != reflect.TypeOf(SymbolTable.ReturnType{}) {
				fmt.Println("Error en funcion, se esperaba un retorno valido")
				return SymbolTable.ReturnType{
					Type:  SymbolTable.NULL,
					Value: -1,
				}
			}

			valRet := valInstr.(SymbolTable.ReturnType)
			cmpTypes := typeDef[typeRet][valRet.Type]

			if cmpTypes == SymbolTable.NULL {
				fmt.Println("Error de tipos de datos")
				return SymbolTable.ReturnType{
					Type:  SymbolTable.NULL,
					Value: -1,
				}
			}

			return valRet
		}
	}

	return SymbolTable.ReturnType{
		Type:  SymbolTable.NULL,
		Value: -1,
	}
}
