# Tiny Register-Based Virtual Machine in Go
This project is a minimal register-based virtual machine (VM) written in Go.
It simulates a tiny CPU with 16 registers  of 8 bits and a 16-bit instruction format.
# Features
- 16 general-purpose registers (R0–R15)
- 16-bit instructions with fixed layout:
```
[ 4 bits opcode ][ 4 bits dest ][ 4 bits src1 ][ 4 bits src2/imm ]
Binary: 0010 0010 0000 0001
         |    |    |    |
         |    |    |    └── src2 = 1 (R1)
         |    |    └────── src1 = 0 (R0)
         |    └─────────── dest = 2 (R2)
         └──────────────── opcode = 2 (ADD)
```
- Supported instructions: HALT, LOAD, ADD, PRINT

# Example program 
```go
vm := VM{
    program: []uint16{
        0x1064, // LOAD R0, 100
        0x1164, // LOAD R1, 100
        0x2201, // ADD  R2 = R0 + R1
        0x3200, // PRINT R2
        0x0000, // HALT
    },
    stopped: false,
}
for !vm.stopped {
    vm.run()
}
```
Output 
```
Instruction 0x1064
Loading 100 to reg 1
Instruction 0x1164
Loading 100 to reg 1
Instruction 0x2201
Adding 100 + 100 to reg 2
Instruction 0x3200
Printing register 2
200
Operation halt
```
# What I Learned
- How virtual machines work at a low level.
- Difference between stack-based and register-based VMs.
- How to design an instruction format using bit-fields.
- The role of endianness and immediate values in instruction encoding.
