package decryption

import (
	"reflect"
	"testing"
)

// TestTransposeBlocks tests the transposeBlocks function with various input scenarios.
func TestTransposeBlocks(t *testing.T) {
	tests := []struct {
		name    string
		blocks  [][]byte
		keySize int
		want    [][]byte
	}{
		{
			name: "Even-sized blocks",
			blocks: [][]byte{
				{0x00, 0x01, 0x02},
				{0x10, 0x11, 0x12},
				{0x20, 0x21, 0x22},
			},
			keySize: 3,
			want: [][]byte{
				{0x00, 0x10, 0x20},
				{0x01, 0x11, 0x21},
				{0x02, 0x12, 0x22},
			},
		},
		{
			name: "Last block shorter",
			blocks: [][]byte{
				{0x00, 0x01, 0x02},
				{0x10, 0x11, 0x12},
				{0x20, 0x21}, // Shorter block
			},
			keySize: 3,
			want: [][]byte{
				{0x00, 0x10, 0x20},
				{0x01, 0x11, 0x21},
				{0x02, 0x12},
			},
		},
		{
			name: "Single block",
			blocks: [][]byte{
				{0x00, 0x01, 0x02},
			},
			keySize: 3,
			want: [][]byte{
				{0x00},
				{0x01},
				{0x02},
			},
		},
		{
			name: "Varying key sizes",
			blocks: [][]byte{
				{0x00, 0x01},
				{0x10, 0x11},
			},
			keySize: 2,
			want: [][]byte{
				{0x00, 0x10},
				{0x01, 0x11},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := transposeBlocks(tt.blocks, tt.keySize)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transposeBlocks() = %v, want %v", got, tt.want)
			}
		})
	}
}
