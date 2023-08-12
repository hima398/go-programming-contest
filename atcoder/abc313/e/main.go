package main

import (
	"bufio"
	"errors"
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
	s := nextString()
	ans, err := solve(n, s)
	if err != nil {
		PrintInt(-1)
	} else {
		PrintInt(ans)
	}
}

func solve(n int, s string) (int, error) {
	for i := 0; i < n-1; i++ {
		if s[i] != '1' && s[i+1] != '1' {
			return -1, errors.New("Impossible")
		}
	}
	const p = 998244353
	//後ろからsi(0<=i<n)を消すために必要な操作回数
	dp := make([]int, n)
	//dp[n]=0
	dp[n-1] = 1

	for i := n - 1; i > 0; i-- {
		if s[i] == '1' {
			//次の操作でs[i-1]が1個に変換されてs[i]を消す
			//操作回数は1回だけで良い
			dp[i-1] = (dp[i] + 1) % p
		} else {
			//次の操作で1がs[i]個に変換されてs[i]を消す(1回)
			//s[i]を消すことで出てきたs[i]個の1を消す((s[i]-1)*dp[i]回)
			w := (1 + (int(s[i]-'0')-1)*dp[i]) % p
			dp[i-1] = (dp[i] + w) % p
		}
	}
	//s0は削除しなくて良いのでdp[0]ではなくdp[1]が答え
	return dp[1], nil
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
