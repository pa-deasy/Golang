package main

import (
	"fmt"

	"example.com/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello to the world")
	fmt.Println(morestrings.ReverseRunes("oG olleH"))
	fmt.Println(cmp.Diff("Hello to the world", "Hello Go"))
}
