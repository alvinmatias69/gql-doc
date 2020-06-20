package extractor

import (
	"reflect"
	"testing"

	"github.com/alvinmatias69/gql-doc/entity"
)

func TestExtractor_Union(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		e       *Extractor
		args    args
		want    entity.Definition
		wantErr bool
	}{
		{
			name: "failure on invalid union",
			args: args{
				input: "union TestUnion",
			},
			wantErr: true,
		},
		{
			name: "success extract union",
			args: args{
				input: "union TestUnion = first | second",
			},
			want: entity.Definition{
				Name:    "TestUnion",
				Variant: entity.Union,
				Properties: []entity.Property{
					{
						Name: "first",
					},
					{
						Name: "second",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Extractor{}
			got, err := e.Union(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Extractor.Union() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extractor.Union() = %v, want %v", got, tt.want)
			}
		})
	}
}
