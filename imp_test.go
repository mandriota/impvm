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

func BenchmarkMachineRun(b *testing.B) {
	m := NewMachine(os.Stdout, os.Stdin)

	m.Text = []uint8{
		OP_CODE_REDO,
		OP_CODE_REDO,
		OP_CODE_LINC,
		OP_CODE_SWAP,
		OP_CODE_UNDO,
	}

	for i := 0; i < b.N; i++ {
		m.regs.TextP = 0
		m.Run()
	}
}
