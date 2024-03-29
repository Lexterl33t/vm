package game

func GenerateBytecode() [][]byte {
	var bytecodes [][]byte = [][]byte{

		{
			0x4A, 0x20, 0x0A, 0x00, 0x00, 0x00, 0x00, 0x00, // SET QOX, 10
			0x4A, 0x21, 0x12, 0x00, 0x00, 0x00, 0x00, 0x00, // SET QCX, 0x12
			0x01, 0x46, 0x21, 0x13, 0x00, 0x00, 0x00, 0x00, // COMP 10, 10
			0x47, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // EQ 4
			0x48, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // NEQ 5
			0x44, 0x41, 0x41, 0x41, 0x00, 0x00, 0x00, 0x00, // PUSH 0x414141
			0x44, 0x43, 0x43, 0x43, 0x00, 0x00, 0x00, 0x00, // PUSH 0x434343
		},
		
		{
			0x44, 0xde, 0xad, 0xbe, 0xef, 0x00, 0x00, 0x00, // PUSH 0xdeadbeef
			0x4A, 0x23, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00,
		},

	}

	/*
		bytecodes[0][2] = byte(rand.Intn(50))
		bytecodes[0][10] = byte(rand.Intn(50))
		bytecodes[0][18] = byte(rand.Intn(50))
		bytecodes[0][19] = bytecodes[0][18]*/

	return bytecodes
}
