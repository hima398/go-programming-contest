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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	r := nextString()
	c := nextString()

	ans, err := solve(n, r, c)
	if err != nil {
		Print("No")
	} else {
		Print("Yes")
		for _, v := range ans {
			Print(strings.Join(v, ""))
		}
	}
}

// /nが3-5なので力技で構築
func buildIndexes(n int) [][]int {
	var res [][]int
	if n == 3 {
		return [][]int{{0, 1, 2}}
	} else if n == 4 {
		return [][]int{{0, 1, 2}, {0, 1, 3}, {1, 2, 3}}
	} else {
		//n==5
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				for k := j + 1; k < n; k++ {
					res = append(res, []int{i, j, k})
				}
			}
		}
		return res
	}
}

func solve(n int, r, c string) ([][]string, error) {
	pattern := []string{"ABC", "ACB", "BAC", "BCA", "CAB", "CBA"}
	indexes := buildIndexes(n)

	var candidates [][][]string
	var dfs func(i int, candidate [][]string, used []int)
	dfs = func(i int, candidate [][]string, used []int) {
		if i == 5 {
			candidates = append(candidates, candidate)
			return
		}
		row := strings.Split(strings.Repeat(".", n), "")
		for _, idx := range indexes {
			for _, j := range idx {
				row[j] = "#"
				used[j]++
				if used[j] > 3 {
					return
				}
			}
			cc := make([][]string, len(candidate))
			for i := range cc {
				copy(cc[i], candidate[i])
			}
			cc = append(cc, row)
			dfs(i+1, cc, used)
		}
	}
	dfs(0, make([][]string, 0), make([]int, n))

	Print(len(candidates))
	var dfs2 func(i int, template [][]string) [][]string
	dfs2 = func(i int, template [][]string) [][]string {
		if i == 5 {
			return template
		}
		idx := 0
		tt := make([][]string, n)
		for i := range tt {
			copy(tt[i], template[i])
		}
		for _, pat := range pattern {
			for j := 0; j < n; j++ {
				if tt[i][j] == "#" {
					tt[i][j] = string(pat[idx])
					idx++
				}
			}
			dfs2(i+1, tt)
		}
		return nil
	}
	printA := func(a [][]string) {
		for _, v := range a {
			Print(strings.Join(v, ""))
		}
		Print("")
	}
	//ABCのパターンを埋める
	for _, candidate := range candidates {
		a := dfs2(0, candidate)
		printA(a)
		//チェックをする
		ok := true
		for j := 0; j < n; j++ {
			m := make(map[string]struct{})
			for i := 0; i < n; i++ {
				if a[i][j] != "." {
					m[a[i][j]] = struct{}{}
				}
			}
			ok = ok && len(m) == 3
		}
		if !ok {
			continue
		}
		ok2 := true
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if a[i][j] == "." {
					continue
				} else {
					if a[i][j] != string(r[i]) {
						ok2 = false
					}
					break
				}
			}
		}
		if !ok2 {
			continue
		}
		ok3 := true
		for j := 0; j < n; j++ {
			for i := 0; i < n; i++ {
				if a[i][j] == "." {
					continue
				} else {
					if a[i][j] != string(c[j]) {
						ok3 = false
					}
					break
				}
			}
		}
		if !ok3 {
			continue
		}
		return a, nil
	}
	return nil, errors.New("Impossible")
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
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

func DivideSlice(A []int, K int) ([]int, []int, error) {

	if len(A) < K {
		return nil, nil, errors.New("")
	}
	return A[:K+1], A[K:], nil
}
