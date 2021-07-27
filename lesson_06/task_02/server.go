package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Launching server...")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("Message received: ", message)
		buf := []rune(message)
		buf = buf[:len(buf)-1]
		val, err := strconv.Atoi(string(buf))
		if err != nil {
			conn.Write([]byte(strings.ToUpper(message) + "\n"))
		} else {
			conn.Write([]byte(strconv.Itoa(val*2) + "\n"))
		}
	}

}
