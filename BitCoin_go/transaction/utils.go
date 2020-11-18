package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// IntToHex 将一个Int64转化成一个字节数组(byte array)
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.LittleEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
