package basex

import (
	"bytes"
	"math/rand"
	"testing"
	"time"

	"github.com/fufuok/basex/base58"
)

var (
	raw5k     = bytes.Repeat([]byte{0xff}, 5000)
	encoded5k = base58.Encode(raw5k)
	n         = 5000000
	testPairs = make([]testValues, 0, n)
)

type testValues struct {
	dec []byte
	enc string
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < n; i++ {
		data := make([]byte, 32)
		rand.Read(data)
		testPairs = append(testPairs, testValues{dec: data, enc: base58.Encode(data)})
	}
}

func BenchmarkB58Encode_5K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		B58Encode(raw5k)
	}
}

func BenchmarkBase58Encode_5K(b *testing.B) {
	b.SetBytes(int64(len(raw5k)))
	for i := 0; i < b.N; i++ {
		base58.Encode(raw5k)
	}
}

func BenchmarkB58Encode_N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		B58Encode(testPairs[i].dec)
	}
}

func BenchmarkBase58Encode_N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		base58.Encode(testPairs[i].dec)
	}
}

func BenchmarkB58Decode_5K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		B58Decode(encoded5k)
	}
}

func BenchmarkBase58Decode_5K(b *testing.B) {
	b.SetBytes(int64(len(encoded5k)))
	for i := 0; i < b.N; i++ {
		base58.Decode(encoded5k)
	}
}

func BenchmarkB58Decode_N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		B58Decode(testPairs[i].enc)
	}
}

func BenchmarkBase58Decode_N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		base58.Decode(testPairs[i].enc)
	}
}

// BenchmarkB58Encode_5K-8      	       3	 341248767 ns/op	  240032 B/op	      26 allocs/op
// BenchmarkBase58Encode_5K-8   	     218	   5654490 ns/op	   0.88 MB/s	   19208 B/op	       4 allocs/op
// BenchmarkB58Encode_N-8       	   75786	     16426 ns/op	    1120 B/op	       8 allocs/op
// BenchmarkBase58Encode_N-8    	 1504614	       881 ns/op	     168 B/op	       4 allocs/op
// BenchmarkB58Decode_5K-8      	      64	  17330180 ns/op	   48640 B/op	      16 allocs/op
// BenchmarkBase58Decode_5K-8   	    3694	    485801 ns/op	  14.06 MB/s	  348800 B/op	     129 allocs/op
// BenchmarkB58Decode_N-8       	  411217	      2853 ns/op	     240 B/op	       5 allocs/op
// BenchmarkBase58Decode_N-8    	 1997832	       612 ns/op	     128 B/op	       5 allocs/op