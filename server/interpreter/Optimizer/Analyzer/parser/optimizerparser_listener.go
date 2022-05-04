// Code generated from OptimizerParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // OptimizerParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OptimizerParserListener is a complete listener for a parse tree produced by OptimizerParser.
type OptimizerParserListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterHeaders is called when entering the headers production.
	EnterHeaders(c *HeadersContext)

	// EnterList_temps is called when entering the list_temps production.
	EnterList_temps(c *List_tempsContext)

	// EnterTemps is called when entering the temps production.
	EnterTemps(c *TempsContext)

	// EnterListFuncs is called when entering the listFuncs production.
	EnterListFuncs(c *ListFuncsContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterInstructions is called when entering the instructions production.
	EnterInstructions(c *InstructionsContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterAssign_stack is called when entering the assign_stack production.
	EnterAssign_stack(c *Assign_stackContext)

	// EnterAssign_heap is called when entering the assign_heap production.
	EnterAssign_heap(c *Assign_heapContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterIf_instr is called when entering the if_instr production.
	EnterIf_instr(c *If_instrContext)

	// EnterGoto_instr is called when entering the goto_instr production.
	EnterGoto_instr(c *Goto_instrContext)

	// EnterLabel_instr is called when entering the label_instr production.
	EnterLabel_instr(c *Label_instrContext)

	// EnterPrintf_instr is called when entering the printf_instr production.
	EnterPrintf_instr(c *Printf_instrContext)

	// EnterReturn_instr is called when entering the return_instr production.
	EnterReturn_instr(c *Return_instrContext)

	// EnterGetHeap_instr is called when entering the getHeap_instr production.
	EnterGetHeap_instr(c *GetHeap_instrContext)

	// EnterGetStack_instr is called when entering the getStack_instr production.
	EnterGetStack_instr(c *GetStack_instrContext)

	// EnterCall_instr is called when entering the call_instr production.
	EnterCall_instr(c *Call_instrContext)

	// EnterExpr_print is called when entering the expr_print production.
	EnterExpr_print(c *Expr_printContext)

	// EnterConvert is called when entering the convert production.
	EnterConvert(c *ConvertContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpr_rel is called when entering the expr_rel production.
	EnterExpr_rel(c *Expr_relContext)

	// EnterExpr_arit is called when entering the expr_arit production.
	EnterExpr_arit(c *Expr_aritContext)

	// EnterData_type is called when entering the data_type production.
	EnterData_type(c *Data_typeContext)

	// EnterExpr_valor is called when entering the expr_valor production.
	EnterExpr_valor(c *Expr_valorContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitHeaders is called when exiting the headers production.
	ExitHeaders(c *HeadersContext)

	// ExitList_temps is called when exiting the list_temps production.
	ExitList_temps(c *List_tempsContext)

	// ExitTemps is called when exiting the temps production.
	ExitTemps(c *TempsContext)

	// ExitListFuncs is called when exiting the listFuncs production.
	ExitListFuncs(c *ListFuncsContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitInstructions is called when exiting the instructions production.
	ExitInstructions(c *InstructionsContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitAssign_stack is called when exiting the assign_stack production.
	ExitAssign_stack(c *Assign_stackContext)

	// ExitAssign_heap is called when exiting the assign_heap production.
	ExitAssign_heap(c *Assign_heapContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitIf_instr is called when exiting the if_instr production.
	ExitIf_instr(c *If_instrContext)

	// ExitGoto_instr is called when exiting the goto_instr production.
	ExitGoto_instr(c *Goto_instrContext)

	// ExitLabel_instr is called when exiting the label_instr production.
	ExitLabel_instr(c *Label_instrContext)

	// ExitPrintf_instr is called when exiting the printf_instr production.
	ExitPrintf_instr(c *Printf_instrContext)

	// ExitReturn_instr is called when exiting the return_instr production.
	ExitReturn_instr(c *Return_instrContext)

	// ExitGetHeap_instr is called when exiting the getHeap_instr production.
	ExitGetHeap_instr(c *GetHeap_instrContext)

	// ExitGetStack_instr is called when exiting the getStack_instr production.
	ExitGetStack_instr(c *GetStack_instrContext)

	// ExitCall_instr is called when exiting the call_instr production.
	ExitCall_instr(c *Call_instrContext)

	// ExitExpr_print is called when exiting the expr_print production.
	ExitExpr_print(c *Expr_printContext)

	// ExitConvert is called when exiting the convert production.
	ExitConvert(c *ConvertContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpr_rel is called when exiting the expr_rel production.
	ExitExpr_rel(c *Expr_relContext)

	// ExitExpr_arit is called when exiting the expr_arit production.
	ExitExpr_arit(c *Expr_aritContext)

	// ExitData_type is called when exiting the data_type production.
	ExitData_type(c *Data_typeContext)

	// ExitExpr_valor is called when exiting the expr_valor production.
	ExitExpr_valor(c *Expr_valorContext)
}
