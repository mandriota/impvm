package imp

import (
	"math"
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

func TestMathModf(t *testing.T) {
	const FMASK = ^uint64(0) >> 12
	const EMASK = ^FMASK &^ (1 << 63)

	bits := math.Float64bits(1. / 5)

	f := bits & FMASK
	e := (bits & EMASK) >> 52

	t.Logf("%d %d", f, e)

	fnum := math.Pow(2, float64(e-1023))

	t.Logf("%f", fnum)
}
