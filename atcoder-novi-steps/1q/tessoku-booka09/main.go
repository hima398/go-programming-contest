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

	h, w, n := nextInt(), nextInt(), nextInt()
	var a, b, c, d []int
	for i := 0; i < n; i++ {
		//0-index化もあわせて行なっておく
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		c = append(c, nextInt()-1)
		d = append(d, nextInt()-1)
	}
	ans := solve(h, w, n, a, b, c, d)
	for _, v := range ans {
		PrintHorizonaly(v)
	}
}

func solve(h, w, n int, a, b, c, d []int) [][]int {
	ans := make([][]int, h)
	for i := range ans {
		ans[i] = make([]int, w)
	}
	for i := 0; i < n; i++ {
		ans[a[i]][b[i]]++
		if c[i]+1 < h {
			ans[c[i]+1][b[i]]--
		}
		if d[i]+1 < w {
			ans[a[i]][d[i]+1]--
		}
		if c[i]+1 < h && d[i]+1 < w {
			ans[c[i]+1][d[i]+1]++
		}
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if j+1 == w {
				break
			}
			ans[i][j+1] += ans[i][j]
		}
	}
	for j := 0; j < w; j++ {
		for i := 0; i < h; i++ {
			if i+1 == h {
				break
			}
			ans[i+1][j] += ans[i][j]
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
