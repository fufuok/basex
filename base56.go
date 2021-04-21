package basex

var enc56, _ = New("23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz")

// Base56 编码
func B56Encode(b []byte) string {
	return enc56.Encode(b)
}

// Base56 解码
func B56Decode(s string) []byte {
	if b, err := enc56.Decode(s); err == nil {
		return b
	}

	return nil
}
