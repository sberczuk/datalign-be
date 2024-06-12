package main

import (
	"github.com/expr-lang/expr"
)

// evaluateExpression evaluates a string expression using the expressopm library
func evaluateExpression(expression string) (float64, error) {
	program, err := expr.Compile(expression)
	if err != nil {
		return 0, err
	}

	output, err := expr.Run(program, nil)
	if err != nil {
		return 0, err
	}

	if err != nil {
		return 0, err
	}
	switch output.(type) {
	case float64:
		v := output.(float64)
		return v, nil
	case int:
		v := output.(int)
		return float64(v), nil
	}
	return 0, nil
}
