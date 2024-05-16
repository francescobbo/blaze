grammar Calculator;

root: expression;

expression
    : lhs=expression op=(POW | PPOW) rhs=expression
    | lhs=expression op=(MUL | DIV | MOD) rhs=expression
    | lhs=expression op=PCT rhss=factor // Modulo, disambiguate from percentage
    | lhs=expression op=(ADD | SUB) rhs=expression
    | pct=percentage 'of' rhs=expression
    | uop=(ADD | SUB)* unary=factor
    | EOF;

factor: primary unit?;

primary
    : NUMBER (constant | fn)?
    | fn
    | constant
    | percentage
    | currency
    | LPAREN sub=expression RPAREN;

fn: FUNC_NAME factor;
constant: CONSTANT;
percentage: NUMBER '%';
currency: CURRENCY_SYMBOL NUMBER;

unit
    : unit POW? NUMBER
    | unit MUL unit
    | unit DIV unit
    | LPAREN unit RPAREN
    | UNIT_NAME;

FUNC_NAME
    : 'sqrt'
    | 'log'
    | 'ln'
    | 'sin'
    | 'cos'
    | 'tan'
    ;

CONSTANT: 'pi' | 'π' | 'phi' | 'φ' | 'e';
CURRENCY_SYMBOL: '$' | '€' | '£';

ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/';
POW: '^';
PPOW: '**';
MOD: 'mod';
LPAREN: '(';
RPAREN: ')';
PCT: '%';

NUMBER: [0-9.,]+;
UNIT_NAME: [a-z]+;

WS: [ \t\r\n]+ -> skip;