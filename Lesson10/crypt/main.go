package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

// crypt

func main() {

	// MD5ハッシュ値を生成
	h := md5.New()

	io.WriteString(h, "abcdef")
	fmt.Println(h.Sum(nil))

	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)

	// SHA1ハッシュ値を生成
	a := sha1.New()

	io.WriteString(a, "abcxyz")
	fmt.Printf("%x\n", a.Sum(nil))

	// SHA256ハッシュ値を生成
	b := sha256.New()

	io.WriteString(b, "abcxyz")
	fmt.Printf("%x\n", b.Sum(nil))

	// SHA512ハッシュ値を生成
	c := sha512.New()

	io.WriteString(c, "abcxyz")
	fmt.Printf("%x\n", c.Sum(nil))

}