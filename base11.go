package basex

var enc11, _ = New("0123456789a")

// Base11 编码
func B11Encode(b []byte) string {
	return enc11.Encode(b)
}

// Base11 解码
func B11Decode(s string) []byte {
	if b, err := enc11.Decode(s); err == nil {
		return b
	}

	return nil
}
