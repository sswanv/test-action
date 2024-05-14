package main

import (
	"fmt"
	"github.com/duke-git/lancet/v2/cryptor"
)

func main() {
	str := "hello"

	md5Str := cryptor.Md5String(str)
	fmt.Println(md5Str)

	// Output:
	// 5d41402abc4b2a76b9719d911017c592
}
