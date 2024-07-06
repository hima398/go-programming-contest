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

	n := nextInt()
	a := nextIntSlice(n)

	ans := solve(n, a)

	Print(ans)
}

func computeWeight(x int) int {
	m := len(strconv.Itoa(x))
	res := 1
	for i := 0; i < m; i++ {
		res *= 10
	}
	return res
}

func solve(n int, a []int) int {
	const p = 998244353
	//a[n-1]からa[0]までの総和
	s := make([]int, n)
	s[n-1] = a[n-1]
	for i := n - 1; i >= 1; i-- {
		s[i-1] = (s[i] + a[i-1]) % p
	}

	var ans int
	for i := n - 1; i >= 1; i-- {
		w := computeWeight(a[i])
		ans += ((s[0] - s[i] + p) % p) * w
		ans %= p
		ans += a[i] * i
		ans %= p
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
