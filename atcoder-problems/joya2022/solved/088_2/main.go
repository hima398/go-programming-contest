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

	n, m := nextInt(), nextInt()
	ans := solve(n, m)
	PrintVertically(ans)
}

type cell struct {
	i, j int
}

func solve(n, m int) [][]int {
	var ws [][2]int
	for k := 0; k < n; k++ {
		for l := 0; l < n; l++ {
			if k*k+l*l == m {
				//fmt.Println(k, l)
				ws = append(ws, [2]int{k, l})
			}
		}
	}
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = -1
		}
	}
	var q []cell
	q = append(q, cell{0, 0})
	ans[0][0] = 0
	di := []int{-1, -1, 1, 1}
	dj := []int{-1, 1, -1, 1}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for dir := range di {
			for _, w := range ws {
				ni, nj := cur.i+di[dir]*w[0], cur.j+dj[dir]*w[1]
				if ni < 0 || ni >= n || nj < 0 || nj >= n {
					continue
				}
				if ans[ni][nj] >= 0 {
					continue
				}
				q = append(q, cell{ni, nj})
				ans[ni][nj] = ans[cur.i][cur.j] + 1
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func PrintVertically(x [][]int) {
	defer out.Flush()
	for _, v := range x {
		//fmt.Fprintln(out, v)
		PrintHorizonaly(v)
	}
}
