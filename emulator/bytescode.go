package emulator

import "fmt"

type Bytecode byte

const (
	SUM  Bytecode = 0x40
	PROD Bytecode = 0x41
	QUOT Bytecode = 0x42
	DIFF Bytecode = 0x43
	PUSH Bytecode = 0x44
	POP  Bytecode = 0x45
	COMP Bytecode = 0x46
	EQ   Bytecode = 0x47
	NEQ  Bytecode = 0x48
	RES  Bytecode = 0x49
	SET  Bytecode = 0x4A
	JMP  Bytecode = 0x4B
)

func (runtime *Runtime) Sum() error {

	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
		return fmt.Errorf("[ %X ] Unknow register on line %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
	}

	runtime.AdvanceNOpcode(1)

	var source_operand Register = Register(runtime.Opcodes[runtime.QIP])

	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
		return fmt.Errorf("[ %X ] Unknow register on line %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
	}

	runtime.AdvanceNOpcode(1)

	var destination_operand Register = Register(runtime.Opcodes[runtime.QIP])

	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; ok {
		runtime.AdvanceNOpcode(1)

		var third_operand_destination Register = Register(runtime.Opcodes[runtime.QIP])

		runtime.Registers[third_operand_destination] += (runtime.Registers[source_operand] + runtime.Registers[destination_operand])

		if third_operand_destination == QSP {
			runtime.Stack.Push((runtime.Registers[source_operand] + runtime.Registers[destination_operand]))
		}

		runtime.AdvanceNOpcode(5)

	} else {

		runtime.Registers[QOX] += (runtime.Registers[source_operand] + runtime.Registers[destination_operand])

		runtime.AdvanceNOpcode(6)

	}

	return nil
}

func (runtime *Runtime) Prod() error {

	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
		return fmt.Errorf("[ %X ] Unknow register on line %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
	}

	runtime.AdvanceNOpcode(1)

	var source_operand Register = Register(runtime.Opcodes[runtime.QIP])

	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
		return fmt.Errorf("[ %X ] Unknow register on line %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
	}

	runtime.AdvanceNOpcode(1)

	var destination_operand Register = Register(runtime.Opcodes[runtime.QIP])

	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; ok {
		runtime.AdvanceNOpcode(1)

		var third_operand_destination Register = Register(runtime.Opcodes[runtime.QIP])

		runtime.Registers[third_operand_destination] *= (runtime.Registers[source_operand] * runtime.Registers[destination_operand])

		if third_operand_destination == QSP {
			runtime.Stack.Push((runtime.Registers[source_operand] * runtime.Registers[destination_operand]))
		}

		runtime.AdvanceNOpcode(5)

	} else {

		runtime.Registers[QOX] *= (runtime.Registers[source_operand] * runtime.Registers[destination_operand])

		runtime.AdvanceNOpcode(6)

	}

	return nil
}



func (runtime *Runtime) Quot() error {

	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
		return fmt.Errorf("[ %X ] Unknow register on line %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
	}

	runtime.AdvanceNOpcode(1)

	var source_operand Register = Register(runtime.Opcodes[runtime.QIP])

	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
		return fmt.Errorf("[ %X ] Unknow register on line %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
	}

	runtime.AdvanceNOpcode(1)

	var destination_operand Register = Register(runtime.Opcodes[runtime.QIP])

	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; ok {
		runtime.AdvanceNOpcode(1)

		var third_operand_destination Register = Register(runtime.Opcodes[runtime.QIP])

		runtime.Registers[third_operand_destination] /= (runtime.Registers[source_operand] / runtime.Registers[destination_operand])

		if third_operand_destination == QSP {
			runtime.Stack.Push((runtime.Registers[source_operand] / runtime.Registers[destination_operand]))
		}

		runtime.AdvanceNOpcode(5)

	} else {

		runtime.Registers[QOX] /= (runtime.Registers[source_operand] / runtime.Registers[destination_operand])

		runtime.AdvanceNOpcode(6)

	}

	return nil
}

