# postfix
--
    import "github.com/darylnwk/postfix"


## Usage

#### func  Evaluate

```go
func Evaluate(postfix Postfix, variables map[string]float64) float64
```
Evaluate takes `Postfix` and a map of variables and returns the result of the
expression See
https://en.wikipedia.org/wiki/Reverse_Polish_notation#Postfix_evaluation_algorithm
for more details.

#### type Postfix

```go
type Postfix []mathtoken.Token
```

Postfix defines a list of tokens

#### func  ParseInfix

```go
func ParseInfix(infix string) (postfix Postfix, err error)
```
ParseInfix parses a mathematical expression in infix notation to postfix
notation (or RPN) using Shunting-yard algorithm. See
https://en.wikipedia.org/wiki/Shunting-yard_algorithm#The_algorithm_in_detail
for more details.
