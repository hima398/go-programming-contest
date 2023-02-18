package main

import (
	"bufio"
	"fmt"
	"math/big"
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
	a := nextString()
	b := nextString()
	ans := solve(n, a, b)
	PrintInt(ans)
}

func solve(n int, a, b string) int {
	if a > b {
		a, b = b, a
	}
	ta := strings.Split(a, "")
	tb := strings.Split(b, "")
	for i := range a {
		if ta[i] > tb[i] {
			ta[i], tb[i] = tb[i], ta[i]
		}
	}
	sa, sb := strings.Join(ta, ""), strings.Join(tb, "")
	bp := big.NewInt(998244353)
	ba := big.NewInt(0)
	ba.SetString(sa, 10)
	bb := big.NewInt(0)
	bb.SetString(sb, 10)
	ans := ba.Mul(ba, bb)
	ans = ans.Mod(ans, bp)
	return int(ans.Int64())
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
