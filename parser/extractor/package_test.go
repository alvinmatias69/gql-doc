package extractor

import "testing"

func TestExtractor_Package(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success get package name",
			args: args{
				input: "package test_hello",
			},
			want: "test_hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Extractor{}
			if got := e.Package(tt.args.input); got != tt.want {
				t.Errorf("Extractor.Package() = %v, want %v", got, tt.want)
			}
		})
	}
}
