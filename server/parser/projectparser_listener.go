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

	// EnterBloq_match is called when entering the bloq_match production.
	EnterBloq_match(c *Bloq_matchContext)

	// EnterListParams is called when entering the listParams production.
	EnterListParams(c *ListParamsContext)

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

	// EnterIf_prod is called when entering the if_prod production.
	EnterIf_prod(c *If_prodContext)

	// EnterList_else_if is called when entering the list_else_if production.
	EnterList_else_if(c *List_else_ifContext)

	// EnterElse_if is called when entering the else_if production.
	EnterElse_if(c *Else_ifContext)

	// EnterMatch_prod is called when entering the match_prod production.
	EnterMatch_prod(c *Match_prodContext)

	// EnterList_instr_match is called when entering the list_instr_match production.
	EnterList_instr_match(c *List_instr_matchContext)

	// EnterInstr_match is called when entering the instr_match production.
	EnterInstr_match(c *Instr_matchContext)

	// EnterExpr_match is called when entering the expr_match production.
	EnterExpr_match(c *Expr_matchContext)

	// EnterBucle_prod is called when entering the bucle_prod production.
	EnterBucle_prod(c *Bucle_prodContext)

	// EnterWhile_prod is called when entering the while_prod production.
	EnterWhile_prod(c *While_prodContext)

	// EnterLoop_prod is called when entering the loop_prod production.
	EnterLoop_prod(c *Loop_prodContext)

	// EnterCalled_func is called when entering the called_func production.
	EnterCalled_func(c *Called_funcContext)

	// EnterListExpressions is called when entering the listExpressions production.
	EnterListExpressions(c *ListExpressionsContext)

	// EnterDec_arr is called when entering the dec_arr production.
	EnterDec_arr(c *Dec_arrContext)

	// EnterListDim is called when entering the listDim production.
	EnterListDim(c *ListDimContext)

	// EnterTransfer_prod is called when entering the transfer_prod production.
	EnterTransfer_prod(c *Transfer_prodContext)

	// EnterBreak_instr is called when entering the break_instr production.
	EnterBreak_instr(c *Break_instrContext)

	// EnterContinue_instr is called when entering the continue_instr production.
	EnterContinue_instr(c *Continue_instrContext)

	// EnterReturn_instr is called when entering the return_instr production.
	EnterReturn_instr(c *Return_instrContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterArraydata is called when entering the arraydata production.
	EnterArraydata(c *ArraydataContext)

	// EnterAccess_array is called when entering the access_array production.
	EnterAccess_array(c *Access_arrayContext)

	// EnterListInArray is called when entering the listInArray production.
	EnterListInArray(c *ListInArrayContext)

	// EnterInArray is called when entering the inArray production.
	EnterInArray(c *InArrayContext)

	// EnterExpr_rel is called when entering the expr_rel production.
	EnterExpr_rel(c *Expr_relContext)

	// EnterExpr_arit is called when entering the expr_arit production.
	EnterExpr_arit(c *Expr_aritContext)

	// EnterExpr_valor is called when entering the expr_valor production.
	EnterExpr_valor(c *Expr_valorContext)

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

	// ExitBloq_match is called when exiting the bloq_match production.
	ExitBloq_match(c *Bloq_matchContext)

	// ExitListParams is called when exiting the listParams production.
	ExitListParams(c *ListParamsContext)

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

	// ExitIf_prod is called when exiting the if_prod production.
	ExitIf_prod(c *If_prodContext)

	// ExitList_else_if is called when exiting the list_else_if production.
	ExitList_else_if(c *List_else_ifContext)

	// ExitElse_if is called when exiting the else_if production.
	ExitElse_if(c *Else_ifContext)

	// ExitMatch_prod is called when exiting the match_prod production.
	ExitMatch_prod(c *Match_prodContext)

	// ExitList_instr_match is called when exiting the list_instr_match production.
	ExitList_instr_match(c *List_instr_matchContext)

	// ExitInstr_match is called when exiting the instr_match production.
	ExitInstr_match(c *Instr_matchContext)

	// ExitExpr_match is called when exiting the expr_match production.
	ExitExpr_match(c *Expr_matchContext)

	// ExitBucle_prod is called when exiting the bucle_prod production.
	ExitBucle_prod(c *Bucle_prodContext)

	// ExitWhile_prod is called when exiting the while_prod production.
	ExitWhile_prod(c *While_prodContext)

	// ExitLoop_prod is called when exiting the loop_prod production.
	ExitLoop_prod(c *Loop_prodContext)

	// ExitCalled_func is called when exiting the called_func production.
	ExitCalled_func(c *Called_funcContext)

	// ExitListExpressions is called when exiting the listExpressions production.
	ExitListExpressions(c *ListExpressionsContext)

	// ExitDec_arr is called when exiting the dec_arr production.
	ExitDec_arr(c *Dec_arrContext)

	// ExitListDim is called when exiting the listDim production.
	ExitListDim(c *ListDimContext)

	// ExitTransfer_prod is called when exiting the transfer_prod production.
	ExitTransfer_prod(c *Transfer_prodContext)

	// ExitBreak_instr is called when exiting the break_instr production.
	ExitBreak_instr(c *Break_instrContext)

	// ExitContinue_instr is called when exiting the continue_instr production.
	ExitContinue_instr(c *Continue_instrContext)

	// ExitReturn_instr is called when exiting the return_instr production.
	ExitReturn_instr(c *Return_instrContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitArraydata is called when exiting the arraydata production.
	ExitArraydata(c *ArraydataContext)

	// ExitAccess_array is called when exiting the access_array production.
	ExitAccess_array(c *Access_arrayContext)

	// ExitListInArray is called when exiting the listInArray production.
	ExitListInArray(c *ListInArrayContext)

	// ExitInArray is called when exiting the inArray production.
	ExitInArray(c *InArrayContext)

	// ExitExpr_rel is called when exiting the expr_rel production.
	ExitExpr_rel(c *Expr_relContext)

	// ExitExpr_arit is called when exiting the expr_arit production.
	ExitExpr_arit(c *Expr_aritContext)

	// ExitExpr_valor is called when exiting the expr_valor production.
	ExitExpr_valor(c *Expr_valorContext)

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
