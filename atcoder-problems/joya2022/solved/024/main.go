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

	h, w := nextInt(), nextInt()
	var s []string
	for i := 0; i < h; i++ {
		s = append(s, nextString())
	}
	var p [][2]int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == 'o' {
				p = append(p, [2]int{i, j})
			}
		}
	}
	computeDist := func(i1, j1, i2, j2 int) int {
		return Abs(i2-i1) + Abs(j2-j1)
	}
	ans := computeDist(p[0][0], p[0][1], p[1][0], p[1][1])
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
