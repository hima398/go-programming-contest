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

	_ = nextInt()
	c := nextString()
	var nw, nr int
	for i := range c {
		switch c[i] {
		case 'W':
			nw++
		case 'R':
			nr++
		}
	}
	ans := Min(nw, nr)
	var cnt int
	for i := 0; i < nr; i++ {
		if c[i] == 'W' {
			cnt++
		}
	}
	ans = Min(ans, cnt)
	PrintInt(ans)
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
