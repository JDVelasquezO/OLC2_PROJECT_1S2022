package Environment

import (
	"OLC2_Project1/server/Generator"
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

func (f Function) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	newTable := SymbolTable.NewSymbolTable("function", symbolTable)
	returnLabel := generator.NewLabel()
	newTable.ReturnLabel = returnLabel
	newTable.SizeTable = 1
	generator.FreeAllTemps()
	generator.AddBeginFunc(f.Id, f.DataType)
	for i := 0; i < f.ListInstructs.Len(); i++ {
		instr := f.ListInstructs.GetValue(i)
		instr.(Abstract.Instruction).Compile(&newTable, generator)
	}
	if f.DataType != SymbolTable.NULL {
		generator.SetLabel(returnLabel)
	}
	generator.AddEndFunc()
	generator.FreeAllTemps()

	return nil
}

func NewFunction(line int, col int, name string, listParams *arrayList.List, listInstructs *arrayList.List,
	dataType string) Function {

	var newDataType SymbolTable.DataType
	switch dataType {
	case "i64":
		newDataType = SymbolTable.INTEGER
	case "f64":
		newDataType = SymbolTable.FLOAT
	case "String":
		newDataType = SymbolTable.STRING
	case "&str":
		newDataType = SymbolTable.STR
	case "char":
		newDataType = SymbolTable.CHAR
	case "bool":
		newDataType = SymbolTable.BOOLEAN
	case "void":
		newDataType = SymbolTable.VOID
	}

	funcSymbol := SymbolTable.NewSymbolFunction(line, col, name, newDataType, listParams)

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

		if res != nil {
			if res.(SymbolTable.ReturnType).Type == SymbolTable.ERROR {
				return false
			}
		}
	}

	return true
}

func (f Function) Execute(table SymbolTable.SymbolTable) interface{} {
	for i := 0; i < f.ListInstructs.Len(); i++ {
		instrPivot := f.ListInstructs.GetValue(i).(Abstract.Instruction)

		if typeof(instrPivot) == "Natives.Break" || typeof(instrPivot) == "Natives.Continue" {
			errors.CounterError += 1

			var row int
			var col int
			var transfer string
			if typeof(instrPivot) == "Natives.Break" {
				row = instrPivot.(Natives.Break).Row
				col = instrPivot.(Natives.Break).Col
				transfer = "Break"
			}

			if typeof(instrPivot) == "Natives.Continue" {
				row = instrPivot.(Natives.Continue).Row
				col = instrPivot.(Natives.Continue).Col
				transfer = "Continue"
			}

			msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") Error: No se puede declarar " + transfer + " sin un ciclo \n"
			err := errors.NewError(errors.CounterError, 0, 0, msg, interpreter.GlobalTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			interpreter.Console += fmt.Sprintf("%v", msg)
			continue
		}

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
				errors.CounterError += 1
				row := f.Row
				col := f.Col
				msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") Error: La funcion \"" + f.Id + "\" No retorna tipo de dato válido \n"
				err := errors.NewError(errors.CounterError, row, col, msg, table.Name)
				errors.TypeError = append(errors.TypeError, err)
				interpreter.Console += fmt.Sprintf("%v", err.Msg)
				return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
			}

			return valRet
		}
	}

	return SymbolTable.ReturnType{
		Type:  SymbolTable.NULL,
		Value: -1,
	}
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
