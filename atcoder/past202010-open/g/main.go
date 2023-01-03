package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = nextString()
	}
	ans := solve(n, m, s)
	PrintInt(ans)
}

func bfs(n, m, si, sj int, f [][]string) bool {
	v := make([][]bool, n)
	for i := 0; i < n; i++ {
		v[i] = make([]bool, m)
	}
	type cell struct {
		i, j int
	}
	var q []cell
	q = append(q, cell{si, sj})
	v[si][sj] = true
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for k := 0; k < 4; k++ {
			ni, nj := cur.i+di[k], cur.j+dj[k]
			if ni < 0 || ni >= n || nj < 0 || nj >= m {
				continue
			}
			if f[ni][nj] == "#" || v[ni][nj] {
				continue
			}
			q = append(q, cell{ni, nj})
			v[ni][nj] = true
		}
	}
	ok := true
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if f[i][j] == "#" {
				continue
			}
			ok = ok && v[i][j]
		}
	}
	return ok
}

func solve(n, m int, s []string) int {
	var f [][]string
	si, sj := -1, -1
	for i := 0; i < n; i++ {
		f = append(f, strings.Split(s[i], ""))
		for j := 0; j < m; j++ {
			if f[i][j] == "." {
				si, sj = i, j
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if f[i][j] == "." {
				continue
			}
			f[i][j] = "."
			if bfs(n, m, si, sj, f) {
				ans++
			}
			f[i][j] = "#"
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
