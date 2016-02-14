package main

import (
	"bufio"
	"fmt"
	"os"
	//"encoding/hex" //not sure if I should even use this
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

	realBytes := []byte(stringBytes)

	twoByte := getXbytePair(realBytes, 0)

	fmt.Printf("this is 1 + 2: %x\n", twoByte)

	if twoByte == 0x31fe {
		fmt.Println("now were in business")
	}

}

func getXbytePair(byteslice []byte, pos int) int {
	return int(byteslice[0*2])*256 + int(byteslice[1*2])
}
