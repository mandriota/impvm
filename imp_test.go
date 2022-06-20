package imp

import (
	"bufio"
	"math"
	"os"
	"testing"
)

func TestMachineRun(t *testing.T) {
	m := &Machine{
		w:    bufio.NewWriter(os.Stdout),
		r:    bufio.NewReader(os.Stdin),
		Memo: *NewStack(),
		Text: []uint8{
			OP_CODE_IPUT,
			OP_CODE_QINC | OP_FLAG_LOOP,
			OP_CODE_OPUT,
		},
	}

	m.Memo.AddHead(2)

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
