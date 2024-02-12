package decryption

import "github.com/nronzel/xoracle/utils"

// Scores the resulting decrypted data after attempting to decrypt with the guessed
// key and keySize. Higher score means it is more likely English text and it will
// return that result, removing the false positives.
//
// If the scoring results in a tie; the first result will be returned as the best.
func ScoreResults(results []DecryptionResult) DecryptionResult {
	var highScore float64
	var best DecryptionResult
	for _, result := range results {
		score := utils.ScoreText([]byte(result.DecryptedData))
		if score > highScore {
			highScore = score
			best.DecryptedData = result.DecryptedData
			best.KeySize = result.KeySize
			best.Key = result.Key
		}
	}

	return best
}
