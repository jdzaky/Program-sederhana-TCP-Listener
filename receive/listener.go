package main

// julian daffa dzaky 2602176465

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleServerConn(clientConn)
	}
}

func handleServerConn(client net.Conn) {
	var size uint32
	err := binary.Read(client, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}
	bytMsg := make([]byte, size)
	_, err = client.Read(bytMsg)
	if err != nil {
		panic(err)
	}
	strMsg := string(bytMsg)
	fmt.Printf("%s\n", strMsg)

	reply := "Pesan telah diterima"

	err = binary.Write(client, binary.LittleEndian, uint32(len(reply)))
	if err != nil {
		panic(err)
	}
	_, err = client.Write([]byte(reply))
	if err != nil {
		panic(err)
	}
}
