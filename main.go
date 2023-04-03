package main

import (
	_"vm/server"
	"vm/emulator"
	"fmt"
)

func main() {
	var bytecodes []byte = []byte{
	}
 /*	
	serve, err := server.NewServer("1337")
	if err != nil {
		panic(err)
	}

	serve.Run()
*/

	res := emulator.Exec(bytecodes)
	
	fmt.Println(res)
}
