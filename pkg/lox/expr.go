package lox

type Expr interface {
	Accept(visitor ExprVisitor) interface{}
}

type ExprVisitor interface {
	VisitBinaryExpr(expr *ExprBinary) interface{}
	VisitGroupingExpr(expr *ExprGrouping) interface{}
	VisitLiteralExpr(expr *ExprLiteral) interface{}
	VisitUnaryExpr(expr *ExprUnary) interface{}
}

type ExprBinary struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (expr *ExprBinary) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitBinaryExpr(expr)
}

type ExprGrouping struct {
	Expr Expr
}

func (expr *ExprGrouping) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitGroupingExpr(expr)
}

type ExprLiteral struct {
	Value interface{}
}

func (expr *ExprLiteral) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitLiteralExpr(expr)
}

type ExprUnary struct {
	Operator Token
	Right    Expr
}

func (expr *ExprUnary) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitUnaryExpr(expr)
}
