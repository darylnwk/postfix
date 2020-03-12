package postfix_test

import (
	"testing"

	"github.com/darylnwk/postfix"
	"github.com/stretchr/testify/assert"
)

func TestPostfix_Parse(t *testing.T) {
	var (
		postfix, err = postfix.ParseInfix("1 +2 * 3 / 4")
	)

	assert.NoError(t, err)
	assert.Equal(t, "1", postfix[0].Value)
	assert.Equal(t, "2", postfix[1].Value)
	assert.Equal(t, "3", postfix[2].Value)
	assert.Equal(t, "*", postfix[3].Value)
	assert.Equal(t, "4", postfix[4].Value)
	assert.Equal(t, "/", postfix[5].Value)
	assert.Equal(t, "+", postfix[6].Value)
}

func TestPostfix_Parse_WithParenthesis(t *testing.T) {
	var (
		postfix, err = postfix.ParseInfix("(1 +2) * 3 / 4")
	)

	assert.NoError(t, err)
	assert.Equal(t, "1", postfix[0].Value)
	assert.Equal(t, "2", postfix[1].Value)
	assert.Equal(t, "+", postfix[2].Value)
	assert.Equal(t, "3", postfix[3].Value)
	assert.Equal(t, "*", postfix[4].Value)
	assert.Equal(t, "4", postfix[5].Value)
	assert.Equal(t, "/", postfix[6].Value)
}

func TestPostfix_Evaluate(t *testing.T) {
	var (
		pf, err = postfix.ParseInfix("((15 / (7 - (1 + 1))) * 3) - (2 + (1 + 1))")
		result  = postfix.Evaluate(pf, nil)
	)

	assert.NoError(t, err)
	assert.Equal(t, float64(5), result)
}

func TestPostfix_EvaluateWithVariables(t *testing.T) {
	var (
		pf, err = postfix.ParseInfix("((a / (b - (c + c))) * d) - (e + (c + c))")
		result  = postfix.Evaluate(pf, map[string]float64{
			"a": 15,
			"b": 7,
			"c": 1,
			"d": 3,
			"e": 2,
		})
	)

	assert.NoError(t, err)
	assert.Equal(t, float64(5), result)
}
