package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, h := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	ans := solve(n, h, a, b)
	Print(ans)
}

func solve(n, h int, a, b []int) int {
	var maxA int
	for _, ai := range a {
		maxA = Max(maxA, ai)
	}
	sort.Ints(b)

	var ans int
	for i := n - 1; i >= 0; i-- {
		if h <= 0 || maxA >= b[i] {
			break
		}
		h -= b[i]
		ans++
	}
	if h > 0 {
		ans += Ceil(h, maxA)
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
