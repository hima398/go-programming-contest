package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	var s []string
	for i := 0; i < 10; i++ {
		s = append(s, nextString())
	}
	var a, b, c, d int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if s[i][j] == '#' {
				b, d = i+1, j+1
			}
		}
	}
	for i := 9; i >= 0; i-- {
		for j := 9; j >= 0; j-- {
			if s[i][j] == '#' {
				a, c = i+1, j+1
			}
		}
	}
	fmt.Println(a, b)
	fmt.Println(c, d)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
