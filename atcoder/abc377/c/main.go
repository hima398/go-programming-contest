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
	var a, b []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}

	ans := solve(n, m, a, b)

	Print(ans)
}

func solve(n, m int, a, b []int) int {
	di := []int{2, 1, -1, -2, -2, -1, 1, 2}
	dj := []int{1, 2, 2, 1, -1, -2, -2, -1}

	ngCell := make(map[int]map[int]struct{})
	for k := 0; k < m; k++ {
		i, j := a[k], b[k]
		if ngCell[i] == nil {
			ngCell[i] = make(map[int]struct{})
		}
		ngCell[i][j] = struct{}{}

		for l := 0; l < 8; l++ {
			ni, nj := i+di[l], j+dj[l]
			if ni < 1 || ni > n || nj < 1 || nj > n {
				continue
			}
			if ngCell[ni] == nil {
				ngCell[ni] = make(map[int]struct{})
			}
			ngCell[ni][nj] = struct{}{}
		}
	}

	ans := n * n
	for i := range ngCell {
		ans -= len(ngCell[i])
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
