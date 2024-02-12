package encoding

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
