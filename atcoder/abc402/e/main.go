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

	n, x := nextInt(), nextInt()
	var s, c, p []int
	for i := 0; i < n; i++ {
		s = append(s, nextInt())
		c = append(c, nextInt())
		p = append(p, nextInt())
	}

	ans := solve(n, x, s, c, p)

	Print(ans)
}

func solve(n, x int, s, c, p []int) float64 {
	set := (1 << n) - 1
	memo := make([]map[int]float64, set+1)
	for subSet := range memo {
		memo[subSet] = make(map[int]float64)
	}

	var dfs func(subSet, money int) float64
	dfs = func(subSet, money int) float64 {
		if v, found := memo[subSet][money]; found {
			return v
		}
		var e float64
		for i := 0; i < n; i++ {
			if (subSet>>i)&1 > 0 {
				continue
			}
			if money-c[i] < 0 {
				continue
			}
			nextSet := subSet | (1 << i)
			nextMoney := money - c[i]
			v := (dfs(nextSet, nextMoney)+float64(s[i]))*float64(p[i])/100.0 + dfs(subSet, nextMoney)*float64(100-p[i])/100.0
			e = math.Max(e, v)
		}
		memo[subSet][money] = e
		return e
	}

	ans := dfs(0, x)
	//for set, e := range memo {
	//	Print(fmt.Sprintf("%03b: %v", set, e))
	//}
	return ans
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
