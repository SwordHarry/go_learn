package main

import (
	"fmt"
	"math"
)

// 表示一个变量
type Var string

// 数字常量
type literal float64

// unary 表示一元操作符表达式，比如 -x
type unary struct {
	op rune // +, - 中的一个
	x  Expr
}

// binary 表示二元操作符表达式，如 x+y
type binary struct {
	op   rune
	x, y Expr
}

// call 表示函数调用表达式，比如 sin(x)
type call struct {
	fn   string
	args []Expr
}

type Env map[Var]float64

// Expr: 算数表达式
type Expr interface {
	// Eval 返回表达式在 env 上下文下的值
	Eval(env Env) float64
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsuppoerted unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsuppoerted binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
