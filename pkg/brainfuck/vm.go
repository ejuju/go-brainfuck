package brainfuck

import (
	"bufio"
	"io"
)

type VM struct {
	memptr int // data pointer
	mem    []byte
	stdin  *bufio.Reader
	stdout io.Writer
}

func NewVM(stdin io.Reader, stdout io.Writer) *VM {
	return &VM{stdin: bufio.NewReader(stdin), stdout: stdout}
}

func (vm *VM) Exec(src []byte) error {
	for i := 0; i < len(src); i++ {
		if vm.memptr >= len(vm.mem) {
			vm.mem = append(vm.mem, make([]byte, 1)...)
		}
		switch src[i] {
		case '>':
			vm.memptr++
		case '<':
			vm.memptr--
		case '+':
			vm.mem[vm.memptr]++
		case '-':
			vm.mem[vm.memptr]--
		case '[':
			if vm.mem[vm.memptr] == 0 {
				// Jump after next ']'
				loops := 1
				for loops > 0 {
					i++
					switch src[i] {
					case '[':
						loops++
					case ']':
						loops--
					}
				}
			}
		case ']':
			if vm.mem[vm.memptr] != 0 {
				// Go back after previous '['
				loops := 1
				for loops > 0 {
					i--
					switch src[i] {
					case ']':
						loops++
					case '[':
						loops--
					}
				}
			}
		case '.':
			_, err := vm.stdout.Write([]byte{vm.mem[vm.memptr]})
			if err != nil {
				return err
			}
		case ',':
			c, err := vm.stdin.ReadByte()
			if err != nil {
				return err
			}
			vm.mem[vm.memptr] = c
		}
	}
	return nil
}
