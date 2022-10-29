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

func main() {
	exepath := os.Args[0]
  exebase := exepath[:len(exepath)-len(filepath.Ext(exepath))]
	exefname := filepath.Base(exebase)

	if len(os.Args) < 3 {
		fmt.Printf("%s [ip]:port /path/to/engine\n", exefname)
		fmt.Printf("  ex. %s :7979 stockfish\n", exefname)
		os.Exit(1)
	}

	addr := os.Args[1]
	engine, err := filepath.Abs(os.Args[2])
	check(err)

	log.Printf("Listening on %s\n", addr)
	log.Printf("Using engine %s\n", engine)
	l, err := net.Listen("tcp", addr)
	check(err)
	for {
		conn, err := l.Accept()
		check(err)
		log.Printf("New connection from %s\n", conn.RemoteAddr())
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
					log.Printf("Connection lost from %s\n", conn.RemoteAddr())
					conn.Close()
					p.Process.Kill()
					break
				}
				stdin.Write(buff[:n])
			}
		}()

		go io.Copy(conn, stdout)
		if err := p.Wait(); err != nil {
			log.Printf("Engine was killed due to: %s\n", err.Error())
		}
	}
}
