package main

import (
	"bufio"
	"fmt"
	"math/bits"
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

	var a [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = nextInt()
		}
	}

	ok := solve(a)

	if ok {
		Print("Takahashi")
	} else {
		Print("Aoki")
	}
}

func solve(a [3][3]int) bool {
	const n = 3

	f := make([]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			idx := n*i + j
			f[idx] = a[i][j]
		}
	}

	var pattern []int
	pattern = append(pattern, 1|(1<<1)|(1<<2))
	pattern = append(pattern, 1|(1<<3)|(1<<6))
	pattern = append(pattern, 1|(1<<4)|(1<<8))
	pattern = append(pattern, (1<<1)|(1<<4)|(1<<7))
	pattern = append(pattern, (1<<2)|(1<<5)|(1<<8))
	pattern = append(pattern, (1<<2)|(1<<4)|(1<<6))
	pattern = append(pattern, (1<<3)|(1<<4)|(1<<5))
	pattern = append(pattern, (1<<6)|(1<<7)|(1<<8))

	//高橋くんの勝ちを判定する
	judge := func(f []int, t, a int) int {
		for _, p := range pattern {
			if t&p == p {
				return 1
			}
		}
		for _, p := range pattern {
			if a&p == p {
				return 2
			}
		}

		//同じマスに赤と青が塗られたデータが入ってきている
		if t&a > 0 {
			return 0
		}
		if bits.OnesCount(uint(t|a)) < n*n {
			return 0
		}

		//高橋くん、青木くんのスコア
		var st, sa int
		for i := 0; i < n*n; i++ {
			st += f[i] * (t >> i & 1)
			sa += f[i] * (a >> i & 1)
		}
		if st > sa {
			return 1
		} else if st < sa {
			return 2
		} else {
			return 0
		}
	}
	var memo [2]map[int]map[int]bool
	for i := 0; i < 2; i++ {
		memo[i] = make(map[int]map[int]bool)
	}
	//盤面がf、配色がvで高橋くんのスコアがt、青木くんのスコアがaでturn目の時に高橋くんが必勝であるかを返す
	var dfs func(t, a, turn int) bool
	dfs = func(t, a, turn int) bool {
		if memo[turn%2][t] != nil {
			if _, found := memo[turn%2][t][a]; found {
				return memo[turn%2][t][a]
			}
		}
		e := judge(f, t, a)
		if e > 0 {
			if memo[turn%2][t] == nil {
				memo[turn%2][t] = make(map[int]bool)
			}
			switch e {
			case 1:
				memo[turn%2][t][a] = true
			case 2:
				memo[turn%2][t][a] = false
			}
			return memo[turn%2][t][a]
		}

		var ok2 bool
		if turn%2 == 1 {
			ok2 = true
		}
		for k := 0; k < n*n; k++ {
			if (t>>k)&1 > 0 || (a>>k)&1 > 0 {
				continue
			}
			if turn%2 == 0 { //高橋くんのターン
				//次の手に高橋くんが必勝の盤面があれば良い
				ok2 = ok2 || dfs(t|(1<<k), a, turn+1)
			} else { //青木くんのターン
				//次の手が全て高橋くんが必勝の盤面である必要がある
				ok2 = ok2 && dfs(t, a|(1<<k), turn+1)
			}
		}
		if memo[turn%2][t] == nil {
			memo[turn%2][t] = make(map[int]bool)
		}
		memo[turn%2][t][a] = ok2
		return memo[turn%2][t][a]
	}

	return dfs(0, 0, 0)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
