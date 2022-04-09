package DecArrays

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Array"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	"reflect"
	"strconv"
)

type DecArray struct {
	Length       int
	Id           string
	ValueInitial Abstract.Expression
	Type         SymbolTable.DataType
	IsMut        bool
	Positions    Abstract.Expression
	Row          int
	Col          int
}

func (d DecArray) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	//TODO implement me
	panic("implement me")
}

func NewDecArray(length int, id string, init Abstract.Expression,
	dataType SymbolTable.DataType, isMut bool, positions Abstract.Expression, row int, col int) DecArray {
	return DecArray{
		Length:       length,
		Id:           id,
		ValueInitial: init,
		Type:         dataType,
		IsMut:        isMut,
		Positions:    positions,
		Row:          row,
		Col:          col,
	}
}

func (d DecArray) Execute(table SymbolTable.SymbolTable) interface{} {
	intPos := d.Positions.GetValue(table).Value.(int)
	fmt.Println(intPos)
	valueDec := d.ValueInitial.GetValue(table)

	//if valueDec.Type != SymbolTable.VOID {
	//	fmt.Printf("%v", valueDec)
	//}

	if valueDec.Type != SymbolTable.ARRAY {
		fmt.Println("Err1")
		return nil
	}

	if reflect.TypeOf(valueDec.Value) != reflect.TypeOf(SymbolTable.ReturnType{}) {
		fmt.Println("Err2")
		return nil
	}

	objectArray := valueDec.Value.(SymbolTable.ReturnType)

	if objectArray.Type != d.Type {
		fmt.Println("Err3")
		return nil
	}

	if objectArray.Value.(Array.Array).ListIntDim.Len() > d.Length {
		errors.CounterError += 1
		msg := "(" + strconv.Itoa(0) + ", " + strconv.Itoa(0) + ")  Longitud de Arreglo Incorrecta \n"
		err := errors.NewError(errors.CounterError, 0, 0, msg, table.Name)
		errors.TypeError = append(errors.TypeError, err)
		interpreter.Console += fmt.Sprintf("%v", msg)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
	}

	if objectArray.Value.(Array.Array).ListIntDim.GetValue(0) != d.Positions.(Expression.Primitive).Value {
		errors.CounterError += 1
		msg := "(" + strconv.Itoa(0) + ", " + strconv.Itoa(0) + ")  No cumplen las posiciones \n"
		err := errors.NewError(errors.CounterError, 0, 0, msg, table.Name)
		errors.TypeError = append(errors.TypeError, err)
		interpreter.Console += fmt.Sprintf("%v", msg)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
	}

	if table.ExistsSymbol(d.Id) {
		fmt.Printf("Error, variable %s ya declarada", d.Id)
	} else {
		symbol := objectArray.Value.(Array.Array)
		symbol.Id = d.Id
		symbol.Row = d.Row
		symbol.Col = d.Col

		if !d.IsMut {
			symbol.IsConst = true
		}

		table.AddArray(d.Id, symbol)
	}

	//fmt.Printf("%v", table.Table)
	return nil
}
