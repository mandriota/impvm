// Copyright 2022 Mark Mandriota
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package impvm

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := new(Stack)
	stack.Reset()

	stack.AddHead(13)

	t.Log(stack.PopHead())
	t.Log(stack.PopHead())
}

func BenchmarkAddHead(b *testing.B) {
	stack := new(Stack)
	stack.Reset()

	for i := 0; i < b.N; i++ {
		stack.AddHead(13)
	}
}

func BenchmarkAddTail(b *testing.B) {
	stack := new(Stack)
	stack.Reset()

	for i := 0; i < b.N; i++ {
		stack.AddTail(13)
	}
}

func BenchmarkPopHead(b *testing.B) {
	stack := new(Stack)
	stack.Reset()

	for i := 0; i < b.N; i++ {
		stack.PopHead()
	}
}

func BenchmarkPopTail(b *testing.B) {
	stack := new(Stack)
	stack.Reset()

	for i := 0; i < b.N; i++ {
		stack.PopTail()
	}
}

func BenchmarkAddHeadPopHead(b *testing.B) {
	stack := new(Stack)
	stack.Reset()

	for i := 0; i < b.N; i++ {
		stack.AddHead(13)
		stack.PopHead()
	}
}

func BenchmarkAddHeadPopTail(b *testing.B) {
	stack := new(Stack)
	stack.Reset()

	for i := 0; i < b.N; i++ {
		stack.AddHead(13)
		stack.PopTail()
	}
}

func BenchmarkAddTailPopHead(b *testing.B) {
	stack := new(Stack)
	stack.Reset()

	for i := 0; i < b.N; i++ {
		stack.AddTail(13)
		stack.PopHead()
	}
}

func BenchmarkAddTailPopTail(b *testing.B) {
	stack := new(Stack)
	stack.Reset()

	for i := 0; i < b.N; i++ {
		stack.AddTail(13)
		stack.PopTail()
	}
}
func BenchmarkPopHeadAddHeadAddHead(b *testing.B) {
	stack := new(Stack)
	stack.Reset()

	for i := 0; i < b.N; i++ {
		el := stack.PopHead()
		stack.AddHead(el)
		stack.AddHead(el)
	}
}
