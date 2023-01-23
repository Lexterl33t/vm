package emulator

type Runtime struct {
	Registers map[Register]int
	QIP       int
	Stack     *Stack
	Opcodes   []byte
}

func NewRuntime(opcodes []byte) *Runtime {
	return &Runtime{
		Registers: make(map[Register]int),
		QIP:       0,
		Stack:     NewStack(),
		Opcodes:   opcodes,
	}
}

func (runtime *Runtime) NextOpcode() {
	runtime.QIP += 8
}

func (runtime *Runtime) AdvanceNOpcode(n int) {
	runtime.QIP += n
}

func Exec(opcodes []byte) map[Register]int {
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
		case byte(COMP):
			runtime.Comp()
		case byte(EQ):
			runtime.Eq()
		case byte(NEQ):
			runtime.Neq()
		default:
			panic("Unknow opcode")
		}

	}

	return runtime.Registers
}
