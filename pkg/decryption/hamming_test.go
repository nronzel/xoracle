package decryption

import (
	"testing"
)

func TestHamming(t *testing.T) {
	tests := []struct {
		name    string
		s1      string
		s2      string
		want    int
		wantErr bool
	}{
		{
			name:    "valid input",
			s1:      "this is a test",
			s2:      "wokka wokka!!!",
			want:    37,
			wantErr: false,
		},
		{
			name:    "invalid input",
			s1:      "test",
			s2:      "aa",
			want:    0,
			wantErr: true,
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
			if got != tc.want {
				t.Fatalf("want: %v, got: %v", tc.want, got)
			}
		})
	}
}

func TestAverageHammingDistance(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		keySize int
		want    float64
		wantErr bool
	}{
		{
			name:    "sufficient data, no error",
			data:    []byte("this is a test string with enough length"),
			keySize: 4,
			want:    2.6666666666666665,
			wantErr: false,
		},
		{
			name:    "insufficient data, error expected",
			data:    []byte("short"),
			keySize: 10,
			want:    0,
			wantErr: true,
		},
		{
			name:    "exact data length for comparison, no error",
			data:    []byte("1234567890123456"), // 16 bytes, 4 comparisons with keySize 4
			keySize: 4,
			want:    1.3333333333333333,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := averageHammingDistance(tt.data, tt.keySize)
			if err == nil && tt.wantErr {
				t.Errorf("expected an error, got none")
			}
			if got != tt.want {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}
