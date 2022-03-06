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

// EnterListFuncs is called when production listFuncs is entered.
func (s *BaseProjectParserListener) EnterListFuncs(ctx *ListFuncsContext) {}

// ExitListFuncs is called when production listFuncs is exited.
func (s *BaseProjectParserListener) ExitListFuncs(ctx *ListFuncsContext) {}

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

// EnterList_else_if is called when production list_else_if is entered.
func (s *BaseProjectParserListener) EnterList_else_if(ctx *List_else_ifContext) {}

// ExitList_else_if is called when production list_else_if is exited.
func (s *BaseProjectParserListener) ExitList_else_if(ctx *List_else_ifContext) {}

// EnterElse_if is called when production else_if is entered.
func (s *BaseProjectParserListener) EnterElse_if(ctx *Else_ifContext) {}

// ExitElse_if is called when production else_if is exited.
func (s *BaseProjectParserListener) ExitElse_if(ctx *Else_ifContext) {}

// EnterBucle_prod is called when production bucle_prod is entered.
func (s *BaseProjectParserListener) EnterBucle_prod(ctx *Bucle_prodContext) {}

// ExitBucle_prod is called when production bucle_prod is exited.
func (s *BaseProjectParserListener) ExitBucle_prod(ctx *Bucle_prodContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseProjectParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseProjectParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpr_rel is called when production expr_rel is entered.
func (s *BaseProjectParserListener) EnterExpr_rel(ctx *Expr_relContext) {}

// ExitExpr_rel is called when production expr_rel is exited.
func (s *BaseProjectParserListener) ExitExpr_rel(ctx *Expr_relContext) {}

// EnterExpr_arit is called when production expr_arit is entered.
func (s *BaseProjectParserListener) EnterExpr_arit(ctx *Expr_aritContext) {}

// ExitExpr_arit is called when production expr_arit is exited.
func (s *BaseProjectParserListener) ExitExpr_arit(ctx *Expr_aritContext) {}

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
