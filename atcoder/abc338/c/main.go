package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	q := nextIntSlice(n)
	a := nextIntSlice(n)
	b := nextIntSlice(n)

	ans := solve(n, q, a, b)

	Print(ans)
}

func solve(n int, q, a, b []int) int {
	//料理Aだけを作れる最大人数分
	maxA := math.MaxInt
	for i := 0; i < n; i++ {
		if a[i] == 0 {
			continue
		}
		maxA = Min(maxA, q[i]/a[i])
	}
	//fmt.Println("maxA = ", maxA)
	var ans int
	for j := 0; j <= maxA; j++ {
		var rem []int
		for i := 0; i < n; i++ {
			rem = append(rem, q[i]-j*a[i])
		}
		maxB := math.MaxInt
		for i := 0; i < n; i++ {
			if b[i] == 0 {
				continue
			}
			maxB = Min(maxB, rem[i]/b[i])
		}
		//fmt.Println("A = ", j, " B = ", maxB)
		ans = Max(ans, j+maxB)
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Print(x any) {
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
