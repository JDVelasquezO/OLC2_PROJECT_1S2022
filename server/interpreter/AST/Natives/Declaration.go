package Natives

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/errors"
	"encoding/json"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"strconv"
)

var typeDef = [6][6]SymbolTable.DataType{
	{SymbolTable.INTEGER, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.FLOAT, SymbolTable.FLOAT, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.STR, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.CHAR, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.BOOLEAN, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
}

type Declaration struct {
	InitVal  Abstract.Expression
	DataType SymbolTable.DataType
	ListIds  *arrayList.List
	IsMut    bool
}

func NewDeclaration(listIds *arrayList.List, isMut bool, dataType string) *Declaration {

	newDec := Declaration{
		ListIds: listIds,
		InitVal: nil,
		IsMut:   isMut,
	}

	if dataType != "" {
		switch dataType {
		case "i64":
			newDec.DataType = SymbolTable.INTEGER
		case "f64":
			newDec.DataType = SymbolTable.FLOAT
		case "String":
			newDec.DataType = SymbolTable.STRING
		case "char":
			newDec.DataType = SymbolTable.CHAR
		case "bool":
			newDec.DataType = SymbolTable.BOOLEAN
		}
	} else {
		newDec.DataType = SymbolTable.NULL
	}

	return &newDec
}

func NewDeclarationInit(listIds *arrayList.List, valInit Abstract.Expression, isMut bool, dataType string) *Declaration {
	newDec := Declaration{
		ListIds: listIds,
		InitVal: valInit,
		IsMut:   isMut,
	}

	if dataType != "" {
		switch dataType {
		case "i64":
			newDec.DataType = SymbolTable.INTEGER
		case "f64":
			newDec.DataType = SymbolTable.FLOAT
		case "String":
			newDec.DataType = SymbolTable.STRING
		case "char":
			newDec.DataType = SymbolTable.CHAR
		case "bool":
			newDec.DataType = SymbolTable.BOOLEAN
		}
	} else {
		newDec.DataType = SymbolTable.NULL
	}

	return &newDec
}

func (d *Declaration) IsInitialized() bool {
	return d.InitVal != nil
}

func (d *Declaration) Execute(table SymbolTable.SymbolTable) interface{} {
	if d.IsInitialized() {
		if d.ListIds.Len() > 1 {
			return nil
		}

		dataOrigin := d.InitVal.GetValue(table).Type
		if (dataOrigin != d.DataType) && (d.DataType != SymbolTable.NULL) {
			row := d.InitVal.(Expression.Primitive).Row
			col := d.InitVal.(Expression.Primitive).Col
			errors.CounterError += 1
			msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") Tipos de datos incorrectos. \n"
			err := errors.NewError(errors.CounterError, row, col, msg, table.Name)
			errors.TypeError = append(errors.TypeError, err)
			interpreter.Console += fmt.Sprintf("%v", err.Msg)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
		}

		switch d.InitVal.GetValue(table).Type {
		case SymbolTable.INTEGER:
			d.DataType = SymbolTable.INTEGER
		case SymbolTable.FLOAT:
			d.DataType = SymbolTable.FLOAT
		case SymbolTable.STR:
			d.DataType = SymbolTable.STR
		case SymbolTable.CHAR:
			d.DataType = SymbolTable.CHAR
		case SymbolTable.BOOLEAN:
			d.DataType = SymbolTable.BOOLEAN
		}

		retExpr := d.InitVal.GetValue(table)
		typeExpr := retExpr.Type
		typeDec := d.DataType
		typeRes := typeDef[typeDec][typeExpr]

		if typeRes == SymbolTable.NULL {
			row := d.InitVal.(Expression.Primitive).Row
			col := d.InitVal.(Expression.Primitive).Col
			errors.CounterError += 1
			msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No concuerdan los tipos de datos \n"
			err := errors.NewError(errors.CounterError, row, col, msg, table.Name)
			errors.TypeError = append(errors.TypeError, err)
			interpreter.Console += fmt.Sprintf("%v", err.Msg)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
		}

		for i := 0; i < d.ListIds.Len(); i++ {
			varDec := d.ListIds.GetValue(i).(Expression.Identifier)

			if table.ExistsSymbol(varDec.Id) {
				fmt.Println("Error: Variable declarada")
			} else {
				symbol := SymbolTable.NewSymbolId(varDec.Id, 0, 0, typeRes, retExpr.Value, d.IsMut)
				table.AddNewSymbol(varDec.Id, symbol)
			}
		}
	} else {
		typeDec := d.DataType
		for i := 0; i < d.ListIds.Len(); i++ {
			varDec := d.ListIds.GetValue(i).(Expression.Identifier)

			if table.ExistsSymbol(varDec.Id) {
				fmt.Println("Error: Variable declarada")
			} else {
				symbol := SymbolTable.NewSymbolId(varDec.Id, 0, 0, typeDec, nil, d.IsMut)
				table.AddNewSymbol(varDec.Id, symbol)
			}
		}
	}

	data, err := json.MarshalIndent(table, "", "  ")
	if err != nil {
		panic(err)
	}

	stringQuery := string(data)
	fmt.Println(stringQuery)
	fmt.Printf("%v\n", d.ListIds)

	return nil
}
