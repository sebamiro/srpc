package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type Cards struct {
	Spade, Clubs int
}
type Data struct {
	Name  string
	Cards Cards
}

func main() {
	l, err := net.Listen("tcp", ":8421")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer l.Close()
	c, err := l.Accept()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadBytes('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting")
			return
		}

		Data := &Data{}

		err = json.Unmarshal(netData, Data)
		if err != nil {
			fmt.Println("->", err)
		} else {
			fmt.Println("->", Data)
		}
		fmt.Println("->", string(netData))
		t := time.Now()
		c.Write([]byte(t.Format(time.RFC3339) + "\n"))
	}
}
