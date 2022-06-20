package main

import (
	_ "embed"
	"os"

	imp "github.com/MarkMandriota/ImpVM"
)

func main() {
	m := imp.NewMachine(os.Stdout, os.Stdin)

	m.Memo = *imp.NewStack()
	m.Text = []uint8{
		imp.OP_CODE_IPUT,
		imp.OP_CODE_REDO,
		imp.OP_CODE_REDO,
		imp.OP_CODE_IPUT,
		imp.OP_CODE_SWAP,
		imp.OP_CODE_LDEC,
		imp.OP_CODE_REDO,
		imp.OP_CODE_HROT,
		imp.OP_CODE_REDO | imp.OP_FLAG_LOOP,
		imp.OP_CODE_TROT,
		imp.OP_CODE_QINC | imp.OP_FLAG_LOOP,
		imp.OP_CODE_OPUT,
		imp.OP_CODE_OPUT,
	}

	m.Run()
}
