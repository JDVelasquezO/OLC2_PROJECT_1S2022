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
    : PRINT ADMIRATION LEFT_PAR expression RIGHT_PAR SEMICOLON { $instr = Natives.NewPrint($expression.p, false, $LEFT_PAR.line, localctx.(*Print_prodContext).Get_LEFT_PAR().GetColumn()) }
    | PRINT ADMIRATION LEFT_PAR opBefore = expression COMMA listVars RIGHT_PAR SEMICOLON { $instr = Natives.NewPrintWithAfter($opBefore.p, $listVars.list, false, $LEFT_PAR.line, localctx.(*Print_prodContext).Get_LEFT_PAR().GetColumn()) }
    | PRINTLN ADMIRATION LEFT_PAR expression RIGHT_PAR SEMICOLON { $instr = Natives.NewPrint($expression.p, true, $LEFT_PAR.line, localctx.(*Print_prodContext).Get_LEFT_PAR().GetColumn()) }
    | PRINTLN ADMIRATION LEFT_PAR opBefore = expression COMMA listVars RIGHT_PAR SEMICOLON { $instr = Natives.NewPrintWithAfter($opBefore.p, $listVars.list, true, $LEFT_PAR.line, localctx.(*Print_prodContext).Get_LEFT_PAR().GetColumn()) }
    ;

listVars returns [*arrayList.List list]
    @init {
        $list = arrayList.New()
    }
    : sub = listVars COMMA expression {
        $sub.list.Add($expression.p)
        $list = $sub.list
    }
    | expression {
        $list.Add($expression.p)
    }
    ;

declaration_prod returns [Abstract.Instruction instr]
    : DECLARAR MUT? ids_type (COLON data_type)? EQUAL expression SEMICOLON {
        if ($MUT.text != "") {
            if ($data_type.data != nil) {
                $instr = Natives.NewDeclarationInit($ids_type.list, $expression.p, true, $data_type.data)
            } else {
                $instr = Natives.NewDeclarationInit($ids_type.list, $expression.p, true, "")
            }
        } else {
            if ($data_type.data != nil) {
                $instr = Natives.NewDeclarationInit($ids_type.list, $expression.p, false, $data_type.data)
            } else {
                $instr = Natives.NewDeclarationInit($ids_type.list, $expression.p, false, "")
            }
        }
    }
    | DECLARAR MUT? ids_type (COLON data_type)? SEMICOLON {
        if ($MUT.text != "") {
            $instr = Natives.NewDeclaration($ids_type.list, true, $data_type.data)
        } else {
            $instr = Natives.NewDeclaration($ids_type.list, false, $data_type.data)
        }
    }
    ;

assign_prod returns [Abstract.Instruction instr]
    : listIds EQUAL expression SEMICOLON {
        $instr = Natives.NewAssign($listIds.list, $expression.p)
    }
    ;

ids_type returns [*arrayList.List list]
    : listIds { $list = $listIds.list }
    | listIds COLON RSTRING { $list = $listIds.list }
    | listIds COLON REFERENCE RSTR { $list = $listIds.list }
    ;

listIds returns [*arrayList.List list]
    @init {
        $list = arrayList.New()
    }
    : sub = listIds COMMA ID  {
                                $sub.list.Add(Expression.NewIdentifier($ID.text, $ID.line, localctx.(*ListIdsContext).Get_ID().GetColumn()))
                                $list = $sub.list
                              }
    | ID  {  $list.Add(Expression.NewIdentifier($ID.text, $ID.line, localctx.(*ListIdsContext).Get_ID().GetColumn())) }
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

expression returns [Abstract.Expression p]
    : expr_rel    { $p = $expr_rel.p }
    | expr_arit   { $p = $expr_arit.p }
    | expr_logic  { $p = $expr_logic.p }
    | expr_cast   { $p = $expr_cast.p }
;

