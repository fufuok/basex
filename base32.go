package basex

var (
	enc32, _ = New("0123456789ABCDEFGHJKMNPQRSTVWXYZ")
	enc32z, _ = New("ybndrfg8ejkmcpqxot1uwisza345h769")
)

// Base32 编码
func B32Encode(b []byte) string {
	return enc32.Encode(b)
}

// Base32 解码
func B32Decode(s string) []byte {
	if b, err := enc32.Decode(s); err == nil {
		return b
	}

	return nil
}

// Base32 编码 (z-base-32)
func B32zEncode(b []byte) string {
	return enc32z.Encode(b)
}

// Base32 解码 (z-base-32)
func B32zDecode(s string) []byte {
	if b, err := enc32z.Decode(s); err == nil {
		return b
	}

	return nil
}
