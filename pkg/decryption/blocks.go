package decryption

// transposeBlocks rearranges the bytes in blocks of data into a new slice of
// byte slices where each slice contains the nth byte of every block, with n
// being the index of the slice. This is useful in breaking repeating-key XOR
// ciphers, where you want to analyze each nth byte under the assumption they
// were XOR'd with the same byte of the key.
func transposeBlocks(blocks [][]byte, keySize int) [][]byte {
	// Initialize a slice of byte slices with length equal to the keySize.
	// This will hold the transposed bytes. Each index i of 'transposed'
	// will contain the ith byte from each of the input blocks.
	transposed := make([][]byte, keySize)

	for i := 0; i < keySize; i++ {
		for _, block := range blocks {
			// Check if the current block has enough bytes to include an ith
			// byte. This is necessary because the last block might be shorter
			// than the others.
			if i < len(block) {
				// If the current block has an ith byte, append that byte to
				// the i'th position in the 'transposed' slice. This effectively
				// groups all nth bytes together in the same slice, facilitating
				// further analysis or manipulation based on those bytes alone.
				transposed[i] = append(transposed[i], block[i])
			}
		}
	}
	return transposed
}
