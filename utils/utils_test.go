package utils

import (
	"testing"
)

func TestScoreText(t *testing.T) {
	tests := []struct {
		name  string
		text  string
		score float64
	}{
		{name: "Valid score", text: "What Do You Want From Me", score: 158.66299999999998},
		{name: "No text", text: "", score: 0.0},
		{name: "Single letter", text: "a", score: 8.167},
		{name: "Space", text: " ", score: 13.000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ScoreText([]byte(tt.text))
			if got != tt.score {
				t.Fatalf("expected: %f, got: %f", tt.score, got)
			}
		})
	}
}

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
