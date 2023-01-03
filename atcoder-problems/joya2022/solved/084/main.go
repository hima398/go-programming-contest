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

	k := nextInt()

	ans := solve(k)

	PrintInt(ans)
}

func solve(k int) int {
	m := make(map[int]int)
	for i := 2; i*i <= k; i++ {
		for k%i == 0 {
			m[i]++
			k /= i
		}
	}
	ans := 1
	for p, x := range m {
		y := 0
		for x > 0 {
			y += p
			z := y
			for z%p == 0 {
				z /= p
				x--
			}
		}
		ans = Max(ans, y)
	}
	ans = Max(ans, k)

	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
