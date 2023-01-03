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

	n, m, t := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n - 1)
	var x, y []int
	for i := 0; i < m; i++ {
		x = append(x, nextInt()-1)
		y = append(y, nextInt())
	}
	mx := make(map[int]int)
	for i := range x {
		mx[x[i]] = y[i]
	}
	for i, ai := range a {
		t += mx[i]
		t -= ai
		if t <= 0 {
			PrintString("No")
			return
		}
	}
	PrintString("Yes")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
