// Code generated from ProjectParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ProjectParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ProjectParserListener is a complete listener for a parse tree produced by ProjectParser.
type ProjectParserListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterListFuncs is called when entering the listFuncs production.
	EnterListFuncs(c *ListFuncsContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFuncMain is called when entering the funcMain production.
	EnterFuncMain(c *FuncMainContext)

	// EnterBloq is called when entering the bloq production.
	EnterBloq(c *BloqContext)

	// EnterCalled_func is called when entering the called_func production.
	EnterCalled_func(c *Called_funcContext)

	// EnterListExpressions is called when entering the listExpressions production.
	EnterListExpressions(c *ListExpressionsContext)

	// EnterInstructions is called when entering the instructions production.
	EnterInstructions(c *InstructionsContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterPrint_prod is called when entering the print_prod production.
	EnterPrint_prod(c *Print_prodContext)

	// EnterListVars is called when entering the listVars production.
	EnterListVars(c *ListVarsContext)

	// EnterDeclaration_prod is called when entering the declaration_prod production.
	EnterDeclaration_prod(c *Declaration_prodContext)

	// EnterAssign_prod is called when entering the assign_prod production.
	EnterAssign_prod(c *Assign_prodContext)

	// EnterIds_type is called when entering the ids_type production.
	EnterIds_type(c *Ids_typeContext)

	// EnterListIds is called when entering the listIds production.
	EnterListIds(c *ListIdsContext)

	// EnterConditional_prod is called when entering the conditional_prod production.
	EnterConditional_prod(c *Conditional_prodContext)

	// EnterList_else_if is called when entering the list_else_if production.
	EnterList_else_if(c *List_else_ifContext)

	// EnterElse_if is called when entering the else_if production.
	EnterElse_if(c *Else_ifContext)

	// EnterBucle_prod is called when entering the bucle_prod production.
	EnterBucle_prod(c *Bucle_prodContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpr_rel is called when entering the expr_rel production.
	EnterExpr_rel(c *Expr_relContext)

	// EnterExpr_arit is called when entering the expr_arit production.
	EnterExpr_arit(c *Expr_aritContext)

	// EnterPow_op is called when entering the pow_op production.
	EnterPow_op(c *Pow_opContext)

	// EnterExpr_logic is called when entering the expr_logic production.
	EnterExpr_logic(c *Expr_logicContext)

	// EnterExpr_cast is called when entering the expr_cast production.
	EnterExpr_cast(c *Expr_castContext)

	// EnterData_type is called when entering the data_type production.
	EnterData_type(c *Data_typeContext)

	// EnterPrimitive is called when entering the primitive production.
	EnterPrimitive(c *PrimitiveContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitListFuncs is called when exiting the listFuncs production.
	ExitListFuncs(c *ListFuncsContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFuncMain is called when exiting the funcMain production.
	ExitFuncMain(c *FuncMainContext)

	// ExitBloq is called when exiting the bloq production.
	ExitBloq(c *BloqContext)

	// ExitCalled_func is called when exiting the called_func production.
	ExitCalled_func(c *Called_funcContext)

	// ExitListExpressions is called when exiting the listExpressions production.
	ExitListExpressions(c *ListExpressionsContext)

	// ExitInstructions is called when exiting the instructions production.
	ExitInstructions(c *InstructionsContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitPrint_prod is called when exiting the print_prod production.
	ExitPrint_prod(c *Print_prodContext)

	// ExitListVars is called when exiting the listVars production.
	ExitListVars(c *ListVarsContext)

	// ExitDeclaration_prod is called when exiting the declaration_prod production.
	ExitDeclaration_prod(c *Declaration_prodContext)

	// ExitAssign_prod is called when exiting the assign_prod production.
	ExitAssign_prod(c *Assign_prodContext)

	// ExitIds_type is called when exiting the ids_type production.
	ExitIds_type(c *Ids_typeContext)

	// ExitListIds is called when exiting the listIds production.
	ExitListIds(c *ListIdsContext)

	// ExitConditional_prod is called when exiting the conditional_prod production.
	ExitConditional_prod(c *Conditional_prodContext)

	// ExitList_else_if is called when exiting the list_else_if production.
	ExitList_else_if(c *List_else_ifContext)

	// ExitElse_if is called when exiting the else_if production.
	ExitElse_if(c *Else_ifContext)

	// ExitBucle_prod is called when exiting the bucle_prod production.
	ExitBucle_prod(c *Bucle_prodContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpr_rel is called when exiting the expr_rel production.
	ExitExpr_rel(c *Expr_relContext)

	// ExitExpr_arit is called when exiting the expr_arit production.
	ExitExpr_arit(c *Expr_aritContext)

	// ExitPow_op is called when exiting the pow_op production.
	ExitPow_op(c *Pow_opContext)

	// ExitExpr_logic is called when exiting the expr_logic production.
	ExitExpr_logic(c *Expr_logicContext)

	// ExitExpr_cast is called when exiting the expr_cast production.
	ExitExpr_cast(c *Expr_castContext)

	// ExitData_type is called when exiting the data_type production.
	ExitData_type(c *Data_typeContext)

	// ExitPrimitive is called when exiting the primitive production.
	ExitPrimitive(c *PrimitiveContext)
}
