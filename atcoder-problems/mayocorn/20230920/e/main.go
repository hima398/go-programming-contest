package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/liyue201/gostl/ds/priorityqueue"
	"github.com/liyue201/gostl/utils/comparator"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	v := nextIntSlice(n)

	ans := solve(n, k, v)

	Print(ans)
}

func solve(n, k int, v []int) int {
	var ans int
	for l := 0; l <= n; l++ {
		for r := 0; r <= n; r++ {
			if l+r > n {
				continue
			}
			rem := k - (l + r)
			if rem < 0 {
				continue
			}
			var s int
			q := priorityqueue.New[int](comparator.IntComparator)
			for i := 0; i < l; i++ {
				q.Push(v[i])
				s += v[i]
			}
			for i := 0; i < r; i++ {
				q.Push(v[n-i-1])
				s += v[n-i-1]
			}
			for i := 0; i < rem && !q.Empty(); i++ {
				v := q.Pop()
				if v > 0 {
					break
				}
				s -= v
			}
			ans = Max(ans, s)
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
