parser grammar ProjectParser;

options {
    tokenVocab = ProjectLexer;
}

@header {
    import "OLC2_Project1/server/interpreter/AST"
    import "OLC2_Project1/server/interpreter/Abstract"
    import "OLC2_Project1/server/interpreter/AST/Expression"
    import "OLC2_Project1/server/interpreter/AST/Natives"
    import "OLC2_Project1/server/interpreter/SymbolTable"
    import arrayList "github.com/colegno/arraylist"
}

@members {
    type BinaryOperation struct {
        Op1 int
        Operator string
        Op2 int
    }
}

// Rules
start returns [AST.Tree tree]
    : instructions { $tree = AST.NewTree($instructions.l) }
;

instructions returns [*arrayList.List l]
    @init{
        $l = arrayList.New()
    }
    : e += instruction* {
        listInt := localctx.(*InstructionsContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
        fmt.Print("tipo %T", localctx.(*InstructionsContext).GetE())
    }
;

instruction returns [Abstract.Instruction instr]
    : print_prod        { $instr = $print_prod.instr }
    | declaration_prod  { $instr = $declaration_prod.instr }
    | assign_prod       { $instr = $assign_prod.instr }
    | conditional_prod  { $instr = $conditional_prod.instr }
    ;

print_prod returns [Abstract.Instruction instr]
    : SENTENCIA DOT CONSOLA LEFT_PAR expression RIGHT_PAR SEMICOLON { $instr = Natives.NewPrint($expression.p, false, $LEFT_PAR.line, localctx.(*Print_prodContext).Get_LEFT_PAR().GetColumn()) }
    | SENTENCIA DOT CONSOLALN LEFT_PAR expression RIGHT_PAR SEMICOLON { $instr = Natives.NewPrint($expression.p, true, $LEFT_PAR.line, localctx.(*Print_prodContext).Get_LEFT_PAR().GetColumn()) }
    ;

declaration_prod returns [Abstract.Instruction instr]
    : DECLARAR listIds EQUAL expression SEMICOLON {
        $instr = Natives.NewDeclarationInit($listIds.list, $expression.p)
        }
    | DECLARAR listIds SEMICOLON {
        $instr = Natives.NewDeclaration($listIds.list)
    }
    ;

assign_prod returns [Abstract.Instruction instr]
    : listIds EQUAL expression SEMICOLON {
        $instr = Natives.NewAssign($listIds.list, $expression.p)
    }
    ;

conditional_prod returns [Abstract.Instruction instr]
    : RIF LEFT_PAR expression RIGHT_PAR bloq { $instr = Natives.NewIf($expression.p, $bloq.content, nil, nil) }
//     | RIF LEFT_PAR expression RIGHT_PAR bif=bloq RELSE belse=bloq
//     | RIF LEFT_PAR expression RIGHT_PAR bif=bloq list_else_if
//     | RIF LEFT_PAR expression RIGHT_PAR bif=bloq list_else_if RELSE belse=bloq
    ;

//list_else_if returns [*arrayList.List list]
//    : l += else_if+ {
//        acumulator := localctx.(*List_else_ifContext).GetList()
//
//        for _, e := range acumulator {
//            $l.Add(e.GetInstr)
//        }
//    }
//    ;

//else_if returns[Abstract.Natives instr]
//    : RELSE RIF LEFT_PAR expression RIGHT_PAR bloq { $instr = Natives.NewIf() }
//    ;

bloq returns [*arrayList.List content]
    : LEFT_BRACKET instructions RIGHT_BRACKET   { $content = $instructions.l }
    | LEFT_BRACKET RIGHT_BRACKET                { $content = arrayList.New() }
    ;

listIds returns [*arrayList.List list]
    @init {
        $list = arrayList.New()
    }
    : sub = listIds COMMA ID  {
                                $sub.list.Add(Expression.NewIdentifier($ID.text, $ID.line, localctx.(*ListIdsContext).Get_ID().GetColumn()))
                                $list = $sub.list
                              }
    | ID                    { $list.Add(Expression.NewIdentifier($ID.text, $ID.line, localctx.(*ListIdsContext).Get_ID().GetColumn())) }
    ;

typeDataVar returns [SymbolTable.DataType t]
    : RINTEGER  { $t = SymbolTable.INTEGER }
    | RSTRING   { $t = SymbolTable.STRING }
    | RREAL     { $t = SymbolTable.FLOAT }
    | RBOOLEAN  { $t = SymbolTable.BOOLEAN }
    ;

expression returns [Abstract.Expression p]
    : expr_rel   { $p = $expr_rel.p }
    | expr_arit  { $p = $expr_arit.p }
    | expr_logic { $p = $expr_logic.p }
;

expr_rel returns[Abstract.Expression p]
    : opLeft = expr_rel op=( GREATER_THAN | LESS_THAN | GREATER_EQUALTHAN | LESS_EQUEALTHAN | EQUEAL_EQUAL ) opRight = expr_rel { $p = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_relContext).GetOp().GetColumn() )}
    | expr_arit { $p = $expr_arit.p }
    ;

expr_logic returns[Abstract.Expression p]
    : opLeft = expr_rel op=( RAND | ROR ) opRight = expr_rel { $p = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_logicContext).GetOp().GetColumn() )}
    ;

expr_arit returns [Abstract.Expression p]
    : '-' opU = expression { $p = Expression.NewOperation($opU.p, "-", nil, true, localctx.(*Expr_aritContext).GetOpU().GetStart().GetLine(), localctx.(*Expr_aritContext).GetOpU().GetStart().GetColumn() ) }
    | opLeft = expr_arit op=('*'|'/'|'%') opRight = expr_arit {$p = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_aritContext).GetOp().GetColumn() )}
    | opLeft = expr_arit op=('+'|'-') opRight = expr_arit {$p = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_aritContext).GetOp().GetColumn() )}
    | primitive {$p = $primitive.p}
    | LEFT_PAR expression RIGHT_PAR {$p = $expression.p}
;

primitive returns [Abstract.Expression p]
    :INTEGER{
        num, err := strconv.Atoi($INTEGER.text)
        if err != nil {
            fmt.Println(err)
        }
        $p = Expression.NewPrimitive(num, SymbolTable.INTEGER, $INTEGER.line, localctx.(*PrimitiveContext).Get_INTEGER().GetColumn())
    }
    | FLOAT {
        num, err := strconv.ParseFloat($FLOAT.text, 64)
        if err != nil {
            fmt.Println(err)
        }
        $p = Expression.NewPrimitive(num, SymbolTable.FLOAT, $FLOAT.line, localctx.(*PrimitiveContext).Get_FLOAT().GetColumn())
    }
    | STRING {
        str := $STRING.text[1:len($STRING.text)-1]
        $p = Expression.NewPrimitive(str, SymbolTable.STRING, $STRING.line, localctx.(*PrimitiveContext).Get_STRING().GetColumn())
    }
    | CHAR {
        chr := $CHAR.text
        $p = Expression.NewPrimitive(chr, SymbolTable.CHAR, $CHAR.line, localctx.(*PrimitiveContext).Get_CHAR().GetColumn())
    }
    | BOOLEAN {
        str := $BOOLEAN.text
        fmt.Println(str)
        $p = Expression.NewPrimitive(str, SymbolTable.BOOLEAN, $BOOLEAN.line, localctx.(*PrimitiveContext).Get_BOOLEAN().GetColumn())
    }
    | ID { $p = Expression.NewIdentifier($ID.text, $ID.line, localctx.(*PrimitiveContext).Get_ID().GetColumn() ) }
;