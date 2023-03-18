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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	//ans := solveHonestly(n, m, a, b)
	ans := solveCommentary(n, m, a, b)
	PrintHorizonaly(ans)
}

func solveCommentary(n, m int, a, b []int) []int {
	for i := range a {
		a[i]--
		b[i]--
	}
	var r int
	for _, ai := range a {
		r = Max(r, ai)
	}
	idxes := make([][]int, m)
	for i, ai := range a {
		idxes[ai] = append(idxes[ai], i)
	}
	minB := 2*int(1e5) + 1
	for _, v := range b {
		minB = Min(minB, v)
	}
	ans := make([]int, m+1)
	for l := 0; l <= minB; l++ {
		//r-l+1 <= k <= m-l
		ans[r-l]++
		ans[m-l]--
		for _, idx := range idxes[l] {
			r = Max(r, b[idx])
		}
	}
	for i := 0; i < m; i++ {
		ans[i+1] += ans[i]
	}
	ans = ans[:m]
	return ans
}

func solveHonestly(n, m int, a, b []int) []int {
	ans := make([]int, m)
	for l := 1; l <= m; l++ {
		for r := l; r <= m; r++ {
			ok := true
			for i := 0; i < n; i++ {
				ok = ok && ((l <= a[i] && a[i] <= r) || (l <= b[i] && b[i] <= r))
			}
			if ok {
				//fmt.Println(l, r)
				ans[r-l]++
			}
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
