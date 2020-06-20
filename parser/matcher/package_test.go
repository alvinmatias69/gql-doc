package matcher

import (
	"testing"
)

func TestMatcher_Package(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success match package input",
			args: args{
				input: "package hello",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New()
			if got := m.Package(tt.args.input); got != tt.want {
				t.Errorf("Matcher.Package() = %v, want %v", got, tt.want)
			}
		})
	}
}
