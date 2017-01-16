package main

import (
	"./icmplisten"
	"./icmpsend"
	"fmt"
	"net"
	"os"
	"time"
)

//var Hops = make([]

func getfirstIP(addr string) (*net.IPAddr, error) {
	addrs, err := net.LookupHost(addr)
	if err != nil {
		return nil, err
	}

	ip, err := net.ResolveIPAddr("ip", addrs[0])
	if err != nil {
		return nil, err
	}

	return ip, nil
}

func sendSomething() {
	// Wait for the listening socket to get ready
	time.Sleep(1 * time.Second)

	ip, err := getfirstIP("bing.com")
	if err != nil {
		fmt.Println("in sendSomething: %v", err)
		return
	}
	icmpsend.Send(ip, 2)
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
