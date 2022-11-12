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

	n, c := nextInt(), nextString()
	a := nextString()
	ans := solve(n, c, a)
	PrintString(ans)
}

func solve(n int, c string, a string) string {
	m := map[byte]int{'W': 0, 'B': 1, 'R': 2}
	s := 0
	for i := range a {
		s += m[a[i]]
	}
	if s%3 == m[c[0]] {
		return "Yes"
	} else {
		return "No"
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
