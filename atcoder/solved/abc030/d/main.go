package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, a := nextInt(), nextInt()-1
	//k := nextInt()
	k := nextString()
	b := nextIntSlice(n)
	for i := range b {
		b[i]--
	}

	var ans int
	if len(k) <= 7 {
		ans = solveHonestly(n, a, k, b)
	} else {
		ans = solve(n, a, k, b)
	}
	PrintInt(ans)
}

func solveHonestly(n, a int, k string, b []int) int {
	ans := a
	kd, _ := strconv.Atoi(k)
	for i := 0; i < kd; i++ {
		ans = b[ans]
	}

	return ans + 1
}

func solve(n, a int, k string, b []int) int {
	visited := make([]int, n)
	cur := a
	start := -1
	for visited[cur] < 2 {
		if start < 0 && visited[cur] == 1 {
			start = cur
		}
		visited[cur]++
		cur = b[cur]
	}
	offset, loop := 0, 0
	for _, v := range visited {
		if v == 1 {
			offset++
		} else if v == 2 {
			loop++
		}
	}

	bk := big.NewInt(0)
	bk.SetString(k, 10)
	bo, bl := big.NewInt(int64(offset)), big.NewInt(int64(loop))
	bm := bk.Sub(bk, bo)
	bm = bm.Mod(bm, bl)
	m, _ := strconv.Atoi(bm.String())
	//fmt.Println(offset, loop, start, k, m)
	ans := start
	for i := 0; i < m; i++ {
		ans = b[ans]
	}

	return ans + 1
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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
