package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	SERVER_TYPE = "udp"
)

func startServer(serverPort int) {
	address := ":" + strconv.Itoa(serverPort)

	fmt.Println("Starting server...")
	server, err := net.ListenPacket(SERVER_TYPE, address)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + address)
	fmt.Println("Waiting for client...")

	for {
		buf := make([]byte, 50)
		_, addr, err := server.ReadFrom(buf)
		if err != nil {
			continue
		}
		go response(server, addr, buf)
	}
}

func response(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	//time.Sleep(time.Duration(rand.Intn(2)) * time.Millisecond)
	//flag := rand.Intn(7)
	//if flag == 1 {
	//	return
	//}
	currentTime := time.Now().Format(time.StampMilli)

	responseStr := string(buf) + ";" + currentTime
	responseStr = strings.ReplaceAll(responseStr, "\x00", "")
	//TODO clear unnecessary logs
	fmt.Println("Received a message with timestamp, replying:", responseStr)

	udpServer.WriteTo([]byte(responseStr), addr)
}
