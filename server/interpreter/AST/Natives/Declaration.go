package Natives

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"encoding/json"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

var typeDef = [8][8]SymbolTable.DataType{
	{SymbolTable.INTEGER, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.INTEGER},
	{SymbolTable.NULL, SymbolTable.FLOAT, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.STR, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.STRING, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.CHAR, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.BOOLEAN, SymbolTable.NULL},
	{SymbolTable.INTEGER, SymbolTable.FLOAT, SymbolTable.STR, SymbolTable.STRING, SymbolTable.CHAR, SymbolTable.BOOLEAN, SymbolTable.NULL},
	{SymbolTable.INTEGER, SymbolTable.FLOAT, SymbolTable.STR, SymbolTable.STRING, SymbolTable.CHAR, SymbolTable.BOOLEAN, SymbolTable.NULL},
}

type Declaration struct {
	InitVal  Abstract.Expression
	DataType SymbolTable.DataType
	ListIds  *arrayList.List
	IsMut    bool
}

func (d *Declaration) Compile(symbolTable SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	if d.InitVal == nil {
		switch d.DataType {
		case SymbolTable.INTEGER:
			d.InitVal = Expression.NewPrimitive(0, SymbolTable.INTEGER, 0, 0)
			break
		case SymbolTable.FLOAT:
			d.InitVal = Expression.NewPrimitive(0.0, SymbolTable.FLOAT, 0, 0)
			break
		case SymbolTable.CHAR:
			d.InitVal = Expression.NewPrimitive("", SymbolTable.CHAR, 0, 0)
			break
		case SymbolTable.BOOLEAN:
			d.InitVal = Expression.NewPrimitive(true, SymbolTable.BOOLEAN, 0, 0)
			break
		}
	}

	value := d.InitVal.(Abstract.Instruction).Compile(symbolTable, generator)
	fmt.Println(value)
	newVar := symbolTable.GetSymbol(d.ListIds.GetValue(0).(Expression.Identifier).Id)
	var tempPos int
	if newVar == nil {
		/*inHeap := value.(Abstract.Value).Type == SymbolTable.STRING
		newSymbol :=*/
	} else {
		tempPos = newVar.(SymbolTable.Symbol).Pos
	}

	if value.(Abstract.Value).Type == SymbolTable.BOOLEAN {
		ValueBoolean(value, tempPos, generator)
	} else {
		generator.SetStack(tempPos, value.(Abstract.Value).Value, true)
	}

	return nil
}

func ValueBoolean(value interface{}, tempPos int, generator *Generator.Generator) {
	tempLabel := generator.NewLabel()
	generator.SetLabel(value.(Abstract.Value).TrueLabel)
	generator.SetStack(tempPos, 1, true)
	generator.AddGoTo(tempLabel)

	generator.SetLabel(value.(Abstract.Value).FalseLabel)
	generator.SetStack(tempPos, 0, true)
	generator.SetLabel(tempLabel)
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
		case "&str":
			newDec.DataType = SymbolTable.STR
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
		case "&str":
			newDec.DataType = SymbolTable.STR
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

		dataOrigin := d.InitVal.GetValue(table)
		dataOriginType := dataOrigin.Type

		switch dataOrigin.Type {
		case SymbolTable.INTEGER:
			d.DataType = SymbolTable.INTEGER
		case SymbolTable.FLOAT:
			d.DataType = SymbolTable.FLOAT
		case SymbolTable.STR:
			d.DataType = SymbolTable.STR
		case SymbolTable.STRING:
			d.DataType = SymbolTable.STRING
		case SymbolTable.CHAR:
			d.DataType = SymbolTable.CHAR
		case SymbolTable.BOOLEAN:
			d.DataType = SymbolTable.BOOLEAN
		case SymbolTable.ARRAY:
			d.DataType = dataOrigin.Type
		case SymbolTable.OBJECT:
			d.DataType = dataOriginType
		case SymbolTable.NULL:
			return dataOrigin
		case SymbolTable.ERROR:
			return dataOrigin
		}

		//typeExpr := dataOrigin.Type
		typeDec := d.DataType
		var typeRes SymbolTable.DataType
		if typeDec == SymbolTable.ARRAY {
			typeRes = dataOriginType
		} else if typeDec == SymbolTable.OBJECT {
			typeRes = dataOriginType
		} else {
			typeRes = typeDef[typeDec][dataOriginType]
		}

		for i := 0; i < d.ListIds.Len(); i++ {
			varDec := d.ListIds.GetValue(i).(Expression.Identifier)
			if dataOriginType == SymbolTable.OBJECT {
				symbol := SymbolTable.NewSymbolId(varDec.Id, varDec.Row, varDec.Col, typeRes, dataOrigin.Value, !d.IsMut, true, "", "")
				table.AddObject(symbol.Id, symbol)
				break
			}

			symbol := SymbolTable.NewSymbolId(varDec.Id, varDec.Row, varDec.Col, typeRes, dataOrigin.Value, !d.IsMut, false, "", "")
			table.AddNewSymbol(varDec.Id, symbol)
		}
	} else {
		typeDec := d.DataType
		for i := 0; i < d.ListIds.Len(); i++ {
			varDec := d.ListIds.GetValue(i).(Expression.Identifier)

			if table.ExistsSymbol(varDec.Id) {
				fmt.Println("Error: Variable declarada")
			} else {
				symbol := SymbolTable.NewSymbolId(varDec.Id, 0, 0, typeDec, nil, !d.IsMut, false, "", "")
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

	return nil
}
