package main

import (
	"bufio"
	"errors"
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

	n, k := nextInt(), nextInt()
	ans, err := solve(n, k)
	if err != nil {
		PrintInt(-1)
		return
	}
	PrintInt(len(ans))
	for _, uv := range ans {
		PrintHorizonaly(uv)
	}
}

func solve(n, k int) ([][]int, error) {
	m := (n - 1) * (n - 2) / 2
	if m < k {
		return nil, errors.New("Impossible")
	}
	var ans [][]int
	for i := 1; i < n; i++ {
		ans = append(ans, []int{1, i + 1})
	}
	for i := 2; i <= n-1; i++ {
		for j := i + 1; j <= n; j++ {
			if m == k {
				return ans, nil
			}
			ans = append(ans, []int{i, j})
			m--
		}
	}
	return ans, nil
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
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
