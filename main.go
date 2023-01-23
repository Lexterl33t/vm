package main

import (
	"emulator/emulator"
	"fmt"
)

func main() {
	var bytecodes []byte = []byte{
		0x4A, 0x20, 0x0A, 0x00, 0x00, 0x00, 0x00, 0x00, // SET QOX, 10
		0x40, 0x21, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, // SUM QCX, QOX
		0x40, 0x21, 0x20, 0x21, 0x00, 0x00, 0x00, 0x00, // SUM QCX, QOX, QCX
		0x44, 0x13, 0x37, 0x00, 0x00, 0x00, 0x00, 0x00, // PUSH 0x1337
		0x45, 0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // POP QCX
		0x44, 0x13, 0x37, 0x00, 0x00, 0x00, 0x00, 0x00, // PUSH 0x1337
		0x40, 0x21, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, // SUM QCX, QOX
		0x46, 0x0A, 0x0A, 0x00, 0x00, 0x00, 0x00, 0x00, // COMP 10, 10
		0x47, 0x0A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // EQ 10
		0x48, 0x0B, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // NEQ 11
		0x44, 0x41, 0x41, 0x41, 0x00, 0x00, 0x00, 0x00, // PUSH 0x414141
		0x44, 0x43, 0x43, 0x43, 0x00, 0x00, 0x00, 0x00, // PUSH 0x434343
	}

	registers := emulator.Exec(bytecodes)

	fmt.Println(registers)
}