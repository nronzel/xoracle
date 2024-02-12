package decryption

type DecryptionResult struct {
	KeySize       int
	Key           []byte
	DecryptedData string
}
