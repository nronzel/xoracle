package encoding

import "encoding/base64"

// Decodes Base64 encoded string. Returns bytes.
func DecodeBase64(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
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
