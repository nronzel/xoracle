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
