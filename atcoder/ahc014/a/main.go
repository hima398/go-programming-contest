package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

var di = []int{-1, 0, 1, 0}
var dj = []int{0, -1, 0, 1}

//var dir = []string{"F", "B", "L", "R"}
//問題文の並びより、"L"が0の方が都合が良いのでこの順番にしておく
var dir = []string{"L", "R", "F", "B"}

var d = [4]int{}
var d2 int

const numCandies = 100

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	f := nextIntSlice(100)
	for i := 0; i < numCandies; i++ {
		d[f[i]]++
	}
	for k := 1; k <= 3; k++ {
		d2 += d[k] * d[k]
	}
	//fmt.Println("d = ", d)
	//fmt.Println("d2 = ", d2)
	//solveRandom(f)
	//solveGreedily(f)
	solve(f)
}

func solve(f []int) {
	//s := make([][]int, 10)
	//for i := 0; i < 10; i++ {
	//	s[i] = make([]int, 100)
	//}
	const numTry = 2000
	counts := make([][4]int, 100)
	scores := make([][4]int, 100)
	for try := 0; try < numTry; try++ {
		box := make([][]int, 10)
		for i := 0; i < 10; i++ {
			box[i] = make([]int, 10)
		}
		var idxs []int
		for k := 0; k < numCandies; k++ {
			var candidate []int
			for i := 0; i < 10; i++ {
				for j := 0; j < 10; j++ {
					if box[i][j] == 0 {
						candidate = append(candidate, i*10+j)
					}
				}
			}
			pi := rand.Intn(len(candidate))
			p := candidate[pi]
			box[p/10][p%10] = f[k]
			idx := rand.Intn(4)
			idxs = append(idxs, idx)
			box = simulate(idx, box)
		}
		s := computeScore(box, d2)
		for i := 0; i < 100; i++ {
			scores[i][idxs[i]] += s
			counts[i][idxs[i]]++
		}
	}
	for k := 0; k < numCandies; k++ {
		_ = nextInt()
		idx := -1
		s := -1
		for kk := 0; kk < 4; kk++ {
			if counts[k][kk] == 0 {
				continue
			}
			if s < scores[k][kk]/counts[k][kk] {
				idx = kk
				s = scores[k][kk] / counts[k][kk]
			}
		}
		PrintString(dir[idx])
	}
}

func computeUnion(box [][]int) int {
	buf := make([]int, 100)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			buf[10*i+j] = box[i][j]
		}
	}
	uf := NewUnionFind(100)
	for idx := 0; idx < 100; idx++ {
		i, j := idx/10, idx%10
		for k := 0; k < 4; k++ {
			ni, nj := i+di[k], j+dj[k]
			if ni < 0 || ni >= 10 || nj < 0 || nj >= 10 {
				continue
			}
			nIdx := 10*ni + nj
			if uf.ExistSameUnion(idx, nIdx) {
				continue
			}
			if buf[idx] != 0 && buf[idx] == buf[nIdx] {
				uf.Unite(idx, nIdx)
			}
		}
	}
	var res int
	for i := 0; i < 100; i++ {
		if i == uf.parent[i] {
			res += uf.Size(i) * uf.Size(i)
		}
	}
	return res

}

func computeScore(box [][]int, d2 int) int {
	n2 := computeUnion(box)
	res := float64(n2) / float64(d2)
	return int(math.Round(1e6 * res))
}

func simulateLeft(box [][]int) [][]int {
	newBox := make([][]int, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if box[i][j] != 0 {
				newBox[i] = append(newBox[i], box[i][j])
			}
		}
	}
	for i := 0; i < 10; i++ {
		for len(newBox[i]) < 10 {
			newBox[i] = append(newBox[i], 0)
		}
	}
	return newBox
}

//箱を時計回りに90度回転
func LotateBox(box [][]int) [][]int {
	res := make([][]int, 10)
	for i := 0; i < 10; i++ {
		res[i] = make([]int, 10)
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			res[j][9-i] = box[i][j]
		}
	}
	return res
}

func simulate(dir int, box [][]int) [][]int {
	newBox := box
	switch dir {
	case 0: // "L"
		newBox = simulateLeft(newBox)
	case 1: // "R"
		for k := 0; k < 2; k++ {
			newBox = LotateBox(newBox)
		}
		newBox = simulateLeft(newBox)
		for k := 0; k < 2; k++ {
			newBox = LotateBox(newBox)
		}
	case 2: // "F"
		for k := 0; k < 3; k++ {
			newBox = LotateBox(newBox)
		}
		newBox = simulateLeft(newBox)
		newBox = LotateBox(newBox)
	case 3: // "B"
		newBox = LotateBox(newBox)
		newBox = simulateLeft(newBox)
		for k := 0; k < 3; k++ {
			newBox = LotateBox(newBox)
		}
	}
	return newBox
}

func solveGreedily(f []int) {
	box := make([][]int, 10)
	for i := 0; i < 10; i++ {
		box[i] = make([]int, 10)
	}
	idx := -1
	s := 0
	//fmt.Println("d2 = ", d2)
	for i := 0; i < 100; i++ {
		p := nextInt()
		box[p/10][p%10] = f[i]
		copyBox := make([][]int, 100)
		for i := 0; i < 10; i++ {
			copyBox[i] = make([]int, 10)
			copy(copyBox[i], box[i])
		}
		var scores []int
		for k := 0; k < 4; k++ {
			box = simulate(k, copyBox)
			cs := computeScore(box, d2)
			scores = append(scores, cs)
			if s < cs {
				idx = k
				s = cs
			}
		}
		if scores[0] == scores[1] && scores[1] == scores[2] && scores[2] == scores[3] && scores[3] == scores[1] {
			idx = rand.Intn(4)
		}
		PrintString(dir[idx])
		box = simulate(idx, copyBox)
	}
}

func solveRandom(f []int) {
	_ = nextInt()
	for i := 0; i < 100; i++ {
		idx := rand.Intn(4)
		PrintString(dir[idx])
	}
}

type UnionFind struct {
	parent []int // parentent numbers
	rank   []int // height of tree
	size   []int
}

func New(n int) *UnionFind {
	return NewUnionFind(n)
}

func NewUnionFind(n int) *UnionFind {
	if n <= 0 {
		return nil
	}
	u := new(UnionFind)
	// for accessing index without minus 1
	u.parent = make([]int, n+1)
	u.rank = make([]int, n+1)
	u.size = make([]int, n+1)
	for i := 0; i <= n; i++ {
		u.parent[i] = i
		u.rank[i] = 0
		u.size[i] = 1
	}
	return u
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] == x {
		return x
	} else {
		// compress path
		// ex. Find(4)
		// 1 - 2 - 3 - 4
		// 1 - 2
		//  L-3
		//  L 4
		uf.parent[x] = uf.Find(uf.parent[x])
		return uf.parent[x]
	}
}

func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.Find(x)]
}

func (uf *UnionFind) ExistSameUnion(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind) Unite(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}
	// rank
	if uf.rank[x] < uf.rank[y] {
		//yがrootの木にxがrootの木を結合する
		uf.parent[x] = y
		uf.size[y] += uf.size[x]
	} else {
		// uf.rank[x] >= uf.rank[y]
		//xがrootの木にyがrootの木を結合する
		uf.parent[y] = x
		uf.size[x] += uf.size[y]
		if uf.rank[x] == uf.rank[y] {
			uf.rank[x]++
		}
	}
}

func PrintUnionFind(u *UnionFind) {
	// for debuging. not optimize.
	fmt.Println(u.parent)
	fmt.Println(u.rank)
	fmt.Println(u.size)
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
