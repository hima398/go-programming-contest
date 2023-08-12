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

	n, m := nextInt(), nextInt()
	h := nextIntSlice(n)
	w := nextIntSlice(m)

	ans := solve(n, m, h, w)

	PrintInt(ans)
}

func solve(n, m int, h, w []int) int {
	const INF = 1 << 60

	sort.Ints(h)

	sl := []int{0}
	var sr []int
	for i := 1; i < n; i += 2 {
		sl = append(sl, h[i]-h[i-1])
	}
	for i := 2; i < n; i += 2 {
		sr = append(sr, h[i]-h[i-1])
	}
	sr = append(sr, 0)
	for i := 1; i < len(sl); i++ {
		sl[i] += sl[i-1]
	}
	for i := len(sr) - 2; i >= 0; i-- {
		sr[i] += sr[i+1]
	}

	ans := INF
	for _, wi := range w {
		idx := sort.Search(len(h), func(i int) bool {
			return h[i] >= wi
		})
		s := sl[idx/2] + Abs(wi-h[idx/2*2]) + sr[idx/2]
		ans = Min(ans, s)
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
