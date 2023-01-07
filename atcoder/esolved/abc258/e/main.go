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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q, x := nextInt(), nextInt(), nextInt()
	w := nextIntSlice(n)
	var k []int
	for i := 0; i < q; i++ {
		k = append(k, nextInt()-1)
	}
	ans := solveCommentary(n, q, x, w, k)
	PrintVertically(ans)
}

func solveCommentary(n, q, x int, w, k []int) []int {
	//s := 0
	sw := make([]int, 2*n+1)
	for i := 0; i < 2*n; i++ {
		sw[i+1] = sw[i] + w[i%n]
	}

	//i番目のジャガイモから詰め始めたらj番目のジャガイモを詰めて終わる時、インデックスの差分
	count := make([]int, n)
	for i := range count {
		count[i] = (x / sw[n]) * n
	}

	x %= sw[n]
	for i := 0; i < n; i++ {
		j := sort.Search(2*n, func(idx int) bool {
			return sw[idx]-sw[i] >= x
		})
		count[i] += (j - i)
	}

	order := make([]int, n)
	for i := range order {
		order[i] = -1
	}
	var path []int
	loop := -1
	var offset, u int
	for {
		if order[u] != -1 {
			loop = offset - order[u]
			break
		}
		order[u] = offset
		path = append(path, u)
		u = (u + count[u]) % n
		offset++
	}
	nonLoop := len(path) - loop

	//fmt.Println(nonLoop, loop, path)
	var ans []int
	for i := 0; i < q; i++ {
		var idx int
		if k[i] < nonLoop {
			idx = path[k[i]]
		} else {
			idx = path[nonLoop+(k[i]-nonLoop)%loop]
		}
		ans = append(ans, count[idx])
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
