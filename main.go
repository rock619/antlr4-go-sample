package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/antlr4-go-sample/parsing"
)

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

type Listener struct {
	*parsing.BaseExprListener
	stack Stack
}

func NewListener() *Listener {
	return &Listener{}
}

func (l *Listener) ExitMulDiv(ctx *parsing.MulDivContext) {
	right, left := l.stack.Pop(), l.stack.Pop()
	switch ctx.GetMulOp().GetTokenType() {
	case parsing.ExprParserMULTIPLY:
		l.stack.Push(left * right)
	case parsing.ExprParserDIVIDE:
		l.stack.Push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", ctx.GetMulOp().GetText()))
	}
}

func (l *Listener) ExitAddSub(ctx *parsing.AddSubContext) {
	right, left := l.stack.Pop(), l.stack.Pop()
	switch ctx.GetAddOp().GetTokenType() {
	case parsing.ExprParserADD:
		l.stack.Push(left + right)
	case parsing.ExprParserSUB:
		l.stack.Push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", ctx.GetAddOp().GetText()))
	}
}

func (l *Listener) ExitInt(ctx *parsing.IntContext) {
	i, _ := strconv.Atoi(ctx.GetText())
	l.stack.Push(i)
}

func main() {
	input := antlr.NewInputStream(os.Args[1])
	lexer := parsing.NewExprLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parsing.NewExprParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Prog()
	l := NewListener()
	antlr.NewParseTreeWalker().Walk(l, tree)
	fmt.Println(l.stack.Pop())
}
