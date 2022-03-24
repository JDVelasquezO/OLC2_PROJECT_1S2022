parser grammar ProjectParser;

options {
    tokenVocab = ProjectLexer;
}

@header {
    import "OLC2_Project1/server/interpreter/AST"
    import "OLC2_Project1/server/interpreter/Abstract"
    import "OLC2_Project1/server/interpreter/AST/Expression"
    import "OLC2_Project1/server/interpreter/AST/Natives"
    import "OLC2_Project1/server/interpreter/AST/Natives/BucleForIn"
    import "OLC2_Project1/server/interpreter/SymbolTable"
    import "OLC2_Project1/server/interpreter/SymbolTable/Environment"
    import "OLC2_Project1/server/interpreter/AST/ExpressionSpecial"
    import "OLC2_Project1/server/interpreter/AST/Natives/DecArrays"
    import "OLC2_Project1/server/interpreter/AST/Natives/DecVectors"
    import "OLC2_Project1/server/interpreter/AST/Natives/DecStructs"
    import "OLC2_Project1/server/interpreter/AST/Natives/DecObjects"
    import "OLC2_Project1/server/interpreter/AST/Expression/Access"
    import "OLC2_Project1/server/interpreter/AST/Expression/Objects"
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
    : listInits { $tree = AST.NewTree($listInits.l) }
;

listInits returns[*arrayList.List l]
    : listFuncs { $l = $listFuncs.l }
    ;

listFuncs returns[*arrayList.List l]
    @init {
        $l = arrayList.New()
    }
    : subList = listFuncs function {
        $subList.l.Add($function.instr)
        $l = $subList.l
    }
    | listArrays { $l = $listArrays.l }
    | function  { $l.Add($function.instr) }
    ;

listArrays returns[*arrayList.List l]
    @init {
            $l = arrayList.New()
    }
    : subList = listArrays dec_arr {
        $subList.l.Add($dec_arr.instr)
        $l = $subList.l
    }
    | listStructs { $l = $listStructs.l }
    | dec_arr  { $l.Add($dec_arr.instr) }
    ;

listStructs returns [*arrayList.List l]
    @init {
        $l = arrayList.New()
    }
    : subList = listStructs dec_struct {
        $subList.l.Add($dec_struct.instr)
        $l = $subList.l
    }
    | dec_struct  { $l.Add($dec_struct.instr) }
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

bloq_match returns [*arrayList.List content]
    : LEFT_KEY instructions RIGHT_KEY    { $content = $instructions.l }
    | instructions                       { $content = $instructions.l }
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
    | expr_rel COMMA?   { $instr = $expr_rel.instr }
    | expr_logic        { $instr = $expr_logic.instr }
    | dec_arr           { $instr = $dec_arr.instr }
    | dec_struct        { $instr = $dec_struct.instr }
    | dec_object        { $instr = $dec_object.instr }
    | vector_instr      { $instr = $vector_instr.instr }
    | transfer_prod     { $instr = $transfer_prod.instr }
    ;

print_prod returns [Abstract.Instruction instr]
    : PRINT ADMIRATION LEFT_PAR expression RIGHT_PAR (SEMICOLON|COMMA) { $instr = Natives.NewPrint($expression.p, false, $LEFT_PAR.line, localctx.(*Print_prodContext).Get_LEFT_PAR().GetColumn()) }
    | PRINT ADMIRATION LEFT_PAR opBefore = expression COMMA listVars RIGHT_PAR (SEMICOLON|COMMA) { $instr = Natives.NewPrintWithAfter($opBefore.p, $listVars.list, false, $LEFT_PAR.line, localctx.(*Print_prodContext).Get_LEFT_PAR().GetColumn()) }
    | PRINTLN ADMIRATION LEFT_PAR expression RIGHT_PAR (SEMICOLON|COMMA) { $instr = Natives.NewPrint($expression.p, true, $LEFT_PAR.line, localctx.(*Print_prodContext).Get_LEFT_PAR().GetColumn()) }
    | PRINTLN ADMIRATION LEFT_PAR opBefore = expression COMMA listVars RIGHT_PAR (SEMICOLON|COMMA) { $instr = Natives.NewPrintWithAfter($opBefore.p, $listVars.list, true, $LEFT_PAR.line, localctx.(*Print_prodContext).Get_LEFT_PAR().GetColumn()) }
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
    : DECLARAR MUT? ids_type EQUAL expression (SEMICOLON|COMMA)? {
        if ($MUT.text != "") {
            $instr = Natives.NewDeclarationInit($ids_type.list, $expression.p, true, "")
        } else {
            $instr = Natives.NewDeclarationInit($ids_type.list, $expression.p, false, "")
        }
    }
    | DECLARAR MUT? ids_type COLON data_type EQUAL expression (SEMICOLON|COMMA)? {
        if ($MUT.text != "") {
            $instr = Natives.NewDeclarationInit($ids_type.list, $expression.p, true, $data_type.data)
        } else {
            $instr = Natives.NewDeclarationInit($ids_type.list, $expression.p, false, $data_type.data)
        }
    }
    | DECLARAR MUT? ids_type (SEMICOLON|COMMA)? {
        if ($MUT.text != "") {
            $instr = Natives.NewDeclaration($ids_type.list, true, "")
        } else {
            $instr = Natives.NewDeclaration($ids_type.list, false, "")
        }
    }
    | DECLARAR MUT? ids_type COLON data_type (SEMICOLON|COMMA)? {
        if ($MUT.text != "") {
            $instr = Natives.NewDeclaration($ids_type.list, true, $data_type.data)
        } else {
            $instr = Natives.NewDeclaration($ids_type.list, false, $data_type.data)
        }
    }
    ;

assign_prod returns [Abstract.Instruction instr]
    : listIds EQUAL expression (SEMICOLON|COMMA)? {
        $instr = Natives.NewAssign($listIds.list, $expression.p)
    }
    | ID listInArray EQUAL expression (SEMICOLON|COMMA)? {
        $instr = DecArrays.NewAssignArray($ID.text, $listInArray.l, $expression.p, $EQUAL.line, localctx.(*Assign_prodContext).Get_EQUAL().GetColumn())
    }
    | access_object EQUAL expression SEMICOLON {
        $instr = DecObjects.NewAssignObject($access_object.p, $expression.p)
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
    : if_prod {
        $instr = $if_prod.instr
        $p = $if_prod.p
    }
    | match_prod {
        $instr = $match_prod.instr
        $p = $match_prod.p
    }
    ;

if_prod returns [Abstract.Instruction instr, Abstract.Expression p]
    : RIF expression bloq {
        $instr = Natives.NewIf($expression.p, $bloq.content, nil, nil, $RIF.line, localctx.(*If_prodContext).Get_RIF().GetColumn())
        $p = Natives.NewIf($expression.p, $bloq.content, nil, nil, $RIF.line, localctx.(*If_prodContext).Get_RIF().GetColumn())
    }
    | RIF expression bif=bloq RELSE belse=bloq {
        $instr = Natives.NewIf($expression.p, $bif.content, nil, $belse.content, $RIF.line, localctx.(*If_prodContext).Get_RIF().GetColumn())
        $p = Natives.NewIf($expression.p, $bif.content, nil, $belse.content, $RIF.line, localctx.(*If_prodContext).Get_RIF().GetColumn())
    }
    | RIF expression bif=bloq list_else_if RELSE belse=bloq {
        $instr = Natives.NewIf($expression.p, $bif.content, $list_else_if.list, $belse.content, $RIF.line, localctx.(*If_prodContext).Get_RIF().GetColumn() )
        $p = Natives.NewIf($expression.p, $bif.content, $list_else_if.list, $belse.content, $RIF.line, localctx.(*If_prodContext).Get_RIF().GetColumn() )
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

match_prod returns [Abstract.Instruction instr, Abstract.Expression p]
    : RMATCH expression LEFT_KEY list_instr_match RIGHT_KEY {
        $instr = Natives.NewMatch($expression.p, $list_instr_match.list)
        $p = Natives.NewMatch($expression.p, $list_instr_match.list)
    }
    ;

list_instr_match returns [*arrayList.List list]
    @init {
        $list = arrayList.New()
    }
    : sub = list_instr_match instr_match {
        $sub.list.Add($instr_match.instr)
        $list = $sub.list
     }
    | instr_match { $list.Add($instr_match.instr) }
    ;

instr_match returns [Abstract.Instruction instr]
    : expr_match EQUAL_ARROW bloq_match {
        $instr = Natives.NewCase($expr_match.list, $bloq_match.content, $EQUAL_ARROW.line, localctx.(*Instr_matchContext).Get_EQUAL_ARROW().GetColumn())
    }
    ;

expr_match returns [*arrayList.List list]
    @init {
        $list = arrayList.New()
    }
    : sub = expr_match PIPE expression {
            $sub.list.Add($expression.p)
            $list = $sub.list
    }
    | expression {
            $list.Add($expression.p)
    }
    ;

bucle_prod returns [Abstract.Instruction instr]
    : while_prod { $instr = $while_prod.instr }
    | loop_prod  { $instr = $loop_prod.instr }
    | forin_prod { $instr = $forin_prod.instr }
    ;

while_prod returns[Abstract.Instruction instr]
    : RWHILE expression bloq { $instr = Natives.NewWhile($expression.p, $bloq.content, $RWHILE.line, localctx.(*While_prodContext).Get_RWHILE().GetColumn()) }
    ;

loop_prod returns[Abstract.Instruction instr, Abstract.Expression p]
    : RLOOP bloq {
        $instr = Natives.NewLoop($bloq.content)
        $p = Natives.NewLoop($bloq.content)
    }
    ;

forin_prod returns[Abstract.Instruction instr]
    : RFOR expression RIN range_prod bloq { $instr = BucleForIn.NewForIn($expression.p, $range_prod.p, $bloq.content) }
    | RFOR expression RIN arraydata bloq { $instr = BucleForIn.NewForIn($expression.p, $arraydata.p, $bloq.content) }
    ;

range_prod returns[Abstract.Expression p]
    : e1 = expression RANGE e2 = expression {
        $p = Expression.NewRange($RANGE.line, localctx.(*Range_prodContext).Get_RANGE().GetColumn(), $e1.p, $e2.p)
     }
    ;

called_func returns [Abstract.Instruction instr, Abstract.Expression p]
    : ID LEFT_PAR RIGHT_PAR {
        $instr = ExpressionSpecial.NewCallFunction($ID.text, arrayList.New(), $ID.line, localctx.(*Called_funcContext).Get_ID().GetColumn())
        $p = ExpressionSpecial.NewCallFunction($ID.text, arrayList.New(), $ID.line, localctx.(*Called_funcContext).Get_ID().GetColumn())
    }
    | ID LEFT_PAR listExpressions RIGHT_PAR {
        $instr = ExpressionSpecial.NewCallFunction($ID.text, $listExpressions.l, $ID.line, localctx.(*Called_funcContext).Get_ID().GetColumn())
        $p = ExpressionSpecial.NewCallFunction($ID.text, $listExpressions.l, $ID.line, localctx.(*Called_funcContext).Get_ID().GetColumn())
    }
    ;

listExpressions returns [*arrayList.List l]
    @init {
        $l = arrayList.New()
    }
    : List = listExpressions (COMMA) expression {
        $List.l.Add($expression.p)
        $l = $List.l
    }
    | expression { $l.Add($expression.p) }
    ;

dec_arr returns [Abstract.Instruction instr]
    : DECLARAR MUT? ID (COLON listDim)? EQUAL expression (SEMICOLON|COMMA)
    {
        if $MUT.text != "" {
            $instr = DecArrays.NewDecArray($listDim.length, $ID.text, $expression.p, $listDim.data, true, $listDim.pos, $ID.line, localctx.(*Dec_arrContext).Get_ID().GetColumn())
        } else {
            $instr = DecArrays.NewDecArray($listDim.length, $ID.text, $expression.p, $listDim.data, false, $listDim.pos, $ID.line, localctx.(*Dec_arrContext).Get_ID().GetColumn())
        }
    }
    ;

listDim returns[int length, SymbolTable.DataType data, Abstract.Expression pos]
    @init { $length = 0 }
    : LEFT_BRACKET l = listDim SEMICOLON expression RIGHT_BRACKET { $length = $l.length + 1
                                                                    $data = $l.data
                                                                    $pos = $expression.p
                                                               }
    | LEFT_BRACKET data_type SEMICOLON expression RIGHT_BRACKET {
        $length = 1
        $pos = $expression.p
        switch ($data_type.data) {
            case "i64":
                $data = SymbolTable.INTEGER
            case "f64":
                $data = SymbolTable.FLOAT
            case "&str":
                $data = SymbolTable.STR
            case "String":
                $data = SymbolTable.STRING
            case "bool":
                $data = SymbolTable.BOOLEAN
        }
    }
    ;

vector_instr returns[Abstract.Instruction instr, Abstract.Expression p]
    : dec_vector {
        $instr = $dec_vector.instr
    }
    | natives_vector {
        $instr = $natives_vector.instr
        $p = $natives_vector.p
    }
    ;

dec_vector returns[Abstract.Instruction instr]
    : DECLARAR MUT? ID COLON RVECMayus LESS_THAN data_type GREATER_THAN EQUAL REVECTORNEW SEMICOLON {
        var data SymbolTable.DataType
        switch ($data_type.data) {
            case "i64":
                data = SymbolTable.INTEGER
            case "f64":
                data = SymbolTable.FLOAT
            case "&str":
                data = SymbolTable.STR
            case "String":
                data = SymbolTable.STRING
            case "bool":
                data = SymbolTable.BOOLEAN
        }

        if $MUT.text != "" {
            $instr = DecVectors.NewDecVector(1, nil, $ID.text, nil, data, true, false)
        } else {
            $instr = DecVectors.NewDecVector(1, nil, $ID.text, nil, data, false, false)
        }
    }
    | DECLARAR MUT? ID EQUAL expr_vector (SEMICOLON|COMMA) {
        if $MUT.text != "" {
            $instr = DecVectors.NewDecVector(1, nil, $ID.text, $expr_vector.p, SymbolTable.NULL, true, false)
        } else {
            $instr = DecVectors.NewDecVector(1, nil, $ID.text, $expr_vector.p, SymbolTable.NULL, false, false)
        }
    }
    | DECLARAR MUT? ID COLON RVECMayus LESS_THAN data_type GREATER_THAN EQUAL REVECCAPACITY LEFT_PAR expression RIGHT_PAR SEMICOLON {
        var data SymbolTable.DataType
        switch ($data_type.data) {
            case "i64":
                data = SymbolTable.INTEGER
            case "f64":
                data = SymbolTable.FLOAT
            case "&str":
                data = SymbolTable.STR
            case "String":
                data = SymbolTable.STRING
            case "bool":
                data = SymbolTable.BOOLEAN
        }

        if $MUT.text != "" {
            $instr = DecVectors.NewDecVector(0, $expression.p, $ID.text, nil, data, true, true)
        } else {
            $instr = DecVectors.NewDecVector(0, $expression.p, $ID.text, nil, data, false, true)
        }
    }
    | DECLARAR MUT? ID EQUAL RVEC ADMIRATION LEFT_BRACKET e1=expression SEMICOLON e2=expression RIGHT_BRACKET SEMICOLON {
        if $MUT.text != "" {
            $instr = DecVectors.NewDecVector(0, $e2.p, $ID.text, $e1.p, SymbolTable.NULL, true, false)
        } else {
            $instr = DecVectors.NewDecVector(0, $e2.p, $ID.text, $e1.p, SymbolTable.NULL, false, false)
        }
    }
    ;

expr_vector returns [Abstract.Expression p]
    : RVEC ADMIRATION expression {
        $p = $expression.p
    }
    ;

natives_vector returns[Abstract.Instruction instr, Abstract.Expression p]
    : ID DOT op=ID LEFT_PAR expression RIGHT_PAR SEMICOLON? {
        $instr = DecVectors.NewOperation($ID.text, $expression.p, $op.text, nil)
        $p = DecVectors.NewOperation($ID.text, $expression.p, $op.text, nil)
    }
    | ID DOT op=ID LEFT_PAR RIGHT_PAR SEMICOLON? {
        $instr = DecVectors.NewOperation($ID.text, nil, $op.text, nil)
        $p = DecVectors.NewOperation($ID.text, nil, $op.text, nil)
    }
    | ID DOT op=ID LEFT_PAR e1=expression COMMA e2=expression RIGHT_PAR SEMICOLON? {
        $instr = DecVectors.NewOperation($ID.text, $e2.p, $op.text, $e1.p)
        $p = DecVectors.NewOperation($ID.text, $e2.p, $op.text, $e1.p)
    }
    | ID DOT op=ID LEFT_PAR REFERENCE expression RIGHT_PAR SEMICOLON? {
        $instr = DecVectors.NewOperation($ID.text, $expression.p, $op.text, nil)
        $p = DecVectors.NewOperation($ID.text, $expression.p, $op.text, nil)
    }
    ;

transfer_prod returns[Abstract.Instruction instr]
    : break_instr                { $instr = $break_instr.instr }
    | continue_instr             { $instr = $continue_instr.instr }
    | return_instr               { $instr = $return_instr.instr }
    ;

break_instr returns[Abstract.Instruction instr]
    : RBREAK SEMICOLON { $instr = Natives.NewBreak($RBREAK.line, localctx.(*Break_instrContext).Get_RBREAK().GetColumn(), nil) }
    | RBREAK expression SEMICOLON { $instr = Natives.NewBreak($RBREAK.line, localctx.(*Break_instrContext).Get_RBREAK().GetColumn(), $expression.p) }
    ;

continue_instr returns[Abstract.Instruction instr]
    : RCONTINUE SEMICOLON { $instr = Natives.NewContinue($RCONTINUE.line, localctx.(*Continue_instrContext).Get_RCONTINUE().GetColumn()) }
    ;

return_instr returns[Abstract.Instruction instr]
    : RRETURN SEMICOLON { $instr = Natives.NewReturn($RRETURN.line, localctx.(*Return_instrContext).Get_RRETURN().GetColumn(), nil) }
    | RRETURN expression SEMICOLON? { $instr = Natives.NewReturn($RRETURN.line, localctx.(*Return_instrContext).Get_RRETURN().GetColumn(), $expression.p) }
    ;

dec_struct returns [Abstract.Instruction instr]
    : RSTRUCT ID bloq_struct   {
        $instr = DecStructs.NewDecStruct($ID.text, $bloq_struct.l)
    }
    ;

bloq_struct returns [*arrayList.List l]
    : LEFT_KEY content_struct RIGHT_KEY {
        $l = $content_struct.l
    }
    ;

content_struct returns [*arrayList.List l]
    @init {
        $l = arrayList.New()
    }
    : subList = content_struct item_struct {
        $subList.l.Add($item_struct.p)
        $l = $subList.l
    }
    | item_struct {
        $l.Add($item_struct.p)
     }
    ;

item_struct returns [Abstract.Expression p]
    : ID COLON data_type COMMA? {
        var data SymbolTable.DataType
        switch ($data_type.data) {
            case "i64":
                data = SymbolTable.INTEGER
            case "f64":
                data = SymbolTable.FLOAT
            case "&str":
                data = SymbolTable.STR
            case "String":
                data = SymbolTable.STRING
            case "bool":
                data = SymbolTable.BOOLEAN
        }

        $p = Expression.NewItemStruct($ID.text, data, $ID.line, localctx.(*Item_structContext).Get_ID().GetColumn())
    }
    ;

dec_object returns [Abstract.Instruction instr]
    : DECLARAR MUT? ID EQUAL ob=ID LEFT_KEY def_items RIGHT_KEY {
        $instr = DecObjects.NewDecObjects($ID.text, $ob.text, $def_items.l)
     }
    ;

expression returns [Abstract.Expression p]
    : conditional_prod           { $p = $conditional_prod.p }
    | loop_prod                  { $p = $loop_prod.p }
    | expr_logic                 { $p = $expr_logic.p }
    | expr_rel                   { $p = $expr_rel.p }
    | arraydata                  { $p = $arraydata.p }
    | natives_vector             { $p = $natives_vector.p }
    | type_struct                { $p = $type_struct.p }
    | access_object              { $p = $access_object.p }
;

arraydata returns [Abstract.Expression p]
    : LEFT_BRACKET listExpressions RIGHT_BRACKET {
        $p = ExpressionSpecial.NewValueArray($listExpressions.l, nil, nil)
    }
    | LEFT_BRACKET e1=expression SEMICOLON e2=expression RIGHT_BRACKET {
        $p = ExpressionSpecial.NewValueArray(nil, $e1.p, $e2.p)
    }
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

access_vector returns[Abstract.Expression p, Abstract.Instruction instr]
    : ID listInVector {
        $p = Access.NewAccessVector($ID.text, $listInVector.l)
        $instr = Access.NewAccessVector($ID.text, $listInVector.l)
    }
    ;

listInVector returns [*arrayList.List l]
    @init {
            $l = arrayList.New()
    }
    : sublist = listInVector inVector {
        $sublist.l.Add($inVector.p)
        $l = $sublist.l
    }
    | inVector { $l.Add($inVector.p) }
    ;

inVector returns [Abstract.Expression p]
    : LEFT_BRACKET expression RIGHT_BRACKET     { $p = $expression.p }
    ;

type_struct returns [Abstract.Expression p]
    : ID LEFT_KEY def_items RIGHT_KEY SEMICOLON { $p = Objects.NewObject($ID.text, $def_items.l) }
    ;

def_items returns[*arrayList.List l]
    @init{
        $l = arrayList.New()
    }
    : subList = def_items item {
        $subList.l.Add($item.p)
        $l = $subList.l
    }
    | item { $l.Add($item.p) }
    ;

item returns[Abstract.Expression p]
    : ID COLON expression COMMA? { $p = Objects.NewAttribute($ID.text, $expression.p); }
    ;

access_object returns [Abstract.Expression p]
    : listAccess { $p = Access.NewObjectAccess($listAccess.l) }
    ;

listAccess returns [*arrayList.List l]
    @init {
        $l = arrayList.New()
    }
    : subList = listAccess DOT access {
        $subList.l.Add($access.p)
        $l = $subList.l
    }
    | access { $l.Add($access.p) }
    ;

access returns [Abstract.Expression p]
    : ID { $p = Expression.NewIdentifier($ID.text, $ID.line, localctx.(*AccessContext).Get_ID().GetColumn()) }
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
    | access_vector {
        $p = $access_vector.p
        $instr = $access_vector.instr
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