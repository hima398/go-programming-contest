package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintField(a [][]int) {
	for _, sub := range a {
		fmt.Println(sub)
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := make([][]int, 2*n)
	for i := 0; i < 2*n; i++ {
		a[i] = make([]int, 2*n)
	}
	for i := 0; i < 2*n; i++ {
		for j := i + 1; j < 2*n; j++ {
			v := nextInt()
			a[i][j] = v
		}
	}

	ans := solve(n, a)
	fmt.Println(ans)
}

func solve(n int, a [][]int) int {
	n2 := 2 * n
	mx := (1 << n2) - 1
	//fmt.Printf("%b\n", mx)

	var dfs func(idx, pat, score int)
	ans := 0
	dfs = func(idx, pat, score int) {
		if pat == mx {
			ans = Max(ans, score)
			return
		}
		if pat>>idx&1 == 0 {
			for i := idx + 1; i < n2; i++ {
				if pat>>i&1 == 0 {
					dfs(idx+1, pat|(1<<idx)|(1<<i), score^a[idx][i])
				}
			}
		} else {
			dfs(idx+1, pat, score)
		}
	}
	dfs(0, 0, 0)
	return ans
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
