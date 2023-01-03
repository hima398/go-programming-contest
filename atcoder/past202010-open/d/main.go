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

	n := nextInt()
	s := nextString()
	x, y := solve(n, s)
	fmt.Println(x, y)
}

func solve(n int, s string) (x, y int) {
	d := 0
	var nd []int
	for i := 0; i < n; i++ {
		if s[i] == '#' {
			nd = append(nd, d)
			d = 0
		} else {
			d++
		}
	}
	nd = append(nd, d)
	x = nd[0]
	y = nd[len(nd)-1]
	mx := 0
	for i := 1; i < len(nd)-1; i++ {
		mx = Max(mx, nd[i]-(x+y))
	}
	x += mx
	return x, y
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
