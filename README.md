# Go-Basex

使用给定的任意字母表编码和解码数据

*forked from eknkc/basex*

包含: `base2` `base8` `base11` `base16` `base32` `base36` `base56` `base57` `base58` `base62` `base67` 

以及任意自定义字母表

## base58

包含: `Bitcoin` `IPFS` `Flickr` `Ripple`

`Bitcoin` 推荐使用:

- `base58.Encode()`
- `base58.Decode()`
- `base58.CheckEncode()`
- `base58.CheckDecode()`

*base58, bech32 forked from btcsuite/btcutil*

## 使用

```go
go get github.com/fufuok/basex
```

## 示例

### base58

```go
// base58 `Bitcoin`
fmt.Println(basex.B58Encode([]byte("中")), base58.Encode([]byte("中")))
fmt.Println(string(basex.B58Decode("2KprQ")), string(base58.Decode("2KprQ")))
fmt.Println(basex.B58Encode([]byte("密~&#1x")), base58.Encode([]byte("密~&#1x")))
fmt.Println(string(basex.B58Decode("fREqFdERUAK")), string(base58.Decode("fREqFdERUAK")))
fmt.Println(base58.CheckEncode([]byte("中"), 20))
txt, ver, _ := base58.CheckDecode("4VhAoVNjGNT")
fmt.Println(string(txt), ver)
// base58 `Flickr`
fmt.Println(basex.B58fEncode([]byte("中")))
fmt.Println(string(basex.B58fDecode("2jPRp")))
fmt.Println(basex.B58fEncode([]byte("密~&#1x")))
fmt.Println(string(basex.B58fDecode("EqeQfCeqtaj")))
// base58 `Ripple`
fmt.Println(basex.B58rEncode([]byte("中")))
fmt.Println(string(basex.B58rDecode("pKFiQ")))
fmt.Println(basex.B58rEncode([]byte("密~&#1x")))
fmt.Println(string(basex.B58rDecode("CRNqEdNR7wK")))
```

### base2, base8, base11, base16, base32, base36, base56, base57, base62, base67

```go
fmt.Println(basex.B2Encode([]byte("中")))
fmt.Println(string(basex.B2Decode("111001001011100010101101")))
fmt.Println(basex.B2Encode([]byte("密~&#1x")))
fmt.Println(string(basex.B2Decode("1110010110101111100001100111111000100110001000110011000101111000")))

fmt.Println(basex.B8Encode([]byte("中")))
fmt.Println(string(basex.B8Decode("71134255")))
fmt.Println(basex.B8Encode([]byte("密~&#1x")))
fmt.Println(string(basex.B8Decode("1626574147704610630570")))

fmt.Println(basex.B11Encode([]byte("中")))
fmt.Println(string(basex.B11Decode("8508905")))
fmt.Println(basex.B11Encode([]byte("密~&#1x")))
fmt.Println(string(basex.B11Decode("2a820986a07867a2533")))

fmt.Println(basex.B16Encode([]byte("中")))
fmt.Println(string(basex.B16Decode("e4b8ad")))
fmt.Println(basex.B16Encode([]byte("密~&#1x")))
fmt.Println(string(basex.B16Decode("e5af867e26233178")))

fmt.Println(basex.B32Encode([]byte("中")))
fmt.Println(string(basex.B32Decode("E9E5D")))
fmt.Println(basex.B32Encode([]byte("密~&#1x")))
fmt.Println(string(basex.B32Decode("EBBW6FRK26CBR")))
fmt.Println(basex.B32zEncode([]byte("中")))
fmt.Println(string(basex.B32zDecode("qjqfp")))
fmt.Println(basex.B32zEncode([]byte("密~&#1x")))
fmt.Println(string(basex.B32zDecode("qmmhgxaungcma")))

fmt.Println(basex.B36Encode([]byte("中")))
fmt.Println(string(basex.B36Decode("8x9yl")))
fmt.Println(basex.B36Encode([]byte("密~&#1x")))
fmt.Println(string(basex.B36Decode("3hqrsahs0q9a0")))

fmt.Println(basex.B56Encode([]byte("中")))
fmt.Println(string(basex.B56Decode("3XMpP")))
fmt.Println(basex.B56Encode([]byte("密~&#1x")))
fmt.Println(string(basex.B56Decode("yZmQK7eFEgi")))

fmt.Println(basex.B57Encode([]byte("中")))
fmt.Println(string(basex.B57Decode("3RwaS")))
fmt.Println(basex.B57Encode([]byte("密~&#1x")))
fmt.Println(string(basex.B57Decode("oioDRyadoBe")))

fmt.Println(basex.B62Encode([]byte("中")))
fmt.Println(string(basex.B62Decode("10TrT")))
fmt.Println(basex.B62Encode([]byte("密~&#1x")))
fmt.Println(string(basex.B62Decode("jIBTQOl1Xcs")))

fmt.Println(basex.B67Encode([]byte("中")))
fmt.Println(string(basex.B67Decode("x4Ks")))
fmt.Println(basex.B67Encode([]byte("密~&#1x")))
fmt.Println(string(basex.B67Decode("JFWOp05dgWR")))
```

### 任意字母表

```go
alphabet := "x@-_*2"
enc, _ := basex.New(alphabet)
fmt.Println(enc.Encode([]byte("密~&#1x")))
src, _ := enc.Decode("_-2*-*_**@*-2**xx*-@_@*xx")
fmt.Println(string(src))
```







*ff*