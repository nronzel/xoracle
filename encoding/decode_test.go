package encoding

import "testing"

func TestDecode(t *testing.T) {
	tests := []struct {
		name    string
		encoded string
		want    string
		wantErr bool
	}{
		{name: "Hex Encoded Text", encoded: "7465737420696e707574", want: "test input", wantErr: false},
		{name: "Non-Encoded Text", encoded: "not encoded text", want: "not encoded text", wantErr: false},
		{name: "Base64 Encoded Text", encoded: "dGVzdCBpbnB1dA==", want: "test input", wantErr: false},
		{name: "Invalid Hex", encoded: "7465737420696e70757", want: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.encoded)
			if err == nil && tt.wantErr {
				t.Errorf("expected error and didn't get one")
			}
			if string(got) != tt.want {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}
