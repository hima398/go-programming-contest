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
	a := nextIntSlice(n)

	//ans := solveHonestly(n, a)
	ans := solve(n, a)

	Print(ans)
	//generateWorstTestCase()
}

func generateWorstTestCase() {
	const n = 2 * int(1e5)
	Print(n)
	var a []int
	for i := 0; i < n/2; i++ {
		a = append(a, 448)
	}
	for i := 0; i < n/2; i++ {
		a = append(a, 448)
	}
	PrintHorizonaly(a)
}

func solveHonestly(n int, a []int) int {
	const p = 998244353
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := i + a[i]; j < n; j += a[i] {
			dp[j] += dp[i]
			dp[j] %= p
		}
	}
	//fmt.Println(dp)
	var ans int
	for _, v := range dp {
		ans += v
		ans %= p
	}
	return ans
}

func solve(n int, a []int) int {
	const p = 998244353

	maxA := 0
	for _, ai := range a {
		maxA = Max(maxA, ai)
	}
	boundary := Sqrt(2 * int(1e5))
	//fmt.Println("boundary = ", boundary)

	dp1 := make([]int, n)
	dp2 := make([][]int, boundary+1)
	for i := range dp2 {
		dp2[i] = make([]int, boundary)
	}
	dp1[0] = 1
	for i := 0; i < n; i++ {
		for j := 1; j <= boundary; j++ {
			dp1[i] += dp2[j][i%j]
			dp1[i] %= p
		}

		if a[i] > boundary {
			for j := i + a[i]; j < n; j += a[i] {
				dp1[j] += dp1[i]
				dp1[j] %= p
			}
		} else {
			dp2[a[i]][i%a[i]] += dp1[i]
			dp2[a[i]][i%a[i]] %= p
		}
	}
	//fmt.Println(dp)
	var ans int
	for _, v := range dp1 {
		ans += v
		ans %= p
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Sqrt(x int) int {
	x2 := int(math.Sqrt(float64(x))) - 1
	for (x2+1)*(x2+1) <= x {
		x2++
	}
	return x2
}
