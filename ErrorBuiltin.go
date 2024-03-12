package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(check())
}

func check() error {
	return errors.New("ini error ya")
}
