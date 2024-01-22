package main

import (
	log "github.com/sirupsen/logrus"
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

	log.Info("Starting server...")
	server, err := net.ListenPacket(SERVER_TYPE, address)
	if err != nil {
		log.Error("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	log.Info("Listening on " + address)
	log.Info("Waiting for client...")

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
	currentTime := time.Now().Format(time.StampMilli)

	responseStr := string(buf) + ";" + currentTime
	responseStr = strings.ReplaceAll(responseStr, "\x00", "")
	log.Debug("Received a message with timestamp, replying:", responseStr)

	udpServer.WriteTo([]byte(responseStr), addr)
}
