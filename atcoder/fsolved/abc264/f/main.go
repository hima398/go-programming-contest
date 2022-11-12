package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

const INF = 1 << 60

func computeMinimumCost(h, w int, r, c []int, f [][]int) int {
	dp := make([][][2][2]int, h)
	for i := 0; i < h; i++ {
		dp[i] = make([][2][2]int, w)
		for j := 0; j < w; j++ {
			for d := 0; d < 2; d++ {
				for k := 0; k < 2; k++ {
					dp[i][j][d][k] = INF
				}
			}
		}
	}
	for i := 0; i < 2; i++ {
		j := f[0][0] ^ i
		dp[0][0][0][j] = i*r[0] + j*c[0]
		dp[0][0][1][i] = i*r[0] + j*c[0]
	}

	di := [2]int{1, 0}
	dj := [2]int{0, 1}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for d := 0; d < 2; d++ {
				ni, nj := i+di[d], j+dj[d]
				if ni >= h || nj >= w {
					continue
				}
				for k := 0; k < 2; k++ {
					flip := f[ni][nj] ^ k
					cost := 0
					if d == 0 {
						cost = flip * r[ni]
					} else {
						cost = flip * c[nj]
					}
					for nd := 0; nd < 2; nd++ {
						nk := k
						if d != nd {
							nk = flip
						}
						dp[ni][nj][nd][nk] = Min(dp[ni][nj][nd][nk], dp[i][j][d][k]+cost)
					}
				}
			}
		}
	}

	ans := INF
	for d := 0; d < 2; d++ {
		for k := 0; k < 2; k++ {
			ans = Min(ans, dp[h-1][w-1][d][k])
		}
	}
	return ans
}

func solve(h, w int, r, c []int, s []string) int {
	f := make([][]int, h)
	for i := 0; i < h; i++ {
		f[i] = make([]int, w)
		for j := 0; j < w; j++ {
			if s[i][j] == '1' {
				f[i][j] = 1
			}
		}
	}
	ans := computeMinimumCost(h, w, r, c, f)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			f[i][j] ^= 1
		}
	}
	ans = Min(ans, computeMinimumCost(h, w, r, c, f))
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	r := nextIntSlice(h)
	c := nextIntSlice(w)
	var s []string
	for i := 0; i < h; i++ {
		s = append(s, nextString())
	}
	ans := solve(h, w, r, c, s)
	PrintInt(ans)
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
