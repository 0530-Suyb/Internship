package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 解释器模式
// 解释器模式是一种设计模式，用于定义一种语言的文法，并提供一个解释器来解释该语言中的句子。
// 这种模式通常用于编写脚本语言、查询语言等场景。
// 解释器模式的核心思想是将一个复杂的表达式分解为简单的表达式，并通过递归的方式来解释这些表达式。
// 解释器模式通常包含以下几个角色：
// 1. 抽象表达式（Expression）：定义一个抽象的解释操作，所有的具体表达式都需要实现这个接口。
// 2. 终结符表达式（TerminalExpression）：实现了抽象表达式接口，表示文法中的终结符。
// 3. 非终结符表达式（NonTerminalExpression）：实现了抽象表达式接口，表示文法中的非终结符。
// 4. 上下文（Context）：包含解释器所需的全局信息，通常是一个数据结构，用于存储变量、常量等信息。
// 5. 客户端（Client）：使用解释器来解释一个句子，通常是一个字符串。

// 例子：
// 假设我们要解释一个简单的数学表达式，例如 "3 + 5 - 2"。
// 我们可以使用解释器模式来实现这个功能。

// 抽象表达式接口
type Expression interface {
	Interpret() int
}

// 终结符表达式：数字
type NumberExpression struct {
	value int
}

func (n *NumberExpression) Interpret() int {
	return n.value
}

// 非终结符表达式：加法
type AddExpression struct {
	left, right Expression
}

func (a *AddExpression) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}

// 非终结符表达式：减法
type SubtractExpression struct {
	left, right Expression
}

func (s *SubtractExpression) Interpret() int {
	return s.left.Interpret() - s.right.Interpret()
}

// 上下文：解析输入表达式
func parseExpression(tokens []string) Expression {
	if len(tokens) == 1 {
		// 只有一个数字
		value, _ := strconv.Atoi(tokens[0])
		return &NumberExpression{value: value}
	}

	// 找到最后一个运算符
	for i := len(tokens) - 1; i >= 0; i-- {
		if tokens[i] == "+" || tokens[i] == "-" {
			left := parseExpression(tokens[:i])
			right := parseExpression(tokens[i+1:])
			if tokens[i] == "+" {
				return &AddExpression{left: left, right: right}
			} else {
				return &SubtractExpression{left: left, right: right}
			}
		}
	}

	return nil
}

func main() {
	// 输入表达式
	input := "3 + 5 - 2"
	tokens := strings.Split(input, " ")

	// 解析表达式
	expression := parseExpression(tokens)

	// 解释表达式
	result := expression.Interpret()

	fmt.Printf("表达式 '%s' 的结果是: %d\n", input, result)
}
