// Code generated from ProjectParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ProjectParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseProjectParserListener is a complete listener for a parse tree produced by ProjectParser.
type BaseProjectParserListener struct{}

var _ ProjectParserListener = &BaseProjectParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProjectParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProjectParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProjectParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProjectParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseProjectParserListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseProjectParserListener) ExitStart(ctx *StartContext) {}

// EnterListInits is called when production listInits is entered.
func (s *BaseProjectParserListener) EnterListInits(ctx *ListInitsContext) {}

// ExitListInits is called when production listInits is exited.
func (s *BaseProjectParserListener) ExitListInits(ctx *ListInitsContext) {}

// EnterListFuncs is called when production listFuncs is entered.
func (s *BaseProjectParserListener) EnterListFuncs(ctx *ListFuncsContext) {}

// ExitListFuncs is called when production listFuncs is exited.
func (s *BaseProjectParserListener) ExitListFuncs(ctx *ListFuncsContext) {}

// EnterListArrays is called when production listArrays is entered.
func (s *BaseProjectParserListener) EnterListArrays(ctx *ListArraysContext) {}

// ExitListArrays is called when production listArrays is exited.
func (s *BaseProjectParserListener) ExitListArrays(ctx *ListArraysContext) {}

// EnterListStructs is called when production listStructs is entered.
func (s *BaseProjectParserListener) EnterListStructs(ctx *ListStructsContext) {}

