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

	n, x := nextInt(), nextInt()
	t := nextIntSlice(n)

	ans := solve(n, x, t)

	Print(ans)
}

func solve(n, x int, t []int) int {
	p := 998244353
	//i+0.5秒後に曲が切り替わる確率
	dp := make([]int, x+1)
	dp[0] = 1
	inverseN := Inv(n, p)
	for k := 1; k <= x; k++ {
		for i := 0; i < n; i++ {
			if k >= t[i] {
				dp[k] += dp[k-t[i]]
				dp[k] %= p
			}
		}
		dp[k] *= inverseN
		dp[k] %= p
	}
	//fmt.Println(dp)
	var ans int
	if t[0] > x {
		ans += inverseN
	} else {
		for i := x - t[0] + 1; i <= x; i++ {
			ans += dp[i] * inverseN
			ans %= p
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}
