package main

import (
	"bufio"
	"fmt"
	"math/big"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var a []string
	for i := 0; i < n; i++ {
		a = append(a, nextString())
	}
	ans := solve(n, a)
	PrintInt(ans)
}

func Reverse(s string) string {
	var res []string
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, string(s[i]))
	}
	return strings.Join(res, "")
}

func solve(n int, a []string) int {
	g := make([]*big.Int, n)
	for i, ai := range a {
		g[i] = big.NewInt(0)
		g[i], _ = g[i].SetString(Reverse(ai), 2)
	}
	var ans int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			//i-jが連結していない
			if g[i].Bit(j) == 0 {
				continue
			}

			k := big.NewInt(0)
			//i-jから連結されている頂点をbitで管理
			k = k.And(g[i], g[j])
			for _, v := range k.Bits() {
				ans += bits.OnesCount(uint(v))
			}
		}
	}
	ans /= 3
	return ans
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
