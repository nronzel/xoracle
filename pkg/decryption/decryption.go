package decryption

import (
	"fmt"
	"github.com/nronzel/xoracle/utils"
	"math"
	"sort"
)

// singleByteXORCipher attempts to decrypt a message that has been XOR'ed
// against a single byte. The function uses a scoring system to identify the
// most likely original message by iterating through all possible single-byte
// keys (0-255), applying each key to the encoded message, and evaluating the
// "decrypted" message with a scoring function. The key and message with the
// highest score are considered the best candidates for the original encryption
// key and message.
func singleByteXORCipher(encoded []byte) (byte, []byte) {
	var key byte
	maxScore := 0.0
	var message []byte

	// Iterate over all possible single-byte keys (0 to 255).
	for k := 0; k <= 255; k++ {
		decoded := make([]byte, len(encoded))

		// Decrypt the message with the current key by XOR'ing each byte of the
		// encoded message.
		for i, b := range encoded {
			decoded[i] = b ^ byte(k)
		}

		// Evaluate the decrypted message using a scoring function.
		score := utils.ScoreText(decoded)

		// If the current message's score is higher than the highest score found
		// so far, update maxScore, key, and message with the current values.
		if score > maxScore {
			maxScore = score
			key = byte(k)
			message = decoded
		}
	}

	return key, message
}

// guessKeySizes attempts to guess the key size used in an encryption algorithm
// based on the average Hamming distance between blocks of bytes in the
// encrypted data. It returns a slice of the top candidate key sizes that have
// the lowest average Hamming distances, suggesting these sizes are likely
// candidates for the actual key size.
func GuessKeySizes(data []byte) ([]int, error) {
	const maxKeySize = 40      // The maximum key size to test.
	const maxKeysToCompare = 2 // The number of top key sizes to return.

	type keySizeScore struct {
		KeySize  int
		Distance float64
	}

	var scores []keySizeScore

	// Loop through each possible key size from 2 to maxKeySize (inclusive).
	for keySize := 2; keySize <= maxKeySize; keySize++ {
		// Calculate the average Hamming distance for this key size.
		if keySize*4 > len(data) {
			continue
		}
		averageDistance, err := averageHammingDistance(data, keySize)
		if err != nil {
			return nil, fmt.Errorf("calculating average distance: %w\n", err)
		}
		scores = append(scores, keySizeScore{keySize, averageDistance})
	}

	// Sort the scores slice based on the average distance, with the lowest
	// distances first. This prioritizes key sizes with the smallest differences
	// between blocks, suggesting a better fit.
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Distance < scores[j].Distance
	})

	topKeySizes := make([]int, 0, maxKeysToCompare)
	// Select up to maxKeysToCompare key sizes with the lowest average distances.
	for i := 0; i < maxKeysToCompare && i < len(scores); i++ {
		topKeySizes = append(topKeySizes, scores[i].KeySize)
	}

	return topKeySizes, nil
}

type DecryptionResult struct {
	KeySize       int
	Key           []byte
	DecryptedData string
}

// processKeySizes attempts to decrypt the provided byte slice (data) for each
// of the top key sizes found. It attempts to break a repeating-key XOR cipher
// without directly knowing the key.
func ProcessKeySizes(topKeySizes []int, data []byte) []DecryptionResult {
	var results []DecryptionResult
	for _, keySize := range topKeySizes {
		// Break the ciphertext into blocks of KEYSIZE length to manage the
		// analysis in chunks. This approach is critical for transposing the
		// bytes correctly in a later step.
		blocks := make([][]byte, int(math.Ceil(float64(len(data))/float64(keySize))))
		for i := 0; i < len(blocks); i++ {
			start := i * keySize
			end := start + keySize
			if end > len(data) {
				end = len(data) // Stay in bounds of data slice.
			}
			blocks[i] = data[start:end]
		}

		// Transpose the blocks: Make N new blocks where each block is composed
		// of the nth byte of every original block.
		transposed := transposeBlocks(blocks, keySize)

		// Attempt to solve each transposed block as if it was encrypted with
		// a single-character XOR cipher. The idea is that each byte of the key
		// is used to XOR the same position across multiple blocks.
		key := make([]byte, keySize)
		for i, block := range transposed {
			k, _ := singleByteXORCipher(block)
			// Collect each byte of the key from the solved single-byte XOR ciphers.
			key[i] = k
		}

		// Use the derived key to decrypt the entire message, XORing the data
		// with the repeating key pattern.
		decrypted := make([]byte, len(data))
		for i, b := range data {
			// Apply the key in a repeating pattern across the data.
			decrypted[i] = b ^ key[i%len(key)]
		}
		results = append(results, DecryptionResult{
			KeySize:       keySize,
			Key:           key,
			DecryptedData: string(decrypted),
		})
	}
	return results
}
