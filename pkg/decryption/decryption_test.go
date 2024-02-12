package decryption

import (
	"reflect"
	"testing"

	"github.com/nronzel/xoracle/encoding"
)

func TestProcessKeySizes(t *testing.T) {
	decodedBase64, _ := encoding.DecodeBase64("MiciMCQ2YTYkOjViJTclJyQ=")
	decodedHex, _ := encoding.DecodeHex("3227223024366136243a35622537252724")
	tests := []struct {
		name        string
		topKeySizes []int
		data        []byte
		want        []DecryptionResult
	}{
		{name: "Base64 Decoded - Should Decrypt",
			topKeySizes: []int{2},
			data:        decodedBase64,
			want: []DecryptionResult{
				{KeySize: 2,
					Key:           []byte("AB"),
					DecryptedData: "secret text dudee",
				},
			},
		},
		{name: "Base64 Decoded - Should Decrypt",
			topKeySizes: []int{2},
			data:        decodedHex,
			want: []DecryptionResult{
				{KeySize: 2,
					Key:           []byte("AB"),
					DecryptedData: "secret text dudee",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProcessKeySizes(tt.topKeySizes, tt.data)
			if !reflect.DeepEqual(got[0], tt.want[0]) {
				t.Errorf("did not receive expected result: %v", got[0])
			}
		})
	}
}
