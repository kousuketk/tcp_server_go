package main

import (
	"io"
	"log"
	"net"
	"os"
)

func echo_handler(conn net.Conn) {
	defer conn.Close()
	w := io.MultiWriter(os.Stdout, conn)
	io.Copy(w, conn)
	os.Stdout.Write([]byte("\n"))
}

func main() {
	psock, e := net.Listen("tcp", ":5000")
	if e != nil {
		log.Fatal(e)
		return
	}
	for {
		conn, e := psock.Accept()
		if e != nil {
			log.Fatal(e)
			return
		}
		go echo_handler(conn)
	}
}
