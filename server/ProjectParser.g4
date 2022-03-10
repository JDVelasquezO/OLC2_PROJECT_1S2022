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
    import "OLC2_Project1/server/interpreter/SymbolTable/Environment"
    import "OLC2_Project1/server/interpreter/AST/ExpressionSpecial"
    import "OLC2_Project1/server/interpreter/AST/Natives/DecArrays"
    import "OLC2_Project1/server/interpreter/AST/Expression/Access"
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
    : listFuncs { $tree = AST.NewTree($listFuncs.l) }
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

function returns[Abstract.Instruction instr]
    @init { listParams := arrayList.New() }
    : funcMain      { $instr = $funcMain.instr }
    | RFN ID LEFT_PAR RIGHT_PAR bloq   {
        $instr = Environment.NewFunction($RFN.line, localctx.(*FunctionContext).Get_RFN().GetColumn(), $ID.text, listParams, $bloq.content, "void")
     }
    | RFN ID LEFT_PAR RIGHT_PAR ARROW data_type bloq {
        $instr = Environment.NewFunction($RFN.line, localctx.(*FunctionContext).Get_RFN().GetColumn(), $ID.text, listParams, $bloq.content, $data_type.data)
    }
    | RFN ID LEFT_PAR listParams RIGHT_PAR bloq   {
         $instr = Environment.NewFunction($RFN.line, localctx.(*FunctionContext).Get_RFN().GetColumn(), $ID.text, $listParams.l, $bloq.content, "void")
    }
    | RFN ID LEFT_PAR listParams RIGHT_PAR ARROW data_type bloq {
        $instr = Environment.NewFunction($RFN.line, localctx.(*FunctionContext).Get_RFN().GetColumn(), $ID.text, $listParams.l, $bloq.content, $data_type.data)
    }
    ;

funcMain returns [Abstract.Instruction instr]
    @init { listParams := arrayList.New() }
    : RFN RMAIN LEFT_PAR RIGHT_PAR bloq { $instr = Environment.NewFunction($RFN.line, localctx.(*FuncMainContext).Get_RFN().GetColumn(), "main", listParams, $bloq.content, "void") }
    ;

bloq returns [*arrayList.List content]
    : LEFT_KEY instructions RIGHT_KEY   { $content = $instructions.l }
    | LEFT_KEY RIGHT_KEY                { $content = arrayList.New() }
    ;

listParams returns [*arrayList.List l]
    @init {
        $l = arrayList.New()
    }
    : List = listParams COMMA ID COLON data_type {
        listIds := arrayList.New()
        listIds.Add(Expression.NewIdentifier($ID.text, $ID.line, localctx.(*ListParamsContext).Get_ID().GetColumn()))
        newDec := Natives.NewDeclaration(listIds, false, $data_type.data)
        $List.l.Add(newDec)
        $l = $List.l
    }
    | ID COLON data_type {
        listIds := arrayList.New()
        listIds.Add(Expression.NewIdentifier($ID.text, $ID.line, localctx.(*ListParamsContext).Get_ID().GetColumn()))
        newDec := Natives.NewDeclaration(listIds, false, $data_type.data)
        $l.Add(newDec)
    }
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
    | bucle_prod        { $instr = $bucle_prod.instr }
    | called_func       { $instr = $called_func.instr }
    | expr_arit         { $instr = $expr_arit.instr }
    | primitive         { $instr = $primitive.instr }
    | expr_cast         { $instr = $expr_cast.instr }
    | expr_rel          { $instr = $expr_rel.instr }
    | expr_logic        { $instr = $expr_logic.instr }
    | dec_arr           { $instr = $dec_arr.instr }
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
    : DECLARAR MUT? ids_type EQUAL expression SEMICOLON? {
        if ($MUT.text != "") {
            $instr = Natives.NewDeclarationInit($ids_type.list, $expression.p, true, "")
        } else {
            $instr = Natives.NewDeclarationInit($ids_type.list, $expression.p, false, "")
        }
    }
    | DECLARAR MUT? ids_type COLON data_type EQUAL expression SEMICOLON? {
        if ($MUT.text != "") {
            $instr = Natives.NewDeclarationInit($ids_type.list, $expression.p, true, $data_type.data)
        } else {
            $instr = Natives.NewDeclarationInit($ids_type.list, $expression.p, false, $data_type.data)
        }
    }
    | DECLARAR MUT? ids_type SEMICOLON? {
        if ($MUT.text != "") {
            $instr = Natives.NewDeclaration($ids_type.list, true, "")
        } else {
            $instr = Natives.NewDeclaration($ids_type.list, true, "")
        }
    }
    | DECLARAR MUT? ids_type COLON data_type SEMICOLON? {
        if ($MUT.text != "") {
            $instr = Natives.NewDeclaration($ids_type.list, true, $data_type.data)
        } else {
            $instr = Natives.NewDeclaration($ids_type.list, true, $data_type.data)
        }
    }
    ;

