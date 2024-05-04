package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"math/bits"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w, t := nextInt(), nextInt(), nextInt()
	var a []string
	for i := 0; i < h; i++ {
		a = append(a, nextString())
	}
	ans := solve(h, w, t, a)
	PrintInt(ans)
}

func solve(h, w, t int, a []string) int {
	const INF = 1 << 60
	type cell struct {
		i, j int
	}

	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	var si, sj, gi, gj int
	var candies []cell
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if a[i][j] == 'S' {
				si, sj = i, j
			} else if a[i][j] == 'G' {
				gi, gj = i, j
			} else if a[i][j] == 'o' {
				candies = append(candies, cell{i, j})
			}
		}
	}
	//bfs
	bfs := func(si, sj int) [][]int {
		dist := make([][]int, h)
		for i := range dist {
			dist[i] = make([]int, w)
			for j := range dist[i] {
				dist[i][j] = INF
			}
		}
		var q []cell
		q = append(q, cell{si, sj})
		//お菓子を0個取ってたどり着ける最小の距離
		dist[si][sj] = 0
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for k := 0; k < 4; k++ {
				ni, nj := cur.i+di[k], cur.j+dj[k]
				if ni < 0 || ni >= h || nj < 0 || nj >= w {
					continue
				}
				if a[ni][nj] == '#' {
					continue
				}
				if dist[ni][nj] != INF {
					continue
				}
				q = append(q, cell{ni, nj})

				if dist[ni][nj] < INF {
					continue
				}
				dist[ni][nj] = dist[cur.i][cur.j] + 1
				q = append(q, cell{ni, nj})

				/*
					if dist[ni][nj] > dist[cur.i][cur.j]+1 {
						dist[ni][nj] = dist[cur.i][cur.j] + 1
						q = append(q, cell{ni, nj})
					}
				*/

			}
		}
		return dist
	}
	checkDist := bfs(si, sj)

	//これ以降答えが0以上である
	n := len(candies)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i, c := range candies {
		dist := bfs(c.i, c.j)
		for j := range candies {
			matrix[i][j] = dist[candies[j].i][candies[j].j]
			matrix[j][i] = matrix[i][j]
		}
	}

	mask := (1 << n)
	var ps []int
	for i := 1; i < mask; i++ {
		ps = append(ps, i)
	}
	//fmt.Println(ps)
	sort.Slice(ps, func(i, j int) bool {
		if bits.OnesCount(uint(ps[i])) == bits.OnesCount(uint(ps[j])) {
			return ps[i] < ps[j]
		}
		return bits.OnesCount(uint(ps[i])) < bits.OnesCount(uint(ps[j]))
	})
	dp := make([][]int, mask)
	for pat := range dp {
		dp[pat] = make([]int, n)
		for i := range dp[pat] {
			dp[pat][i] = INF
		}
	}
	for i, c := range candies {
		dp[1<<i][i] = checkDist[c.i][c.j]
	}

	for _, pat := range ps {
		for i := 0; i < n; i++ {
			//パターンと今いる位置が不一致
			if (pat>>i)&1 == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				//訪問済み
				if (pat>>j)&1 == 1 {
					continue
				}
				np := pat | (1 << j)

				dp[np][j] = Min(dp[np][j], dp[pat][i]+matrix[i][j])
			}
		}
	}

	//お菓子をi個取ってゴールに向かう距離
	minDist := make([]int, n+1)
	for i := range minDist {
		minDist[i] = INF
	}
	dist := bfs(gi, gj)
	minDist[0] = dist[si][sj]
	for _, pat := range ps {
		idx := bits.OnesCount(uint(pat))
		for i := range dp[pat] {
			minDist[idx] = Min(minDist[idx], dp[pat][i]+dist[candies[i].i][candies[i].j])
		}
	}

	//fmt.Println(minDist)
	ans := -1
	for i := range minDist {
		if minDist[i] <= t {
			ans = Max(ans, i)
		}
	}
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
