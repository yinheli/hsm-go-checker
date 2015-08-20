package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var (
	host = flag.String("h", "192.168.1.200", "HSM host")
	port = flag.Int("p", 8000, "HSM port")
)

func main() {
	flag.Parse()

	if *host == "" {
		flag.Usage()
		fmt.Println()
		return
	}

	hsmAddr := *host + ":" + strconv.Itoa(*port)
	fmt.Println("connecting...", hsmAddr)
	timeout := time.Second * 3
	conn, err := net.DialTimeout("tcp", hsmAddr, timeout)
	if err != nil {
		fmt.Println("connect error", err)
		return
	}

	conn.SetReadDeadline(time.Now().Add(timeout))
	conn.SetWriteDeadline(time.Now().Add(timeout))

	defer func() {
		conn.Close()
	}()

	data, _ := hex.DecodeString("00024E43") // send `NC` command
	_, err = conn.Write(data)

	if err != nil {
		fmt.Println("send data error", err)
		return
	}

	resp := make([]byte, 128)
	n, err := conn.Read(resp)
	if err != nil {
		fmt.Println("read data error", err)
		return
	}

	if strings.Contains(hex.EncodeToString(resp[:n]), "4e443030") {
		fmt.Println("OK")
	} else {
		fmt.Println("error")
	}

}
