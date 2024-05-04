package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/liyue201/gostl/ds/queue"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	ans, err := solve(n, s)
	if err != nil {
		Print(-1)
	} else {
		Print(ans)
	}
}

func solve(n int, s []string) (int, error) {
	di := []int{-1, 0, 1, 0}
	dj := []int{0, 1, 0, -1}
	const INF = math.MaxInt
	var dist [60][60][60][60]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for ii := 0; ii < n; ii++ {
				for jj := 0; jj < n; jj++ {
					dist[i][j][ii][jj] = INF
				}
			}
		}
	}
	type player struct {
		i, j int
	}
	type node []player

	var start node
	for i := range s {
		for j := range s[i] {
			if s[i][j] == 'P' {
				start = append(start, player{i, j})
			}
		}
	}
	q := queue.New[node]()
	q.Push(start)
	dist[start[0].i][start[0].j][start[1].i][start[1].j] = 0
	for !q.Empty() {
		cur := q.Pop()
		for k := 0; k < 4; k++ {
			var next []player
			for l := 0; l < 2; l++ {
				ni, nj := cur[l].i+di[k], cur[l].j+dj[k]
				if ni < 0 || ni >= n || nj < 0 || nj >= n || s[ni][nj] == '#' {
					ni, nj = cur[l].i, cur[l].j
				}
				next = append(next, player{ni, nj})
			}
			if dist[next[0].i][next[0].j][next[1].i][next[1].j] != INF {
				continue
			}
			q.Push(next)
			dist[next[0].i][next[0].j][next[1].i][next[1].j] = dist[cur[0].i][cur[0].j][cur[1].i][cur[1].j] + 1
			if next[0].i == next[1].i && next[0].j == next[1].j {
				return dist[next[0].i][next[0].j][next[1].i][next[1].j], nil
			}
		}
	}
	//fmt.Println("size = ", size)
	ans := INF
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans = Min(ans, dist[i][j][i][j])
		}
	}
	if ans == INF {
		return -1, errors.New("Impossible")
	} else {
		return ans, nil
	}
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
