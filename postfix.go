package postfix

import (
	"errors"
	"strconv"

	"github.com/darylnwk/mathtoken"
	"github.com/darylnwk/stack"
)

// Postfix defines a list of tokens
type Postfix []mathtoken.Token

// ParseInfix parses a mathematical expression in infix notation to postfix notation (or RPN)
// using Shunting-yard algorithm.
// See https://en.wikipedia.org/wiki/Shunting-yard_algorithm#The_algorithm_in_detail for more details.
func ParseInfix(infix string) (postfix Postfix, err error) {
	operators := stack.New()
	tokens, err := mathtoken.Parse(infix)
	if err != nil {
		return postfix, errors.New("postfix: invalid mathematical expression")
	}

	for _, token := range tokens {
		switch token.Type {
		case mathtoken.TypeConstant, mathtoken.TypeVariable:
			postfix = append(postfix, token)
		case mathtoken.TypeOperator:
			for operator := operators.Peek(); operator != nil &&
				((operator.(mathtoken.Token).Precedence > token.Precedence) ||
					(operator.(mathtoken.Token).Precedence == token.Precedence && token.Associativity == mathtoken.AssociativityLeft)) &&
				operator.(mathtoken.Token).Type != mathtoken.TypeLParent; operator = operators.Peek() {
				postfix = append(postfix, operators.Pop().(mathtoken.Token))
			}

			operators.Push(token)
		case mathtoken.TypeLParent:
			operators.Push(token)
		case mathtoken.TypeRParent:
			for operators.Peek().(mathtoken.Token).Type != mathtoken.TypeLParent {
				postfix = append(postfix, operators.Pop().(mathtoken.Token))
			}

			operators.Pop()
		}
	}

	for {
		if operators.Size() == 0 {
			break
		}

		postfix = append(postfix, operators.Pop().(mathtoken.Token))
	}

	return postfix, nil
}

// Evaluate takes `Postfix` and a map of variables and returns the result of the expression
// See https://en.wikipedia.org/wiki/Reverse_Polish_notation#Postfix_evaluation_algorithm for more details.
func Evaluate(postfix Postfix, variables map[string]float64) float64 {
	result := stack.New()

	for _, token := range postfix {
		switch token.Type {
		case mathtoken.TypeConstant:
			value, _ := strconv.ParseFloat(token.Value, 64)
			result.Push(value)
		case mathtoken.TypeVariable:
			result.Push(variables[token.Value])
		case mathtoken.TypeOperator:
			value1 := result.Pop().(float64)
			value2 := result.Pop().(float64)

			switch token.Value {
			case "/":
				result.Push(value1 / value2)
			case "*":
				result.Push(value1 * value2)
			case "+":
				result.Push(value1 + value2)
			case "-":
				result.Push(value1 - value2)
			}

		}
	}

	return result.Pop().(float64)
}
