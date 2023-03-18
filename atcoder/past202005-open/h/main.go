package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, l := nextInt(), nextInt()
	x := nextIntSlice(n)
	t := nextIntSlice(3)
	ans := solve(n, l, x, t)
	fmt.Println(ans)
}

func solve(n, l int, x, t []int) int {
	const INF = 1 << 60
	hurdles := make([]bool, l+10)
	for _, v := range x {
		hurdles[v] = true
	}
	dp := make([][2]int, l+10)
	for i := 0; i < l+10; i++ {
		for j := 0; j < 2; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	for i := 0; i < l; i++ {
		//行動1
		action1 := i + 1
		if hurdles[action1] {
			dp[action1][0] = Min(dp[action1][0], dp[i][0]+t[0]+t[2])
		} else {
			dp[action1][0] = Min(dp[action1][0], dp[i][0]+t[0])
		}
		//行動2
		dp[i+1][1] = Min(dp[i+1][1], dp[i][0]+t[0]/2+t[1]/2)
		if hurdles[i+2] {
			dp[i+2][0] = Min(dp[i+2][0], dp[i][0]+t[0]+t[1]+t[2])
		} else {
			dp[i+2][0] = Min(dp[i+2][0], dp[i][0]+t[0]+t[1])
		}
		//行動3
		for j := 1; j < 4; j++ {
			dp[i+j][1] = Min(dp[i+j][1], dp[i][0]+t[0]/2+(j-1)*t[1]+t[1]/2)
		}
		if hurdles[i+4] {
			dp[i+4][0] = Min(dp[i+4][0], dp[i][0]+t[0]+3*t[1]+t[2])
		} else {
			dp[i+4][0] = Min(dp[i+4][0], dp[i][0]+t[0]+3*t[1])
		}
	}
	ans := Min(dp[l][0], dp[l][1])
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
