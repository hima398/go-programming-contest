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
	f := nextIntSlice(n)
	//ans := solveHonestly(n, m, f)
	ans := solve(n, m, f)
	PrintInt(ans)
}

func solveHonestly(n, m int, f []int) int {
	const p = int(1e9) + 7
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dup := make(map[int]struct{})
		dup[f[i-1]] = struct{}{}
		k := 1
		for i-k-1 >= 0 && k <= i {
			if _, found := dup[f[i-k-1]]; found {
				break
			}
			dup[f[i-k-1]] = struct{}{}
			k++
		}
		for j := 1; i-j >= 0 && j <= k; j++ {
			dp[i] += dp[i-j]
			dp[i] %= p
		}
	}
	//fmt.Println(dp)
	return dp[n]
}

func solve(n, m int, f []int) int {
	const p = int(1e9) + 7
	dp := make([]int, n+1)
	dps := make([]int, n+1)
	dp[0] = 1
	dps[0] = dp[0]

	l := 0
	dup := make(map[int]struct{})
	for r := 1; r <= n; r++ {
		if _, found := dup[f[r-1]]; found {
			for l < r {
				l++
				delete(dup, f[l-1])
				if _, found := dup[f[r-1]]; !found {
					break
				}
			}
		}
		dup[f[r-1]] = struct{}{}
		//fmt.Println(l, r, dup)
		if l == 0 {
			dp[r] = dps[r-1]
		} else {
			dp[r] = (dps[r-1] - dps[l-1] + p) % p
		}
		dps[r] = (dps[r-1] + dp[r]) % p
	}
	//fmt.Println(dp)
	//fmt.Println(dps)
	return dp[n]
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
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
