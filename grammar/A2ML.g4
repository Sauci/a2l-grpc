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
   | name = 'double'
   | name = 'float'
   )
   ;

/* chapter 5.2: block_definition "block" tag type_name. A tag is "a character sequence enclosed
   within double inverted commas", i.e. the STRING token of the A2L lexer; the third alternative
   is not part of the metalanguage, it is owed to Vector Informatik */
blockDefinition:
   'block' tag = stringValue tn = a2mlTypeName
   | 'block' tag = stringValue '(' mem = member ')' (mult = '*')?
   ;

enumTypeName:
    ('enum' identifier = a2mlIdentifier? '{' enumerators = enumeratorList '}' )
    | ('enum' identifier = a2mlIdentifier)
    ;

enumeratorList:
   ids += enumerator (',' ids += enumerator )*
   ;

enumerator:
   keyword = stringValue ('=' constant = numericValue)?
   ;

structTypeName:
      'struct' identifier = a2mlIdentifier? '{' members += structMember* '}'
    | 'struct' identifier = a2mlIdentifier
    ;

structMember:
     m = member ';'
   | '(' mstar = member ')' (m0 = '*')? ';'
   ;

member:
    typeName = a2mlTypeName (dimension += a2mlArraySpecifier)*
    ;

/* chapter 5.2: array_specifier "[" constant "]". The rule carries its own name, because the
   arraySpecifier of the A2L grammar describes the index of a partial identifier (chapter 3.2) and
   also accepts an alphabetic string, which is not a constant */
a2mlArraySpecifier:
   '[' value = integerValue ']'
   ;

taggedStructTypeName:
   /* 'taggedstruct' t1 = identifierValue */
   'taggedstruct' identifier = a2mlIdentifier?
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

/* chapter 5.2: taggedstruct_definition tag [ member ] | tag "(" member ")*;". The tag identifies
   the element and is mandatory in both forms; the repeated form is listed first, so that its
   opening parenthesis is matched here rather than left to the enclosing member */
taggedStructDefinition:
     tag = stringValue star = '(' mem = member ')' '*'
   | tag = stringValue mem = member?
   ;

taggedUnionTypeName:
    (('taggedunion' identifier = a2mlIdentifier? '{' members += taggedUnionMember* '}')
    | ('taggedunion' identifier = a2mlIdentifier))
    ;

taggedUnionMember:
   (
     tag = stringValue  m = member?  ';'
   | block = blockDefinition ';'
   )
   ;

/* This grammar is imported by A2L.g4 and is never generated on its own, so it declares no token
   of its own: integerValue, numericValue, stringValue and identifierValue, and the tokens they
   are built from, are the ones of the importing grammar. Declaring them here again only hid the
   fact that the importing grammar overrides them, and with it the fact that its STRING rule
   precedes the TAG rule declared here, so that a tag was always matched as a STRING and every
   alternative expecting a TAG was dead. The tags above are therefore stringValue. */
