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

func replaceSpecialCharactersInURI(in string) string {
	if strings.Contains(in, "uri=\"") {
		first := strings.Index(in, "uri=\"") + len("uri=\"")
		last := strings.Index(in, "\",method=") // TODO this is very bad!!

		toChange := in[first:last]
		new := ""
		for _, char := range toChange {
			if char == '=' {
				new += "\\="
			} else {
				new += string(char)
			}
		}
		in = strings.Replace(in, toChange, new, 1)
	}
	return in
}

func report(str string) string {
	parts := strings.Split(str, " ")
	lastIdx := len(parts) - 1
	layout := time.RFC3339
	t, err := time.Parse(layout, parts[lastIdx])
	if err != nil {
		fmt.Errorf("error")
	}
	parts[lastIdx] = strconv.FormatInt(t.UnixNano(), 10)
	res := strings.Join(parts, " ")
	if len(res) < 22 {
		return ""
	}
	return replaceSpecialCharactersInURI(res)
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
				ss := report(part)
				if len(ss) > 0 {
					fmt.Println(ss)
				}
			}
			if len(parts[len(parts)-1]) == len(parts[0]) {
				ss := report(string(parts[len(parts)-1]))
				if len(ss) > 0 {
					fmt.Println(ss)
				}
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
