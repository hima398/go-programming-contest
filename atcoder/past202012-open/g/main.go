package main

import (
	"bufio"
	"errors"
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

	h, w := nextInt(), nextInt()
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = nextString()
	}
	ans := solve(h, w, s)
	fmt.Println(len(ans))
	for _, v := range ans {
		fmt.Println(v[0], v[1])
	}
}

func solve(h, w int, s []string) [][2]int {
	k := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				k++
			}
		}
	}
	visited := make([][]bool, h)
	for i := 0; i < h; i++ {
		visited[i] = make([]bool, w)
	}
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	var dfs func(ci, cj, d int) ([][2]int, error)
	dfs = func(ci, cj, d int) ([][2]int, error) {
		//fmt.Println(ci, cj, d)
		visited[ci][cj] = true
		if d == k {
			visited[ci][cj] = false
			return [][2]int{{ci + 1, cj + 1}}, nil
		}
		for ii := range di {
			ni, nj := ci+di[ii], cj+dj[ii]
			if ni < 0 || ni >= h || nj < 0 || nj >= w {
				continue
			}
			if visited[ni][nj] || s[ni][nj] == '.' {
				continue
			}
			snake, err := dfs(ni, nj, d+1)
			if err != nil {
				//visited[ci][cj] = false
				//return nil, errors.New("Impossible")
				continue
			}
			snake = append(snake, [2]int{ci + 1, cj + 1})
			visited[ci][cj] = false
			return snake, nil
		}
		visited[ci][cj] = false
		return nil, errors.New("Impossible")
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				ans, err := dfs(i, j, 1)
				if err != nil {
					continue
				}
				return ans
			}
		}
	}
	return nil
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
