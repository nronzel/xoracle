package decryption

import (
	"testing"
)

func TestScoreResults(t *testing.T) {
	tests := []struct {
		name     string
		results  []DecryptionResult
		expected DecryptionResult
	}{
		{
			name: "A should win",
			results: []DecryptionResult{
				{KeySize: 4, Key: []byte("A"), DecryptedData: "This is valid English"},
				{KeySize: 4, Key: []byte("A"), DecryptedData: "asldkjfaiocnlajeizpg"},
			},
			expected: DecryptionResult{KeySize: 4, Key: []byte("A"), DecryptedData: "This is valid English"},
		},
		{
			name: "multiple same score, first one wins",
			results: []DecryptionResult{
				{KeySize: 3, Key: []byte{1, 2, 3}, DecryptedData: "excellent"},
				{KeySize: 5, Key: []byte{2, 3, 4}, DecryptedData: "excellent"},
			},
			expected: DecryptionResult{KeySize: 3, Key: []byte{1, 2, 3}, DecryptedData: "excellent"},
		},
		{
			name:     "empty input",
			results:  []DecryptionResult{},
			expected: DecryptionResult{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ScoreResults(tt.results)
			if got.DecryptedData != tt.expected.DecryptedData {
				t.Fatalf("Expected: %v, got: %v", tt.expected.DecryptedData, got.DecryptedData)
			}
		})
	}
}
