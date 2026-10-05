package main

const (
	A_REG     = 0
	B_REG     = 1
	C_REG     = 2
	X_REG     = 3
	Y_REG     = 4
	Z_REG     = 5
	I_REG     = 6
	J_REG     = 7
	PC_REG    = 8
	INT_REG   = 9
	mem_total = 65553
	reg_count = 10
)

type virtcpuState struct {
	memory    [65553]uint16
	registers [10]uint16
}

func loadProgram(state virtcpuState) {

}

func stepProcessor(state virtcpuState) {

}
