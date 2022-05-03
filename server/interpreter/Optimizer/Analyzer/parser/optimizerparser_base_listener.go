// Code generated from OptimizerParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // OptimizerParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOptimizerParserListener is a complete listener for a parse tree produced by OptimizerParser.
type BaseOptimizerParserListener struct{}

var _ OptimizerParserListener = &BaseOptimizerParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOptimizerParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOptimizerParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOptimizerParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOptimizerParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseOptimizerParserListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseOptimizerParserListener) ExitStart(ctx *StartContext) {}

// EnterHeaders is called when production headers is entered.
func (s *BaseOptimizerParserListener) EnterHeaders(ctx *HeadersContext) {}

// ExitHeaders is called when production headers is exited.
func (s *BaseOptimizerParserListener) ExitHeaders(ctx *HeadersContext) {}

// EnterList_temps is called when production list_temps is entered.
func (s *BaseOptimizerParserListener) EnterList_temps(ctx *List_tempsContext) {}

// ExitList_temps is called when production list_temps is exited.
func (s *BaseOptimizerParserListener) ExitList_temps(ctx *List_tempsContext) {}

// EnterTemps is called when production temps is entered.
func (s *BaseOptimizerParserListener) EnterTemps(ctx *TempsContext) {}

// ExitTemps is called when production temps is exited.
func (s *BaseOptimizerParserListener) ExitTemps(ctx *TempsContext) {}

// EnterListFuncs is called when production listFuncs is entered.
func (s *BaseOptimizerParserListener) EnterListFuncs(ctx *ListFuncsContext) {}

// ExitListFuncs is called when production listFuncs is exited.
func (s *BaseOptimizerParserListener) ExitListFuncs(ctx *ListFuncsContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseOptimizerParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseOptimizerParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterInstructions is called when production instructions is entered.
func (s *BaseOptimizerParserListener) EnterInstructions(ctx *InstructionsContext) {}

// ExitInstructions is called when production instructions is exited.
func (s *BaseOptimizerParserListener) ExitInstructions(ctx *InstructionsContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseOptimizerParserListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseOptimizerParserListener) ExitInstruction(ctx *InstructionContext) {}

// EnterAssign_stack is called when production assign_stack is entered.
func (s *BaseOptimizerParserListener) EnterAssign_stack(ctx *Assign_stackContext) {}

// ExitAssign_stack is called when production assign_stack is exited.
func (s *BaseOptimizerParserListener) ExitAssign_stack(ctx *Assign_stackContext) {}

// EnterAssign_heap is called when production assign_heap is entered.
func (s *BaseOptimizerParserListener) EnterAssign_heap(ctx *Assign_heapContext) {}

// ExitAssign_heap is called when production assign_heap is exited.
func (s *BaseOptimizerParserListener) ExitAssign_heap(ctx *Assign_heapContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseOptimizerParserListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseOptimizerParserListener) ExitAssign(ctx *AssignContext) {}

// EnterIf_instr is called when production if_instr is entered.
func (s *BaseOptimizerParserListener) EnterIf_instr(ctx *If_instrContext) {}

// ExitIf_instr is called when production if_instr is exited.
func (s *BaseOptimizerParserListener) ExitIf_instr(ctx *If_instrContext) {}

// EnterGoto_instr is called when production goto_instr is entered.
func (s *BaseOptimizerParserListener) EnterGoto_instr(ctx *Goto_instrContext) {}

// ExitGoto_instr is called when production goto_instr is exited.
func (s *BaseOptimizerParserListener) ExitGoto_instr(ctx *Goto_instrContext) {}

// EnterLabel_instr is called when production label_instr is entered.
func (s *BaseOptimizerParserListener) EnterLabel_instr(ctx *Label_instrContext) {}

// ExitLabel_instr is called when production label_instr is exited.
func (s *BaseOptimizerParserListener) ExitLabel_instr(ctx *Label_instrContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOptimizerParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOptimizerParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpr_rel is called when production expr_rel is entered.
func (s *BaseOptimizerParserListener) EnterExpr_rel(ctx *Expr_relContext) {}

// ExitExpr_rel is called when production expr_rel is exited.
func (s *BaseOptimizerParserListener) ExitExpr_rel(ctx *Expr_relContext) {}

// EnterExpr_arit is called when production expr_arit is entered.
func (s *BaseOptimizerParserListener) EnterExpr_arit(ctx *Expr_aritContext) {}

// ExitExpr_arit is called when production expr_arit is exited.
func (s *BaseOptimizerParserListener) ExitExpr_arit(ctx *Expr_aritContext) {}

// EnterData_type is called when production data_type is entered.
func (s *BaseOptimizerParserListener) EnterData_type(ctx *Data_typeContext) {}

// ExitData_type is called when production data_type is exited.
func (s *BaseOptimizerParserListener) ExitData_type(ctx *Data_typeContext) {}

// EnterExpr_valor is called when production expr_valor is entered.
func (s *BaseOptimizerParserListener) EnterExpr_valor(ctx *Expr_valorContext) {}

// ExitExpr_valor is called when production expr_valor is exited.
func (s *BaseOptimizerParserListener) ExitExpr_valor(ctx *Expr_valorContext) {}
