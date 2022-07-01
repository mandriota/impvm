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

const StackLen = 1 << 20

type Kind = int64

type Stack struct {
	data [StackLen]Kind

	headP, tailP uint64
}

func (s *Stack) Reset() {
	s.headP = StackLen / 2
	s.tailP = StackLen / 2
}

func (s *Stack) GetHead(i uint64) *Kind {
	return &s.data[(s.headP-i)%StackLen]
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
