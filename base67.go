package basex

var enc67, _ = New("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_.!~")

// Base67 编码
func B67Encode(b []byte) string {
	return enc67.Encode(b)
}

// Base67 解码
func B67Decode(s string) []byte {
	if b, err := enc67.Decode(s); err == nil {
		return b
	}

	return nil
}
