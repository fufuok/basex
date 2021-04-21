package basex

var enc2, _ = New("01")

// Base2 编码
func B2Encode(b []byte) string {
	return enc2.Encode(b)
}

// Base2 解码
func B2Decode(s string) []byte {
	if b, err := enc2.Decode(s); err == nil {
		return b
	}

	return nil
}
