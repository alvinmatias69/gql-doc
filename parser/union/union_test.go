package union

import (
	"reflect"
	"testing"

	"github.com/alvinmatias69/gql-doc/entity"
)

func TestMatch(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success test for union",
			args: args{
				line: "union ImageQuery = min | complete",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Match(tt.args.line); got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtract(t *testing.T) {
	type args struct {
		line    string
		comment string
	}
	tests := []struct {
		name string
		args args
		want entity.VarType
	}{
		{
			name: "success extract union data",
			args: args{
				line:    "union ImageQuery = min | complete",
				comment: "ImageQuery unionized",
			},
			want: entity.VarType{
				Name:    "ImageQuery",
				Comment: "ImageQuery unionized",
				Parameters: []entity.Parameter{
					{
						Name: "min",
					},
					{
						Name: "complete",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Extract(tt.args.line, tt.args.comment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extract() = %+v,\nwant = %+v", got, tt.want)
			}
		})
	}
}
