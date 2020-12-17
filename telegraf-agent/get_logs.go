package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	namespace = "nginx_request"
)

func report(str string) string {
	parts := strings.Split(str, " ")
	lastIdx := len(parts) - 1
	layout := time.RFC3339
	t, err := time.Parse(layout, parts[lastIdx])
	if err != nil {
		fmt.Errorf("error")
	}
	parts[lastIdx] = strconv.FormatInt(t.UnixNano(), 10)
	return strings.Join(parts, " ")
}

func main() {
	pattern := `<[0-9]*>[A-Z][a-z]* [0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9] ([a-z]*): `
	re := regexp.MustCompile(pattern)
	port, err := strconv.Atoi(os.Getenv("SYSLOG_SERVER_PORT"))
	p := make([]byte, 2048)

	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("*"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("error %v\n", err)
		return
	}
	last := make([]byte, 0)
	for {
		n, _, err := ser.ReadFromUDP(p)
		last = append(last, p[:n]...)
		for {
			matched, err := regexp.Match(pattern, last)
			if err != nil {
				fmt.Printf("error %v\n", err)
			}
			if !matched {
				break
			}
			parts := re.Split(string(last), -1)
			for _, part := range parts[:len(parts)-1] {
				fmt.Println(report(part))
			}
			if len(parts[len(parts)-1]) == len(parts[0]) {
				fmt.Println(report(string(parts[len(parts)-1])))
				last = make([]byte, 0)
			} else {
				last = []byte(parts[len(parts)-1])
			}
		}

		if err != nil {
			fmt.Printf("error  %v", err)
			continue
		}
	}
}
