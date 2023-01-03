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

	a := nextIntSlice(3)
	d := nextIntSlice(3)
	x := nextInt()
	t1, t2 := a[0]+a[2], d[0]+d[2]
	tk := (x / t1) * a[1] * a[0]
	tk += a[1] * Min(x%t1, a[0])
	ao := (x / t2) * d[1] * d[0]
	ao += d[1] * Min(x%t2, d[0])

	//fmt.Println(tk, ao)
	if tk == ao {
		PrintString("Draw")
	} else if tk > ao {
		PrintString("Takahashi")
	} else {
		PrintString("Aoki")
	}
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
