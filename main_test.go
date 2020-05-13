package main

import (
	"os"
	"testing"
)

func Test_resolvePath(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		prepare func()
	}{
		{
			name: "success return absolute path",
			args: args{
				param: "/home/gopher",
			},
			want: "/home/gopher",
		},
		{
			name: "success return absolute path with environtment variable",
			args: args{
				param: "$CUSTOM_HOME/gopher",
			},
			want: "/home/gopher",
			prepare: func() {
				os.Setenv("CUSTOM_HOME", "/home")
			},
		},
		{
			name: "success return relative path",
			args: args{
				param: "./gopher",
			},
			want: "gopher",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepare != nil {
				tt.prepare()
			}
			if got := resolvePath(tt.args.param); got != tt.want {
				t.Errorf("resolvePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
