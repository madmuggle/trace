package icmpsend

import (
	"../checksum"
	"net"
	"syscall"
)

func setTTL(conn *net.IPConn, ttl int) error {
	// Get the file descriptor copy of the connection.
	f, err := conn.File()
	if err != nil {
		return err
	}
	// You can close the new created fd
	defer f.Close()

	fd := int(f.Fd())
	err = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_TTL, ttl)
	return err
}

func reqdata() []byte {
	msg := []byte{8, 0, 0, 0, 0, 13, 0, 37, 99}

	// Calculate checksum and fill it into buffer
	chk := checksum.Calc(msg)
	msg[2] = byte(chk >> 8)
	msg[3] = byte(chk)

	return msg[:]
}

func Send(target *net.IPAddr, ttl int) error {
	conn, err := net.DialIP("ip4:1", nil, target)
	if err != nil {
		return err
	}

	setTTL(conn, ttl)
	_, err = conn.Write(reqdata())
	return err
}
