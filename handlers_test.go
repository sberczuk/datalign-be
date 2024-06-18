package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_validateExpression(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "pass",
			args: args{
				input: "1+2*3-4 /8",
			},
			wantErr: assert.NoError,
		},
		{
			name: "newlines at end",
			args: args{
				input: "1+2*3-4 /8\n",
			},
			wantErr: assert.NoError,
		},
		{
			name: "embedded newlines",
			args: args{
				input: "1+2*3\n-4 /8\n",
			},
			wantErr: assert.NoError,
		},
		{
			name: "pass-parens",
			args: args{
				input: "(1+2*3-4 /8)",
			},
			wantErr: assert.NoError,
		},
		{
			name: "fail",
			args: args{
				input: "(1+2*3-4 /8$$)",
			},
			wantErr: assert.Error,
		},
		{
			name: "decimal",
			args: args{
				input: "(1+2*3.2-4 /8$$)",
			},
			wantErr: assert.Error,
		},
		//{ // not yet implemented. We need a better way to check for valid inputs when parenthenses are involved
		//	name: "does not end with number with parens",
		//	args: args{
		//		input: "(1+2*3.2-4 /8+)",
		//	},
		//	wantErr: assert.Error,
		//},
		{
			name: "does not end with number",
			args: args{
				input: "1+2*3.2-4 /8+",
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, validateExpression(tt.args.input), fmt.Sprintf("validateExpression(%v)", tt.args.input))
		})
	}
}
