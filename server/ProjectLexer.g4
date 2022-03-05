// CalcLexer.g4
lexer grammar ProjectLexer;

// Reserve words
PRINTLN:   'println';
PRINT:     'print';

// Variables
DECLARAR:  'let';
MUT:       'mut';

// Data Type of Vars
RSTRING:    'String';
RINTEGER:   'i64';
RREAL:      'f64';
RBOOLEAN:   'bool';
RCHAR:      'char';
RSTR:       '&str';
RAS:        'as';

// Conditional Structures
RIF:        'if';
RELSE:      'else';

// Bucle Structures
RWHILE:     'while';

// Mathematic Operators
POWI:                   'pow';
POWF:                   'powf';

// Text Strings
TOSTRING:                'to_string';
TOOWNED:                 'to_owned';

// Comments
MULTILINE:      '/*' (~[/])+ '*/' -> skip;
INLINE:         '//'(~[\n])+ -> skip;

// Regular expressions
INTEGER:                [0-9]+;
FLOAT:                  [0-9]+'.'[0-9]+;
CHAR:                   '\''~["]'\'';
STRING:                 '"'~["]*'"';
BOOLEAN:                ('true' | 'false');
ID:                     [a-zñA-ZÑ_][a-zñA-ZÑ0-9]*;

// General Operators
EQUAL:                  '=';
DOT:                    '.';
SEMICOLON:              ';';
COMMA:                  ',';
COLON:                  ':';
ADMIRATION:             '!';
REFERENCE:              '&';
HERITAGE:               '::';

// Grouping Operators
LEFT_PAR:               '(';
RIGHT_PAR:              ')';
LEFT_KEY:               '{';
RIGHT_KEY:              '}';
LEFT_BRACKET:           '[';
RIGHT_BRACKET:          ']';

// Relational Operators
EQUEAL_EQUAL:           '==';
NOTEQUAL:               '!=';
GREATER_THAN:           '>';
LESS_THAN:              '<';
GREATER_EQUALTHAN:      '>=';
LESS_EQUEALTHAN:        '<=';

// Aritmethic Operators
MUL:                    '*';
DIV:                    '/';
MOD:                    '%';
ADD:                    '+';
SUB:                    '-';

// Logic Operators
AND:                    '&&';
OR:                     '||';

// White spaces
WHITESPACE: [ \r\n\t]+ -> skip;

// Escape characters
fragment
ESC_SEQ
    : '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ');
