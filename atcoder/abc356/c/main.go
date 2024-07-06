package main

import (
	"bufio"
	"fmt"
	"math/bits"
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

	n, m, k := nextInt(), nextInt(), nextInt()
	var c []int
	var a [][]int
	var r []string
	for i := 0; i < m; i++ {
		c = append(c, nextInt())
		a = append(a, nextIntSlice(c[i]))
		r = append(r, nextString())
	}

	ans := solve(n, m, k, c, a, r)

	Print(ans)
}

func solve(n, m, k int, c []int, a [][]int, r []string) int {
	l := 1 << n
	var ans int
	for pat := 0; pat < l; pat++ {
		ok := true
		for i := 0; i < m; i++ {
			var testPattern int
			for _, aij := range a[i] {
				testPattern |= 1 << (aij - 1)
			}
			if r[i] == "o" {
				//fmt.Printf("%v, %b, %b, %b, %v\n", r[i], pat, testPattern, pat&testPattern, bits.OnesCount(uint(pat&testPattern)))
				ok = ok && bits.OnesCount(uint(pat&testPattern)) >= k
			} else if r[i] == "x" {
				//fmt.Printf("%v, %b, %b, %b, %v\n", r[i], pat, testPattern, pat&testPattern, bits.OnesCount(uint(pat&testPattern)))
				ok = ok && bits.OnesCount(uint(pat&testPattern)) < k
			}
		}
		if ok {
			ans++
		}
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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
