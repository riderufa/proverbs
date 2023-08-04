package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

const addr = "0.0.0.0:12345"

const proto = "tcp4"

var proverbs [19]string

func main() {
	fillProverbs()
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		conn.Write([]byte(proverbs[rand.Intn(18)] + "\n"))
		time.Sleep(3 * time.Second)
	}
}

func fillProverbs() {
	f, err := os.OpenFile("./proverbs.txt", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("Такого файла не существует")
		return
	}
	defer f.Close()
	fileReader := bufio.NewReader(f)
	for i := 0; i < 19; i++ {
		line, _, err := fileReader.ReadLine()
		if err != nil {
			break
		}
		proverbs[i] = string(line)
	}
}
