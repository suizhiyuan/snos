package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := errors.New("inner error")
	erro := errors.Wrap(err, "out err")
	fmt.Println(erro)
}