expr_rel returns[Abstract.Expression p]
    : opLeft = expr_rel op=( GREATER_THAN | LESS_THAN | GREATER_EQUALTHAN | LESS_EQUEALTHAN | EQUEAL_EQUAL ) opRight = expr_rel { $p = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_relContext).GetOp().GetColumn() )}
    | expr_arit { $p = $expr_arit.p }
    ;

expr_arit returns [Abstract.Expression p]
    : '-' opU = expression { $p = Expression.NewOperation($opU.p, "-", nil, true, localctx.(*Expr_aritContext).GetOpU().GetStart().GetLine(), localctx.(*Expr_aritContext).GetOpU().GetStart().GetColumn() ) }
    | pow_op LEFT_PAR opLeft = expr_arit COMMA opRight = expr_arit RIGHT_PAR { $p = Expression.NewOperation($opLeft.p, $pow_op.op, $opRight.p, false, localctx.(*Expr_aritContext).Get_pow_op().GetStart().GetLine(), localctx.(*Expr_aritContext).Get_pow_op().GetStart().GetColumn()) }
    | opLeft = expr_arit op=('*'|'/'|'%') opRight = expr_arit {$p = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_aritContext).GetOp().GetColumn() )}
    | opLeft = expr_arit op=('+'|'-') opRight = expr_arit {$p = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_aritContext).GetOp().GetColumn() )}
    | primitive {$p = $primitive.p}
    | expr_cast   { $p = $expr_cast.p }
    | LEFT_PAR expression RIGHT_PAR {$p = $expression.p}
;

pow_op returns [string op]
    : RINTEGER HERITAGE POWI { $op = $POWI.text }
    | RREAL HERITAGE POWF { $op = $POWF.text }
    ;

expr_logic returns[Abstract.Expression p]
    : ADMIRATION opU=expression { $p = Expression.NewOperation($opU.p, "!", nil, true, localctx.(*Expr_logicContext).GetOpU().GetStart().GetLine(), localctx.(*Expr_logicContext).GetOpU().GetStart().GetColumn()) }
    | opLeft = expr_rel op=( AND | OR ) opRight = expr_rel { $p = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_logicContext).GetOp().GetColumn() )}
    ;

expr_cast returns[Abstract.Expression p]
    : LEFT_PAR expression RAS data_type RIGHT_PAR {
        if $data_type.data == "i64" {
            $p = Expression.NewCast($expression.p, SymbolTable.INTEGER)
        } else if $data_type.data == "f64" {
            $p = Expression.NewCast($expression.p, SymbolTable.FLOAT)
        }
    }
    ;

data_type returns[string data]
    : RINTEGER { $data = $RINTEGER.text }
    | RREAL { $data = $RREAL.text }
    | RSTRING { $data = $RSTRING.text }
    | RBOOLEAN { $data = $RBOOLEAN.text }
    | RCHAR { $data = $RCHAR.text }
    |
    ;

//type_number returns [string type_num]
//    : RINTEGER { type_num = $RINTEGER.text }
//    | RREAL { type_num = $RREAL.text }
//    ;

//type_power returns [string mod_pow]
//    : POWI { mod_pow = $POWI.text }
//    | POWF { mod_pow = $POWF.text }
//    ;

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
        chr := $CHAR.text[1:len($CHAR.text)-1]
        $p = Expression.NewPrimitive(chr, SymbolTable.CHAR, $CHAR.line, localctx.(*PrimitiveContext).Get_CHAR().GetColumn())
    }
    | BOOLEAN {
        str := $BOOLEAN.text
        fmt.Println(str)
        $p = Expression.NewPrimitive(str, SymbolTable.BOOLEAN, $BOOLEAN.line, localctx.(*PrimitiveContext).Get_BOOLEAN().GetColumn())
    }
    | ID { $p = Expression.NewIdentifier($ID.text, $ID.line, localctx.(*PrimitiveContext).Get_ID().GetColumn() ) }
;