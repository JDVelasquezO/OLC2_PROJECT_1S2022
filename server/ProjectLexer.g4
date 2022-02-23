// CalcLexer.g4
lexer grammar ProjectLexer;

// Reserve words
PRINTLN:   'println';
PRINT:     'print';

// Variables
DECLARAR:  'let';

// Data Type of Vars
RSTRING:    'String';
RINTEGER:   'Integer';
RREAL:      'Real';
RBOOLEAN:   'Boolean';
RSTR:       'str';

// Conditional Structures
RIF:        'If';
RELSE:      'entonces';

// Logical Operators
RAND:       'AND';
ROR:        'OR';

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
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
MOD: '%';

// White spaces
WHITESPACE: [ \r\n\t]+ -> skip;

// Escape characters
fragment
ESC_SEQ
    : '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ');
