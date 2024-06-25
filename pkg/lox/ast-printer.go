package lox

import (
	"fmt"
)

func PrintAST(expr Expr) {
	printer := AstPrinter{}
	fmt.Println(printer.Print(expr))
}

type AstPrinter struct {
	depth int
}

func (printer *AstPrinter) padding() string {
	return fmt.Sprintf("%*s", printer.depth*2, "")
}

func (printer *AstPrinter) Print(expr Expr) string {
	return expr.Accept(printer).(string)
}

func (printer *AstPrinter) VisitBinaryExpr(expr *ExprBinary) interface{} {
	out := printer.padding() + "Binary{\n"
	printer.depth++
	out += expr.Left.Accept(printer).(string) + "\n"
	out += printer.padding() + expr.Operator.Lexeme + "\n"
	out += expr.Right.Accept(printer).(string) + "\n"
	printer.depth--
	out += printer.padding() + "}"
	return out
}

func (printer *AstPrinter) VisitGroupingExpr(expr *ExprGrouping) interface{} {
	out := printer.padding() + "Group{\n"
	printer.depth++
	out += expr.Expr.Accept(printer).(string) + "\n"
	printer.depth--
	out += printer.padding() + "}"
	return out
}

func (printer *AstPrinter) VisitLiteralExpr(expr *ExprLiteral) interface{} {
	return printer.padding() + fmt.Sprintf("Literal{%v}", expr.Value)
}

func (printer *AstPrinter) VisitUnaryExpr(expr *ExprUnary) interface{} {
	out := printer.padding() + "Unary{\n"
	printer.depth++
	out += printer.padding() + expr.Operator.Lexeme + "\n"
	out += expr.Right.Accept(printer).(string) + "\n"
	printer.depth--
	out += printer.padding() + "}"
	return out
}
