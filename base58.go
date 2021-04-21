package basex

var (
	// Bitcoin, IPFS
	enc58, _       = New("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	enc58Flickr, _ = New("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
	enc58Ripple, _ = New("rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz")
)

// Base58 编码, 推荐使用: base58.Encode()
func B58Encode(b []byte) string {
	return enc58.Encode(b)
}

// Base58 解码, 推荐使用: base58.Decode()
func B58Decode(s string) []byte {
	if b, err := enc58.Decode(s); err == nil {
		return b
	}

	return nil
}

// Base58 Flickr 编码
func B58fEncode(b []byte) string {
	return enc58Flickr.Encode(b)
}

// Base56 Flickr 解码
func B58fDecode(s string) []byte {
	if b, err := enc58Flickr.Decode(s); err == nil {
		return b
	}

	return nil
}

// Base58 Ripple 编码
func B58rEncode(b []byte) string {
	return enc58Ripple.Encode(b)
}

// Base56 Ripple 解码
func B58rDecode(s string) []byte {
	if b, err := enc58Ripple.Decode(s); err == nil {
		return b
	}

	return nil
}
