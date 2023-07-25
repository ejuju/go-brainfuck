# Simple Brainfuck interpreter / VM in Go

Instruction set
- `>`: go to next cell
- `<`: go to previous cell
- `+`: increment the byte at current cell
- `-`: decrement the byte at current cell
- `[`: jump after next `]` if current byte is zero, else move forward
- `]`: jump to previous `[` if current byte is not zero, else move forward
- `.`: write byte at current cell to stdout
- `,`: read one byte from stdin and set to current cell
