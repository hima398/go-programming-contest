package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, m int, a, b, c []int) int {
	const INF = 1 << 60
	d := make([][]int, n)
	e := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		e[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			d[i][j] = INF
		}
	}
	for i := 0; i < m; i++ {
		a[i]--
		b[i]--
		d[a[i]][b[i]] = c[i]
		d[b[i]][a[i]] = c[i]
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				d[i][j] = Min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}
	ans := m
	for i := 0; i < m; i++ {
		var isInShotest bool
		for j := 0; j < n; j++ {
			isInShotest = isInShotest || d[j][a[i]]+c[i] == d[j][b[i]]
		}
		if isInShotest {
			ans--
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var a, b, c []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
		c = append(c, nextInt())
	}
	ans := solve(n, m, a, b, c)
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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
