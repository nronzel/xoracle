package encoding

import (
	"encoding/hex"
	"fmt"
	"regexp"
)

// Decodes hex encoded string. Returns bytes.
func DecodeHex(encoded string) ([]byte, error) {
	return hex.DecodeString(encoded)
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
