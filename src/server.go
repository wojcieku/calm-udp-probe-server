package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	SERVER_PORT = ":1501"
	SERVER_TYPE = "udp"
)

func main() {
	fmt.Println("Server Running...")
	server, err := net.ListenPacket(SERVER_TYPE, SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_PORT)
	fmt.Println("Waiting for client...")

	for {
		buf := make([]byte, 1024)
		_, addr, err := server.ReadFrom(buf)
		if err != nil {
			continue
		}
		go response(server, addr, buf)
	}
}

func response(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	currentTime := time.Now().Format(time.StampMilli)

	responseStr := string(currentTime) + ";" + string(buf)
	fmt.Println("Received a message with timestamp, replying:", responseStr)

	udpServer.WriteTo([]byte(responseStr), addr)
}
