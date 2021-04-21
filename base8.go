package basex

var enc8, _ = New("01234567")

// Base8 编码
func B8Encode(b []byte) string {
	return enc8.Encode(b)
}

// Base8 解码
func B8Decode(s string) []byte {
	if b, err := enc8.Decode(s); err == nil {
		return b
	}

	return nil
}
