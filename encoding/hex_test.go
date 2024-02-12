package encoding

import "testing"

func TestIsHexEncoded(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "Valid Hex", input: "7465737420696e707574", want: true},
		{name: "Invalid Hex", input: "test input", want: false},
		{name: "Empty String", input: "", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isHexEncoded(tt.input)
			if got != tt.want {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}
