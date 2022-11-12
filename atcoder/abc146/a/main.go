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

	s := nextString()
	switch s {
	case "MON":
		PrintInt(6)
	case "TUE":
		PrintInt(5)
	case "WED":
		PrintInt(4)
	case "THU":
		PrintInt(3)
	case "FRI":
		PrintInt(2)
	case "SAT":
		PrintInt(1)
	case "SUN":
		PrintInt(7)
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
