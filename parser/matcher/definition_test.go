package matcher

import (
	"testing"
)

func TestMatcher_Object(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success match simple object",
			args: args{
				input: "type Data {",
			},
			want: true,
		},
		{
			name: "success match object with implement",
			args: args{
				input: "type Data implements SomeInterface {",
			},
			want: true,
		},
		{
			name: "success match object with multiple implement",
			args: args{
				input: "type Data implements SomeInterface & AnotherInterface {",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New()
			if got := m.Definition(tt.args.input); got != tt.want {
				t.Errorf("Matcher.Object() = %v, want %v", got, tt.want)
			}
		})
	}
}
