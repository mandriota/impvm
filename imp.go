package imp

import (
	"bufio"
	"fmt"
	"io"
)

func NewMachine(w io.Writer, r io.Reader) *Machine {
	m := &Machine{
		w: bufio.NewWriter(w),
		r: bufio.NewReader(r),
	}
	m.Memo.Reset()

	return m
}

type Registers struct {
	TextP Kind
	Count Kind
	CondF uint8
}

type Machine struct {
	regs Registers

	w *bufio.Writer
	r *bufio.Reader

	Text []uint8
	Memo Stack
}

func (m *Machine) Run() {
	for m.regs.TextP < Kind(len(m.Text)) {
		cuop := m.Text[m.regs.TextP]

		m.regs.TextP++

		if cuop&m.regs.CondF != cuop&OP_MASK_FLAG_COND {
			continue
		}

		m.regs.Count = 1
		if cuop&OP_FLAG_LOOP == OP_FLAG_LOOP {
			m.regs.Count = m.Memo.PopHead()
		}

		switch cuop & OP_MASK_CODE {
		case OP_CODE_REDO:
			m.OpREDO()
		case OP_CODE_UNDO:
			m.OpUNDO()
		case OP_CODE_HROT:
			m.OpHROT()
		case OP_CODE_TROT:
			m.OpTROT()
		case OP_CODE_SWAP:
			m.OpSWAP()
		case OP_CODE_JUMP:
			m.OpJUMP()
		case OP_CODE_TEST:
			m.OpTEST()
		case OP_CODE_OPUT:
			m.OpOPUT()
		case OP_CODE_IPUT:
			m.OpIPUT()
		case OP_CODE_LINC:
			m.OpLINC()
		case OP_CODE_LDEC:
			m.OpLDEC()
		case OP_CODE_QINC:
			m.OpQINC()
		case OP_CODE_QDEC:
			m.OpQDEC()
		}
	}
	m.w.Flush()
}

func (m *Machine) OpREDO() {
	el := *m.Memo.GetHead(1)

	for i := m.regs.Count; i > 0; i-- {
		m.Memo.AddHead(el)
	}
}

func (m *Machine) OpUNDO() {
	m.Memo.headP -= uint64(m.regs.Count)
}

func (m *Machine) OpHROT() {
	for i := m.regs.Count; i > 0; i-- {
		m.Memo.AddTail(m.Memo.PopHead())
	}
}

func (m *Machine) OpTROT() {
	for i := m.regs.Count; i > 0; i-- {
		m.Memo.AddHead(m.Memo.PopTail())
	}
}

func (m *Machine) OpSWAP() {
	x := m.Memo.PopHead()
	m.Memo.AddHead(m.regs.Count)
	m.Memo.AddHead(x)
}

func (m *Machine) OpJUMP() {
	m.Memo.AddHead(m.regs.TextP)
	m.regs.TextP = m.regs.Count
}

func (m *Machine) OpTEST() {
	switch x := *m.Memo.GetHead(1); {
	case m.regs.Count < x:
		m.regs.CondF |= OP_FLAG_COND_L
	case m.regs.Count == x:
		m.regs.CondF |= OP_FLAG_COND_E
	case m.regs.Count > x:
		m.regs.CondF |= OP_FLAG_COND_G
	}
}

func (m *Machine) OpOPUT() {
	fmt.Fprintf(m.w, "oput: %d\n", m.Memo.PopHead())
}

func (m *Machine) OpIPUT() {
	println("iput called.")
	dst := m.Memo.PopHead()
	fmt.Fscan(m.r, &dst)
	m.Memo.AddHead(dst)
}

func (m *Machine) OpLINC() {
	x := m.Memo.PopHead()

	for i := m.regs.Count; i > 0; i-- {
		x += m.Memo.PopHead()
	}

	m.Memo.AddHead(x)
}

func (m *Machine) OpLDEC() {
	x := m.Memo.PopHead()

	for i := m.regs.Count; i > 0; i-- {
		x -= m.Memo.PopHead()
	}

	m.Memo.AddHead(x)
}

func (m *Machine) OpQINC() {
	x := m.Memo.PopHead()

	for i := m.regs.Count; i > 0; i-- {
		x *= m.Memo.PopHead()
	}

	m.Memo.AddHead(x)
}

func (m *Machine) OpQDEC() {
	x := m.Memo.PopHead()

	for i := m.regs.Count; i > 0; i-- {
		x /= m.Memo.PopHead()
	}

	m.Memo.AddHead(x)
}
