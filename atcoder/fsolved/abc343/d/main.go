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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, t := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < t; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt())
	}

	ans := solve(n, t, a, b)

	for _, v := range ans {
		Print(v)
	}
}

func solve(n, t int, a, b []int) []int {
	s := make([]int, n)
	m := make(map[int]int)
	m[0] = n

	var ans []int
	for i := 0; i < t; i++ {
		m[s[a[i]]]--
		if m[s[a[i]]] == 0 {
			delete(m, s[a[i]])
		}
		s[a[i]] += b[i]
		m[s[a[i]]]++

		ans = append(ans, len(m))
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
