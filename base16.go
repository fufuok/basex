package basex

var enc16, _ = New("0123456789abcdef")

// Base16 编码
func B16Encode(b []byte) string {
	return enc16.Encode(b)
}

// Base16 解码
func B16Decode(s string) []byte {
	if b, err := enc16.Decode(s); err == nil {
		return b
	}

	return nil
}
