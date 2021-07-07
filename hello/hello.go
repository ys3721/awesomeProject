package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"yaodashuai.com/user/hello/morestrings"

)

func main() {
	fmt.Println("Build and install and you first program and wo only know 7 words in english!")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Printf("%s", "123")
}
