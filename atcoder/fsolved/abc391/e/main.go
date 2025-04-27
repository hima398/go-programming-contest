package main

import (
	"bufio"
	"fmt"
	"math"
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
	a := nextString()

	ans := solve(n, a)

	Print(ans)
}

func solve(n int, a string) int {
	// 3**13 = 4782969
	var m int
	for i := 0; i <= n; i++ {
		m += Pow(3, i)
	}
	tree := make([]int, m)
	offset := m - Pow(3, n)
	for i, ai := range a {
		tree[offset+i] = int(ai - '0')
	}
	cnt := make([][2]int, m)
	var dfs func(cur int) int
	dfs = func(cur int) int {
		//fmt.Println(cur)
		if cur >= offset {
			return tree[cur]
		}
		//var cnt [2]int
		for next := cur*3 + 1; next <= cur*3+3; next++ {
			cnt[cur][dfs(next)]++
		}
		if cnt[cur][0] > cnt[cur][1] {
			tree[cur] = 0
		} else {
			tree[cur] = 1
		}
		return tree[cur]
	}
	dfs(0)
	//Print(tree)
	//Print(cnt)
	const INF = 1 << 60

	memo := make([]int, m)
	for i := range memo {
		memo[i] = INF
	}
	for i := 0; i < Pow(3, n); i++ {
		memo[offset+i] = 1
	}

	var dfs2 func(cur int) int
	dfs2 = func(cur int) int {
		//fmt.Println("cur = ", cur)
		if memo[cur] < INF {
			return memo[cur]
		}
		switch cnt[cur][tree[cur]] {
		case 2:
			for next := cur*3 + 1; next <= cur*3+3; next++ {
				dfs2(next)
				if tree[cur] != tree[next] {
					continue
				}
				memo[cur] = Min(memo[cur], dfs2(next))
			}
		case 3:
			n1, n2, n3 := 3*cur+1, 3*cur+2, 3*cur+3
			memo[cur] = Min(dfs2(n1)+dfs2(n2), dfs2(n1)+dfs2(n3), dfs2(n2)+dfs2(n3))
		}
		return memo[cur]
	}
	ans := dfs2(0)
	//fmt.Println(memo)
	return ans
}

func Pow(x, y int) int {
	ret := 1
	for i := 0; i < y; i++ {
		ret *= x
	}
	return ret
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

func Min(x ...int) int {
	f := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	ret := math.MaxInt
	for _, xi := range x {
		ret = f(ret, xi)
	}
	return ret
}
