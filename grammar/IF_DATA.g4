grammar IF_DATA;

ifData:
   '/begin' 'IF_DATA'
        name = identifierValue
        (
            blob += genericParameter
        )*
   '/end' 'IF_DATA'
   ;

genericParameter:
    tag = tagValue
    | sting = stringValue
    | numeric = numericValue
    | generic = genericNode
    | identifier = identifierValue
    ;

genericNode:
    '/begin' name = identifierValue
    (element += genericParameter)*
    '/end' identifierValue
    ;

numericValue:
     i = INT
   | h = HEX
   | f = FLOAT
   ;

stringValue:
    s = STRING
    ;

tagValue:
    s = TAG
    ;

identifierValue:
    i = ID
    ;

/*
** Lexer
*/
INT : '0'..'9'+
    ;

HEX:   '0'('x' | 'X') ('a' .. 'f' | 'A' .. 'F' | '0' .. '9')+
    ;

FLOAT:
    ('+' | '-')?
    (
        ('0'..'9')+ '.' ('0'..'9')* EXPONENT?
    |   '.' ('0'..'9')+ EXPONENT?
    |   ('0'..'9')+ EXPONENT
    )
    ;

ID  : /* ('a'..'z'|'A'..'Z'|'_') */
    ('a'..'z'|'A'..'Z'|'0'..'9'|'_')+
    ;

TAG:  '"' ID '"'  // s. 3.2
   ;

COMMENT
    :   ('//' ~('\n'|'\r')* '\r'? '\n'
    |   '/*' .*? '*/')
        -> channel(HIDDEN)
    ;

WS  :   (' ' | '\t' | '\r' | '\n') -> skip
    ;

STRING
    :  '"' ( ESC_SEQ | ~('\\'|'"') )* '"'
    ;

fragment
EXPONENT : ('e'|'E') ('+'|'-')? ('0'..'9')+ ;

fragment
HEX_DIGIT : ('0'..'9'|'a'..'f'|'A'..'F') ;

fragment
ESC_SEQ
    :   '\\' ('b'|'t'|'n'|'f'|'r'|'\u0022'|'\''|'\\')
    |   OCTAL_ESC
    ;

fragment
OCTAL_ESC
    :   '\\' ('0'..'3') ('0'..'7') ('0'..'7')
    |   '\\' ('0'..'7') ('0'..'7')
    |   '\\' ('0'..'7')
    ;
