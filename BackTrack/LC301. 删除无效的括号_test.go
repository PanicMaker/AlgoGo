package BackTrack

import (
	"reflect"
	"testing"
)

func Test_removeInvalidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{s: "()())()"}, []string{"(())()", "()()()"}},
		{"2", args{s: "(a)())()"}, []string{"(a())()", "(a)()()"}},
		{"3", args{s: ")("}, []string{""}},
		{"4", args{s: ")()"}, []string{"()"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeInvalidParentheses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeInvalidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
