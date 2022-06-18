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
	OP_CODE_SWAP                       // swap head element and k agreat element
	_                                  // ? reserved
	OP_CODE_LGRO                       // rotate stack alow-agreat k times
	OP_CODE_GLRO                       // rotate stack agreat-alow k times
	OP_CODE_TYPE                       // next type position from 0 k times -> type_table[k]
	OP_CODE_CODE                       // next code position from 0 k times -> code_table[k]
	OP_CODE_IPUT                       // put data k times
	OP_CODE_OPUT                       // get data k times
	OP_CODE_INCR                       // increase head element (x) for 1 k times -> x+k
	OP_CODE_DECR                       // decrease head element (x) for 1 k times -> x-k
	OP_CODE_ADD                        // add to head element (x) it self initial value k times -> x*k
	OP_CODE_SUB                        // sub from head element (x) it self initial value and increase counter for 1 -> x/k
	OP_CODE_MUL                        // multiply
	_                                  // ? reserved
)

/*
REDO 3 : 2

*/

// TODO: Rotate in area
