package matcher

import "testing"

func TestComment(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success match comment string",
			args: args{
				input: "# this is a comment",
			},
			want: true,
		},
		{
			name: "success match non-comment string",
			args: args{
				input: "name: Int",
			},
			want: false,
		},
		{
			name: "success match empty string",
			args: args{
				input: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match := New()
			if got := match.Comment(tt.args.input); got != tt.want {
				t.Errorf("Comment() = %v, want %v", got, tt.want)
			}
		})
	}
}
