package matcher

import (
	"testing"
)

func TestMatcher_Property(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success match simple property",
			args: args{
				input: "name: string",
			},
			want: true,
		},
		{
			name: "success match property non-nullable",
			args: args{
				input: "name: string!",
			},
			want: true,
		},
		{
			name: "success match property list",
			args: args{
				input: "name: [string]",
			},
			want: true,
		},
		{
			name: "success match property with argument",
			args: args{
				input: "name(id: int): string",
			},
			want: true,
		},
		{
			name: "success match property with multiple arguments",
			args: args{
				input: "name(id: int, age: int): string",
			},
			want: true,
		},
		{
			name: "success match property with directives",
			args: args{
				input: "name(id: int, age: int): string @deprecated",
			},
			want: true,
		},
		{
			name: "success match property with default arguments",
			args: args{
				input: "name(id: int 15, age: int): string @deprecated",
			},
			want: true,
		},
		{
			name: "success match property with directive arguments",
			args: args{
				input: "name(id: int 15, age: int @deprecated): string @deprecated",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New()
			if got := m.Property(tt.args.input); got != tt.want {
				t.Errorf("Matcher.Property() = %v, want %v", got, tt.want)
			}
		})
	}
}
