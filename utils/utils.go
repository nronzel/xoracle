package utils

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"
)

// Decodes hex encoded string. Returns bytes.
func DecodeHex(encoded string) ([]byte, error) {
	return hex.DecodeString(encoded)
}

// Decodes Base64 encoded string. Returns bytes.
func DecodeBase64(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}

// scoreText evaluates the likelihood that a given byte slice (input) is
// meaningful or coherent English text. It does so by scoring each byte
// according to its frequency of appearance in English, with a higher score
// indicating a higher likelihood of the input being English.
func ScoreText(input []byte) float64 {
	// englishFreq maps English letters and the space character to their
	// relative frequencies in English text
	englishFreq := map[byte]float64{
		'a': 8.167, 'b': 1.492, 'c': 2.782, 'd': 4.253, 'e': 12.702,
		'f': 2.228, 'g': 2.015, 'h': 6.094, 'i': 6.966, 'j': 0.153,
		'k': 0.772, 'l': 4.025, 'm': 2.406, 'n': 6.749, 'o': 7.507,
		'p': 1.929, 'q': 0.095, 'r': 5.987, 's': 6.327, 't': 9.056,
		'u': 2.758, 'v': 0.978, 'w': 2.360, 'x': 0.150, 'y': 1.974,
		'z': 0.074, ' ': 13.000,
	}

	score := 0.0

	for _, b := range input {
		// If the byte is found in the englishFreq map,add its frequency value
		// to the total score.
		if val, ok := englishFreq[b]; ok {
			score += val
		}
	}

	return score
}

// isBase64Encoded checks if the input string is Base64 encoded.
// This function performs a basic check to see if the input is decodable from Base64.
func isBase64Encoded(input string) bool {
	if input == "" {
		return false
	}
	// Attempt to decode the input string from Base64
	_, err := base64.StdEncoding.DecodeString(input)
	return err == nil
}

// isHexEncoded checks if the input string is hex encoded.
// This function uses a regex pattern to ensure the string consists only of hexadecimal characters.
func isHexEncoded(input string) bool {
	if input == "" {
		return false
	}
	// Hex regex pattern to match valid hexadecimal characters
	hexPattern := `^[0-9A-Fa-f]+$`
	matched, err := regexp.MatchString(hexPattern, input)
	if err != nil {
		fmt.Println("Regex match error:", err)
		return false
	}
	return matched
}

func Decode(encodedData string) ([]byte, error) {
	var verifiedData []byte
	var err error
	if isHexEncoded(encodedData) {
		verifiedData, err = DecodeHex(encodedData)
		if err != nil {
			return nil, err
		}
		return verifiedData, nil
	}

	if isBase64Encoded(encodedData) {
		verifiedData, err = DecodeBase64(encodedData)
		if err != nil {
			return nil, err
		}
		return verifiedData, nil
	}

	// plaintext, or some other encoded data type not checked for.
	verifiedData = []byte(encodedData)

	return verifiedData, nil
}
