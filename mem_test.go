package imp

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	for i := int64(0); i < 100; i++ {
		stack.AddHead(Kind{Let: i})
	}

	for i := 0; i < 34; i++ {
		stack.AddTail(stack.PopHead())
	}

	for i := 0; i < 100; i++ {
		t.Log(stack.PopTail())
	}
}

func TestStackDupHead(t *testing.T) {
	stack := NewStack()
	stack.AddHead(Kind{Typ: 8, Let: 13})

	for i := 0; i < 3; i++ {
		stack.DupHead()
	}

	for i := 0; i < 4; i++ {
		t.Log(stack.PopTail())
	}
}

func BenchmarkAddHead(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		stack.AddHead(Kind{Typ: 8, Let: 13})
	}
}

func BenchmarkAddTail(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		stack.AddTail(Kind{Typ: 8, Let: 13})
	}
}

func BenchmarkPopHead(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		stack.PopHead()
	}
}

func BenchmarkPopTail(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		stack.PopTail()
	}
}

func BenchmarkAddHeadPopHead(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		stack.AddHead(Kind{Typ: 8, Let: 13})
		stack.PopHead()
	}
}

func BenchmarkAddHeadPopTail(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		stack.AddHead(Kind{Typ: 8, Let: 13})
		stack.PopTail()
	}
}

func BenchmarkAddTailPopHead(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		stack.AddTail(Kind{Typ: 8, Let: 13})
		stack.PopHead()
	}
}

func BenchmarkAddTailPopTail(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		stack.AddTail(Kind{Typ: 8, Let: 13})
		stack.PopTail()
	}
}
func BenchmarkPopHeadAddHeadAddHead(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		el := stack.PopHead()
		stack.AddHead(el)
		stack.AddHead(el)
	}
}

func BenchmarkDupHead(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		stack.DupHead()
	}
}

func BenchmarkSwpHead(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		stack.SwpHead()
	}
}
