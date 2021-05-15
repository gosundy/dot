// Dot.g4
grammar Dot;
MSGCONTENT: 'msg.content';
IF       : 'if';
INT:[0-9]+;
STRING: '"' (~["\r\n] | '""')* '"';
WORDS:[0-9A-Za-z]+;
WS: (' '|'\t'|'\r'|'\n')+ -> skip;
MUL:'*';
DIV:'/';
ADD:'+';
SUB:'-';
Equals   : '==';
NEquals  : '!=';
GTEquals : '>=';
LTEquals : '<=';
Excl     : '!';
GT       : '>';
LT       : '<';
And      : '&&';
Or       : '||';

NL:'\n';

// Rules
start
    :statment*
    ;

statment
    :ifStatment block #StatmentIf
    |exp #StatmentExp
    ;

ifStatment
   :IF ifPartStatment
   ;

ifPartStatment
    :'('exp')'
    ;
exp
    :msgContentValue op=(DIV|MUL) exp #ContentMulDiv
    |msgContentValue op=(SUB|ADD) exp #ContentAddSub
    |msgContentValue #GetContentValue
    |'!' exp # Not
    |msgContentValue op=(Equals|NEquals|GTEquals|LTEquals|GT|LT) exp #Compare
    |msgContentValue op=(And|Or) exp #AndOr
    |INT #GetInt
    |STRING #GetString
    ;

block
    : '{' exp '}'
    ;

msgContentValue
    : msgContentJsonPath'.'msgValueType
    ;

msgContentJsonPath
    :MSGCONTENT'.'jsonPath
    ;
jsonPath
    : WORDS'.'jsonPath
    | WORDS
    ;

msgValueType
    : 'Number'
    | 'String'
    ;