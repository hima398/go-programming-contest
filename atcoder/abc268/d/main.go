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

func solve01(n, m int, s, t []string) (string, error) {
	var idx []int
	for i := 0; i < n; i++ {
		idx = append(idx, i)
	}
	type Node struct {
		i, mask int
	}
	var bm []Node
	for i := 0; i < n; i++ {
		bm = append(bm, Node{i, 0})
	}
	for k := 0; k < m; k++ {
		tmp := strings.Split(t[k], "_")

		for i := 0; i < n; i++ {
			for j := range tmp {
				// s[i]の文字はTkのj番目にできない
				if tmp[j] == s[i] {
					bm[i].mask = bm[i].mask | (1 << j)
				}
			}
		}
	}
	return "", nil
}

func Join(s []string, mask int) string {
	res := s[0]

	for i := 0; i < len(s)-1; i++ {
		if (mask>>i)&1 > 0 {
			res += "__"
		} else {
			res += "_"
		}
		res += s[i+1]
	}
	return res
}

func solve(n, m int, s, t []string) (string, error) {
	mt := make(map[string]struct{})
	for _, ts := range t {
		mt[ts] = struct{}{}
	}
	var idx []int
	for i := 0; i < n; i++ {
		idx = append(idx, i)
	}
	for {
		var u []string
		for i := 0; i < n; i++ {
			u = append(u, s[idx[i]])
		}
		c := strings.Join(u, "_")
		if _, found := mt[c]; !found {
			if len(c) >= 3 && len(c) <= 16 {
				return c, nil
			}
		}
		pat := 1
		for pat < 1<<(n-1)-1 {
			nc := Join(u, pat)
			_, found := mt[nc]
			if !found && len(nc) >= 3 && len(nc) < 17 {
				return nc, nil
			}
			pat <<= 1
		}
		if !NextPermutation(sort.IntSlice(idx)) {
			break
		}
	}

	return "", errors.New("Impossible")
}

func solveCommentary(n, m int, s, t []string) {
	//mt := make(map[string]struct{})
	mt := make(map[string]bool)
	for _, v := range t {
		//mt[v] = struct{}{}
		mt[v] = true
	}
	var idx []int
	for i := 0; i < n; i++ {
		idx = append(idx, i)
	}
	var ls int
	for _, si := range s {
		ls += len(si)
	}
	pattern := make([]string, n)
	var dfs func(idx, rem int, ans string)
	dfs = func(idx, rem int, ans string) {
		//文字列の候補にこれ以上連結できない
		if rem < 0 {
			return
		}
		//S'[i]を全て使った
		if idx == n {
			found := mt[ans]
			//_, found := mt[ans]
			if len(ans) >= 3 && !found {
				PrintString(ans)
				os.Exit(0)
			}
			return
		}
		if len(ans) > 0 && ans[len(ans)-1] != '_' {
			dfs(idx, rem, ans+"_")
		} else {
			dfs(idx+1, rem, ans+pattern[idx])
			if len(ans) > 0 {
				dfs(idx, rem-1, ans+"_")
			}
		}
	}
	for {
		for i := 0; i < n; i++ {
			pattern[i] = s[idx[i]]
		}
		dfs(0, 16-ls-(n-1), "")
		if !NextPermutation(sort.IntSlice(idx)) {
			break
		}
	}
	PrintInt(-1)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()

	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	var t []string
	for j := 0; j < m; j++ {
		t = append(t, nextString())
	}
	solveCommentary(n, m, s, t)
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
