parser grammar OptimizerParser;

options {
    tokenVocab = OptimizerLexer;
}

@header {
    import "OLC2_Project1/server/interpreter/AST"
    import "OLC2_Project1/server/interpreter/Optimizer/AbstractOptimizer"
    import "OLC2_Project1/server/interpreter/Optimizer/Primitive"
    import "OLC2_Project1/server/interpreter/Optimizer/Assignment"
    import "OLC2_Project1/server/interpreter/Optimizer/Control"
    import "OLC2_Project1/server/interpreter/Optimizer/Function"
    import arrayList "github.com/colegno/arraylist"
}

// Rules
start returns [AST.Tree tree]
    : headers list_temps listFuncs { $tree = AST.NewTreeOptimizer($listFuncs.l, $list_temps.l) }
;

headers returns [string code]
    : RINCLUDE RSTDIO RFLOAT RHEAP LEFT_BRACKET INTEGER RIGHT_BRACKET SEMICOLON RFLOAT RSTACK LEFT_BRACKET INTEGER RIGHT_BRACKET SEMICOLON RFLOAT RPSTACK SEMICOLON RFLOAT RPHEAP SEMICOLON {
            $code = $RINCLUDE.text
        }
;

list_temps returns [*arrayList.List l]
    : RFLOAT temps SEMICOLON { $l = $temps.l }
;

temps returns [*arrayList.List l]
    @init{
        $l = arrayList.New()
    }
    : sub = temps COMMA ID {
        $sub.l.Add($ID.text)
        $l = $sub.l
    }
    | ID { $l.Add($ID.text) }
;

listFuncs returns[*arrayList.List l]
    @init {
        $l = arrayList.New()
    }
    : subList = listFuncs function {
        $subList.l.Add($function.instr)
        $l = $subList.l
    }
    | function  { $l.Add($function.instr) }
;

function returns[AbstractOptimizer.Instruction instr]
    : RVOID ID LEFT_PAR RIGHT_PAR LEFT_KEY instructions RIGHT_KEY {
        $instr = Function.NewFunction($ID.line, localctx.(*FunctionContext).Get_ID().GetColumn(), $ID.text, $instructions.l)
    }
;

instructions returns[*arrayList.List l]
    @init{
        $l = arrayList.New()
    }
    : e += instruction* {
        listInt := localctx.(*InstructionsContext).GetE()
        for _, e := range listInt {
            if e.GetInstr() != nil {
                $l.Add(e.GetInstr())
            }
        }
    }
;

instruction returns [AbstractOptimizer.Instruction instr]
    : assign_stack { $instr = $assign_stack.instr }
    | assign_heap  { $instr = $assign_heap.instr }
    | assign       { $instr = $assign.instr }
    | if_instr     { $instr = $if_instr.instr }
    | goto_instr   { $instr = $goto_instr.instr }
    | label_instr  { $instr = $label_instr.instr }
;

assign_stack returns [AbstractOptimizer.Instruction instr]
    : RSTACK LEFT_BRACKET LEFT_PAR data_type RIGHT_PAR expr_valor RIGHT_BRACKET EQUAL expression SEMICOLON {
        $instr = Assignment.NewAssign($expr_valor.p, $expression.p, $EQUAL.line, localctx.(*Assign_stackContext).Get_EQUAL().GetColumn())
    }
;

assign_heap returns [AbstractOptimizer.Instruction instr]
    : RHEAP LEFT_BRACKET expr_valor RIGHT_BRACKET EQUAL expression SEMICOLON {
        $instr = Assignment.NewAssign($expr_valor.p, $expression.p, $EQUAL.line, localctx.(*Assign_heapContext).Get_EQUAL().GetColumn())
    }
;

assign returns [AbstractOptimizer.Instruction instr]
    : expr_valor EQUAL expression SEMICOLON {
        $instr = Assignment.NewAssign($expr_valor.p, $expression.p, $EQUAL.line, localctx.(*AssignContext).Get_EQUAL().GetColumn())
    }
;

if_instr returns [AbstractOptimizer.Instruction instr]
    : RIF LEFT_PAR expression RIGHT_PAR RGOTO ID SEMICOLON {
        $instr = Control.NewIf($expression.p, $ID.text, $RIF.line, localctx.(*If_instrContext).Get_ID().GetColumn())
    }
;

goto_instr returns [AbstractOptimizer.Instruction instr]
    : RGOTO ID SEMICOLON {
        $instr = Control.NewGoTo($ID.text, $ID.line, localctx.(*Goto_instrContext).Get_ID().GetColumn())
    }
;

label_instr returns [AbstractOptimizer.Instruction instr]
    : ID COLON {
        $instr = Control.NewLabel($ID.text, $ID.line, localctx.(*Label_instrContext).Get_ID().GetColumn())
    }
;

expression returns[AbstractOptimizer.Expression p]
    : expr_rel  { $p = $expr_rel.p }
;

expr_rel returns[AbstractOptimizer.Expression p]
    : opLeft = expr_rel op=( GREATER_THAN | LESS_THAN | GREATER_EQUALTHAN | LESS_EQUEALTHAN | EQUEAL_EQUAL | NOTEQUAL ) opRight = expr_rel {
        $p = Primitive.NewOperation($opLeft.p, $op.text, $opRight.p, $op.line, localctx.(*Expr_relContext).GetOp().GetColumn())
    }
    | expr_arit {
        $p = $expr_arit.p
    }
;

expr_arit returns [AbstractOptimizer.Expression p]
    : '-' opU = expression {

    }
    | opLeft = expr_arit op=('*'|'/') opRight = expr_arit {
        $p = Primitive.NewOperation($opLeft.p, $op.text, $opRight.p, $op.line, localctx.(*Expr_aritContext).GetOp().GetColumn())
    }
    | opLeft = expr_arit op=('+'|'-') opRight = expr_arit {
        $p = Primitive.NewOperation($opLeft.p, $op.text, $opRight.p, $op.line, localctx.(*Expr_aritContext).GetOp().GetColumn())
    }
    | expr_valor {
        $p = $expr_valor.p
    }
    | LEFT_PAR expression RIGHT_PAR {

    }
;

data_type returns[string data]
    : RINT { $data = $RINT.text }
    | RFLOAT { $data = $RFLOAT.text }
    | RCHAR { $data = $RCHAR.text }
    ;

expr_valor returns [AbstractOptimizer.Expression p]
    : FLOAT {
        num, err := strconv.ParseFloat($FLOAT.text, 64)
        if err != nil {
            fmt.Println(err)
        }
        $p = Primitive.NewNumber(num, $FLOAT.line, localctx.(*Expr_valorContext).Get_FLOAT().GetColumn())
    }
    | INTEGER {
        num, err := strconv.Atoi($INTEGER.text)
        if err != nil {
            fmt.Println(err)
        }
        $p = Primitive.NewNumber(num, $INTEGER.line, localctx.(*Expr_valorContext).Get_INTEGER().GetColumn())
    }
    | ID {
         $p = Primitive.NewTemp($ID.text, $ID.line, localctx.(*Expr_valorContext).Get_ID().GetColumn() )
     }
    | RPSTACK {
        $p = Primitive.NewTemp($RPSTACK.text, $RPSTACK.line, localctx.(*Expr_valorContext).Get_RPSTACK().GetColumn() )
    }
    | RPHEAP {
        $p = Primitive.NewTemp($RPHEAP.text, $RPHEAP.line, localctx.(*Expr_valorContext).Get_RPHEAP().GetColumn() )
    }
;