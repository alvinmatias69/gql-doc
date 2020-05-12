package comment

import "testing"

func TestExtract(t *testing.T) {
	type args struct {
		strComment string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success return trailing slash on empty comment",
			args: args{
				strComment: "#",
			},
			want: "\n",
		},
		{
			name: "success remove multiple comment tag",
			args: args{
				strComment: "#######",
			},
			want: "\n",
		},
		{
			name: "success extract comment",
			args: args{
				strComment: "# this is comment",
			},
			want: "this is comment\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Extract(tt.args.strComment); got != tt.want {
				t.Errorf("Extract() = %v, want %v", got, tt.want)
			}
		})
	}
}
