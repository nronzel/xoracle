package decryption

import (
	"testing"
)

func TestHamming(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected int
		wantErr  bool
	}{
		{
			name:     "valid input",
			s1:       "this is a test",
			s2:       "wokka wokka!!!",
			expected: 37,
			wantErr:  false,
		},
		{
			name:     "invalid input",
			s1:       "test",
			s2:       "aa",
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b1 := []byte(tc.s1)
			b2 := []byte(tc.s2)
			got, err := hammingDistance(b1, b2)
			if err != nil && !tc.wantErr {
				t.Errorf("got an unexpected error: %v", err)
			}
			if got != tc.expected {
				t.Fatalf("expected: %d, got: %d", tc.expected, got)
			}
		})
	}
}

func TestAverageHammingDistance(t *testing.T) {
	tests := []struct {
		name          string
		data          []byte
		keySize       int
		expected      float64
		expectedError bool
	}{
		{
			name:          "sufficient data, no error",
			data:          []byte("this is a test string with enough length"),
			keySize:       4,
			expected:      2.6666666666666665,
			expectedError: false,
		},
		{
			name:          "insufficient data, error expected",
			data:          []byte("short"),
			keySize:       10,
			expected:      0,
			expectedError: true,
		},
		{
			name:          "exact data length for comparison, no error",
			data:          []byte("1234567890123456"), // 16 bytes, 4 comparisons with keySize 4
			keySize:       4,
			expected:      1.3333333333333333,
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := averageHammingDistance(tt.data, tt.keySize)
			if err == nil && tt.expectedError {
				t.Errorf("expected an error, got none")
			}
			if got != tt.expected {
				t.Errorf("got: %v, expected: %v", got, tt.expected)
			}
		})
	}
}
