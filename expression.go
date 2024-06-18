package main

import (
	"fmt"
	"github.com/expr-lang/expr"
	"reflect"
	"strings"
)

// evaluateExpression evaluates a string expression using the expressopm library
func evaluateExpression(expression string) (float64, error) {

	expression = cleanString(expression)
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
	case float32:
		v := output.(float64)
		return v, nil
	case int16:
		v := output.(int16)
		return float64(v), nil
	case int32:
		v := output.(int32)
		return float64(v), nil
	case int64:
		v := output.(int64)
		return float64(v), nil
	case int:
		v := output.(int)
		return float64(v), nil
	}
	return 0, fmt.Errorf("unable to cast return value of type %v to float64", reflect.TypeOf(output))
}

// cleanString replaces newlines with spaces. Not strictly needed, but it seems cleaner
func cleanString(expression string) string {
	return strings.Replace(expression, "\n", " ", -1)
}
