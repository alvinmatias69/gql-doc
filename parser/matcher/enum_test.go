package matcher

import (
	"testing"
)

func TestMatcher_Enum(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "failure on invalid enum",
			args: args{
				input: "dat a",
			},
			want: false,
		},
		{
			name: "success match enum data",
			args: args{
				input: "DATA",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New()
			if got := m.Enum(tt.args.input); got != tt.want {
				t.Errorf("Matcher.Enum() = %v, want %v", got, tt.want)
			}
		})
	}
}
