package main

// julian daffa dzaky 2602176465
import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func menu() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("1. Kirim pesan ke server\n2. Keluar\n>>")
		scanner.Scan()
		s := scanner.Text()
		if s == "1" {
			TulisPesan()
		} else if s == "2" {
			fmt.Println("Terima kasih telah memakai program ini")
			break
		}

	}
}

func TulisPesan() {
	scanner := bufio.NewScanner(os.Stdin)
	var pesan string
	for {
		fmt.Print("masukkan pesan: ")
		scanner.Scan()
		pesan = scanner.Text()
		if len(pesan) < 1 {
			fmt.Println("Pesan minimal 1 karakter")
		} else if strings.Contains(pesan, "kasar") {
			fmt.Println("pesan tidak boleh mengandung kata kasar")
		} else {
			break
		}
	}
	KirimPesan(pesan)
}

func KirimPesan(pesan string) {
	serverConn, err := net.DialTimeout("tcp", "127.0.0.1:1234", 3*time.Second)
	if err != nil {
		panic(err)
	}
	defer serverConn.Close()

	err = serverConn.SetWriteDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		fmt.Println("Error setting write deadline:", err)
	}

	err = binary.Write(serverConn, binary.LittleEndian, uint32(len(pesan)))
	if err != nil {
		panic(err)
	}

	err = serverConn.SetWriteDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		fmt.Println("Error setting write deadline:", err)
	}

	_, err = serverConn.Write([]byte(pesan))
	if err != nil {
		panic(err)
	}

	err = serverConn.SetReadDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		fmt.Println("Error setting read deadline:", err)
	}

	var size uint32
	err = binary.Read(serverConn, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}

	err = serverConn.SetReadDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		fmt.Println("Error setting read deadline:", err)
	}

	bytReply := make([]byte, size)
	_, err = serverConn.Read(bytReply)
	if err != nil {
		panic(err)
	}
	fmt.Printf("replied: %s\n", string(bytReply))
}

func main() {
	menu()
}
