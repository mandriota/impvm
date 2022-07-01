package impvm

const (
	OP_FLAG_LOOP   = 1 << 4 << iota // sets k value from stack
	OP_FLAG_COND_L                  // will executed if is lower -> x<y
	OP_FLAG_COND_E                  // will executed if is equal -> x=y
	OP_FLAG_COND_G                  // will executed if is greater -> x>y
)

const (
	OP_MASK_CODE      = OP_FLAG_LOOP - 1
	OP_MASK_FLAG      = ^OP_MASK_CODE
	OP_MASK_FLAG_COND = OP_FLAG_COND_L | OP_FLAG_COND_E | OP_FLAG_COND_G
)

const (
	OP_CODE_REDO = iota & OP_MASK_CODE // redo head element k times
	OP_CODE_UNDO                       // undo head element k times
	OP_CODE_HROT                       // rotate stack agreat-alow k times
	OP_CODE_TROT                       // rotate stack alow-agreat k times
	OP_CODE_SWAP                       // swap heads elements
	OP_CODE_JUMP                       // jump to next text position from 0 k times -> code_table[k]
	OP_CODE_TEST                       // test head element and sets flags
	OP_CODE_IPUT                       // put data k times
	OP_CODE_OPUT                       // get data k times
	OP_CODE_LINC                       // linear increment -> x+k
	OP_CODE_LDEC                       // linear decrement -> x-k
	OP_CODE_QINC                       // quadratic increment -> x*k
	OP_CODE_QDEC                       // quadratic decrement with rest -> x/k
)
