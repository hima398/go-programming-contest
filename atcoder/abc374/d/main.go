package main

import (
	"bufio"
	"fmt"
	"math"
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

	n, s, t := nextInt(), nextInt(), nextInt()
	var a, b, c, d []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
		c = append(c, nextInt())
		d = append(d, nextInt())
	}

	ans := solve(n, s, t, a, b, c, d)

	Print(ans)
}

func computeDist(x1, y1, x2, y2 int) float64 {
	xx, yy := x2-x1, y2-y1
	res := float64(xx*xx + yy*yy)
	return math.Sqrt(res)
}

func solve(n, s, t int, a, b, c, d []int) float64 {
	// 6!=720
	// 2**6 = 64
	var p []int
	for i := 0; i < n; i++ {
		p = append(p, i)
	}

	var times []float64
	for i := 0; i < n; i++ {
		times = append(times, computeDist(a[i], b[i], c[i], d[i])/float64(t))
	}

	ans := 1e18
	for {
		//fmt.Println(p)
		for pat := 0; pat < (1 << n); pat++ {
			var time float64
			var cx, cy int
			for _, idx := range p {
				if (pat>>idx)&1 == 0 {
					time += computeDist(cx, cy, a[idx], b[idx]) / float64(s)
					cx, cy = c[idx], d[idx]
				} else {
					time += computeDist(cx, cy, c[idx], d[idx]) / float64(s)
					cx, cy = a[idx], b[idx]
				}
				time += times[idx]
			}
			//Print(time)
			ans = math.Min(ans, time)
		}
		if !NextPermutation(sort.IntSlice(p)) {
			break
		}
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

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
