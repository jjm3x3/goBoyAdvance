package main

import (
	"fmt"
	"os"
	"bufio"
	//"encoding/hex"
)

func check(e error){
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
	twoByte := int(realBytes[0])*256+int(realBytes[1])
	fmt.Printf("this is 1 + 2: %x\n", twoByte)

	if twoByte == 0x31fe {
		fmt.Println("now were in business")
	}

}
