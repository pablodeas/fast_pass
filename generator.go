package main

import (
	"fmt"

	"github.com/sethvargo/go-password/password"
)

func main() {
	rand := password.MustGenerate(10, 0, 3, false, false)

	fmt.Println(rand)
}
