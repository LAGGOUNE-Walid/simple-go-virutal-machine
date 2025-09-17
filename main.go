package main

import "fmt"

const MAX_REGISTERS = 16

// 16-bit instruction format:
// [ 4 bits opcode ][ 4 bits dest ][ 4 bits src1 ][ 4 bits src2/imm ]

const (
	OP_HALT uint16 = iota
	OP_LOAD
	OP_ADD
	OP_PRINT
)

type VM struct {
	pc      int
	program []uint16
	regs    [MAX_REGISTERS]uint8
	stopped bool
}

func (vm *VM) run() {
	instr := vm.program[vm.pc]
	fmt.Printf("Instruction %#x \n", instr)
	op := (instr & 0xF000) >> 12

	switch op {
	case OP_HALT:
		fmt.Println("Operation halt")
		vm.stopped = true
	case OP_LOAD:

		reg := (instr & 0x0F00) >> 8
		if reg >= MAX_REGISTERS {
			fmt.Printf("register out of bound")
			vm.stopped = true
		}
		imm := (instr & 0x00FF)
		fmt.Printf("Loading %d to reg %d \n", imm, reg)
		vm.regs[reg] = uint8(imm)
	case OP_ADD:
		dest := (instr & 0x0F00) >> 8
		source_right := (instr & 0x00F0) >> 4
		source_left := (instr & 0x000F)
		if dest >= MAX_REGISTERS || source_right >= MAX_REGISTERS || source_left >= MAX_REGISTERS {
			fmt.Printf("register out of bound")
			vm.stopped = true
		}
		fmt.Printf("Adding %d + %d to reg %d \n", vm.regs[source_right], vm.regs[source_left], dest)
		vm.regs[dest] = vm.regs[source_right] + vm.regs[source_left]
	case OP_PRINT:
		reg := (instr & 0x0F00) >> 8
		if reg >= MAX_REGISTERS {
			fmt.Printf("register out of bound")
			vm.stopped = true
		}
		fmt.Printf("Printing register %d \n", reg)
		fmt.Printf("%d \n", vm.regs[reg])
	default:
		fmt.Printf("Unknown opcode: %#x\n", op)
		vm.stopped = true
	}

	vm.pc++
}

func main() {
	vm := VM{program: []uint16{0x1064, 0x1164, 0x2201, 0x3200, 0x0000}, stopped: false}
	for !vm.stopped {
		vm.run()
	}
}
