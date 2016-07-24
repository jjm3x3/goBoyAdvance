package main

import (
	"bufio"
	"fmt"
	"os"
	//"encoding/hex" //not sure if I should even use this
)

type Z80 struct {
	reg_sp int16
	reg_pc int
}

type SystemState struct {
	bootRom []byte
}

var (
	copc  Z80
	state SystemState
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("start my gBA run")

	initRom, err := os.Open("/home/jmeixner/Downloads/DMG_ROM.bin")
	check(err)

	bufferedReader := bufio.NewReader(initRom)
	stringBytes, err := bufferedReader.ReadString('P')
	check(err)

	state.bootRom = []byte(stringBytes)

	for copc.reg_pc = 0; copc.reg_pc < len(state.bootRom); copc.reg_pc++ {
		byte := getXByte(copc.reg_pc)
		execute(byte)
	}
}

func execute(byte int16) {
	if byte == 0x11 || byte == 0x21 || byte == 0x31 || byte == 0x41 {
		fmt.Printf("this is ld dd,nn (load imidate)")
		copc.reg_pc++
		secondPart := getXByte(copc.reg_pc)
		copc.reg_pc++
		firstPart := getXByte(copc.reg_pc)
		fmt.Printf("loading this address: %x %x\n", firstPart, secondPart)
	} else {
		fmt.Printf("this is: %x\n", byte)
	}
}

func getXByte(pos int) int16 {
	oneByte := int16(state.bootRom[pos])
	return oneByte
}
