grammar Expr;
prog: expr EOF;
expr:
	expr mulOp = (MULTIPLY | DIVIDE) expr	# MulDiv
	| expr addOp = (ADD | SUB) expr			# AddSub
	| INT									# Int
	| '(' expr ')'							# Parens;
INT: [0-9]+;
ADD: '+';
SUB: '-';
MULTIPLY: '*';
DIVIDE: '/';
WS: [ \n\t\r]+ -> skip;
