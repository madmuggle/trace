package icmplisten

import (
	"fmt"
	"syscall"
)

// The ICMP Receive Buffer size, small but enough
const ICMPbufsize = 512

type Sock struct {
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
