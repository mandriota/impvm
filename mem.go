package imp

const StackLen = 1 << 20

type Kind = int64

type Stack struct {
	data [StackLen]Kind

	headP, tailP uint32
}

func NewStack() *Stack {
	return &Stack{
		headP: StackLen / 2,
		tailP: StackLen / 2,
	}
}

func (s *Stack) GetHead() *Kind {
	return &s.data[(s.headP-1)%StackLen]
}

func (s *Stack) AddHead(el Kind) {
	s.data[s.headP%StackLen] = el
	s.headP++
}

func (s *Stack) AddTail(el Kind) {
	s.data[(StackLen-1-s.tailP)%StackLen] = el
	s.tailP++
}

func (s *Stack) PopHead() Kind {
	s.headP--
	return s.data[s.headP%StackLen]
}

func (s *Stack) PopTail() Kind {
	s.tailP--
	return s.data[(StackLen-1-s.tailP)%StackLen]
}
