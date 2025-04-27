package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"math/bits"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	m := nextInt()
	s1 := nextString()
	s2 := nextString()
	s3 := nextString()
	ans, err := solve(m, s1, s2, s3)
	if err != nil {
		Print(-1)
	} else {
		Print(ans)
	}
}

func solve(m int, s1, s2, s3 string) (int, error) {
	const INF = math.MaxInt
	s := []string{s1, s2, s3}
	v := 1<<3 - 1
	var w []int
	for u := 0; u < v; u++ {
		w = append(w, u)
	}
	sort.Slice(w, func(i, j int) bool {
		return bits.OnesCount(uint(w[i])) < bits.OnesCount(uint(w[j]))
	})
	//dp := make([][]int, v+1)
	dp := make(map[int]map[int]int)
	//for i := range dp {
	//	dp[i] = make([]int, 10)
	//	for j := range dp[i] {
	//		dp[i][j] = INF
	//	}
	//}
	//t = 0秒に各リールのボタンを押す
	//for k := 0; k < 3; k++ {
	//	if dp[1<<k] == nil {
	//		dp[1<<k] = make(map[int]int)
	//	}
	//	dp[1<<k][int(s[k][0]-'0')] = 0
	//}
	for t := 0; t <= 3*m; t++ {
		//fmt.Println("t = ", t)
		//fmt.Println(dp)
		next := make(map[int]map[int]int, v+1)
		for i := range dp {
			if next[i] == nil {
				next[i] = make(map[int]int)
			}
			for j := range dp[i] {
				next[i][j] = dp[i][j]
			}
		}

		for _, u := range w {
			for k := 0; k < 3; k++ {
				//すでにリールが止めてある
				if (u>>k)&1 > 1 {
					continue
				}
				nu := u | (1 << k)

				idx := t % m
				j := int(s[k][idx] - '0')
				//存在しない状態
				if _, found := dp[u][j]; u > 0 && !found {
					continue
				}
				if next[nu] == nil {
					next[nu] = make(map[int]int)
				}
				if _, found := dp[nu][j]; found {
					next[nu][j] = Min(next[nu][j], t)
				} else {
					next[nu][j] = t
				}
			}
		}
		dp = next
	}
	ans := INF
	for j := range dp[v] {
		ans = Min(ans, dp[v][j])
	}
	if ans == INF {
		return -1, errors.New("Impossible")
	} else {
		return ans, nil
	}
}

func firstsolve(m int, s1, s2, s3 string) (int, error) {
	const INF = math.MaxInt
	s := []string{s1, s2, s3}
	v := 1<<3 - 1
	var w []int
	for u := 1; u < v; u++ {
		w = append(w, u)
	}
	sort.Slice(w, func(i, j int) bool {
		return bits.OnesCount(uint(w[i])) < bits.OnesCount(uint(w[j]))
	})
	dp := make([][]int, v+1)
	for i := range dp {
		dp[i] = make([]int, 10)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	//t = 0秒に各リールのボタンを押す
	for k := 0; k < 3; k++ {
		dp[1<<k][int(s[k][0]-'0')] = 0
	}
	for t := 1; t <= 3*m; t++ {
		next := make([][]int, v+1)
		for i := range next {
			next[i] = make([]int, 10)
			copy(next[i], dp[i])
		}

		for _, u := range w {
			for k := 0; k < 3; k++ {
				//すでにリールが止めてある
				if (u>>k)&1 > 1 {
					continue
				}
				nu := u | (1 << k)

				idx := t % m
				j := int(s[k][idx] - '0')
				//存在しない状態
				if dp[u][j] == INF {
					continue
				}

				next[nu][j] = Min(next[nu][j], t)
			}
		}
		dp = next
	}
	ans := INF
	for j := 0; j < 10; j++ {
		ans = Min(ans, dp[v][j])
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
