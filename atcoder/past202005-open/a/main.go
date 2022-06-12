package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s, t := nextString(), nextString()

	if s == t {
		fmt.Println("same")
	} else if strings.ToLower(s) == strings.ToLower(t) {
		fmt.Println("case-insensitive")
	} else {
		fmt.Println("different")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
