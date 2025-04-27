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
	var u, v []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}

	ans := solve(n, m, u, v)

	Print(ans)
}

func solve(n, m int, u, v []int) int {
	for i := 0; i < m; i++ {
		if u[i] > v[i] {
			u[i], v[i] = v[i], u[i]
		}
	}
	connected := make([]map[int]struct{}, n)
	for i := range connected {
		connected[i] = make(map[int]struct{})
		connected[i][i] = struct{}{}
	}

	var ans int
	for i := 0; i < m; i++ {
		if _, found := connected[u[i]][v[i]]; found {
			ans++
		}
		connected[u[i]][v[i]] = struct{}{}
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
