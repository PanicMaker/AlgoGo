package BackTrack

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{s: "25525511135"}, []string{"255.255.11.135", "255.255.111.35"}},
		{"2", args{s: "0000"}, []string{"0.0.0.0"}},
		{"3", args{s: "101023"}, []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 性能测试函数
func BenchmarkRestoreIpAddresses(b *testing.B) {
	// 定义测试用例
	testCases := []struct {
		input string
	}{
		{"25525511135"},
		{"0000"},
		{"1111"},
		{"010010"},
		{"101023"},
	}

	for _, tc := range testCases {
		b.Run(tc.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				restoreIpAddresses(tc.input)
			}
		})
	}
}
