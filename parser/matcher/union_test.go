package matcher

import (
	"testing"
)

func TestMatcher_Union(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success match union",
			args: args{
				input: "union query = random | sorted",
			},
			want: true,
		},
		{
			name: "success match union with directive",
			args: args{
				input: "union query @deprecated = random | sorted",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New()
			if got := m.Union(tt.args.input); got != tt.want {
				t.Errorf("Matcher.Union() = %v, want %v", got, tt.want)
			}
		})
	}
}
