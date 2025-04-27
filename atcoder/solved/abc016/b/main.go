package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b, c := nextInt(), nextInt(), nextInt()
	if a+b == c {
		if a-b == c {
			PrintString("?")
		} else {
			PrintString("+")
		}
	} else if a-b == c {
		PrintString("-")
	} else {
		PrintString("!")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
