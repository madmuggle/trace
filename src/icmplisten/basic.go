package icmplisten

import (
	"fmt"
	"syscall"
)

// The ICMP Receive Buffer size, no need to be big as icmp package is small.
const ICMPbufsize = 512

type Sock struct {
	// The buffer that stores ICMP packages. New data will override old.
	data [ICMPbufsize]byte
	sock int
	name string
}

func New(name string) (*Sock, error) {
	sock := Sock{}
	s, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		return nil, err
	}
	sock.name = name
	sock.sock = s
	return &sock, nil
}

func (s *Sock) Read() (data []byte, err error) {
	n, from, err := syscall.Recvfrom(s.sock, s.data[:], ICMPbufsize)
	if err != nil {
		return nil, err
	}
	fmt.Println("Peer:", from)
	return s.data[:n], nil
}

func (s *Sock) Close() {
	syscall.Close(s.sock)
}
