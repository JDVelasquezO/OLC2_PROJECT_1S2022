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
    import "OLC2_Project1/server/interpreter/Optimizer/Print"
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
    : assign_stack   { $instr = $assign_stack.instr }
    | assign_heap    { $instr = $assign_heap.instr }
    | assign         { $instr = $assign.instr }
    | if_instr       { $instr = $if_instr.instr }
    | goto_instr     { $instr = $goto_instr.instr }
    | label_instr    { $instr = $label_instr.instr }
    | printf_instr   { $instr = $printf_instr.instr }
    | return_instr   { $instr = $return_instr.instr }
    | getHeap_instr  { $instr = $getHeap_instr.instr }
    | getStack_instr { $instr = $getStack_instr.instr }
    | call_instr     { $instr = $call_instr.instr }
;

assign_stack returns [AbstractOptimizer.Instruction instr]
    : RSTACK LEFT_BRACKET LEFT_PAR data_type RIGHT_PAR expr_valor RIGHT_BRACKET EQUAL expression SEMICOLON {
        $instr = Assignment.NewAssignHeapStack($expr_valor.p, $expression.p, $EQUAL.line, localctx.(*Assign_stackContext).Get_EQUAL().GetColumn(), $RSTACK.text)
    }
;

assign_heap returns [AbstractOptimizer.Instruction instr]
    : RHEAP LEFT_BRACKET LEFT_PAR data_type RIGHT_PAR expr_valor RIGHT_BRACKET EQUAL expression SEMICOLON {
        $instr = Assignment.NewAssignHeapStack($expr_valor.p, $expression.p, $EQUAL.line, localctx.(*Assign_heapContext).Get_EQUAL().GetColumn(), $RHEAP.text)
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

printf_instr returns [AbstractOptimizer.Instruction instr]
    : RPRINTF LEFT_PAR STRING COMMA convert expr_print RIGHT_PAR SEMICOLON {
        $instr = Print.NewPrintf($STRING.text, $convert.dtype, $expr_print.id, $STRING.line, localctx.(*Printf_instrContext).Get_STRING().GetColumn())
    }
;

return_instr returns [AbstractOptimizer.Instruction instr]
    : RRETURN expression SEMICOLON {
        $instr = Control.NewReturn($expression.p, $RRETURN.line, localctx.(*Return_instrContext).Get_RRETURN().GetColumn())
    }
    | RRETURN SEMICOLON {
        $instr = Control.NewReturn(nil, $RRETURN.line, localctx.(*Return_instrContext).Get_RRETURN().GetColumn())
    }
;

getHeap_instr returns [AbstractOptimizer.Instruction instr]
    : expr_valor EQUAL RHEAP LEFT_BRACKET LEFT_PAR data_type RIGHT_PAR expression RIGHT_BRACKET SEMICOLON {
        $instr = Assignment.NewGetHeapStack($expr_valor.p, $expression.p, $EQUAL.line, localctx.(*GetHeap_instrContext).Get_EQUAL().GetColumn(), $RHEAP.text)
    }
;

getStack_instr returns [AbstractOptimizer.Instruction instr]
    : expr_valor EQUAL RSTACK LEFT_BRACKET LEFT_PAR data_type RIGHT_PAR expression RIGHT_BRACKET SEMICOLON {
        $instr = Assignment.NewGetHeapStack($expr_valor.p, $expression.p, $EQUAL.line, localctx.(*GetStack_instrContext).Get_EQUAL().GetColumn(), $RSTACK.text)
    }
;

call_instr returns [AbstractOptimizer.Instruction instr]
    : ID LEFT_PAR RIGHT_PAR SEMICOLON {
        $instr = Function.NewCall($ID.text, $ID.line, localctx.(*Call_instrContext).Get_ID().GetColumn())
    }
;

expr_print returns [AbstractOptimizer.Expression id]
    : expr_valor {
        $id = $expr_valor.p
    }
;

convert returns [string dtype]
    : LEFT_PAR RINT RIGHT_PAR {
        $dtype = $RINT.text
    }
    | LEFT_PAR RFLOAT RIGHT_PAR {
        $dtype = $RFLOAT.text
    }
    | LEFT_PAR RCHAR RIGHT_PAR {
        $dtype = $RCHAR.text
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
    : op='-' opU = expression {
        $p = Primitive.NewOperation($opU.p, "-", nil, $op.line, localctx.(*Expr_aritContext).GetOp().GetColumn())
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
    | TEMP {
         $p = Primitive.NewTemp($TEMP.text, $TEMP.line, localctx.(*Expr_valorContext).Get_TEMP().GetColumn() )
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