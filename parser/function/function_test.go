package function

import (
	"reflect"
	"testing"

	"github.com/alvinmatias69/gql-doc/entity"
)

func TestExtract(t *testing.T) {
	type args struct {
		strFn   string
		comment string
	}
	tests := []struct {
		name string
		args args
		want entity.Function
	}{
		{
			name: "success return no params",
			args: args{
				strFn: "getProduct : Product!",
			},
			want: entity.Function{
				Name:       "getProduct",
				ReturnType: "Product",
			},
		},
		{
			name: "success return one params",
			args: args{
				strFn: "getProduct (limit: Int!) : Product!",
			},
			want: entity.Function{
				Name: "getProduct",
				Parameters: []entity.Parameter{
					{
						Name:          "limit",
						ParamType:     "Int",
						IsMandatory:   true,
						IsBuiltInType: true,
					},
				},
				ReturnType: "Product",
			},
		},
		{
			name: "success return two params",
			args: args{
				strFn: "getProduct (limit: Int!, offset: Int!) : Product!",
			},
			want: entity.Function{
				Name: "getProduct",
				Parameters: []entity.Parameter{
					{
						Name:          "limit",
						ParamType:     "Int",
						IsMandatory:   true,
						IsBuiltInType: true,
					},
					{
						Name:          "offset",
						ParamType:     "Int",
						IsMandatory:   true,
						IsBuiltInType: true,
					},
				},
				ReturnType: "Product",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Extract(tt.args.strFn, tt.args.comment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extract() = %v, want %v", got, tt.want)
			}
		})
	}
}
