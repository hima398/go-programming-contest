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

	n, m, k := nextInt(), nextInt(), nextInt()
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = nextString()
	}
	ans := solve(n, m, k, s)
	PrintInt(ans)
}

func solve(n, m, lk int, s []string) int {
	f := make([][][10]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([][10]int, m)
		for j := 0; j < m; j++ {
			k := int(s[i][j] - '0')
			f[i][j][k] = 1
		}
	}
	sf := make([][][10]int, n+1)
	for i := 0; i <= n; i++ {
		sf[i] = make([][10]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < 10; k++ {
				sf[i+1][j+1][k] = sf[i+1][j][k] + sf[i][j+1][k] - sf[i][j][k] + f[i][j][k]
			}
		}
	}
	ans := 0
	for i := n; i >= 1; i-- {
		for j := m; j >= 1; j-- {
			for l := 1; l <= Min(i, j); l++ {
				for k := 0; k < 10; k++ {
					//l*l内に含まれるkの数
					v := sf[i][j][k] - sf[i-l][j][k] - sf[i][j-l][k] + sf[i-l][j-l][k]
					//正方形の面積
					if l*l-v <= lk {
						ans = Max(ans, l)
					}
				}
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

func nextString() string {
	sc.Scan()
	return sc.Text()
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