assign_prod returns [Abstract.Instruction instr]
    : listIds EQUAL expression SEMICOLON? {
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

conditional_prod returns [Abstract.Instruction instr, Abstract.Expression p]
    : RIF expression bloq {
        $instr = Natives.NewIf($expression.p, $bloq.content, nil, nil, $RIF.line, localctx.(*Conditional_prodContext).Get_RIF().GetColumn())
        $p = Natives.NewIf($expression.p, $bloq.content, nil, nil, $RIF.line, localctx.(*Conditional_prodContext).Get_RIF().GetColumn())
    }
    | RIF expression bif=bloq RELSE belse=bloq {
        $instr = Natives.NewIf($expression.p, $bif.content, nil, $belse.content, $RIF.line, localctx.(*Conditional_prodContext).Get_RIF().GetColumn())
        $p = Natives.NewIf($expression.p, $bif.content, nil, $belse.content, $RIF.line, localctx.(*Conditional_prodContext).Get_RIF().GetColumn())
    }
    | RIF expression bif=bloq list_else_if RELSE belse=bloq {
        $instr = Natives.NewIf($expression.p, $bif.content, $list_else_if.list, $belse.content, $RIF.line, localctx.(*Conditional_prodContext).Get_RIF().GetColumn() )
        $p = Natives.NewIf($expression.p, $bif.content, $list_else_if.list, $belse.content, $RIF.line, localctx.(*Conditional_prodContext).Get_RIF().GetColumn() )
    }
    ;

list_else_if returns [*arrayList.List list]
    @init {
        $list = arrayList.New()
    }
    : newList += else_if+ {
        listInt := localctx.(*List_else_ifContext).GetNewList()
        for _, e := range listInt {
            $list.Add(e.GetInstr())
        }
    }
    ;

else_if returns [Abstract.Instruction instr]
    : RELSE RIF expression bloq { $instr = Natives.NewIf($expression.p, $bloq.content, nil, nil, $RIF.line, localctx.(*Else_ifContext).Get_RIF().GetColumn()) }
    ;

bucle_prod returns [Abstract.Instruction instr]
    : RWHILE expression bloq { $instr = Natives.NewWhile($expression.p, $bloq.content, $RWHILE.line, localctx.(*Bucle_prodContext).Get_RWHILE().GetColumn()) }
    ;

called_func returns [Abstract.Instruction instr, Abstract.Expression p]
    : ID LEFT_PAR RIGHT_PAR {
        $instr = ExpressionSpecial.NewCallFunction($ID.text, arrayList.New())
        $p = ExpressionSpecial.NewCallFunction($ID.text, arrayList.New())
    }
    | ID LEFT_PAR listExpressions RIGHT_PAR {
        $instr = ExpressionSpecial.NewCallFunction($ID.text, $listExpressions.l)
        $p = ExpressionSpecial.NewCallFunction($ID.text, $listExpressions.l)
    }
    ;

listExpressions returns [*arrayList.List l]
    @init {
        $l = arrayList.New()
    }
    : List = listExpressions COMMA expression {
        $List.l.Add($expression.p)
        $l = $List.l
    }
    | expression { $l.Add($expression.p) }
    ;

dec_arr returns [Abstract.Instruction instr]
    : DECLARAR ID COLON listDim EQUAL expression SEMICOLON { $instr = DecArrays.NewDecArray($listDim.length, $ID.text, $expression.p, SymbolTable.INTEGER) }
    ;

listDim returns[int length, string data]
    @init { $length = 0 }
    : LEFT_BRACKET l = listDim SEMICOLON expression RIGHT_BRACKET { $length = $l.length + 1
                                                                 $data = $l.data
                                                               }
    | LEFT_BRACKET data_type SEMICOLON expression RIGHT_BRACKET { $length = 1
                                                              $data = $data_type.data }
    ;

expression returns [Abstract.Expression p]
    : expr_valor                 { $p = $expr_valor.p }
    | conditional_prod           { $p = $conditional_prod.p }
    | expr_rel                   { $p = $expr_rel.p }
    | expr_arit                  { $p = $expr_arit.p }
    | expr_logic                 { $p = $expr_logic.p }
    | arraydata                  { $p = $arraydata.p }
;

arraydata returns [Abstract.Expression p]
    : LEFT_BRACKET listExpressions RIGHT_BRACKET { $p = ExpressionSpecial.NewValueArray($listExpressions.l) }
    ;

access_array returns [Abstract.Expression p, Abstract.Instruction instr]
    : ID listInArray {
        $p = Access.NewAccessArray($ID.text, $listInArray.l)
        $instr = Access.NewAccessArray($ID.text, $listInArray.l)
    }
    ;

listInArray returns [*arrayList.List l]
    @init {
        $l = arrayList.New()
    }
    : sublist = listInArray inArray {
        $sublist.l.Add($inArray.p)
        $l = $sublist.l
    }
    | inArray { $l.Add($inArray.p) }
    ;

inArray returns[Abstract.Expression p]
    : LEFT_BRACKET expression RIGHT_BRACKET     { $p = $expression.p }
    ;

expr_rel returns[Abstract.Expression p, Abstract.Instruction instr]
    : opLeft = expr_rel op=( GREATER_THAN | LESS_THAN | GREATER_EQUALTHAN | LESS_EQUEALTHAN | EQUEAL_EQUAL | NOTEQUAL ) opRight = expr_rel {
        $p = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_relContext).GetOp().GetColumn() )
        $instr = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_relContext).GetOp().GetColumn() )
        }
    | expr_arit {
        $p = $expr_arit.p
        $instr = $expr_arit.instr
        }
    ;

