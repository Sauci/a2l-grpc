grammar IF_DATA;

ifData:
   '/begin' 'IF_DATA'
        name = ifDataIdentifier
        (
            blob += genericParameter
        )*
   '/end' 'IF_DATA'
   ;

/* the content of an IF_DATA block is described by the A2ML of the module, so it is collected as a
   generic blob. A quoted element is matched as a stringValue: the A2L lexer has no separate token
   for a tag */
genericParameter:
    sting = stringValue
    | numeric = numericValue
    | generic = genericNode
    | identifier = ifDataIdentifier
    ;

genericNode:
    '/begin' name = ifDataIdentifier
    (element += genericParameter)*
    '/end' ifDataIdentifier
    ;

/* This grammar is imported by A2L.g4 and is never generated on its own, so it declares no token
   of its own: numericValue, stringValue and identifierValue, and the tokens they are built from,
   are the ones of the importing grammar. */
