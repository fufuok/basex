package basex

var enc36, _ = New("0123456789abcdefghijklmnopqrstuvwxyz")

// Base36 编码
func B36Encode(b []byte) string {
	return enc36.Encode(b)
}

// Base36 解码
func B36Decode(s string) []byte {
	if b, err := enc36.Decode(s); err == nil {
		return b
	}

	return nil
}