expr_arit returns [Abstract.Expression p, Abstract.Instruction instr]
    : '-' opU = expression {
        $p = Expression.NewOperation($opU.p, "-", nil, true, localctx.(*Expr_aritContext).GetOpU().GetStart().GetLine(), localctx.(*Expr_aritContext).GetOpU().GetStart().GetColumn() )
        $instr = Expression.NewOperation($opU.p, "-", nil, true, localctx.(*Expr_aritContext).GetOpU().GetStart().GetLine(), localctx.(*Expr_aritContext).GetOpU().GetStart().GetColumn() )
        }
    | pow_op LEFT_PAR opLeft = expr_arit COMMA opRight = expr_arit RIGHT_PAR {
        $p = Expression.NewOperation($opLeft.p, $pow_op.op, $opRight.p, false, localctx.(*Expr_aritContext).Get_pow_op().GetStart().GetLine(), localctx.(*Expr_aritContext).Get_pow_op().GetStart().GetColumn())
        $instr = Expression.NewOperation($opLeft.p, $pow_op.op, $opRight.p, false, localctx.(*Expr_aritContext).Get_pow_op().GetStart().GetLine(), localctx.(*Expr_aritContext).Get_pow_op().GetStart().GetColumn())
        }
    | opLeft = expr_arit op=('*'|'/'|'%') opRight = expr_arit {
        $p = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_aritContext).GetOp().GetColumn() )
        $instr = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_aritContext).GetOp().GetColumn() )
        }
    | opLeft = expr_arit op=('+'|'-') opRight = expr_arit {
        $p = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_aritContext).GetOp().GetColumn() )
        $instr = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_aritContext).GetOp().GetColumn() )
        }
    | expr_valor {
        $p = $expr_valor.p
        $instr = $expr_valor.instr
    }
    | LEFT_PAR expression RIGHT_PAR {
        $p = $expression.p
        $instr = nil
        }
;

expr_valor returns [Abstract.Expression p, Abstract.Instruction instr]
    : called_func {
        $p = $called_func.p
        $instr = $called_func.instr
    }
    | primitive {
          $p = $primitive.p
          $instr = $primitive.instr
    }
    | expr_cast   {
        $p = $expr_cast.p
        $instr = $expr_cast.instr
    }
    | access_array {
        $p = $access_array.p
        $instr = $access_array.instr
    }
    ;

pow_op returns [string op]
    : RINTEGER HERITAGE POWI { $op = $POWI.text }
    | RREAL HERITAGE POWF { $op = $POWF.text }
    ;

expr_logic returns[Abstract.Expression p, Abstract.Instruction instr]
    : ADMIRATION opU=expression {
        $p = Expression.NewOperation($opU.p, "!", nil, true, localctx.(*Expr_logicContext).GetOpU().GetStart().GetLine(), localctx.(*Expr_logicContext).GetOpU().GetStart().GetColumn())
        $instr = Expression.NewOperation($opU.p, "!", nil, true, localctx.(*Expr_logicContext).GetOpU().GetStart().GetLine(), localctx.(*Expr_logicContext).GetOpU().GetStart().GetColumn())
        }
    | opLeft = expr_rel op=( AND | OR ) opRight = expr_rel {
        $p = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_logicContext).GetOp().GetColumn() )
        $instr = Expression.NewOperation($opLeft.p, $op.text, $opRight.p, false, $op.line, localctx.(*Expr_logicContext).GetOp().GetColumn() )
    }
    ;

