package imp

import (
	"os"
	"testing"
)

func TestMachineRun(t *testing.T) {
	m := NewMachine(os.Stdout, os.Stdin)

	m.Text = []uint8{
		OP_CODE_IPUT,
		OP_CODE_REDO,
		OP_CODE_REDO,
		OP_CODE_IPUT,
		OP_CODE_SWAP,
		OP_CODE_LDEC,
		OP_CODE_REDO,
		OP_CODE_HROT,
		OP_CODE_REDO | OP_FLAG_LOOP,
		OP_CODE_TROT,
		OP_CODE_QINC | OP_FLAG_LOOP,
		OP_CODE_OPUT,
		OP_CODE_OPUT,
	}

	m.Run()
}

func BenchmarkMachineRun(b *testing.B) {
	m := NewMachine(os.Stdout, os.Stdin)

	m.Text = []uint8{
		OP_CODE_REDO,
		OP_CODE_REDO,
		OP_CODE_LINC,
		OP_CODE_SWAP,
		OP_CODE_UNDO,
	}

	for i := 0; i < b.N; i++ {
		m.regs.TextP = 0
		m.Run()
	}
}
