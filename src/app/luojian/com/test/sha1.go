package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func main() {
	s := "sha1 this string"

	//sha1
	//h := sha1.New()
	//h.Write([]byte(s))
	h := sha1.Sum([]byte(s))

	fmt.Println(s)
	fmt.Println(hex.EncodeToString(h[:]))

	//md5
	//m := md5.New()
	//m.Write([]byte(s))
	//md := m.Sum(nil)

	m := md5.Sum([]byte(s))
	fmt.Println(hex.EncodeToString(m[:]))

	ss := "abcdefg"
	fmt.Println(ss)
	fmt.Println([]byte(ss))
	fmt.Println(string([]byte(ss)))

}
