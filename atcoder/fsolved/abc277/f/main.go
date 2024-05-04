package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = nextIntSlice(w)
	}
	ans := solveCommentary(h, w, a)
	PrintString(ans)
}

func solveCommentary(h, w int, a [][]int) string {
	const INF = int(1e12) + 1
	var rows [][2]int

	//行に関する条件の評価
	for i := 0; i < h; i++ {
		min, max := INF, -INF
		for j := 0; j < w; j++ {
			if a[i][j] == 0 {
				continue
			}
			min = Min(min, a[i][j])
			max = Max(max, a[i][j])
		}
		if min <= max {
			rows = append(rows, [2]int{min, max})
		}
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i][0] == rows[j][0] {
			return rows[i][1] < rows[j][1]
		}
		return rows[i][0] < rows[j][0]
	})
	for i := range rows {
		if i == len(rows)-1 {
			break
		}
		if rows[i][1] > rows[i+1][0] {
			return "No"
		}
	}

	g := make([][]int, w+h*w+1)
	//fmt.Println(len(g))
	for r := 0; r < h; r++ {
		var cols [][2]int
		for c := 0; c < w; c++ {
			if a[r][c] == 0 {
				continue
			}
			cols = append(cols, [2]int{a[r][c], c})
		}
		sort.Slice(cols, func(i, j int) bool {
			if cols[i][0] == cols[j][0] {
				return cols[i][1] < cols[j][1]
			}
			return cols[i][0] < cols[j][0]
		})
		for i := 1; i < len(cols); i++ {
			if cols[i-1][0] == cols[i][0] {
				continue
			}
			v := w + cols[i-1][0]
			for j := i - 1; j >= 0; j-- {
				if cols[j][0] != cols[i-1][0] {
					break
				}
				g[cols[j][1]] = append(g[cols[j][1]], v)
			}
			for j := i; j < len(cols); j++ {
				if cols[j][0] != cols[i][0] {
					break
				}
				g[v] = append(g[v], cols[j][1])
			}
		}
	}
	n := w + h*w
	used := make([]bool, n)
	var topo []int
	var dfs func(x int)
	dfs = func(x int) {
		used[x] = true
		for _, next := range g[x] {
			if used[next] {
				continue
			}
			dfs(next)
		}
		topo = append(topo, x)
	}
	for i := 0; i < n; i++ {
		if !used[i] {
			dfs(i)
		}
	}
	for i := 0; i < len(topo)/2; i++ {
		j := len(topo) - 1 - i
		topo[i], topo[j] = topo[j], topo[i]
	}
	for i := 0; i < n; i++ {
		used[i] = false
	}

	for _, v := range topo {
		used[v] = true
		for _, next := range g[v] {
			if used[next] {
				return "No"
			}
		}
	}

	return "Yes"
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

func PrintString(x string) {
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
