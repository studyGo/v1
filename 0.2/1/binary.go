package main

import (
	"bytes"
	"encoding/binary"
	"os"
)

func main() {
	file := "test"
	f, _ := os.Create(file)

	buffer := bytes.NewBuffer(make([]byte, 0, 1024))

	message := []byte("this is a demo")

	messageLen := uint32(len(message))

	mlen := 4 + len(message)

	binary.Write(buffer, binary.LittleEndian, messageLen)
	buffer.Write(message)

	f.Write(buffer.Next(mlen))
}
