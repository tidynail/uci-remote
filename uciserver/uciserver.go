package main

import (
	"fmt"
	"io"
	"path/filepath"
	"log"
	"net"
	"os"
	"os/exec"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handle(conn net.Conn, id int, engine string) {
	log.Printf("[%d] New connection from %s\n", id, conn.RemoteAddr())
	p := exec.Command(engine)
	stdin, err := p.StdinPipe()
	check(err)
	stdout, err := p.StdoutPipe()
	check(err)
	check(p.Start())

	go func() {
		for {
			buff := make([]byte, 2048)
			n, err := conn.Read(buff)
			if err != nil {
				log.Printf("[%d] Connection lost from %s\n", id, conn.RemoteAddr())
				conn.Close()
				p.Process.Kill()
				break
			}
			stdin.Write(buff[:n])
		}
	}()

	go io.Copy(conn, stdout)
}

func main() {
	exepath := os.Args[0]
  exebase := exepath[:len(exepath)-len(filepath.Ext(exepath))]
	exefname := filepath.Base(exebase)

	if len(os.Args) < 3 {
		fmt.Printf("%s [ip]:port /path/to/engine\n", exefname)
		fmt.Printf("  ex. %s :7900 stockfish\n", exefname)
		os.Exit(1)
	}

	addr := os.Args[1]
	engine, err := filepath.Abs(os.Args[2])
	check(err)

	log.Printf("Listening on %s\n", addr)
	log.Printf(" for engine %s\n", engine)
	listener, err := net.Listen("tcp", addr)
	check(err)

	id := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
		  log.Print(err)
			continue
		}

		go handle(conn, id, engine)
		id++
	}

	os.Exit(0)
}
