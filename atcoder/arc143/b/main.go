package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solveHonestly(n int) int {
	fmt.Println(Factional(n * n))
	var s []int
	for i := 1; i <= n*n; i++ {
		s = append(s, i)
	}
	var ans int
	for {
		f := make([][]int, n)
		for i := 0; i < n; i++ {
			f[i] = make([]int, n)
		}
		for idx, v := range s {
			i, j := idx/n, idx%n
			f[i][j] = v
		}
		allOk := true
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				ok := false
				//同じ列のチェック
				for ii := 0; ii < n; ii++ {
					if ii == i {
						continue
					}
					ok = ok || f[i][j] < f[ii][j]
				}
				//同じ行のチェック
				for jj := 0; jj < n; jj++ {
					if jj == j {
						continue
					}
					ok = ok || f[i][j] > f[i][jj]
				}
				allOk = allOk && ok
			}
		}
		if allOk {
			//PrintVertically(f)
			//PrintString("")
			ans++
		} else {
			PrintVertically(f)
			PrintString("")
		}
		if !NextPermutation(sort.IntSlice(s)) {
			break
		}
	}
	fmt.Println(ans%(n*n) == 0, ans/(n*n))
	return ans
}

func solve(n int) int {
	const p = 998244353

	n2 := n * n
	cmb := NewCombination(n2, p)

	ans := cmb.Factional(n2)
	r := n2
	r = r * cmb.Combination(n2, 2*n-1) % p
	r = r * cmb.Factional(n-1) % p
	r = r * cmb.Factional(n-1) % p
	r = r * cmb.Factional((n-1)*(n-1)) % p
	//fmt.Println(ans, r)
	ans = (ans - r + p) % p
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	//ans := solveHonestly(n)
	ans := solve(n)
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
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

func PrintVertically(x [][]int) {
	defer out.Flush()
	for _, v := range x {
		//fmt.Fprintln(out, v)
		PrintHorizonaly(v)
	}
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
	return c.inv[x]4643

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
