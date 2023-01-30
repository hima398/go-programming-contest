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
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	ans := solve(n, x, y)
	PrintInt(ans)
}

func solve(n int, x, y []int) int {
	magic := make(map[int]map[int]struct{})
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			mx, my := x[j]-x[i], y[j]-y[i]
			gcd := Gcd(Abs(mx), Abs(my))
			mx /= gcd
			my /= gcd
			if magic[mx] == nil {
				magic[mx] = make(map[int]struct{})
			}
			magic[mx][my] = struct{}{}
		}
	}
	var ans int
	for _, v := range magic {
		ans += len(v)
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
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

func Gcd(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	/*
		if x < y {
			x, y = y, x
		}
	*/
	return Gcd(y, x%y)
}
