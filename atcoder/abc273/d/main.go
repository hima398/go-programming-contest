package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type Coordinate struct {
	r, c int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w, sr, sc := nextInt(), nextInt(), nextInt(), nextInt()
	n := nextInt()
	var r, c []int
	for i := 0; i < n; i++ {
		r = append(r, nextInt())
		c = append(c, nextInt())
	}
	q := nextInt()
	var d []string
	var l []int
	for i := 0; i < q; i++ {
		d = append(d, nextString())
		l = append(l, nextInt())
	}
	ans := solve(h, w, sr, sc, n, r, c, q, d, l)
	PrintVertically(ans)
}

func solve(h, w, sr, sc, n int, r, c []int, q int, d []string, l []int) []Coordinate {
	row, col := make([][]int, h+1), make([][]int, w+1)
	for i := 0; i < n; i++ {
		row[r[i]] = append(row[r[i]], c[i])
		col[c[i]] = append(col[c[i]], r[i])
	}
	for i := 0; i <= h; i++ {
		row[i] = append(row[i], 0)
		row[i] = append(row[i], w+1)
		sort.Ints(row[i])
	}
	for j := 0; j <= h; j++ {
		col[j] = append(col[j], 0)
		col[j] = append(col[j], h+1)
		sort.Ints(col[j])
	}
	/*
			[0 2 6]
		pos, idx =  -1 0
		pos, idx =  0 0
		pos, idx =  1 1
		pos, idx =  2 1
		pos, idx =  3 2
		pos, idx =  4 2
		pos, idx =  5 2
		pos, idx =  6 2
		pos, idx =  7 3
	*/
	var ans []Coordinate
	curR, curC := sr, sc
	for k := 0; k < q; k++ {
		switch d[k] {
		case "U", "D":
			lec := 1
			if d[k] == "U" {
				lec = -1
			}
			idx1 := sort.Search(len(col[curC]), func(idx int) bool {
				return curR <= col[curC][idx]
			})
			idx2 := sort.Search(len(col[curC]), func(idx int) bool {
				return curR+lec*l[k] <= col[curC][idx]
			})
			fmt.Printf("k = %d, curR = %d, nextR = %d, idx1 = %d, idx2 = %d ", k, curR, curR+lec*l[k], idx1, idx2)
			fmt.Println(col[curC])
			if idx2 == 0 {
				curR = 1
			} else if idx2 == len(col[curC]) {
				curR = h
			} else if idx1 == idx2 {
				curR += lec * l[k]
			} else {
				curR = col[curC][idx2] - lec*1
			}
			if curR > h {
				curR = h
			}
		case "L", "R":
			lec := 1
			if d[k] == "L" {
				lec = -1
			}
			idx1 := sort.Search(len(row[curR]), func(idx int) bool {
				return curC <= row[curR][idx]
			})
			idx2 := sort.Search(len(row[curR]), func(idx int) bool {
				return curC+lec*l[k] <= row[curR][idx]
			})
			fmt.Printf("k = %d, curC = %d, nextC = %d, idx1 = %d, idx2 = %d ", k, curC, curC+lec*l[k], idx1, idx2)
			fmt.Println(row[curR])
			if idx2 == 0 {
				curC = 1
			} else if idx2 == len(row[curR]) {
				curC = w
			} else if idx1 == idx2 {
				curC += lec * l[k]
			} else {
				curC = row[curR][idx2] - lec*1
			}
			if curC > w {
				curC = w
			}
		}
		ans = append(ans, Coordinate{curR, curC})
	}
	return ans
}

func firstsolve(h, w, sr, sc, n int, r, c []int, q int, d []string, l []int) []Coordinate {
	row, col := make([][]int, h+1), make([][]int, w+1)
	for i := 0; i < n; i++ {
		row[r[i]] = append(row[r[i]], c[i])
		col[c[i]] = append(col[c[i]], r[i])
	}
	for i := 0; i <= h; i++ {
		row[i] = append(row[i], 0)
		row[i] = append(row[i], w+1)
		sort.Ints(row[i])
	}
	for j := 0; j <= h; j++ {
		col[j] = append(col[j], 0)
		col[j] = append(col[j], h+1)
		sort.Ints(col[j])
	}
	curR, curC := sr, sc
	var ans []Coordinate
	for k := 0; k < q; k++ {
		//fmt.Println(d[k], l[k])
		if d[k] == "U" || d[k] == "D" {
			//列を基に壁の衝突を検知
			//fmt.Println(curR, curC, col[curC], w, curR+w*l[k])
			idx1 := sort.Search(len(col[curC]), func(idx int) bool {
				return curR < col[curC][idx]
			})
			var idx2 int
			if d[k] == "U" {
				idx2 = sort.Search(len(col[curC]), func(idx int) bool {
					return curR-l[k] < col[curC][idx]
				})
				if idx1 == idx2 {
					curR -= l[k]
				} else if idx2 <= 0 {
					curR = 1
				} else {
					curR = col[curC][idx2] + 1
				}
			} else {
				//d[k]=="D"
				idx2 = sort.Search(len(col[curC]), func(idx int) bool {
					return curR+l[k] < col[curC][idx]
				})
				if idx1 == idx2 {
					curR += l[k]
					if curR > h {
						curR = h
					}
				} else if idx2 >= len(col[curC]) {
					curR = h
				} else {
					curR = col[curC][idx2] - 1
				}
			}
			//fmt.Println(idx1, idx2)
			ans = append(ans, Coordinate{curR, curC})
		} else {
			//d[k]=="L" or d[k]=="R"
			//行をもとに壁の衝突を検知
			//fmt.Println("k = ", k, row[curR], curC, curC+w*l[k])
			idx1 := sort.Search(len(row[curR]), func(idx int) bool {
				return curC < row[curR][idx]
			})
			var idx2 int
			if d[k] == "L" {
				idx2 = sort.Search(len(row[curR]), func(idx int) bool {
					return curC-l[k] < row[curR][idx]
				})
				if idx1 == idx2 {
					curC -= l[k]
				} else if idx2 <= 0 {
					curC = 1
				} else {
					curC = row[curR][idx2] + 1
				}
			} else {
				idx2 = sort.Search(len(row[curR]), func(idx int) bool {
					return curC+l[k] < row[curR][idx]
				})
				if idx1 == idx2 {
					curC += l[k]
					if curC > w {
						curC = w
					}
				} else if idx2 >= len(row[curR]) {
					curC = w
				} else {
					curC = row[curR][idx2] - 1
				}
			}
			//fmt.Println(idx1, idx2)
			ans = append(ans, Coordinate{curR, curC})
		}
		//fmt.Println(curR, curC)
	}
	return ans
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

func PrintVertically(x []Coordinate) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintf(out, "%d %d", v.r, v.c)
		fmt.Fprintln(out)
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
