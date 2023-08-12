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
	x := nextIntSlice(n)
	c := nextIntSlice(n)
	ans := solve(n, x, c)
	PrintInt(ans)
}

func solve(n int, x, c []int) int {
	//0-indexed
	for i := range x {
		x[i]--
	}
	//枝を取り払って、循環だけにする
	//入次数
	d := make([]int, n)
	for _, xi := range x {
		d[xi]++
	}
	var q []int
	for i, di := range d {
		if di == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		d[x[cur]]--
		if d[x[cur]] == 0 {
			q = append(q, x[cur])
		}
	}
	//fmt.Println(d)
	//循環の中で最小値を答えに足していく
	visited := make([]bool, n)
	bfs := func(s int) int {
		var q []int
		q = append(q, s)
		visited[s] = true
		ans := 1 << 60
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			next := x[cur]
			ans = Min(ans, c[cur])
			if visited[next] {
				break
			}
			q = append(q, next)
			visited[next] = true
		}
		return ans
	}
	var ans int
	for i := 0; i < n; i++ {
		if d[i] > 0 && !visited[i] {
			ans += bfs(i)
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

func PrintInt(x int) {
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
