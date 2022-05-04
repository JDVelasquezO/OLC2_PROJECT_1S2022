package Expression

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

var sum = [5][5]SymbolTable.DataType{
	{SymbolTable.INTEGER, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.FLOAT, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.STRING, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
}

var multiDiv = [5][5]SymbolTable.DataType{
	{SymbolTable.INTEGER, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.FLOAT, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
}

var sub = [5][5]SymbolTable.DataType{
	{SymbolTable.INTEGER, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.FLOAT, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
}

var relational = [6][6]SymbolTable.DataType{
	{SymbolTable.INTEGER, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.FLOAT, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.STR, SymbolTable.STRING, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.STRING, SymbolTable.BOOLEAN, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.BOOLEAN},
}

type Operation struct {
	OpLeft     Abstract.Expression
	Operator   string
	OpRight    Abstract.Expression
	Unary      bool
	Row        int
	Col        int
	LabelTrue  string
	LabelFalse string
}

func (o Operation) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	if o.OpRight != nil {
		if o.Operator == "&&" || o.Operator == "||" {
			generator.AddComment("---- Logical ----")

			generator.AddComment("---- Op Left Logic ----")
			left := o.OpLeft.Compile(symbolTable, generator)
			ValueBoolean(left, generator)
			generator.AddComment("---- Op Right Logic ----")
			right := o.OpRight.Compile(symbolTable, generator)
			ValueBoolean(right, generator)
			LT1 := generator.NewLabel()
			LF1 := generator.NewLabel()
			LT2 := generator.NewLabel()
			LF2 := generator.NewLabel()
			LExit := generator.NewLabel()
			//o.CheckLabelsLogic(generator)
			//labelUnion := ""

			//var newVal bool
			var newTemp string
			res := Abstract.NewValue("", SymbolTable.BOOLEAN, false, "")
			switch o.Operator {
			case "&&":
				generator.AddComment("---- Validate Operator && ----")
				generator.AddIf(left.(Abstract.Value).Value.(string), "1", "==", LT1)
				generator.AddGoTo(LF1)
				generator.SetLabel(LT1)

				generator.AddIf(right.(Abstract.Value).Value.(string), "1", "==", LT2)
				generator.AddGoTo(LF2)
				generator.SetLabel(LT2)

				newTemp = generator.AddTemp()
				generator.AddExpression(newTemp, "1", "", "")
				generator.AddGoTo(LExit)

				generator.SetLabel(LF1)
				generator.SetLabel(LF2)
				generator.AddExpression(newTemp, "0", "", "")
				res.IsNegative = true
				generator.AddGoTo(LExit)

				generator.SetLabel(LExit)
				break
			case "||":
				generator.AddIf(left.(Abstract.Value).Value.(string), "1", "==", LT2)
				generator.AddGoTo(LT1)
				generator.SetLabel(LT1)

				generator.AddIf(right.(Abstract.Value).Value.(string), "1", "==", LT2)
				generator.AddGoTo(LF2)
				generator.SetLabel(LT2)

				newTemp = generator.AddTemp()
				generator.AddExpression(newTemp, "1", "", "")
				generator.AddGoTo(LExit)

				generator.SetLabel(LF1)
				generator.SetLabel(LF2)
				generator.AddExpression(newTemp, "0", "", "")
				generator.AddGoTo(LExit)

				generator.SetLabel(LExit)

				break
			}

			res.Value = newTemp
			res.TrueLabel = generator.NewLabel()
			res.FalseLabel = generator.NewLabel()
			res.IsLogical = true
			return res
		}

		left := o.OpLeft.Compile(symbolTable, generator)

		if left.(Abstract.Value).Type == SymbolTable.BOOLEAN {
			gotoRight := generator.NewLabel()
			leftTemp := generator.AddTemp()

			generator.SetLabel(left.(Abstract.Value).TrueLabel)
			generator.AddExpression(leftTemp, "1", "", "")
			generator.AddGoTo(gotoRight)

			generator.SetLabel(left.(Abstract.Value).FalseLabel)
			generator.AddExpression(leftTemp, "0", "", "")

			generator.SetLabel(gotoRight)

			right := o.OpRight.Compile(symbolTable, generator)
			gotoEnd := generator.NewLabel()
			rightTemp := generator.AddTemp()

			generator.SetLabel(right.(Abstract.Value).TrueLabel)
			generator.AddExpression(rightTemp, "1", "", "")
			generator.AddGoTo(gotoEnd)

			generator.SetLabel(right.(Abstract.Value).FalseLabel)
			generator.AddExpression(rightTemp, "0", "", "")

			generator.SetLabel(gotoEnd)
			o.CheckLabelsLogic(generator)
			generator.AddIf(leftTemp, rightTemp, o.Operator, o.LabelTrue)
			generator.AddGoTo(o.LabelFalse)

			res := Abstract.NewValue(true, SymbolTable.BOOLEAN, true, "")
			res.TrueLabel = o.LabelTrue
			res.FalseLabel = o.LabelFalse

			return res
		}

		operation := o.Operator

		if operation == "/" {
			valRight := o.OpRight.GetValue(*symbolTable).Value
			if valRight == 0 {
				//generator.AddError("No es posible dividir por 0", o.Row, o.Col)
				err := errors.NewError(0, o.Row, o.Col, "Error: Division entre 0\n", "local")
				val := Abstract.NewValue(err, SymbolTable.ERROR, false, "")
				return val
			}
		}

		right := o.OpRight.Compile(symbolTable, generator)
		temp := generator.AddTemp()

		if (left.(Abstract.Value).Type == SymbolTable.INTEGER || left.(Abstract.Value).Type == SymbolTable.FLOAT) &&
			(right.(Abstract.Value).Type == SymbolTable.INTEGER || right.(Abstract.Value).Type == SymbolTable.FLOAT) {

			generator.AddComment("---- Arithmetic Operation ----")
			valLeft := LookForDataType(left.(Abstract.Value).Value)
			valRight := LookForDataType(right.(Abstract.Value).Value)

			switch operation {
			case "pow", "powf":
				generator.AddExpression(temp, valLeft, valLeft, "*")
				break
			case ">", "<", ">=", "<=", "==", "!=":
				leftToSend := left.(Abstract.Value)
				CheckLabelsBool(generator, &leftToSend)
				generator.AddIf(valLeft, valRight, operation, leftToSend.TrueLabel)
				generator.AddGoTo(leftToSend.FalseLabel)

				//typeOf := interpreter.DataTypeRes
				leftToSend.Value = temp
				leftToSend.Type = SymbolTable.BOOLEAN
				return leftToSend
			default:
				generator.AddExpression(temp, valLeft, valRight, operation)
				break
			}

			typeOf := sum[left.(Abstract.Value).Type][right.(Abstract.Value).Type]
			return Abstract.NewValue(temp, typeOf, true, "")

		} else if (left.(Abstract.Value).Type == SymbolTable.STRING || left.(Abstract.Value).Type == SymbolTable.STR) &&
			(right.(Abstract.Value).Type == SymbolTable.STRING || right.(Abstract.Value).Type == SymbolTable.STR) {

			switch operation {
			case "+":
				generator.ConcatString()
				paramTemp := generator.AddTemp()
				generator.AddExpression(paramTemp, "P", strconv.Itoa(symbolTable.SizeTable), "+")

				generator.AddExpression(paramTemp, paramTemp, "1", "+")
				generator.SetStack(paramTemp, left.(Abstract.Value).Value, true)

				generator.AddExpression(paramTemp, paramTemp, "1", "+")
				generator.SetStack(paramTemp, right.(Abstract.Value).Value, true)

				generator.NewEnv(symbolTable.SizeTable)
				generator.CallFunc("concat")
				temp := generator.AddTemp()

				generator.GetStack(temp, "P")
				generator.SetEnv(symbolTable.SizeTable)

				retVal := Abstract.NewValue(temp, SymbolTable.STRING, true, "")
				retVal.Size = left.(Abstract.Value).Size + right.(Abstract.Value).Size
				return retVal
			case "<", ">", ">=", "<=", "==", "!=":
				generator.CompareString()
				paramTemp := generator.AddTemp()
				generator.AddExpression(paramTemp, "P", strconv.Itoa(symbolTable.SizeTable), "+")

				generator.AddExpression(paramTemp, paramTemp, "1", "+")
				generator.SetStack(paramTemp, left.(Abstract.Value).Value.(string), true)

				generator.AddExpression(paramTemp, paramTemp, "1", "+")
				generator.SetStack(paramTemp, right.(Abstract.Value).Value.(string), true)

				generator.NewEnv(symbolTable.SizeTable)
				generator.CallFunc("cmp")

				temp := generator.AddTemp()
				generator.GetStack(temp, "P")
				generator.SetEnv(symbolTable.SizeTable)

				trueLabel := generator.NewLabel()
				falseLabel := generator.NewLabel()

				if operation == "==" {
					generator.AddIf(temp, "0", "==", falseLabel)
					generator.AddGoTo(trueLabel)

					res := Abstract.NewValue(temp, SymbolTable.BOOLEAN, true, "")
					res.TrueLabel = trueLabel
					res.FalseLabel = falseLabel

					return res
				}

				if operation == "!=" {
					generator.AddIf(temp, "0", "!=", falseLabel)
					generator.AddGoTo(trueLabel)

					res := Abstract.NewValue(temp, SymbolTable.BOOLEAN, true, "")
					res.TrueLabel = trueLabel
					res.FalseLabel = falseLabel

					return res
				}
			}
		} else {
			valLeft := LookForDataType(left.(Abstract.Value).Value)
			valRight := LookForDataType(right.(Abstract.Value).Value)
			switch operation {
			case "<", ">", ">=", "<=", "==", "!=":
				leftToSend := left.(Abstract.Value)
				CheckLabelsBool(generator, &leftToSend)
				generator.AddIf(valLeft, valRight, operation, leftToSend.TrueLabel)
				generator.AddGoTo(leftToSend.FalseLabel)

				//typeOf := interpreter.DataTypeRes
				leftToSend.Value = temp
				leftToSend.Type = SymbolTable.BOOLEAN
				return leftToSend
			}
		}
	} else {
		left := o.OpLeft.Compile(symbolTable, generator)
		var value string
		if o.Operator == "-" {
			switch left.(Abstract.Value).Type {
			case SymbolTable.INTEGER:
				value = strconv.Itoa(left.(Abstract.Value).Value.(int) * -1)
				break
			case SymbolTable.FLOAT:
				value = fmt.Sprintf("%v", left.(Abstract.Value).Value.(float64)*-1.0)
				break
			}
			return Abstract.NewValue(value, left.(Abstract.Value).Type, false, "")
		} else if o.Operator == "!" {
			//value = fmt.Sprintf("%v", !left.(Abstract.Value).Value.(bool))
			newVal := Abstract.NewValue(value, left.(Abstract.Value).Type, false, "")
			newVal.IsNegative = !left.(Abstract.Value).IsNegative
			newVal.TrueLabel = left.(Abstract.Value).TrueLabel
			newVal.FalseLabel = left.(Abstract.Value).FalseLabel
			return newVal
		}
	}

	return nil
}

func ValueBoolean(value interface{}, generator *Generator.Generator) {
	tempLabel := generator.NewLabel()
	generator.SetLabel(value.(Abstract.Value).TrueLabel)
	//temp := generator.AddTemp()
	generator.AddExpression(value.(Abstract.Value).Value.(string), "1", "", "")
	generator.AddGoTo(tempLabel)

	generator.SetLabel(value.(Abstract.Value).FalseLabel)
	generator.AddExpression(value.(Abstract.Value).Value.(string), "0", "", "")
	generator.SetLabel(tempLabel)
}

func CheckLabelsBool(generator *Generator.Generator, value *Abstract.Value) {
	value.TrueLabel = generator.NewLabel()
	value.FalseLabel = generator.NewLabel()
}

func (o *Operation) CheckLabelsLogic(generator *Generator.Generator) {
	if o.LabelTrue == "" {
		o.LabelTrue = generator.NewLabel()
	}

	if o.LabelFalse == "" {
		o.LabelFalse = generator.NewLabel()
	}
}

func LookForDataType(val interface{}) string {
	var valRet string
	switch typeof(val) {
	case "int":
		valRet = strconv.Itoa(val.(int))
		break
	case "float64":
		valRet = fmt.Sprintf("%v", val)
		break
	case "string":
		valRet = val.(string)
		break
	}

	return valRet
}

func (o Operation) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	res := o.GetValue(symbolTable)
	return res
}

func NewOperation(OpLeft Abstract.Expression, Operator string, OpRight Abstract.Expression,
	unary bool, row int, col int) Operation {
	e := Operation{OpLeft, Operator, OpRight, unary, row, col, "", ""}
	return e
}

func (o Operation) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	retLeft := o.OpLeft.GetValue(symbolTable)
	var retRight SymbolTable.ReturnType

	if retLeft.Value == nil {

		row := strconv.Itoa(o.Row)
		col := strconv.Itoa(o.Col)
		errors.CounterError += 1
		msg := "(" + row + ", " + col + ") ID " + o.OpLeft.(Identifier).Id + " no declarado o asignado \n"
		err := errors.NewError(errors.CounterError, o.Row, o.Col, msg, symbolTable.Name)
		errors.TypeError = append(errors.TypeError, err)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
	}

	if o.Unary {
		retLeft = o.OpLeft.GetValue(symbolTable)
	} else {
		if o.OpRight.GetValue(symbolTable).Value == nil {
			row := strconv.Itoa(o.Row)
			col := strconv.Itoa(o.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") ID " + o.OpRight.(Identifier).Id + " no asignado \n"
			err := errors.NewError(errors.CounterError, o.Row, o.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

		retRight = o.OpRight.GetValue(symbolTable)
	}

	var priority SymbolTable.DataType

	switch o.Operator {
	case "+":
		if retLeft.Type == SymbolTable.BOOLEAN ||
			retLeft.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.CHAR ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.CHAR {

			goto ErrorDataType
		}

		priority = sum[retLeft.Type][retRight.Type]
		interpreter.DataTypeRes = priority

		if priority == SymbolTable.INTEGER {
			fmt.Println(retLeft.Type)
			fmt.Println(retRight.Type)
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(int) + retRight.Value.(int)}

		} else if priority == SymbolTable.FLOAT {
			op1, _ := strconv.ParseFloat(fmt.Sprintf("%v", retLeft.Value), 64)
			op2, _ := strconv.ParseFloat(fmt.Sprintf("%v", retRight.Value), 64)
			return SymbolTable.ReturnType{Type: priority, Value: op1 + op2}

		} else if priority == SymbolTable.STRING {
			r1 := fmt.Sprintf("%v", retLeft.Value)
			r2 := fmt.Sprintf("%v", retRight.Value)

			return SymbolTable.ReturnType{Type: priority, Value: r1 + r2}
		}

	case "-":
		if o.Unary {
			if retLeft.Type == SymbolTable.INTEGER {
				value := retLeft.Value.(int) * -1
				return SymbolTable.ReturnType{Type: retLeft.Type, Value: value}
			} else {
				value := retLeft.Value.(float64) * -1
				return SymbolTable.ReturnType{Type: retLeft.Type, Value: value}
			}
		}

		if retLeft.Type == SymbolTable.STR ||
			retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.CHAR ||
			retLeft.Type == SymbolTable.CHAR {

			goto ErrorDataType
		}

		priority = sub[retLeft.Type][retRight.Type]
		interpreter.DataTypeRes = priority
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(int) - retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(float64) - retRight.Value.(float64)}
		}

	case "*":
		if retLeft.Type == SymbolTable.STR ||
			retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.CHAR ||
			retLeft.Type == SymbolTable.CHAR {

			goto ErrorDataType
		}

		priority = multiDiv[retLeft.Type][retRight.Type]
		interpreter.DataTypeRes = priority
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(int) * retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(float64) * retRight.Value.(float64)}
		}

	case "/":
		if retLeft.Type == SymbolTable.STR ||
			retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.CHAR ||
			retLeft.Type == SymbolTable.CHAR {

			goto ErrorDataType
		}

		priority = multiDiv[retLeft.Type][retRight.Type]
		interpreter.DataTypeRes = priority
		if priority == SymbolTable.INTEGER {

			if retRight.Value.(int) == 0 {
				row := o.Row
				col := o.Col
				errors.CounterError += 1
				msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") Error: No se puede dividir dentro de 0 \n"
				err := errors.NewError(errors.CounterError, row, col, msg, symbolTable.Name)
				errors.TypeError = append(errors.TypeError, err)
				interpreter.Console += fmt.Sprintf("%v", err.Msg)
				return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
			}

			return SymbolTable.ReturnType{Type: SymbolTable.FLOAT, Value: float64(retLeft.Value.(int)) / float64(retRight.Value.(int))}
		} else if priority == SymbolTable.FLOAT {

			if retRight.Value.(float64) == 0 {
				if retRight.Value.(int) == 0 {
					row := 0
					col := 0
					errors.CounterError += 1
					msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No se puede dividir dentro de 0 \n"
					err := errors.NewError(errors.CounterError, row, col, msg, symbolTable.Name)
					errors.TypeError = append(errors.TypeError, err)
					interpreter.Console += fmt.Sprintf("%v", err.Msg)
					return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
				}
			}

			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(float64) / retRight.Value.(float64)}
		}

	case "%":
		if retLeft.Type == SymbolTable.STR ||
			retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.CHAR ||
			retLeft.Type == SymbolTable.CHAR {

			goto ErrorDataType
		}

		priority = multiDiv[retLeft.Type][retRight.Type]
		interpreter.DataTypeRes = priority
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(int) % retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: priority, Value: math.Mod(retLeft.Value.(float64), retRight.Value.(float64))}
		}

	case "pow", "powf":
		if retLeft.Type == SymbolTable.STR ||
			retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.CHAR ||
			retLeft.Type == SymbolTable.CHAR {

			goto ErrorDataType
		}

		priority = relational[retLeft.Type][retRight.Type]
		interpreter.DataTypeRes = priority
		if priority == SymbolTable.INTEGER && o.Operator == "pow" {
			return SymbolTable.ReturnType{Type: SymbolTable.FLOAT, Value: math.Pow(float64(retLeft.Value.(int)), float64(retRight.Value.(int)))}
		} else if priority == SymbolTable.FLOAT && o.Operator == "powf" {
			return SymbolTable.ReturnType{Type: priority, Value: math.Pow(retLeft.Value.(float64), retRight.Value.(float64))}
		} else {

			row := strconv.Itoa(o.Row)
			col := strconv.Itoa(o.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: el operador " + o.Operator + " no puede operarse con este tipo de dato \n"
			err := errors.NewError(errors.CounterError, o.Row, o.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			interpreter.Console += fmt.Sprintf("%v", err.Msg)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

	case ">":
		if retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.CHAR ||
			retLeft.Type == SymbolTable.CHAR {

			goto ErrorDataType
		}

		priority = relational[retLeft.Type][retRight.Type]
		interpreter.DataTypeRes = SymbolTable.BOOLEAN
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: retLeft.Value.(int) > retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: retLeft.Value.(float64) > retRight.Value.(float64)}
		} else if priority == SymbolTable.STR {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: len(retLeft.Value.(string)) > len(retRight.Value.(string))}
		}

	case "<":
		if retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.CHAR ||
			retLeft.Type == SymbolTable.CHAR {

			goto ErrorDataType
		}

		priority = relational[retLeft.Type][retRight.Type]
		interpreter.DataTypeRes = SymbolTable.BOOLEAN
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: retLeft.Value.(int) < retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: retLeft.Value.(float64) < retRight.Value.(float64)}
		} else if priority == SymbolTable.STR {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: len(retLeft.Value.(string)) < len(retRight.Value.(string))}
		}

	case ">=":
		if retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.CHAR ||
			retLeft.Type == SymbolTable.CHAR {

			goto ErrorDataType
		}

		priority = relational[retLeft.Type][retRight.Type]
		interpreter.DataTypeRes = SymbolTable.BOOLEAN
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: retLeft.Value.(int) >= retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: retLeft.Value.(float64) >= retRight.Value.(float64)}
		} else if priority == SymbolTable.STR {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: len(retLeft.Value.(string)) >= len(retRight.Value.(string))}
		}

	case "<=":
		if retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.CHAR ||
			retLeft.Type == SymbolTable.CHAR {

			goto ErrorDataType
		}

		priority = relational[retLeft.Type][retRight.Type]
		interpreter.DataTypeRes = SymbolTable.BOOLEAN
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: retLeft.Value.(int) <= retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: retLeft.Value.(float64) <= retRight.Value.(float64)}
		} else if priority == SymbolTable.STR {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: len(retLeft.Value.(string)) <= len(retRight.Value.(string))}
		}

	case "==":
		if retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.CHAR ||
			retLeft.Type == SymbolTable.CHAR {

			goto ErrorDataType
		}

		priority = relational[retLeft.Type][retRight.Type]
		interpreter.DataTypeRes = SymbolTable.BOOLEAN
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: retLeft.Value.(int) == retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: retLeft.Value.(float64) == retRight.Value.(float64)}
		} else if priority == SymbolTable.STR || priority == SymbolTable.STRING {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: len(retLeft.Value.(string)) == len(retRight.Value.(string))}
		} else if priority == SymbolTable.BOOLEAN {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: retLeft.Value.(bool) == retRight.Value.(bool)}
		}

	case "!=":
		if retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.CHAR ||
			retLeft.Type == SymbolTable.CHAR {

			goto ErrorDataType
		}

		priority = relational[retLeft.Type][retRight.Type]
		interpreter.DataTypeRes = SymbolTable.BOOLEAN
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: retLeft.Value.(int) != retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: retLeft.Value.(float64) != retRight.Value.(float64)}
		} else if priority == SymbolTable.STR {
			return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: len(retLeft.Value.(string)) != len(retRight.Value.(string))}
		}

	case "&&":
		if retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.INTEGER ||
			retRight.Type == SymbolTable.INTEGER ||
			retLeft.Type == SymbolTable.FLOAT ||
			retRight.Type == SymbolTable.FLOAT ||
			retLeft.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.STR {

			row := strconv.Itoa(o.Row)
			col := strconv.Itoa(o.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: Operación 'AND' no soportada \n"
			err := errors.NewError(errors.CounterError, o.Row, o.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

		interpreter.DataTypeRes = SymbolTable.BOOLEAN
		return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: retLeft.Value.(bool) && retRight.Value.(bool)}

	case "||":
		if retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.INTEGER ||
			retRight.Type == SymbolTable.INTEGER ||
			retLeft.Type == SymbolTable.FLOAT ||
			retRight.Type == SymbolTable.FLOAT ||
			retLeft.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.STR {

			row := strconv.Itoa(o.Row)
			col := strconv.Itoa(o.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: Operación 'OR' no soportada \n"
			err := errors.NewError(errors.CounterError, o.Row, o.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

		interpreter.DataTypeRes = SymbolTable.BOOLEAN
		return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: retLeft.Value.(bool) || retRight.Value.(bool)}

	case "!":
		if retLeft.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.INTEGER ||
			retLeft.Type == SymbolTable.FLOAT ||
			retLeft.Type == SymbolTable.STR {

			goto ErrorDataType
		}

		interpreter.DataTypeRes = SymbolTable.BOOLEAN
		return SymbolTable.ReturnType{Type: SymbolTable.BOOLEAN, Value: !retLeft.Value.(bool)}
	}

ErrorDataType:
	row := strconv.Itoa(o.Row)
	col := strconv.Itoa(o.Col)
	errors.CounterError += 1
	msg := "(" + row + ", " + col + ") Error: tipos de datos no soportados \n"
	err := errors.NewError(errors.CounterError, o.Row, o.Col, msg, symbolTable.Name)
	errors.TypeError = append(errors.TypeError, err)
	interpreter.Console += fmt.Sprintf("%v", err.Msg)
	return SymbolTable.ReturnType{Type: SymbolTable.NULL, Value: err}
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
