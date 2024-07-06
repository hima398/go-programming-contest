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

	n, k := nextInt(), nextInt()

	ans := solve(n, k)

	Print(ans)
}

func solve(n, k int) int {
	//if n == 1 {
	//	return 1
	//}

	const mod = 998244353

	n2 := Pow(n, 2, mod)
	//in2 := Inv(Pow(n, p, mod), mod)
	in := Inv(n-1, mod)
	p, q := make([]int, k+1), make([]int, k+1)
	p[1] = ((Pow(n-1, 2, mod) + 1) % mod) * Inv(n2, mod) % mod
	q[1] = (1 - p[1] + mod) % mod
	q[1] = q[1] * in % mod
	for i := 2; i <= k; i++ {
		p[i] = p[i-1]*p[1]%mod + ((n-1)*q[i-1]%mod)*q[1]%mod
		p[i] %= mod
		q[i] = (1 - p[i] + mod) % mod
		q[i] = q[i] * in % mod
	}
	ans := p[k]
	w := ((n + 2) * (n - 1) % mod) * Inv(2, mod) % mod
	w *= q[k]
	w %= mod
	ans += w
	ans %= mod
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
