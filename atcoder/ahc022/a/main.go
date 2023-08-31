package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

//	type Judge interface {
//		send(y, x, p int)
//		receive() int
//	}
type AnotherSpace interface {
	place(p [][]int)
	measure(i, y, x int) int
}

var anotherSpace AnotherSpace

type StandardIO struct {
}

func (anotherSpace StandardIO) place(p [][]int) {
	defer out.Flush()
	for _, v := range p {
		if len(v) <= 0 {
			return
		}
		n := len(v)
		fmt.Fprintf(out, "%d", v[0])
		for i := 1; i < n; i++ {
			fmt.Fprintf(out, " %d", v[i])
		}
		fmt.Fprintln(out)
	}
}

func (anotherSpace StandardIO) measure(i, y, x int) int {
	fmt.Fprintln(out, i, y, x)
	out.Flush()

	res := nextInt()
	return res
}

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	l, n, s := nextInt(), nextInt(), nextInt()
	y, x := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		y[i] = nextInt()
		x[i] = nextInt()
	}

	var st StandardIO
	anotherSpace = st

	ans := solve01(l, n, s, y, x)
	//solve02(l, n, s, y, x)

	if ans != nil {
		PrintVertically(ans)
	}
}

func solve02(l, n, s int, y, x []int) {
	ans := make([]int, n)
	PrintHorizonaly([]int{-1, -1, -1})
	PrintVertically(ans)

}

// サンプル同様、ゲートの場所に機材を置いて測定結果から推測
func solve01(l, n, s int, y, x []int) []int {
	//配置
	p := make([][]int, l)
	for i := range p {
		p[i] = make([]int, l)
	}
	for i := 0; i < n; i++ {
		p[y[i]][x[i]] = 10 * i
		//cnt++
	}
	anotherSpace.place(p)

	//計測
	//n*9<=1e4
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		v := anotherSpace.measure(i, 0, 0)
		if v < 0 {
			return nil
		}
		mn := 1 << 60
		for j := 0; j < n; j++ {
			diff := computeDist(p[y[j]][x[j]], v)
			if mn > diff {
				mn = diff
				ans[i] = j
			}
		}
	}

	//回答
	PrintHorizonaly([]int{-1, -1, -1})
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintString(x string) {
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}

func Sqrt(x int) int {
	x2 := int(math.Sqrt(float64(x))) - 1
	for (x2+1)*(x2+1) <= x {
		x2++
	}
	return x2
}

func computeDist(d1, d2 int) int {
	return Abs(d2 - d1)
}

func computeDist2(y1, x1, y2, x2 int) int {
	return Abs(y2-y1) + Abs(x2-x1)
}

func Gcd(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	/*
		if x < y {
			x, y = y, x
		}
	*/
	return Gcd(y, x%y)
}

func Lcm(x, y int) int {
	// x*yのオーバーフロー対策のため先にGcdで割る
	// Gcd(x, y)はxの約数のため割り切れる
	ret := x / Gcd(x, y)
	ret *= y
	return ret
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

func Permutation(N, K int) int {
	v := 1
	if 0 < K && K <= N {
		for i := 0; i < K; i++ {
			v *= (N - i)
		}
	} else if K > N {
		v = 0
	}
	return v
}

func Factional(N int) int {
	return Permutation(N, N-1)
}

func Combination(N, K int) int {
	if K == 0 {
		return 1
	}
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}

type Comb struct {
	n, p int
	fac  []int // Factional(i) mod p
	finv []int // 1/Factional(i) mod p
	inv  []int // 1/i mod p
}

func NewCombination(n, p int) *Comb {
	c := new(Comb)
	c.n = n
	c.p = p
	c.fac = make([]int, n+1)
	c.finv = make([]int, n+1)
	c.inv = make([]int, n+1)

	c.fac[0] = 1
	c.fac[1] = 1
	c.finv[0] = 1
	c.finv[1] = 1
	c.inv[1] = 1
	for i := 2; i <= n; i++ {
		c.fac[i] = c.fac[i-1] * i % p
		c.inv[i] = p - c.inv[p%i]*(p/i)%p
		c.finv[i] = c.finv[i-1] * c.inv[i] % p
	}
	return c
}

func (c *Comb) Factional(x int) int {
	return c.fac[x]
}

func (c *Comb) Combination(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	ret := c.fac[n] * c.finv[k]
	ret %= c.p
	ret *= c.finv[n-k]
	ret %= c.p
	return ret
}

// 重複組み合わせ H
func (c *Comb) DuplicateCombination(n, k int) int {
	return c.Combination(n+k-1, k)
}
func (c *Comb) Inv(x int) int {
	return c.inv[x]
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