func (runtime *Runtime) Diff() error {

	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
		return fmt.Errorf("[ %X ] Unknow register on line %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
	}

	runtime.AdvanceNOpcode(1)

	var source_operand Register = Register(runtime.Opcodes[runtime.QIP])

	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
		return fmt.Errorf("[ %X ] Unknow register on line %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
	}

	runtime.AdvanceNOpcode(1)

	var destination_operand Register = Register(runtime.Opcodes[runtime.QIP])

	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; ok {
		runtime.AdvanceNOpcode(1)

		var third_operand_destination Register = Register(runtime.Opcodes[runtime.QIP])

		runtime.Registers[third_operand_destination] -= (runtime.Registers[source_operand] - runtime.Registers[destination_operand])

		if third_operand_destination == QSP {
			runtime.Stack.Push((runtime.Registers[source_operand] - runtime.Registers[destination_operand]))
		}

		runtime.AdvanceNOpcode(5)

	} else {

		runtime.Registers[QOX] -= (runtime.Registers[source_operand] - runtime.Registers[destination_operand])

		runtime.AdvanceNOpcode(6)

	}

	return nil
}

func (runtime *Runtime) Set() error {

	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
		return fmt.Errorf("[ %X ] Unknow register on line %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
	}

	runtime.AdvanceNOpcode(1)

	var destination_operand Register = Register(runtime.Opcodes[runtime.QIP])

	var source_operand_value int = int(ByteArrayToInt(runtime.Opcodes[runtime.QIP+1 : runtime.QIP+6]))

	runtime.Registers[destination_operand] = source_operand_value

	runtime.AdvanceNOpcode(7)

	return nil
}

func (runtime *Runtime) SetExtended(prefix byte) error {
	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
		return fmt.Errorf("[ %X ] Unknow register on line %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
	}

	if prefix != 1 {
		return fmt.Errorf("[ %X ] Unknow prefix on QIP %016X", prefix, runtime.QIP)
	}
	runtime.AdvanceNOpcode(1)

	var destination_operand Register = Register(runtime.Opcodes[runtime.QIP])

	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
		return fmt.Errorf("[ %X ] Unknow register on line %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
	}

	runtime.AdvanceNOpcode(1)

	var source_operand Register = Register(runtime.Opcodes[runtime.QIP])

	runtime.Registers[destination_operand] = runtime.Registers[source_operand]

	runtime.AdvanceNOpcode(5)

	return nil

}

func (runtime *Runtime) Pop() error {
	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
		return fmt.Errorf("[ %X ] Unknow register on line %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
	}

	runtime.AdvanceNOpcode(1)

	var destination_operand Register = Register(runtime.Opcodes[runtime.QIP])

	runtime.Registers[destination_operand] = runtime.Stack.Top()

	runtime.Stack.Pop()

	runtime.Registers[QSP] = runtime.Stack.Top()

	runtime.AdvanceNOpcode(7)

	return nil
}

func (runtime *Runtime) Push() error {

	runtime.AdvanceNOpcode(1)

	var source_operand_value int = int(ByteArrayToInt(runtime.Opcodes[runtime.QIP : runtime.QIP+6]))

	fmt.Println(source_operand_value)
	runtime.Stack.Push(source_operand_value)

	runtime.AdvanceNOpcode(7)

	return nil
}

func (runtime *Runtime) PushExtended(prefix byte) error {
	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
		return fmt.Errorf("[ %X ] Unknow register on QIP %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
	}

	if prefix != 0x1 {
		return fmt.Errorf("[ %x ] Unknow prefix for push on QIP %016X")
	}
	runtime.AdvanceNOpcode(1)

	var source_operand Register = Register(runtime.Opcodes[runtime.QIP])

	runtime.Stack.Push(runtime.Registers[source_operand])

	runtime.Registers[QSP] = int(runtime.Stack.Top())

	runtime.AdvanceNOpcode(6)

	return nil
}

