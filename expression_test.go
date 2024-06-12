package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_evaluateExpression(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "simple math ints",
			args: args{
				expression: "1+3",
			},
			want:    4,
			wantErr: assert.NoError,
		},
		{
			name: "simple math float",
			args: args{
				expression: "1.3+3.4",
			},
			want:    4.7,
			wantErr: assert.NoError,
		},
		{
			name: "simple math with spaces ints",
			args: args{
				expression: "1+ 3 ",
			},
			want:    4,
			wantErr: assert.NoError,
		},
		{
			name: "simple math ints",
			args: args{
				expression: "1+3",
			},
			want:    4,
			wantErr: assert.NoError,
		},
		{
			name: "precedence ints",
			args: args{
				expression: "1+ 3 *5  ",
			},
			want:    16,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := evaluateExpression(tt.args.expression)
			if !tt.wantErr(t, err, fmt.Sprintf("evaluateExpression(%v)", tt.args.expression)) {
				return
			}
			assert.Equalf(t, tt.want, got, "evaluateExpression(%v)", tt.args.expression)
		})
	}
}
