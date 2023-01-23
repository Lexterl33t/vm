package main

import (
	"emulator/emulator"
	"fmt"
)

func main() {
	var bytecodes []byte = []byte{
		0x4A, 0x20, 0x0A, 0x00, 0x00, 0x00, 0x00, 0x00, // SET QOX, 10
		0x4A, 0x21, 0x12, 0x00, 0x00, 0x00, 0x00, 0x00, // SET QCX, 0x12
		0x01, 0x46, 0x21, 0x13, 0x00, 0x00, 0x00, 0x00, // COMP 10, 10
		0x47, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // EQ 4
		0x48, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // NEQ 5
		0x44, 0x41, 0x41, 0x41, 0x00, 0x00, 0x00, 0x00, // PUSH 0x414141
		0x44, 0x43, 0x43, 0x43, 0x00, 0x00, 0x00, 0x00, // PUSH 0x434343
	}

	registers := emulator.Exec(bytecodes)

	fmt.Println(registers.Registers)
}
