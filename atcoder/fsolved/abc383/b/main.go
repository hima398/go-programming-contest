package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func computeDist(x1, y1, x2, y2 int) int {
	return Abs(x2-x1) + Abs(y2-y1)
}
func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w, d := nextInt(), nextInt(), nextInt()
	var s []string
	for i := 0; i < h; i++ {
		s = append(s, nextString())
	}
	var ans int
	for u := 0; u < h; u++ {
		for v := 0; v < w; v++ {
			if s[u][v] == '#' {
				continue
			}
			for a := 0; a < h; a++ {
				for b := 0; b < w; b++ {
					if s[a][b] == '#' {
						continue
					}
					if u == a && v == b {
						continue
					}
					var t int
					for i := 0; i < h; i++ {
						for j := 0; j < w; j++ {
							if s[i][j] == '#' {
								continue
							}
							if computeDist(i, j, u, v) <= d || computeDist(i, j, a, b) <= d {
								t++
							}
						}
					}
					ans = Max(ans, t)
				}
			}
		}
	}
	Print(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