func (runtime *Runtime) Comp() {

	runtime.AdvanceNOpcode(1)

	var value_source int = int(runtime.Opcodes[runtime.QIP])

	runtime.AdvanceNOpcode(1)

	var value_dest int = int(runtime.Opcodes[runtime.QIP])

	if value_source == value_dest {
		runtime.Registers[QZF] = 0
	} else {
		runtime.Registers[QZF] = 1
	}

	runtime.AdvanceNOpcode(6)
}

func (runtime *Runtime) CompExtended(prefix byte) error {

	switch prefix {
	case 0x01:
		// register, value

		if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
			return fmt.Errorf("[ %X ] Unknow register on QIP %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
		}

		runtime.AdvanceNOpcode(1)

		var register Register = Register(runtime.Opcodes[runtime.QIP])

		runtime.AdvanceNOpcode(1)

		var value int = int(runtime.Opcodes[runtime.QIP])

		runtime.Stack.Push(runtime.Registers[register])
		runtime.Stack.Push(value)
		runtime.Registers[QSP] = runtime.Stack.Top()
	case 0x02:

		runtime.AdvanceNOpcode(1)

		var value int = int(runtime.Opcodes[runtime.QIP])

		runtime.AdvanceNOpcode(1)

		var register Register = Register(runtime.Opcodes[runtime.QIP])

		if runtime.Registers[register] == value {
			runtime.Registers[QZF] = 0
		} else {
			runtime.Registers[QZF] = 1
		}

	case 0x03:
		if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
			return fmt.Errorf("[ %X ] Unknow register on QIP %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
		}

		runtime.AdvanceNOpcode(1)

		var register1 Register = Register(runtime.Opcodes[runtime.QIP])

		if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
			return fmt.Errorf("[ %X ] Unknow register on QIP %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
		}

		runtime.AdvanceNOpcode(1)

		var register2 Register = Register(runtime.Opcodes[runtime.QIP])

		if runtime.Registers[register1] == runtime.Registers[register2] {
			runtime.Registers[QZF] = 0
		} else {
			runtime.Registers[QZF] = 1
		}

	default:
		return fmt.Errorf("[ %X ] Unknow prefix. QIP %016X", prefix, runtime.QIP)
	}

	runtime.AdvanceNOpcode(5)

	return nil
}

func (runtime *Runtime) Eq() {
	runtime.AdvanceNOpcode(1)

	var ip_dest int = int(runtime.Opcodes[runtime.QIP])
	
	var val1 int = runtime.Stack.Top()
	runtime.Stack.Pop()
	var val2 int = runtime.Stack.Top()

	if val1 == val2 {
		runtime.QIP = ip_dest * 8
		return
	}

	runtime.AdvanceNOpcode(7)
}

func (runtime *Runtime) Neq() {
	runtime.AdvanceNOpcode(1)

	var ip_dest int = int(runtime.Opcodes[runtime.QIP])
	
	var val1 int = runtime.Stack.Top()
	runtime.Stack.Pop()
	var val2 int = runtime.Stack.Top()

	if val1 != val2 {
		runtime.QIP = ip_dest * 8
		return
	}

	runtime.AdvanceNOpcode(7)
}

func (runtime *Runtime) Res() error {
	if _, ok := registers_table[Register(runtime.Opcodes[runtime.QIP+1])]; !ok {
		return fmt.Errorf("[ %X ] Unknow register on QIP %016X", runtime.Opcodes[runtime.QIP+1], runtime.QIP)
	}

	runtime.AdvanceNOpcode(1)

	var destination_operand Register = Register(runtime.Opcodes[runtime.QIP])

	runtime.Registers[destination_operand] = runtime.Registers[destination_operand] ^ runtime.Registers[destination_operand]

	runtime.AdvanceNOpcode(7)

	return nil
}

func (runtime *Runtime) Jmp() {

	runtime.AdvanceNOpcode(1)

	var dest_to_jump int = int(runtime.Opcodes[runtime.QIP])

	runtime.QIP = dest_to_jump * 8

	runtime.AdvanceNOpcode(7)
}
