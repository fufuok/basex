package basex

var enc57, _ = New("23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// Base57 编码
func B57Encode(b []byte) string {
	return enc57.Encode(b)
}

// Base57 解码
func B57Decode(s string) []byte {
	if b, err := enc57.Decode(s); err == nil {
		return b
	}

	return nil
}
