package imp

type Instruction = uint8

const (
	OP_FLAG_LOOP Instruction = 1 << 4 << iota
	OP_FLAG_COND_L
	OP_FLAG_COND_E
	OP_FLAG_COND_G

	OP_MASK_CODE = OP_FLAG_LOOP - 1
	OP_MASK_FLAG = ^OP_MASK_CODE
)

const (
	OP_CODE_REDO = iota & OP_MASK_CODE // redo head element k times
	OP_CODE_UNDO                       // undo head element k times
	OP_CODE_HSWP                       // swap head element with kth agreat element
	OP_CODE_TSWP                       // swap head element with kth alow element
	OP_CODE_HROT                       // rotate stack agreat-alow k times
	OP_CODE_TROT                       // rotate stack alow-agreat k times
	OP_CODE_JTYP                       // jump to next type position from 0 k times -> type_table[k]
	OP_CODE_JCOD                       // jump to next code position from 0 k times -> code_table[k]
	OP_CODE_IPUT                       // put data k times
	OP_CODE_OPUT                       // get data k times
	OP_CODE_AINC                       // increase head element (x) for 1 k times -> x+k
	OP_CODE_ADEC                       // decrease head element (x) for 1 k times -> x-k
	OP_CODE_AADD                       // add to head element (x) it self initial value k times -> x*k
	OP_CODE_ASUB                       // sub from head element (x) it self initial value and increase counter for 1 -> x/k
	OP_CODE_AMUL                       // multiply head element (x) it self initial value k times -> x^k
	OP_CODE_AQUO                       // -> log.k(x)
)
