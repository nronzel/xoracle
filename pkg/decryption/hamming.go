package decryption

import (
	"errors"
	"fmt"
)

// averageHammingDistance calculates the average Hamming distance
// between multiple pairs of blocks of bytes, where each block is of size
// 'keySize'. Helps in estimating the size of the key used in encryption
// algorithms that operate in block modes.
func averageHammingDistance(data []byte, keySize int) (float64, error) {
	// At least 4 blocks of 'keySize' are needed to make a meaningful comparison.
	if len(data) < 4*keySize {
		msg := "not enough data to calculate Hamming distance for keySize: %d"
		return 0, fmt.Errorf(msg, keySize)
	}

	totalDistance := 0
	numComparisons := 0

	for blockNum := 1; blockNum <= 3; blockNum++ {
		block1 := data[:keySize]
		block2 := data[blockNum*keySize : (blockNum+1)*keySize]

		// Calculate the Hamming distance between block1 and block2.
		distance, err := hammingDistance(block1, block2)
		if err != nil {
			return 0.0, err
		}

		totalDistance += distance
		numComparisons++
	}

	// Calculate the average distance, normalized by the key size.
	// This gives a per-bit distance, making the result size-agnostic and
	// easier to compare across different key sizes.
	averageDistance := float64(totalDistance) /
		float64(numComparisons) /
		float64(keySize)

	return averageDistance, nil
}

// hammingDistance calculates the Hamming distance between two byte slices.
// The Hamming distance is the total number of bits that are different
// between two sequences of equal length.
func hammingDistance(s1, s2 []byte) (int, error) {
	if len(s1) != len(s2) {
		return 0, errors.New("error: byte slices must be of equal length")
	}

	distance := 0
	for i := range s1 {
		xor := s1[i] ^ s2[i] // XOR to find differing bits.

		// Count the number of set bits (1s) in the xor result.
		// Each set bit contributes to the Hamming distance.
		for xor != 0 {
			// Increment distance by the least significant bit of xor.
			// If the bit is 1, it means there is a difference in this position.
			distance += int(xor & 1)
			// Right shift xor by one bit to check the next bit in the next
			// iteration.
			xor >>= 1
		}
	}
	return distance, nil
}
