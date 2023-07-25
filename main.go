package main

import (
	"os"

	"github.com/ejuju/go-brainfuck/pkg/brainfuck"
)

func main() {
	vm := brainfuck.NewVM(os.Stdin, os.Stdout)
	err := vm.Exec(mustReadFile("helloworld.bf"))
	if err != nil {
		panic(err)
	}
}

func mustReadFile(fpath string) []byte {
	b, err := os.ReadFile(fpath)
	if err != nil {
		panic(err)
	}
	return b
}
