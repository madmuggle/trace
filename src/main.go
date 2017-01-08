package main

import (
	"./icmplisten"
	"fmt"
	"net"
	"os"
	"syscall"
	"time"
)

func sendSomething() {
	addr, err := net.ResolveUDPAddr("udp", "47.88.20.73:80")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	defer conn.Close()

	f, err := conn.File()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed conn.File")
	}
	// You can close the new created fd
	defer f.Close()

	fd := int(f.Fd())
	err = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_TTL, 2)

	// Wait for the listening socket to get ready
	time.Sleep(1 * time.Second)
	conn.Write([]byte{1, 2, 3})
}

func main() {
	sock, err := icmplisten.New("a")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed creating listener:", err.Error())
		return
	}
	defer sock.Close()

	go sendSomething()

	for {
		data, _ := sock.Read()
		fmt.Println("data:", data)
	}
}
