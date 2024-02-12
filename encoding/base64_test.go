package encoding

import "testing"

func TestIsBase64Encoded(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "Valid Base64", input: "dGVzdCBpbnB1dA==", want: true},
		{name: "Invalid Base64", input: "test input", want: false},
		{name: "Empty String", input: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isBase64Encoded(tt.input)
			if got != tt.want {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}
