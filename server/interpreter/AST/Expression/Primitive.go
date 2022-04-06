package Expression

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type Primitive struct {
	Value      interface{}
	Type       SymbolTable.DataType
	Row        int
	Col        int
	TrueLabel  string
	FalseLabel string
}

func (p Primitive) Compile(symbolTable SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	if p.Type == SymbolTable.INTEGER || p.Type == SymbolTable.FLOAT || p.Type == SymbolTable.CHAR {
		return Abstract.NewValue(p.Value, p.Type, false, "")

	} else if p.Type == SymbolTable.BOOLEAN {
		//res := Abstract.NewValue(p.Value, p.Type, false, "")
		p.CheckLabels(generator)

		if p.Value.(bool) {
			generator.AddGoTo(p.TrueLabel)
			//generator.AddGoTo(p.FalseLabel)
		} else {
			generator.AddGoTo(p.FalseLabel)
			//generator.AddGoTo(p.TrueLabel)
		}

		res := Abstract.NewValue(p.Value, p.Type, false, "")
		res.TrueLabel = p.TrueLabel
		res.FalseLabel = p.FalseLabel

		return res

	} else if p.Type == SymbolTable.STR || p.Type == SymbolTable.STRING {
		retTemp := generator.AddTemp()
		generator.AddExpression(retTemp, "H", "", "")
		counter := 0
		for _, char := range p.Value.(string) {
			generator.SetHeap("H", int(char))
			generator.NextHeap()
			counter += 1
		}

		generator.SetHeap("H", -1)
		generator.NextHeap()

		retVal := Abstract.NewValue(retTemp, p.Type, true, "")
		retVal.Size = counter
		return retVal
	}

	return nil
}

func (p *Primitive) CheckLabels(generator *Generator.Generator) {
	if p.TrueLabel == "" {
		p.TrueLabel = generator.NewLabel()
	}

	if p.FalseLabel == "" {
		p.FalseLabel = generator.NewLabel()
	}
}

func (p Primitive) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	res := p.GetValue(symbolTable)

	return res
}

func (p Primitive) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	return SymbolTable.ReturnType{
		Type:  p.Type,
		Value: p.Value,
	}
}

func NewPrimitive(val interface{}, types SymbolTable.DataType, row int, col int) Primitive {
	e := Primitive{val, types, row, col, "", ""}
	return e
}
