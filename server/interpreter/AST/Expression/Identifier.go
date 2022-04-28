package Expression

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Array"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Vector"
	"strconv"
)

type Identifier struct {
	Id  string
	Row int
	Col int
}

func (id Identifier) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	generator.AddComment("---- Identifier ----")
	value := symbolTable.GetSymbol(id.Id)
	if value.(*SymbolTable.Symbol) == nil {
		return nil
	}

	temp := generator.AddTemp()
	//temp := "t" + strconv.Itoa(value.(*SymbolTable.Symbol).Pos)
	var tempPos string
	if value.(*SymbolTable.Symbol).Id != "" {
		tempPos = generator.AddTemp()
		//newPos := strconv.Itoa(value.(*SymbolTable.Symbol).Pos - 1)
		generator.AddExpression(tempPos, "P", strconv.Itoa(value.(*SymbolTable.Symbol).Pos), "+")
	}

	generator.GetStack(temp, tempPos)

	if value.(*SymbolTable.Symbol).DataType != SymbolTable.BOOLEAN {
		return Abstract.NewValue(temp, value.(*SymbolTable.Symbol).DataType, true, "")
	}

	valToSend := value.(*SymbolTable.Symbol)
	CheckLabelsId(generator, valToSend)
	generator.AddIf(temp, "1", "==", valToSend.LabelTrue)
	generator.AddGoTo(valToSend.LabelFalse)

	valRet := Abstract.NewValue(temp, SymbolTable.BOOLEAN, false, "")
	valRet.TrueLabel = valToSend.LabelTrue
	valRet.FalseLabel = valToSend.LabelFalse

	return valRet
}

func CheckLabelsId(generator *Generator.Generator, value *SymbolTable.Symbol) {
	value.LabelTrue = generator.NewLabel()
	value.LabelFalse = generator.NewLabel()
}

func (id Identifier) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	res := id.GetValue(symbolTable)

	return res
}

func NewIdentifier(id string, row int, col int) Identifier {
	return Identifier{
		Id:  id,
		Row: row,
		Col: col,
	}
}

func (id Identifier) GetValue(table SymbolTable.SymbolTable) SymbolTable.ReturnType {
	var founded = table.ExistsSymbol(id.Id)

	if !founded {
		var founded2 = table.ExistsArray(id.Id)
		if !founded2 {
			return SymbolTable.ReturnType{Value: nil, Type: SymbolTable.NULL}
		}
		symbol := table.GetSymbolArray(id.Id)

		if typeof(symbol) == "Array.Array" {
			dataType := symbol.(Array.Array).Symbol.DataType
			values := symbol.(Array.Array).Values
			return SymbolTable.ReturnType{
				Value: values,
				Type:  dataType,
			}
		}

		if typeof(symbol) == "Vector.Vector" {
			dataType := symbol.(Vector.Vector).Symbol.DataType
			values := symbol.(Vector.Vector).Value

			if len(values.([]interface{})) == 0 {
				values = append(values.([]interface{}), "")
			}

			return SymbolTable.ReturnType{
				Value: values,
				Type:  dataType,
			}
		}
	}

	symbol := table.GetSymbol(id.Id)

	if typeof(symbol) == "*SymbolTable.Symbol" {
		data := symbol.(*SymbolTable.Symbol)
		return SymbolTable.ReturnType{
			Value: data.Value,
			Type:  data.DataType,
		}
	}

	if typeof(symbol) == "Vector.Vector" {
		dataType := symbol.(Vector.Vector).Symbol.DataType
		values := symbol.(Vector.Vector).Value
		return SymbolTable.ReturnType{
			Value: values,
			Type:  dataType,
		}
	}

	return SymbolTable.ReturnType{Value: -1, Type: SymbolTable.NULL}
}