// ExitListStructs is called when production listStructs is exited.
func (s *BaseProjectParserListener) ExitListStructs(ctx *ListStructsContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseProjectParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseProjectParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterFuncMain is called when production funcMain is entered.
func (s *BaseProjectParserListener) EnterFuncMain(ctx *FuncMainContext) {}

// ExitFuncMain is called when production funcMain is exited.
func (s *BaseProjectParserListener) ExitFuncMain(ctx *FuncMainContext) {}

// EnterBloq is called when production bloq is entered.
func (s *BaseProjectParserListener) EnterBloq(ctx *BloqContext) {}

// ExitBloq is called when production bloq is exited.
func (s *BaseProjectParserListener) ExitBloq(ctx *BloqContext) {}

// EnterBloq_match is called when production bloq_match is entered.
func (s *BaseProjectParserListener) EnterBloq_match(ctx *Bloq_matchContext) {}

// ExitBloq_match is called when production bloq_match is exited.
func (s *BaseProjectParserListener) ExitBloq_match(ctx *Bloq_matchContext) {}

// EnterListParams is called when production listParams is entered.
func (s *BaseProjectParserListener) EnterListParams(ctx *ListParamsContext) {}

// ExitListParams is called when production listParams is exited.
func (s *BaseProjectParserListener) ExitListParams(ctx *ListParamsContext) {}

// EnterInstructions is called when production instructions is entered.
func (s *BaseProjectParserListener) EnterInstructions(ctx *InstructionsContext) {}

// ExitInstructions is called when production instructions is exited.
func (s *BaseProjectParserListener) ExitInstructions(ctx *InstructionsContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseProjectParserListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseProjectParserListener) ExitInstruction(ctx *InstructionContext) {}

// EnterPrint_prod is called when production print_prod is entered.
func (s *BaseProjectParserListener) EnterPrint_prod(ctx *Print_prodContext) {}

// ExitPrint_prod is called when production print_prod is exited.
func (s *BaseProjectParserListener) ExitPrint_prod(ctx *Print_prodContext) {}

// EnterListVars is called when production listVars is entered.
func (s *BaseProjectParserListener) EnterListVars(ctx *ListVarsContext) {}

// ExitListVars is called when production listVars is exited.
func (s *BaseProjectParserListener) ExitListVars(ctx *ListVarsContext) {}

// EnterDeclaration_prod is called when production declaration_prod is entered.
func (s *BaseProjectParserListener) EnterDeclaration_prod(ctx *Declaration_prodContext) {}

// ExitDeclaration_prod is called when production declaration_prod is exited.
func (s *BaseProjectParserListener) ExitDeclaration_prod(ctx *Declaration_prodContext) {}

// EnterAssign_prod is called when production assign_prod is entered.
func (s *BaseProjectParserListener) EnterAssign_prod(ctx *Assign_prodContext) {}

// ExitAssign_prod is called when production assign_prod is exited.
func (s *BaseProjectParserListener) ExitAssign_prod(ctx *Assign_prodContext) {}

// EnterIds_type is called when production ids_type is entered.
func (s *BaseProjectParserListener) EnterIds_type(ctx *Ids_typeContext) {}

// ExitIds_type is called when production ids_type is exited.
func (s *BaseProjectParserListener) ExitIds_type(ctx *Ids_typeContext) {}

// EnterListIds is called when production listIds is entered.
func (s *BaseProjectParserListener) EnterListIds(ctx *ListIdsContext) {}

// ExitListIds is called when production listIds is exited.
func (s *BaseProjectParserListener) ExitListIds(ctx *ListIdsContext) {}

// EnterConditional_prod is called when production conditional_prod is entered.
func (s *BaseProjectParserListener) EnterConditional_prod(ctx *Conditional_prodContext) {}

// ExitConditional_prod is called when production conditional_prod is exited.
func (s *BaseProjectParserListener) ExitConditional_prod(ctx *Conditional_prodContext) {}

// EnterIf_prod is called when production if_prod is entered.
func (s *BaseProjectParserListener) EnterIf_prod(ctx *If_prodContext) {}

// ExitIf_prod is called when production if_prod is exited.
func (s *BaseProjectParserListener) ExitIf_prod(ctx *If_prodContext) {}

// EnterList_else_if is called when production list_else_if is entered.
func (s *BaseProjectParserListener) EnterList_else_if(ctx *List_else_ifContext) {}

// ExitList_else_if is called when production list_else_if is exited.
func (s *BaseProjectParserListener) ExitList_else_if(ctx *List_else_ifContext) {}

// EnterElse_if is called when production else_if is entered.
func (s *BaseProjectParserListener) EnterElse_if(ctx *Else_ifContext) {}

// ExitElse_if is called when production else_if is exited.
func (s *BaseProjectParserListener) ExitElse_if(ctx *Else_ifContext) {}

// EnterMatch_prod is called when production match_prod is entered.
func (s *BaseProjectParserListener) EnterMatch_prod(ctx *Match_prodContext) {}

// ExitMatch_prod is called when production match_prod is exited.
func (s *BaseProjectParserListener) ExitMatch_prod(ctx *Match_prodContext) {}

// EnterList_instr_match is called when production list_instr_match is entered.
func (s *BaseProjectParserListener) EnterList_instr_match(ctx *List_instr_matchContext) {}

// ExitList_instr_match is called when production list_instr_match is exited.
func (s *BaseProjectParserListener) ExitList_instr_match(ctx *List_instr_matchContext) {}

// EnterInstr_match is called when production instr_match is entered.
func (s *BaseProjectParserListener) EnterInstr_match(ctx *Instr_matchContext) {}

// ExitInstr_match is called when production instr_match is exited.
func (s *BaseProjectParserListener) ExitInstr_match(ctx *Instr_matchContext) {}

// EnterExpr_match is called when production expr_match is entered.
func (s *BaseProjectParserListener) EnterExpr_match(ctx *Expr_matchContext) {}

// ExitExpr_match is called when production expr_match is exited.
func (s *BaseProjectParserListener) ExitExpr_match(ctx *Expr_matchContext) {}

// EnterBucle_prod is called when production bucle_prod is entered.
func (s *BaseProjectParserListener) EnterBucle_prod(ctx *Bucle_prodContext) {}

// ExitBucle_prod is called when production bucle_prod is exited.
func (s *BaseProjectParserListener) ExitBucle_prod(ctx *Bucle_prodContext) {}

// EnterWhile_prod is called when production while_prod is entered.
func (s *BaseProjectParserListener) EnterWhile_prod(ctx *While_prodContext) {}

// ExitWhile_prod is called when production while_prod is exited.
func (s *BaseProjectParserListener) ExitWhile_prod(ctx *While_prodContext) {}

// EnterLoop_prod is called when production loop_prod is entered.
func (s *BaseProjectParserListener) EnterLoop_prod(ctx *Loop_prodContext) {}

// ExitLoop_prod is called when production loop_prod is exited.
func (s *BaseProjectParserListener) ExitLoop_prod(ctx *Loop_prodContext) {}

// EnterForin_prod is called when production forin_prod is entered.
func (s *BaseProjectParserListener) EnterForin_prod(ctx *Forin_prodContext) {}

// ExitForin_prod is called when production forin_prod is exited.
func (s *BaseProjectParserListener) ExitForin_prod(ctx *Forin_prodContext) {}

// EnterRange_prod is called when production range_prod is entered.
func (s *BaseProjectParserListener) EnterRange_prod(ctx *Range_prodContext) {}

// ExitRange_prod is called when production range_prod is exited.
func (s *BaseProjectParserListener) ExitRange_prod(ctx *Range_prodContext) {}

// EnterCalled_func is called when production called_func is entered.
func (s *BaseProjectParserListener) EnterCalled_func(ctx *Called_funcContext) {}

// ExitCalled_func is called when production called_func is exited.
func (s *BaseProjectParserListener) ExitCalled_func(ctx *Called_funcContext) {}

// EnterListExpressions is called when production listExpressions is entered.
func (s *BaseProjectParserListener) EnterListExpressions(ctx *ListExpressionsContext) {}

// ExitListExpressions is called when production listExpressions is exited.
func (s *BaseProjectParserListener) ExitListExpressions(ctx *ListExpressionsContext) {}

// EnterDec_arr is called when production dec_arr is entered.
func (s *BaseProjectParserListener) EnterDec_arr(ctx *Dec_arrContext) {}

// ExitDec_arr is called when production dec_arr is exited.
func (s *BaseProjectParserListener) ExitDec_arr(ctx *Dec_arrContext) {}

// EnterListDim is called when production listDim is entered.
func (s *BaseProjectParserListener) EnterListDim(ctx *ListDimContext) {}

// ExitListDim is called when production listDim is exited.
func (s *BaseProjectParserListener) ExitListDim(ctx *ListDimContext) {}

// EnterVector_instr is called when production vector_instr is entered.
func (s *BaseProjectParserListener) EnterVector_instr(ctx *Vector_instrContext) {}

// ExitVector_instr is called when production vector_instr is exited.
func (s *BaseProjectParserListener) ExitVector_instr(ctx *Vector_instrContext) {}

// EnterDec_vector is called when production dec_vector is entered.
func (s *BaseProjectParserListener) EnterDec_vector(ctx *Dec_vectorContext) {}

// ExitDec_vector is called when production dec_vector is exited.
func (s *BaseProjectParserListener) ExitDec_vector(ctx *Dec_vectorContext) {}

// EnterExpr_vector is called when production expr_vector is entered.
func (s *BaseProjectParserListener) EnterExpr_vector(ctx *Expr_vectorContext) {}

// ExitExpr_vector is called when production expr_vector is exited.
func (s *BaseProjectParserListener) ExitExpr_vector(ctx *Expr_vectorContext) {}

// EnterNatives_vector is called when production natives_vector is entered.
func (s *BaseProjectParserListener) EnterNatives_vector(ctx *Natives_vectorContext) {}

// ExitNatives_vector is called when production natives_vector is exited.
func (s *BaseProjectParserListener) ExitNatives_vector(ctx *Natives_vectorContext) {}

// EnterTransfer_prod is called when production transfer_prod is entered.
func (s *BaseProjectParserListener) EnterTransfer_prod(ctx *Transfer_prodContext) {}

// ExitTransfer_prod is called when production transfer_prod is exited.
func (s *BaseProjectParserListener) ExitTransfer_prod(ctx *Transfer_prodContext) {}

// EnterBreak_instr is called when production break_instr is entered.
func (s *BaseProjectParserListener) EnterBreak_instr(ctx *Break_instrContext) {}

// ExitBreak_instr is called when production break_instr is exited.
func (s *BaseProjectParserListener) ExitBreak_instr(ctx *Break_instrContext) {}

// EnterContinue_instr is called when production continue_instr is entered.
func (s *BaseProjectParserListener) EnterContinue_instr(ctx *Continue_instrContext) {}

// ExitContinue_instr is called when production continue_instr is exited.
func (s *BaseProjectParserListener) ExitContinue_instr(ctx *Continue_instrContext) {}

// EnterReturn_instr is called when production return_instr is entered.
func (s *BaseProjectParserListener) EnterReturn_instr(ctx *Return_instrContext) {}

// ExitReturn_instr is called when production return_instr is exited.
func (s *BaseProjectParserListener) ExitReturn_instr(ctx *Return_instrContext) {}

// EnterDec_struct is called when production dec_struct is entered.
func (s *BaseProjectParserListener) EnterDec_struct(ctx *Dec_structContext) {}

// ExitDec_struct is called when production dec_struct is exited.
func (s *BaseProjectParserListener) ExitDec_struct(ctx *Dec_structContext) {}

// EnterBloq_struct is called when production bloq_struct is entered.
func (s *BaseProjectParserListener) EnterBloq_struct(ctx *Bloq_structContext) {}

// ExitBloq_struct is called when production bloq_struct is exited.
func (s *BaseProjectParserListener) ExitBloq_struct(ctx *Bloq_structContext) {}

// EnterContent_struct is called when production content_struct is entered.
func (s *BaseProjectParserListener) EnterContent_struct(ctx *Content_structContext) {}

// ExitContent_struct is called when production content_struct is exited.
func (s *BaseProjectParserListener) ExitContent_struct(ctx *Content_structContext) {}

// EnterItem_struct is called when production item_struct is entered.
func (s *BaseProjectParserListener) EnterItem_struct(ctx *Item_structContext) {}

// ExitItem_struct is called when production item_struct is exited.
func (s *BaseProjectParserListener) ExitItem_struct(ctx *Item_structContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseProjectParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseProjectParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterArraydata is called when production arraydata is entered.
func (s *BaseProjectParserListener) EnterArraydata(ctx *ArraydataContext) {}

// ExitArraydata is called when production arraydata is exited.
func (s *BaseProjectParserListener) ExitArraydata(ctx *ArraydataContext) {}

// EnterAccess_array is called when production access_array is entered.
func (s *BaseProjectParserListener) EnterAccess_array(ctx *Access_arrayContext) {}

// ExitAccess_array is called when production access_array is exited.
func (s *BaseProjectParserListener) ExitAccess_array(ctx *Access_arrayContext) {}

// EnterListInArray is called when production listInArray is entered.
func (s *BaseProjectParserListener) EnterListInArray(ctx *ListInArrayContext) {}

// ExitListInArray is called when production listInArray is exited.
func (s *BaseProjectParserListener) ExitListInArray(ctx *ListInArrayContext) {}

// EnterInArray is called when production inArray is entered.
func (s *BaseProjectParserListener) EnterInArray(ctx *InArrayContext) {}

// ExitInArray is called when production inArray is exited.
func (s *BaseProjectParserListener) ExitInArray(ctx *InArrayContext) {}

// EnterAccess_vector is called when production access_vector is entered.
func (s *BaseProjectParserListener) EnterAccess_vector(ctx *Access_vectorContext) {}

// ExitAccess_vector is called when production access_vector is exited.
func (s *BaseProjectParserListener) ExitAccess_vector(ctx *Access_vectorContext) {}

// EnterListInVector is called when production listInVector is entered.
func (s *BaseProjectParserListener) EnterListInVector(ctx *ListInVectorContext) {}

// ExitListInVector is called when production listInVector is exited.
func (s *BaseProjectParserListener) ExitListInVector(ctx *ListInVectorContext) {}

// EnterInVector is called when production inVector is entered.
func (s *BaseProjectParserListener) EnterInVector(ctx *InVectorContext) {}

// ExitInVector is called when production inVector is exited.
func (s *BaseProjectParserListener) ExitInVector(ctx *InVectorContext) {}

// EnterType_struct is called when production type_struct is entered.
func (s *BaseProjectParserListener) EnterType_struct(ctx *Type_structContext) {}

// ExitType_struct is called when production type_struct is exited.
func (s *BaseProjectParserListener) ExitType_struct(ctx *Type_structContext) {}

// EnterDef_items is called when production def_items is entered.
func (s *BaseProjectParserListener) EnterDef_items(ctx *Def_itemsContext) {}

// ExitDef_items is called when production def_items is exited.
func (s *BaseProjectParserListener) ExitDef_items(ctx *Def_itemsContext) {}

// EnterItem is called when production item is entered.
func (s *BaseProjectParserListener) EnterItem(ctx *ItemContext) {}

// ExitItem is called when production item is exited.
func (s *BaseProjectParserListener) ExitItem(ctx *ItemContext) {}

// EnterExpr_rel is called when production expr_rel is entered.
func (s *BaseProjectParserListener) EnterExpr_rel(ctx *Expr_relContext) {}

// ExitExpr_rel is called when production expr_rel is exited.
func (s *BaseProjectParserListener) ExitExpr_rel(ctx *Expr_relContext) {}

// EnterExpr_arit is called when production expr_arit is entered.
func (s *BaseProjectParserListener) EnterExpr_arit(ctx *Expr_aritContext) {}

// ExitExpr_arit is called when production expr_arit is exited.
func (s *BaseProjectParserListener) ExitExpr_arit(ctx *Expr_aritContext) {}

// EnterExpr_valor is called when production expr_valor is entered.
func (s *BaseProjectParserListener) EnterExpr_valor(ctx *Expr_valorContext) {}

// ExitExpr_valor is called when production expr_valor is exited.
func (s *BaseProjectParserListener) ExitExpr_valor(ctx *Expr_valorContext) {}

// EnterPow_op is called when production pow_op is entered.
func (s *BaseProjectParserListener) EnterPow_op(ctx *Pow_opContext) {}

// ExitPow_op is called when production pow_op is exited.
func (s *BaseProjectParserListener) ExitPow_op(ctx *Pow_opContext) {}

// EnterExpr_logic is called when production expr_logic is entered.
func (s *BaseProjectParserListener) EnterExpr_logic(ctx *Expr_logicContext) {}

// ExitExpr_logic is called when production expr_logic is exited.
func (s *BaseProjectParserListener) ExitExpr_logic(ctx *Expr_logicContext) {}

// EnterExpr_cast is called when production expr_cast is entered.
func (s *BaseProjectParserListener) EnterExpr_cast(ctx *Expr_castContext) {}

// ExitExpr_cast is called when production expr_cast is exited.
func (s *BaseProjectParserListener) ExitExpr_cast(ctx *Expr_castContext) {}

// EnterData_type is called when production data_type is entered.
func (s *BaseProjectParserListener) EnterData_type(ctx *Data_typeContext) {}

// ExitData_type is called when production data_type is exited.
func (s *BaseProjectParserListener) ExitData_type(ctx *Data_typeContext) {}

// EnterPrimitive is called when production primitive is entered.
func (s *BaseProjectParserListener) EnterPrimitive(ctx *PrimitiveContext) {}

// ExitPrimitive is called when production primitive is exited.
func (s *BaseProjectParserListener) ExitPrimitive(ctx *PrimitiveContext) {}
