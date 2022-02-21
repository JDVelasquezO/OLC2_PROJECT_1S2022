// Code generated from ProjectParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ProjectParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ProjectParserListener is a complete listener for a parse tree produced by ProjectParser.
type ProjectParserListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInstructions is called when entering the instructions production.
	EnterInstructions(c *InstructionsContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterDeclaration_prod is called when entering the declaration_prod production.
	EnterDeclaration_prod(c *Declaration_prodContext)

	// EnterAssign_prod is called when entering the assign_prod production.
	EnterAssign_prod(c *Assign_prodContext)

	// EnterTypeDataVar is called when entering the typeDataVar production.
	EnterTypeDataVar(c *TypeDataVarContext)

	// EnterListIds is called when entering the listIds production.
	EnterListIds(c *ListIdsContext)

	// EnterPrint_prod is called when entering the print_prod production.
	EnterPrint_prod(c *Print_prodContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpr_rel is called when entering the expr_rel production.
	EnterExpr_rel(c *Expr_relContext)

	// EnterExpr_logic is called when entering the expr_logic production.
	EnterExpr_logic(c *Expr_logicContext)

	// EnterExpr_arit is called when entering the expr_arit production.
	EnterExpr_arit(c *Expr_aritContext)

	// EnterPrimitive is called when entering the primitive production.
	EnterPrimitive(c *PrimitiveContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInstructions is called when exiting the instructions production.
	ExitInstructions(c *InstructionsContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitDeclaration_prod is called when exiting the declaration_prod production.
	ExitDeclaration_prod(c *Declaration_prodContext)

	// ExitAssign_prod is called when exiting the assign_prod production.
	ExitAssign_prod(c *Assign_prodContext)

	// ExitTypeDataVar is called when exiting the typeDataVar production.
	ExitTypeDataVar(c *TypeDataVarContext)

	// ExitListIds is called when exiting the listIds production.
	ExitListIds(c *ListIdsContext)

	// ExitPrint_prod is called when exiting the print_prod production.
	ExitPrint_prod(c *Print_prodContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpr_rel is called when exiting the expr_rel production.
	ExitExpr_rel(c *Expr_relContext)

	// ExitExpr_logic is called when exiting the expr_logic production.
	ExitExpr_logic(c *Expr_logicContext)

	// ExitExpr_arit is called when exiting the expr_arit production.
	ExitExpr_arit(c *Expr_aritContext)

	// ExitPrimitive is called when exiting the primitive production.
	ExitPrimitive(c *PrimitiveContext)
}
