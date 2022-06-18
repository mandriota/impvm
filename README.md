Virtual Machine : IMPv0 <satan>⛧</satan>
============================
*By Mark Mandriota*

INDEX
-----
- [Virtual Machine : IMPv0 <satan>⛧</satan>](#virtual-machine--impv0-satansatan)
  - [INDEX](#index)
  - [Memory Model](#memory-model)
    - [Type Conversion Convention Set](#type-conversion-convention-set)
  - [Instruction Structure](#instruction-structure)
    - [Instruction Set](#instruction-set)
    - [LOOPF](#loopf)
    - [Combined Conditional Flags](#combined-conditional-flags)

<div class="page"/>

## Memory Model
| organization method | pointer size | value size | type-descriptor size |
| ------------------- | ------------ | ---------- | -------------------- |
| linked stack        | 64-bit       | 64-bit     | 64-bit               |

<!--TODO: Создать универсальных дескриптор типа-->

### Type Conversion Convention Set
| code | name   | synopsis                                              |
| ---- | ------ | ----------------------------------------------------- |
| 0x0  |        | convert to no-type without bit-sequence changing      |
| 0x1  | REPTR  | convert to real pointer / unsigned integer type       |
| 0x2  | VIPTR  | convert to virtual pointer to external function type  |
| 0x3  | FLOAT  | convert to IEEE-754 64-bit floating-point number type | 
| 0x4  | DESYM  | set 32-bit decorated UTF-32 character type            |

<!--TODO: Добавить в таблицу конвенции о прямом преобразовании типов без изменения битовой последовательности-->

## Instruction Structure
| CODE  | LOOPF | LF    | EF    | GF    |
| ----- | ----- | ----- | ----- | ----- |
| 4-bit | 1-bit | 1-bit | 1-bit | 1-bit |
*(total size: 8-bit)*

### Instruction Set
| code | name | arg stack | ret stack | synopsis                    |
| ---- | ---- | --------- | --------- | --------------------------- |
| 0x0  | DUP  | mm        | mm; mm    | duplicate last              |
| 0x1  | DEL  | mm        |           | delete last                 |
| 0x2  | RTU  | mm; mm    | mm; mm    | rotate stack down-up        |
| 0x3  | RTD  | mm; mm    | mm; mm    | rotate stack up-down        |
| 0x4  | CDT  | mm        | mm        | `x→typec[k]`                |
| 0x5  | JMP  |           | mm        | `register.IP=k`             |
| 0x6  | GET  |           | mm        | get user input data         |
| 0x7  | PUT  | mm        | mm        | put user output data        |
| 0x8  | CSR  | mm; mm    | mm        | compare + set + rotate bits |
| 0x9  | SHF  | mm        | mm        | `x=x<<k`                    |
| 0xA  | XCH  | mm        | mm; mm    | exchange lasts              |
| 0xB  | ADD  | mm        | mm        | `x=x+k`                     |
| 0xC  | SUB  | mm        | mm        | `x=x-k`                     |
| 0xD  | MUL  | mm        | mm        | `x=x×k`                     |
| 0xE  | QUO  | mm        | mm; mm?   | `x=x÷k (x:int?→mod)`        |
| 0xF  | POW  | mm        | mm        | `x=x^k`                     |

### LOOPF
If flag is setted, then pops `N` onto the stack:
* If `N=0`, repeats instruction while `condf=1`
* If `N≠0`, repeats instruction while `condf≠0 and k<N`

### Combined Conditional Flags
| code   | name | combination      | bitf   | synopsis                    |
| ------ | ---- | ---------------- | -------| --------------------------- |
| 0b100  | LF   | `LF`             | `OR`   | `x<y` (lower)               |
| 0b011  | NLF  | `GF`; `EF`       | `NOR`  | `x≥y` (not lower)           |               
| 0b010  | EF   | `EF`             | `XOR`  | `x=y` (equal)               |
| 0b101  | NEF  | `LF`; `GF`       | `XNOR` | `x≠y` (not equal)           |
| 0b001  | GF   | `GF`             | `AND`  | `x>y` (greater)             |
| 0b110  | NGF  | `LF`; `EF`       | `NAND` | `x≤y` (not greater)         |
| 0b111  | AF   | `LF`; `GF`; `EF` | `SET`  | `1` (always)                |
| 0b000  | NF   |                  | `NSET` | `PREVIOUS_SUCCESS` (noways) |