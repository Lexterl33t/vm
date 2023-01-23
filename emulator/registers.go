package emulator

type Register int

const (
	QOX Register = 0x20
	QCX Register = 0x21
	QIP Register = 0x22
	QSP Register = 0x23
	QZF Register = 0x24
)

var registers_table map[Register]string = map[Register]string{
	0x20: "QOX",
	0x21: "QCX",
	0x22: "QIP",
	0x23: "QSP",
	0x24: "QZF",
}

func (reg Register) String() string {
	return registers_table[reg]
}
