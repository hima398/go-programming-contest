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

	all, n, m := nextInt(), nextInt(), nextInt()
	var l []int
	for i := 0; i < n; i++ {
		l = append(l, nextInt())
	}
	var x, y []int
	for i := 0; i < m; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	ans := solve(all, n, m, l, x, y)
	PrintVertically(ans)
}

func solve(all, n, m int, l, x, y []int) []int {
	var d []int
	for i := 0; i < n-1; i++ {
		d = append(d, l[i+1]-l[i]-1)
	}
	sort.Ints(d)
	sd := make([]int, n)
	for i := 1; i < n; i++ {
		sd[i] += sd[i-1] + d[i-1]
	}
	//PrintHorizonaly(d)
	var ans []int
	for i := 0; i < m; i++ {
		idx := sort.Search(len(d), func(j int) bool {
			return d[j] > x[i]+y[i]
		})
		//fmt.Printf("x + y = %d, idx = %d\n", x[i]+y[i], idx)
		s := n
		s += Min(l[0]-1, x[i])
		s += Min(all-l[n-1], y[i])
		s += sd[idx] + (x[i]+y[i])*(len(d)-idx)
		ans = append(ans, s)
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
