package parameter

import (
	"reflect"
	"testing"

	"github.com/alvinmatias69/gql-doc/entity"
)

func TestParse(t *testing.T) {
	type args struct {
		strParams string
	}
	tests := []struct {
		name string
		args args
		want entity.Parameter
	}{
		{
			name: "success return name and params",
			args: args{
				strParams: "name: type",
			},
			want: entity.Parameter{
				Name:      "name",
				ParamType: "type",
			},
		},
		{
			name: "success return built in status",
			args: args{
				strParams: "name: String",
			},
			want: entity.Parameter{
				Name:          "name",
				ParamType:     "String",
				IsBuiltInType: true,
			},
		},
		{
			name: "success return mandatory status",
			args: args{
				strParams: "name: String!",
			},
			want: entity.Parameter{
				Name:          "name",
				ParamType:     "String",
				IsBuiltInType: true,
				IsMandatory:   true,
			},
		},
		{
			name: "success return list status",
			args: args{
				strParams: "name: [String]!",
			},
			want: entity.Parameter{
				Name:          "name",
				ParamType:     "String",
				IsBuiltInType: true,
				IsMandatory:   true,
				IsList:        true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.strParams); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtract(t *testing.T) {
	type args struct {
		strParams string
	}
	tests := []struct {
		name string
		args args
		want []entity.Parameter
	}{
		{
			name: "success return nil on empty string",
			args: args{
				strParams: "",
			},
			want: nil,
		},
		{
			name: "success return one param",
			args: args{
				strParams: "(name: String)",
			},
			want: []entity.Parameter{
				{
					Name:          "name",
					ParamType:     "String",
					IsBuiltInType: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Extract(tt.args.strParams); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extract() = %v, want %v", got, tt.want)
			}
		})
	}
}
