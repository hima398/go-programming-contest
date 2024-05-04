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

	n := nextInt()

	ans := solve(n)

	PrintInt(ans)
}

func solve(n int) int {
	const p = 998244353
	var dp [60][60][60]int
	var visited [60][60][60]bool
	dp[0][0][0] = 1
	visited[0][0][0] = true
	i5 := Inv(5, p)
	var x, y, z int
	for n%2 == 0 {
		n /= 2
		x++
	}
	for n%3 == 0 {
		n /= 3
		y++
	}
	for n%5 == 0 {
		n /= 5
		z++
	}
	if n != 1 {
		return 0
	}
	var f func(x, y, z int) int
	f = func(x, y, z int) int {
		if x < 0 || y < 0 || z < 0 {
			return 0
		}
		if visited[x][y][z] {
			return dp[x][y][z]
		}
		visited[x][y][z] = true
		var e int
		e += f(x-1, y, z)
		e += f(x, y-1, z)
		e += f(x-2, y, z)
		e += f(x, y, z-1)
		e += f(x-1, y-1, z)

		e *= i5
		e %= p
		dp[x][y][z] = e
		return dp[x][y][z]
	}
	ans := f(x, y, z)
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
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
