Virtual Machine : IMPv0
=======================
*By Mark Mandriota*

INDEX
-----
- [Virtual Machine : IMPv0](#virtual-machine--impv0)
  - [INDEX](#index)
  - [Memory Model](#memory-model)
  - [Instruction Structure](#instruction-structure)
    - [Instruction Set](#instruction-set)
    - [LOOPF](#loopf)
    - [Combined Conditional Flags](#combined-conditional-flags)

<div class="page"/>

## Memory Model
| organization method | kind size |
| ------------------- | --------- |
| static stack        | 64-bit    |


## Instruction Structure
| CODE  | LOOPF | LF    | EF    | GF    |
| ----- | ----- | ----- | ----- | ----- |
| 4-bit | 1-bit | 1-bit | 1-bit | 1-bit |

*(total size: 8-bit)*

### Instruction Set
| code | name  | arg stack | ret stack | synopsis                    |
| ---- | ----- | --------- | --------- | --------------------------- |
| 0x0  | REDO  | mm        | mm; mm    | duplicate last              |
| 0x1  | UNDO  | mm        |           | delete last                 |
| 0x2  | HROT  | mm; mm    | mm; mm    | rotate stack down-up        |
| 0x3  | TROT  | mm; mm    | mm; mm    | rotate stack up-down        |
| 0x4  | SWAP  | mm; mm    | mm; mm    | exchange 2 elements         |
| 0x5  | JUMP  |           | mm        | jump to code position       |
| 0x6  | TEST  |           | mm        | compare values              |
| 0x7  | IPUT  |           | mm        | get in data                 |
| 0x8  | OPUT  | mm        |           | put out data                |
| 0x9  | LINC  | mm; mm    | mm        | `x+y`                       |
| 0xA  | LDEC  | mm; mm    | mm        | `x-y`                       |
| 0xB  | QINC  | mm; mm    | mm        | `x=x+k`                     |
| 0xC  | QDEC  | mm; mm    | mm        | `x=x-k`                     |
| 0xD  |       |           |           |                             |
| 0xE  |       |           |           |                             |
| 0xF  |       |           |           |                             |

### LOOPF
If flag is setted, then pops `N` from the stack and
  repeats instruction while `condf≠0 and k<N`

### Combined Conditional Flags
| code   | name | combination      | synopsis                    |
| ------ | ---- | ---------------- | --------------------------- |
| 0b100  | LF   | `LF`             | `x<y` (lower)               |
| 0b011  | NLF  | `GF`; `EF`       | `x≥y` (not lower)           |               
| 0b010  | EF   | `EF`             | `x=y` (equal)               |
| 0b101  | NEF  | `LF`; `GF`       | `x≠y` (not equal)           |
| 0b001  | GF   | `GF`             | `x>y` (greater)             |
| 0b110  | NGF  | `LF`; `EF`       | `x≤y` (not greater)         |
| 0b111  | AF   | `LF`; `GF`; `EF` | `1` (always)                |
| 0b000  | AF   |                  | `1` (always)                |