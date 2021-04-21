package basex

var enc62, _ = New("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Base62 编码
func B62Encode(b []byte) string {
	return enc62.Encode(b)
}

// Base62 解码
func B62Decode(s string) []byte {
	if b, err := enc62.Decode(s); err == nil {
		return b
	}

	return nil
}
