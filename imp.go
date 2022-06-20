package imp

import (
	"bufio"
	"fmt"
	"math"
)

const (
	FMODE = 1 << iota
	CMODE
)

type Registers struct {
	TextP Kind
	ModeP Kind
	Count Kind
	Accum Kind
	CondF uint8
}

type Machine struct {
	regs Registers

	w *bufio.Writer
	r *bufio.Reader

	Text []uint8
	Memo Stack
}

func (m *Machine) Compare() {
	switch el := *m.Memo.GetHead(); {
	case el < m.regs.Count:
		m.regs.CondF |= OP_FLAG_COND_L
	case el == m.regs.Count:
		m.regs.CondF |= OP_FLAG_COND_E
	case el > m.regs.Count:
		m.regs.CondF |= OP_FLAG_COND_G
	}
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
			el := *m.Memo.GetHead()

			for i := m.regs.Count; i > 0; i-- {
				m.Memo.AddHead(el)
			}
		case OP_CODE_UNDO:
			m.Memo.headP -= uint32(m.regs.Count)
		case OP_CODE_HSWP:
			panic("currently unsupported opcode.")
		case OP_CODE_TSWP:
			panic("currently unsupported opcode.")
		case OP_CODE_HROT:
			for i := m.regs.Count; i > 0; i-- {
				m.Memo.AddTail(m.Memo.PopHead())
			}
		case OP_CODE_TROT:
			for i := m.regs.Count; i > 0; i-- {
				m.Memo.AddHead(m.Memo.PopTail())
			}
		case OP_CODE_JTXT:
			m.regs.TextP = m.regs.Count
		case OP_CODE_JPAT:
			m.regs.ModeP = m.regs.Count
		case OP_CODE_OPUT:
			fmt.Fprintf(m.w, "oput: %d\n", m.Memo.PopHead())
		case OP_CODE_IPUT:
			m.Memo.AddHead(0)
			fmt.Fscanln(m.r, m.Memo.GetHead())
		case OP_CODE_LINC:
			if m.regs.ModeP&FMODE == FMODE {
				*GetFloat64Ptr(m.Memo.GetHead()) += *GetFloat64Ptr(&m.regs.Count)
			} else {
				*m.Memo.GetHead() += m.regs.Count
			}
		case OP_CODE_LDEC:
			if m.regs.ModeP&FMODE == FMODE {
				*GetFloat64Ptr(m.Memo.GetHead()) -= *GetFloat64Ptr(&m.regs.Count)
			} else {
				*m.Memo.GetHead() -= m.regs.Count
			}
		case OP_CODE_QINC:
			if m.regs.ModeP&FMODE == FMODE {
				*GetFloat64Ptr(m.Memo.GetHead()) *= *GetFloat64Ptr(&m.regs.Count)
			} else {
				*m.Memo.GetHead() *= m.regs.Count
			}
		case OP_CODE_QDEC:
			if m.regs.ModeP&FMODE == FMODE {
				*GetFloat64Ptr(m.Memo.GetHead()) /= *GetFloat64Ptr(&m.regs.Count)
			} else {
				*m.Memo.GetHead() /= m.regs.Count
			}
		case OP_CODE_CINC:
			dst := GetFloat64Ptr(m.Memo.GetHead())
			*dst = math.Pow(*dst, *GetFloat64Ptr(&m.regs.Count))
		case OP_CODE_CDEC:
			dst := GetFloat64Ptr(m.Memo.GetHead())
			*dst = math.Log(*dst) / math.Log(*GetFloat64Ptr(&m.regs.Count))
		}
	}
	m.w.Flush()
}
