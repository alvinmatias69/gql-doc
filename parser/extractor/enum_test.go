package extractor

import (
	"reflect"
	"testing"

	"github.com/alvinmatias69/gql-doc/entity"
)

func TestExtractor_Enum(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want entity.Property
	}{
		{
			name: "success extract enum",
			args: args{
				input: "DATA",
			},
			want: entity.Property{
				Name: "DATA",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New()
			if got := e.Enum(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extractor.Enum() = %v, want %v", got, tt.want)
			}
		})
	}
}
