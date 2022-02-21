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

var typeDef = [5][5]SymbolTable.DataType{
	{SymbolTable.INTEGER, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.FLOAT, SymbolTable.FLOAT, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.STRING, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.BOOLEAN, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
}

type Declaration struct {
	InitVal  Abstract.Expression
	DataType SymbolTable.DataType
	ListIds  *arrayList.List
}

func NewDeclaration(listIds *arrayList.List, dataType SymbolTable.DataType) *Declaration {
	return &Declaration{
		DataType: dataType,
		ListIds:  listIds,
		InitVal:  nil,
	}
}

func NewDeclarationInit(listIds *arrayList.List, dataType SymbolTable.DataType, valInit Abstract.Expression) *Declaration {
	return &Declaration{DataType: dataType, ListIds: listIds, InitVal: valInit}
}

func (d *Declaration) IsInitialized() bool {
	return d.InitVal != nil
}

func (d *Declaration) Execute(table SymbolTable.SymbolTable) interface{} {
	if d.IsInitialized() {
		if d.ListIds.Len() > 1 {
			return nil
		}

		retExpr := d.InitVal.GetValue(table)
		typeExpr := retExpr.Type
		typeDec := d.DataType
		typeRes := typeDef[typeDec][typeExpr]

		if typeRes == SymbolTable.NULL {
			row := d.InitVal.(Expression.Primitive).Row
			col := d.InitVal.(Expression.Primitive).Col
			errors.CounterError += 1
			err := errors.Error{Id: errors.CounterError, Line: row, Col: col, Msg: "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No concuerdan los tipos de datos", Type: "Semantic", Ambit: "global"}
			errors.TypeError = append(errors.TypeError, err)
			interpreter.Console += fmt.Sprintf("%v", err.Msg)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
		}

		for i := 0; i < d.ListIds.Len(); i++ {
			varDec := d.ListIds.GetValue(i).(Expression.Identifier)

			if table.ExistsSymbol(varDec.Id) {
				fmt.Println("Error: Variable declarada")
			} else {
				symbol := SymbolTable.NewSymbolId(varDec.Id, 0, 0, typeRes, retExpr.Value)
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
				symbol := SymbolTable.NewSymbolId(varDec.Id, 0, 0, typeDec, nil)
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
