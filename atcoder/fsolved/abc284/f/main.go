package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	t := nextString()

	//s, i := solveZAlgorithm(n, t)
	s, i := solveRollingHash(n, t)
	if i >= 0 {
		PrintString(s)
	}
	PrintInt(i)
}

func solveZAlgorithm(n int, t string) (string, int) {
	reverse := func(s string) string {
		n := len(s)
		t := strings.Split(s, "")
		for i := 0; i < n/2; i++ {
			j := n - 1 - i
			t[i], t[j] = t[j], t[i]
		}
		return strings.Join(t, "")
	}
	//fmt.Println("n = ", n, len(t))
	a := t[:n]
	b := reverse(t[n:])

	za := ZAlgorithm(a + b)
	za = append(za, 0)

	za2 := ZAlgorithm(b + a)
	za2 = append(za2, 0)

	for i := 0; i <= n; i++ {
		if za[2*n-i] < i {
			continue
		}
		if za2[n+i] < n-i {
			continue
		}
		return t[:i] + t[n+i:], i
	}
	return "", -1
}

func ZAlgorithm(s string) []int {
	n := len(s)
	if n == 0 {
		return []int{}
	}
	z := make([]int, n)
	z[0] = n
	i, j := 1, 0
	for i < n {
		for i+j < n && s[j] == s[i+j] {
			j++
		}
		z[i] = j
		if j == 0 {
			i++
			continue
		}
		k := 1
		for i+k < n && k+z[k] < j {
			//z[i+k] = z[k]
			k++
		}
		i += k
		j -= k
	}
	return z
}

func solveRollingHash(n int, t string) (string, int) {
	const p = 2147483647
	r := make([]string, 2*n)
	for i := 0; i < 2*n; i++ {
		j := 2*n - 1 - i
		r[i] = string(t[j])
	}
	//fmt.Println(t)
	//fmt.Println(strings.Join(r, ""))
	h1 := NewRollingRollingHash()
	h2 := NewRollingRollingHash()

	h1.Init(t, p)
	h2.Init(strings.Join(r, ""), p)
	for i := 0; i < n; i++ {
		hash := h1.computeHash(i, i+n)
		rev := h2.computeExcludedHash(n-i, 2*n-i)
		//fmt.Println("hash, rev = ", hash, rev)
		if hash == rev {
			return strings.Join(r[n-i:2*n-i], ""), i
		}
	}
	return "", -1
}

type RollingHash struct {
	p int
	n int
	s string
	w []int
	h []int
}

func NewRollingRollingHash() *RollingHash {
	return new(RollingHash)
}

func (hash *RollingHash) Init(s string, p int) {
	hash.p = p
	hash.n = len(s)
	hash.s = s
	tt := make([]int, hash.n+1)
	for i := 1; i <= hash.n; i++ {
		tt[i] = int(hash.s[i-1]-'a') + 1
	}
	hash.w = make([]int, hash.n+1)
	hash.w[0] = 1
	for i := 1; i <= hash.n; i++ {
		hash.w[i] = 100 * hash.w[i-1] % p
	}
	hash.h = make([]int, hash.n+1)
	for i := 1; i <= hash.n; i++ {
		hash.h[i] = (100*hash.h[i-1] + tt[i]) % p
	}
}

func (hash *RollingHash) computeHash(l, r int) int {
	res := hash.h[r] - (hash.h[l] * hash.w[r-l] % hash.p)
	if res < 0 {
		res += hash.p
	}
	return res
}

func (hash *RollingHash) computeExcludedHash(l, r int) int {
	t := hash.h[hash.n] - (hash.h[r] * hash.w[hash.n-r] % hash.p)
	//fmt.Println("t = ", t)
	res := hash.h[l] * hash.w[hash.n-r] % hash.p
	res = (res + t) % hash.p
	if res < 0 {
		res += hash.p
	}
	return res
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

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

func nextString() string {
	sc.Scan()
	return sc.Text()
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

//重複組み合わせ H
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

func DivideSlice(A []int, K int) ([]int, []int, error) {

	if len(A) < K {
		return nil, nil, errors.New("")
	}
	return A[:K+1], A[K:], nil
}
