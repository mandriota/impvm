package imp

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	stack.AddHead(13)

	t.Log(stack.PopHead())
	t.Log(stack.PopHead())
}

func BenchmarkAddHead(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		stack.AddHead(13)
	}
}

func BenchmarkAddTail(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		stack.AddTail(13)
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
		stack.AddHead(13)
		stack.PopHead()
	}
}

func BenchmarkAddHeadPopTail(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		stack.AddHead(13)
		stack.PopTail()
	}
}

func BenchmarkAddTailPopHead(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		stack.AddTail(13)
		stack.PopHead()
	}
}

func BenchmarkAddTailPopTail(b *testing.B) {
	stack := NewStack()

	for i := 0; i < b.N; i++ {
		stack.AddTail(13)
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
