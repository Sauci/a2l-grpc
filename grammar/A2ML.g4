grammar A2ML;

a2ml:
   '/begin' 'A2ML'
        (d += declaration)*
   '/end' 'A2ML'
   ;

declaration:
   ( t = typeDefinition
   | b = blockDefinition) ';'
   ;

typeDefinition:
   a2mlTypeName
   ;

a2mlTypeName:
   (
     pr = predefinedTypeName
   | st = structTypeName
   | ts = taggedStructTypeName
   | tu = taggedUnionTypeName
   | en = enumTypeName
   )
   ;

predefinedTypeName:
   (
     name = 'char'
   | name = 'int'
   | name = 'long'
   | name = 'uchar'
   | name = 'uint'
   | name = 'ulong'
   | name = 'int64'
   | name = 'uint64'
   | name = 'double'
   | name = 'float'
   )
   ;

blockDefinition:
   'block' tag0 = tagValue tn = a2mlTypeName
   | 'block' tag1 = stringValue tn = a2mlTypeName
   | 'block' tag1 = stringValue '(' mem = member ')' (mult = '*')? /* Owed to Vector Informatik... */
   ;

enumTypeName:
    ('enum' identifier = identifierValue? '{' enumerators = enumeratorList '}' )
    | ('enum' identifier = identifierValue)
    ;

enumeratorList:
   ids += enumerator (',' ids += enumerator )*
   ;

enumerator:
   tag0 = tagValue ('=' constant = numericValue)? |
   tag1 = stringValue ('=' constant = numericValue)?
   ;

structTypeName:
      'struct' identifier = identifierValue? '{' members += structMember* '}'
    | 'struct' identifier = identifierValue
    ;

structMember:
     m = member ';'
   | '(' mstar = member ')' (m0 = '*')? ';'
   ;

member:
    typeName = a2mlTypeName (dimension += arraySpecifier)*
    ;

arraySpecifier:
   '[' value = numericValue ']'
   ;

taggedStructTypeName:
   /* 'taggedstruct' t1 = identifierValue */
   'taggedstruct' identifier = identifierValue?
   (
        '{' (members += taggedStructMember)* '}' | (members += taggedStructMember)*
   )
   ;

taggedStructMember:
      ('(' ts0 = taggedStructDefinition ';'? ')' '*' ';')
    | ('(' bl0 = blockDefinition ';'? ')' '*' ';')
    | (ts1 = taggedStructDefinition ';')
    | (bl1 = blockDefinition ';')
   ;

taggedStructDefinition:
     tag0 = tagValue? mem = member?
   | tag1 = tagValue? '(' mem = member ')' '*'
   | tag2 = stringValue? mem = member?
   ;

taggedUnionTypeName:
    (('taggedunion' identifier = identifierValue? '{' members += taggedUnionMember* '}')
    | ('taggedunion' identifier = identifierValue))
    ;

taggedUnionMember:
   (
     tag0 = tagValue  m = member?  ';'
   | tag1 = stringValue  m = member?  ';'
   | block = blockDefinition ';'
   )
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
