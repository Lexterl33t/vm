package emulator

import "fmt"
type Runtime struct {
	Registers map[Register]int
	QIP       int
	QSP 	  *int
	Stack     Stack
	Opcodes   []byte
}

func NewRuntime(opcodes []byte) *Runtime {

	stack := NewStack()

	lastElement := &stack.Stack[len(stack.Stack)-1]

	runtime := &Runtime{
		Registers: make(map[Register]int),
		QIP:       0,
		Stack:     stack,
		Opcodes:   opcodes,
		QSP: lastElement,	
	}
		
	return runtime
}

func (runtime *Runtime) NextOpcode() {
	runtime.QIP += 8
}

func (runtime *Runtime) AdvanceNOpcode(n int) {
	runtime.QIP += n
}

func Exec(opcodes []byte) *Runtime {
	var runtime *Runtime = NewRuntime(opcodes)

	for runtime.QIP < len(opcodes) {

		switch opcodes[runtime.QIP] {
		case byte(SUM):
			if err := runtime.Sum(); err != nil {
				panic(err)
			}
		case byte(SET):
			if err := runtime.Set(); err != nil {
				panic(err)
			}
		case byte(POP):
			if err := runtime.Pop(); err != nil {
				panic(err)
			}
		case byte(PUSH):
			if err := runtime.Push(); err != nil {
				panic(err)
			}
		case byte(0x1), byte(0x2), byte(0x3):
			var prefix byte = runtime.Opcodes[runtime.QIP]

			runtime.AdvanceNOpcode(1)

			switch runtime.Opcodes[runtime.QIP] {
			case byte(PUSH):
				if err := runtime.PushExtended(prefix); err != nil {
					panic(err)
				}
			case byte(COMP):
				if err := runtime.CompExtended(prefix); err != nil {
					panic(err)
				}
			case byte(SET):
				if err := runtime.SetExtended(prefix); err != nil {
					panic(err)
				}
			}
		case byte(COMP):
			runtime.Comp()
		case byte(EQ):
			runtime.Eq()
		case byte(NEQ):
			runtime.Neq()
		case byte(RES):
			if err := runtime.Res(); err != nil {
				panic(err)
			}
		default:
			panic("Unknow opcode")
		}

	}

	return runtime
}
