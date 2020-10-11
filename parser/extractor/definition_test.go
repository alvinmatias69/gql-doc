package extractor

import (
	"reflect"
	"testing"

	"github.com/alvinmatias69/gql-doc/entity"
)

func TestExtractor_Object(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Definition
		wantErr bool
	}{
		{
			name: "faiure on malformed object",
			args: args{
				input: "tpe",
			},
			wantErr: true,
		},
		{
			name: "success parse object",
			args: args{
				input: "type Data {",
			},
			want: entity.Definition{
				Name:    "Data",
				Variant: entity.Object,
			},
		},
		{
			name: "success parse interface",
			args: args{
				input: "interface Data {",
			},
			want: entity.Definition{
				Name:    "Data",
				Variant: entity.Interface,
			},
		},
		{
			name: "success parse enum",
			args: args{
				input: "enum Data {",
			},
			want: entity.Definition{
				Name:    "Data",
				Variant: entity.Enum,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New()
			got, err := e.Definition(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Extractor.Object() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extractor.Object() = %v, want %v", got, tt.want)
			}
		})
	}
}
