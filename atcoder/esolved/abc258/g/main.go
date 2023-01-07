package main

import (
	"bufio"
	"fmt"
	"math/big"
	"math/bits"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solveHonestly(n int, a []string) int {
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			g[i][j] = int(a[i][j] - '0')
		}
	}
	var ans int
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			if g[i][j] == 0 {
				continue
			}

			for k := j + 1; k < n; k++ {
				ans += g[i][k] & g[j][k]
			}
		}
	}
	return ans
}

func Reverse(s string) string {
	res := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func solve(n int, a []string) int {
	g := make([]*big.Int, n)
	for i := 0; i < n; i++ {
		g[i] = big.NewInt(0)
		g[i], _ = g[i].SetString(Reverse(a[i]), 2)
	}
	var ans int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if g[i].Bit(j) == 0 {
				continue
			}

			k := big.NewInt(0)
			k = k.And(g[i], g[j])
			for _, v := range k.Bits() {
				ans += bits.OnesCount(uint(v))
			}
		}
	}
	ans /= 3
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = nextString()
	}

	ans := solve(n, a)
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
