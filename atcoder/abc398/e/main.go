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

	n := nextInt()
	var u, v []int
	for i := 0; i < n-1; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}

	solve(n, u, v)
}

func solve(n int, u, v []int) {
	connected := make([][]bool, n)
	for i := range connected {
		connected[i] = make([]bool, n)
		connected[i][i] = true
	}
	e := make([][]int, n)
	for i := 0; i < n-1; i++ {
		connected[u[i]][v[i]] = true
		connected[v[i]][u[i]] = true
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	var group [2][]int
	var dfs func(cur, par, depth int)
	dfs = func(cur, par, depth int) {
		if depth%2 == 0 {
			group[0] = append(group[0], cur)
		} else {
			group[1] = append(group[1], cur)
		}

		for _, next := range e[cur] {
			if next == par {
				continue
			}
			dfs(next, cur, depth+1)
		}
	}
	dfs(0, -1, 0)
	//fmt.Println(group[0])
	//fmt.Println(group[1])
	for _, s := range group[0] {
		for _, t := range group[0] {
			connected[s][t] = true
			connected[t][s] = true
		}
	}
	for _, s := range group[1] {
		for _, t := range group[1] {
			connected[s][t] = true
			connected[t][s] = true
		}
	}
	//for _, v := range connected {
	//	fmt.Println(v)
	//}
	//deck := queue.New[[2]int]()
	var turn int
	for _, s := range group[0] {
		for _, t := range group[1] {
			if connected[s][t] {
				continue
			}
			//if s > t {
			//	s, t = t, s
			//}
			//deck.Push([2]int{s, t})
			turn++
		}
	}
	var isFirst bool
	if turn%2 == 0 {
		fmt.Println("Second")
		//isSecond = true
	} else {
		fmt.Println("First")
		isFirst = true
	}
	if isFirst {
	Loop:
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if !connected[i][j] {
					if i > j {
						i, j = j, i
					}
					fmt.Println(i+1, j+1)
					connected[i][j] = true
					connected[j][i] = true
					break Loop
				}
			}
		}
		turn--
	}
	for turn > 0 {
		i, j := nextInt()-1, nextInt()-1
		connected[i][j] = true
		connected[j][i] = true
		turn--

	Loop2:
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if !connected[i][j] {
					if i > j {
						i, j = j, i
					}
					fmt.Println(i+1, j+1)
					connected[i][j] = true
					connected[j][i] = true
					break Loop2
				}
			}
		}
		turn--
	}
	fmt.Println(-1, -1)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
