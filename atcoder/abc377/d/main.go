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

	n, m := nextInt(), nextInt()
	var l, r []int
	for i := 0; i < n; i++ {
		l = append(l, nextInt()-1)
		r = append(r, nextInt()-1)
	}

	ans := solve(n, m, l, r)

	Print(ans)
}

func solve(n, m int, l, r []int) int {
	//const INF = 1 << 60

	//i以上の左端を持つ区間で最小のr
	prefixMin := make([]int, m)
	for i := range prefixMin {
		prefixMin[i] = m
	}
	for i := 0; i < n; i++ {
		prefixMin[l[i]] = Min(prefixMin[l[i]], r[i])
	}
	for i := m - 1; i > 0; i-- {
		prefixMin[i-1] = Min(prefixMin[i-1], prefixMin[i])
	}
	//fmt.Println(prefixMin)
	var ans int
	for l := 0; l < m; l++ {
		//li以上でri-1の最小値
		//fmt.Println(l, prefixMin[l]-l-1)
		ans += prefixMin[l] - l
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
