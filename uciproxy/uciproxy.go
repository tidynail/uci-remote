package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"io"
	"log"
	"net"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func equals(s string, b []byte) bool {
	bs := []byte(s)
	if len(bs) > len(b) {
		return false
	}
	for i, x := range bs {
		if x != b[i] {
			return false
		}
	}
	return true
}

func read_cfg(path string, addr *string) bool {
  dat, err := os.ReadFile(path)
	if err != nil {
	  return false
	}

	*addr = strings.TrimSpace(string(dat))
	return true
}

func main() {
	exepath := os.Args[0]
  exebase := exepath[:len(exepath)-len(filepath.Ext(exepath))]
	exefname := filepath.Base(exebase)

	var addr string
	if len(os.Args) < 2 {
    if !read_cfg(exebase + ".txt", &addr) {
			fmt.Printf("%s <ip:port>\n", exefname)
			fmt.Printf("  ex. %s 127.0.0.1:7979\n", exefname)
			fmt.Printf("\nto save the argument,\n")
			fmt.Printf("  make %s.txt containing <ip:port>\n", exefname)
			fmt.Printf("    ex. %s.txt\n", exefname)
			fmt.Printf("        127.0.0.1:7979\n")
			fmt.Printf("\nwhen needing multiple proxies,\n")
			fmt.Printf("  rename executable and use <rename>.txt\n")
			os.Exit(1)
		}
	}	else {
		addr = os.Args[1]
	}

	conn, err := net.Dial("tcp", addr)
	check(err)

	go func() {
		for {
			buff := make([]byte, 2048)
			n, err := os.Stdin.Read(buff)
			check(err)

			conn.Write(buff[:n])

			if equals("quit", buff) {
				os.Exit(0)
			}
		}
	}()

	io.Copy(os.Stdout, conn)
}
