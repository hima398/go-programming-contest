package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	s := nextString()
	var x []int
	var c []string
	for i := 0; i < q; i++ {
		x, c = append(x, nextInt()-1), append(c, nextString())
	}

	ans := solve(n, q, s, x, c)

	for _, v := range ans {
		Print(v)
	}
}

func solve(n, q int, s string, x []int, c []string) []int {
	t := strings.Split(s, "")
	var cur int
	for i := 1; i < n-1; i++ {
		if t[i-1] == "A" && t[i] == "B" && t[i+1] == "C" {
			cur++
		}
	}
	var ans []int
	for k := 0; k < q; k++ {
		var prev int
		for i := Max(x[k]-2, 0); i <= Min(x[k], n-3); i++ {
			if t[i] == "A" && t[i+1] == "B" && t[i+2] == "C" {
				prev++
			}
		}
		t[x[k]] = c[k]
		var next int
		for i := Max(x[k]-2, 0); i <= Min(x[k], n-3); i++ {
			if t[i] == "A" && t[i+1] == "B" && t[i+2] == "C" {
				next++
			}
		}
		cur = cur - prev + next
		ans = append(ans, cur)
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
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