expr_cast returns[Abstract.Expression p, Abstract.Instruction instr]
    : LEFT_PAR expression RAS data_type RIGHT_PAR {
        if $data_type.data == "i64" {
            $p = Expression.NewCast($expression.p, SymbolTable.INTEGER, $RAS.line, localctx.(*Expr_castContext).Get_RAS().GetColumn() )
            $instr = Expression.NewCast($expression.p, SymbolTable.INTEGER, $RAS.line, localctx.(*Expr_castContext).Get_RAS().GetColumn())
        } else if $data_type.data == "f64" {
            $p = Expression.NewCast($expression.p, SymbolTable.FLOAT, $RAS.line, localctx.(*Expr_castContext).Get_RAS().GetColumn())
            $instr = Expression.NewCast($expression.p, SymbolTable.FLOAT, $RAS.line, localctx.(*Expr_castContext).Get_RAS().GetColumn())
        }
    }
    ;

data_type returns[string data]
    : RINTEGER { $data = $RINTEGER.text }
    | RREAL { $data = $RREAL.text }
    | RSTR { $data = $RSTR.text }
    | RSTRING { $data = $RSTRING.text }
    | RBOOLEAN { $data = $RBOOLEAN.text }
    | RCHAR { $data = $RCHAR.text }
    ;

primitive returns [Abstract.Expression p, Abstract.Instruction instr]
    :INTEGER{
        num, err := strconv.Atoi($INTEGER.text)
        if err != nil {
            fmt.Println(err)
        }
        $p = Expression.NewPrimitive(num, SymbolTable.INTEGER, $INTEGER.line, localctx.(*PrimitiveContext).Get_INTEGER().GetColumn())
        $instr = Expression.NewPrimitive(num, SymbolTable.INTEGER, $INTEGER.line, localctx.(*PrimitiveContext).Get_INTEGER().GetColumn())
    }
    | FLOAT {
        num, err := strconv.ParseFloat($FLOAT.text, 64)
        if err != nil {
            fmt.Println(err)
        }
        $p = Expression.NewPrimitive(num, SymbolTable.FLOAT, $FLOAT.line, localctx.(*PrimitiveContext).Get_FLOAT().GetColumn())
        $instr = Expression.NewPrimitive(num, SymbolTable.FLOAT, $FLOAT.line, localctx.(*PrimitiveContext).Get_FLOAT().GetColumn())
    }
    | STRING ((DOT TOSTRING | DOT TOOWNED) LEFT_PAR RIGHT_PAR)? {
        str := $STRING.text[1:len($STRING.text)-1]
        if $TOSTRING.text != "" || $TOOWNED.text != "" {
            $p = Expression.NewPrimitive(str, SymbolTable.STRING, $STRING.line, localctx.(*PrimitiveContext).Get_STRING().GetColumn())
            $instr = Expression.NewPrimitive(str, SymbolTable.STRING, $STRING.line, localctx.(*PrimitiveContext).Get_STRING().GetColumn())
        } else {
            $p = Expression.NewPrimitive(str, SymbolTable.STR, $STRING.line, localctx.(*PrimitiveContext).Get_STRING().GetColumn())
            $instr = Expression.NewPrimitive(str, SymbolTable.STR, $STRING.line, localctx.(*PrimitiveContext).Get_STRING().GetColumn())
        }
    }
    | CHAR {
        chr := $CHAR.text[1:len($CHAR.text)-1]
        $p = Expression.NewPrimitive(chr, SymbolTable.CHAR, $CHAR.line, localctx.(*PrimitiveContext).Get_CHAR().GetColumn())
        $instr = Expression.NewPrimitive(chr, SymbolTable.CHAR, $CHAR.line, localctx.(*PrimitiveContext).Get_CHAR().GetColumn())
    }
    | BOOLEAN {
        str := $BOOLEAN.text
        fmt.Println(str)
        if str == "true" {
            $p = Expression.NewPrimitive(true, SymbolTable.BOOLEAN, $BOOLEAN.line, localctx.(*PrimitiveContext).Get_BOOLEAN().GetColumn())
            $instr = Expression.NewPrimitive(true, SymbolTable.BOOLEAN, $BOOLEAN.line, localctx.(*PrimitiveContext).Get_BOOLEAN().GetColumn())
        } else {
            $p = Expression.NewPrimitive(false, SymbolTable.BOOLEAN, $BOOLEAN.line, localctx.(*PrimitiveContext).Get_BOOLEAN().GetColumn())
            $instr = Expression.NewPrimitive(false, SymbolTable.BOOLEAN, $BOOLEAN.line, localctx.(*PrimitiveContext).Get_BOOLEAN().GetColumn())
        }
    }
    | ID {
        $p = Expression.NewIdentifier($ID.text, $ID.line, localctx.(*PrimitiveContext).Get_ID().GetColumn() )
        $instr = Expression.NewIdentifier($ID.text, $ID.line, localctx.(*PrimitiveContext).Get_ID().GetColumn() )
    }
;