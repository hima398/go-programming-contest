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

	n, q := nextInt(), nextInt()
	var t, x []int
	for i := 0; i < q; i++ {
		t = append(t, nextInt())
		x = append(x, nextInt())
	}
	ans := solve(n, q, t, x)
	PrintInt(ans)
}

func solve(n, q int, t, x []int) int {
	h, w := n, n
	//row[i]:i行め(2<=i<=n-1)に白い石を置いたとき置き換えられる黒い石の数
	//col[i]:i行め(2<=i<=n-1)に白い石を置いたとき置き換えられる黒い石の数
	row, col := make([]int, n+1), make([]int, n+1)
	for i := 0; i <= n; i++ {
		row[i] = n
		col[i] = n
	}
	ans := (n - 2) * (n - 2)
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			//(1, x)に白い石を置く
			if x[i] < w {
				for i := x[i]; i < w; i++ {
					col[i] = h
				}
				w = x[i]
			}
			ans -= col[x[i]] - 2
		case 2:
			//(x, 1)に白い石を置く
			if x[i] < h {
				for i := x[i]; i < h; i++ {
					row[i] = w
				}
				h = x[i]
			}
			ans -= row[x[i]] - 2
		}
	}
	return ans
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
