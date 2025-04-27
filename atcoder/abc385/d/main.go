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

	n, m, sx, sy := nextInt(), nextInt(), nextInt(), nextInt()
	var x, y []int
	for i := 0; i < n; i++ {
		x, y = append(x, nextInt()), append(y, nextInt())
	}
	var d []string
	var c []int
	for i := 0; i < m; i++ {
		d = append(d, nextString())
		c = append(c, nextInt())
	}

	gx, gy, ans := solve(n, m, sx, sy, x, y, d, c)

	fmt.Println(gx, gy, ans)
}

func solve(n, m, sx, sy int, x, y []int, d []string, c []int) (int, int, int) {
	rows, cols := make(map[int]map[int]struct{}), make(map[int]map[int]struct{})
	for i := 0; i < n; i++ {
		if rows[y[i]] == nil {
			rows[y[i]] = make(map[int]struct{})
		}
		rows[y[i]][x[i]] = struct{}{}
		if cols[x[i]] == nil {
			cols[x[i]] = make(map[int]struct{})
		}
		cols[x[i]][y[i]] = struct{}{}
	}

	cx, cy := sx, sy
	var ans int
	for i := 0; i < m; i++ {
		switch d[i] {
		case "U":
			if cols[cx] == nil {
				cy += c[i]
				continue
			}
			for k := range cols[cx] {
				if cy <= k && k <= cy+c[i] {
					delete(cols[cx], k)
					if rows[k] != nil {
						delete(rows[k], cx)
					}
					ans++
				}
			}
			cy += c[i]
		case "D":
			if cols[cx] != nil {
				for k := range cols[cx] {
					if cy-c[i] <= k && k <= cy {
						delete(cols[cx], k)
						if rows[k] != nil {
							delete(rows[k], cx)
						}
						ans++
					}
				}
			}
			cy -= c[i]
		case "L":
			if rows[cy] != nil {
				for k := range rows[cy] {
					if cx-c[i] <= k && k <= cx {
						delete(rows[cy], k)
						if cols[k] != nil {
							delete(cols[k], cy)
						}
						ans++
					}
				}
			}
			cx -= c[i]
		case "R":
			if rows[cy] != nil {
				for k := range rows[cy] {
					if cx <= k && k <= cx+c[i] {
						delete(rows[cy], k)
						if cols[k] != nil {
							delete(cols[k], cy)
						}
						ans++
					}
				}
			}
			cx += c[i]
		}
	}
	return cx, cy, ans
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
