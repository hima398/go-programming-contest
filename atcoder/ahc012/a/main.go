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
	"time"
)

const INF = 1 << 60

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

var start int64

type Strawberry struct {
	x, y  int
	group string
}

type Line struct {
	i              int
	x1, y1, x2, y2 int
	w              int
}

var ls []Line

type Remain struct {
	x, y int
	w    int
}

var rs []Remain

var ss []Strawberry

func solve01(n, k int, a, x, y []int) (int, [][]int) {
	rand.Seed(time.Now().UnixNano())
	var detail [][]int
	mx := int(1e4)
	for i := 0; i < k; i++ {
		x1 := rand.Intn(2*mx) - mx
		y1 := rand.Intn(2*mx) - mx
		x2 := rand.Intn(2*mx) - mx
		y2 := rand.Intn(2*mx) - mx
		detail = append(detail, []int{x1, y1, x2, y2})
	}
	return k, detail
}

//O(10**1)
func ComputeScore(a, b []int) int {
	var deno, mole int
	for i := 0; i < 10; i++ {
		mole += Min(a[i], b[i])
		deno += a[i]
	}
	return mole
}

func isUpper(x, y, x1, y1, x2, y2 int) bool {
	yl := (y2-y1)*(x-x1) + (x2-x1)*y1
	return y*(x2-x1) > yl
}

//O(1000*100+マッピングのコスト)
func ComputeCake(n int, detail [][]int) []int {
	cs := make([]Strawberry, n)
	copy(cs, ss)
	//苺の数だけ直線を評価
	for i, berry := range cs {
		for j, line := range detail {
			var group string //:= strconv.Itoa(j)
			if isUpper(berry.x, berry.y, line[0], line[1], line[2], line[3]) {
				group += strconv.Itoa(j)
			} else {
				//group += "D"
			}
			cs[i].group += group
		}
	}
	m := make(map[string]int)
	for _, berry := range cs {
		m[berry.group]++
	}

	res := make([]int, 10)
	for _, v := range m {
		if v < 10 {
			res[v]++
		}
	}
	return res
}

func Rotate(x1, y1, x2, y2 int, rad float64) (int, int, int, int) {
	fx1, fy1, fx2, fy2 := float64(x1), float64(y1), float64(x2), float64(y2)
	nx1 := math.Cos(rad)*fx1 - math.Sin(rad)*fy1
	ny1 := math.Sin(rad)*fx1 + math.Cos(rad)*fy1
	nx2 := math.Cos(rad)*fx2 - math.Sin(rad)*fy2
	ny2 := math.Sin(rad)*fx2 + math.Cos(rad)*fy2
	return int(nx1), int(ny1), int(nx2), int(ny2)
}

func solve02(n, k int, a, x, y []int) (int, [][]int) {
	for i := 0; i < n; i++ {
		ss = append(ss, Strawberry{x[i], y[i], ""})
	}
	rand.Seed(time.Now().UnixNano())

	ans := k
	var detail [][]int
	mx := int(1e4)
	for i := 0; i < ans; i++ {
		x1 := rand.Intn(2*mx) - mx
		y1 := rand.Intn(2*mx) - mx
		x2 := rand.Intn(2*mx) - mx
		y2 := rand.Intn(2*mx) - mx
		detail = append(detail, []int{x1, y1, x2, y2})
	}
	b := ComputeCake(n, detail)
	score := ComputeScore(a, b)
	now := time.Now().UnixNano()
	for now-start < 5*int64(1e8) {
		nk := rand.Intn(ans)
		var newDetail [][]int
		for i := 0; i < nk; i++ {
			x1 := rand.Intn(2*mx) - mx
			y1 := rand.Intn(2*mx) - mx
			x2 := rand.Intn(2*mx) - mx
			y2 := rand.Intn(2*mx) - mx
			newDetail = append(newDetail, []int{x1, y1, x2, y2})
		}
		nb := ComputeCake(n, newDetail)
		newScore := ComputeScore(a, nb)
		if newScore > score {
			ans = nk
			detail = newDetail
			score = newScore
		}
		now = time.Now().UnixNano()
	}
	//PrintHorizonaly(b)
	//score := ComputeScore(a, b)
	//fmt.Println(score)
	cnt := 0
	dir := []int{1, -1}
	for now-start < 2*int64(1e9)+6*int64(1e8) {
		cnt++
		//fmt.Println("cnt = ", cnt, " time = ", now-start)
		if cnt%2 == 0 {
			var newDetail [][]int
			w := rand.Intn(2)
			if w%2 == 0 {
				if ans == 100 {
					continue
				}
				newDetail = make([][]int, len(detail))
				copy(newDetail, detail)

				x1 := rand.Intn(2*mx) - mx
				y1 := rand.Intn(2*mx) - mx
				x2 := rand.Intn(2*mx) - mx
				y2 := rand.Intn(2*mx) - mx
				newDetail = append(newDetail, []int{x1, y1, x2, y2})
			} else {
				del := rand.Intn(ans)
				for i := range detail {
					if i == del {
						continue
					}
					newDetail = append(newDetail, detail[i])
				}
			}
			nb := ComputeCake(n, newDetail)
			newScore := ComputeScore(a, nb)
			if newScore > score {
				//fmt.Printf("score:%d, new:%d\n", score, newScore)
				score = newScore
				ans += dir[w%2]
				detail = newDetail
			}
		} else {
			idx := rand.Intn(ans)
			bx1, by1, bx2, by2 := detail[idx][0], detail[idx][1], detail[idx][2], detail[idx][3]
			rad := math.Pi / 180.0
			rad *= float64(rand.Intn(180))
			detail[idx][0], detail[idx][1], detail[idx][2], detail[idx][3] = Rotate(bx1, by1, bx2, by2, rad)
			nb := ComputeCake(n, detail)
			newScore := ComputeScore(a, nb)
			if newScore <= score {
				detail[idx][0], detail[idx][1], detail[idx][2], detail[idx][3] = bx1, by1, bx2, by2
			}
		}
		now = time.Now().UnixNano()
	}
	//fmt.Println("cnt = ", cnt)
	return ans, detail
}

func updateRemain(n int, detail [][]int) {
	cs := make([]Strawberry, n)
	copy(cs, ss)

	//苺の数だけ直線を評価
	for i, berry := range cs {
		for j, line := range detail {
			var group string //:= strconv.Itoa(j)
			if isUpper(berry.x, berry.y, line[0], line[1], line[2], line[3]) {
				group += strconv.Itoa(j)
			} else {
				//group += "D"
			}
			cs[i].group += group
		}
	}

	mr := make(map[string]Remain)
	for _, berry := range cs {
		if _, found := mr[berry.group]; found {
			r := Remain{(mr[berry.group].x + berry.x) / 2, (mr[berry.group].y + berry.y) / 2, mr[berry.group].w + 1}
			mr[berry.group] = r
		} else {
			mr[berry.group] = Remain{berry.x, berry.y, 1}
		}
	}

	var nrs []Remain
	for _, v := range mr {
		nrs = append(nrs, v)
	}
	if len(nrs) == 0 {
		return
	}
	sort.Slice(nrs, func(i, j int) bool {
		return nrs[i].w > nrs[j].w
	})
	rs = nrs
}

//O(1000*100+マッピングのコスト)
func ComputeCake2(n int, detail [][]int) []int {
	cs := make([]Strawberry, n)
	copy(cs, ss)
	//苺の数だけ直線を評価
	for i, berry := range cs {
		for j, line := range detail {
			var group string //:= strconv.Itoa(j)
			if isUpper(berry.x, berry.y, line[0], line[1], line[2], line[3]) {
				group += strconv.Itoa(j)
			} else {
				//group += "D"
			}
			cs[i].group += group
		}
	}
	m := make(map[string]int)
	for _, berry := range cs {
		m[berry.group]++
	}
	res := make([]int, 10)
	for _, v := range m {
		if v < 10 {
			res[v]++
		}
	}
	return res
}

func ComputeDist2(x1, y1, x2, y2 int) int {
	xx := x2 - x1
	yy := y2 - y1
	return xx*xx + yy*yy
}

func ComputeSimilarLine(detail [][]int) {
	var nls []Line
	for idx := 0; idx < len(detail); idx++ {
		nls = append(nls, Line{idx, detail[idx][0], detail[idx][1], detail[idx][2], detail[idx][3], INF})
	}
	for i := 0; i < len(nls)-1; i++ {
		for j := i + 1; j < len(nls); j++ {
			if nls[i].x1 > nls[i].x2 {
				nls[i].x1, nls[i].y1, nls[i].x2, nls[i].y2 = nls[i].x2, nls[i].y2, nls[i].x1, nls[i].y1
			}
			if nls[j].x1 > nls[j].x2 {
				nls[j].x1, nls[j].y1, nls[j].x2, nls[j].y2 = nls[j].x2, nls[j].y2, nls[j].x1, nls[j].y1
			}
			nls[i].w = Min(nls[i].w, ComputeDist2(nls[i].x1, nls[i].y1, nls[j].x1, nls[j].y1)+ComputeDist2(nls[i].x2, nls[i].y2, nls[j].x2, nls[j].y2))
			nls[j].w = Min(nls[j].w, nls[i].w)
		}
	}
	sort.Slice(len(nls), func(i, j int) bool {
		return nls[i].w < nls[j].w
	})
	ls = nls
}

func solve03(n, k int, a, x, y []int) (int, [][]int) {
	for i := 0; i < n; i++ {
		ss = append(ss, Strawberry{x[i], y[i], ""})
	}
	rand.Seed(time.Now().UnixNano())

	ans := k
	var detail [][]int
	mx := int(1e4)
	for i := 0; i < ans; i++ {
		x1 := rand.Intn(2*mx) - mx
		y1 := rand.Intn(2*mx) - mx
		x2 := rand.Intn(2*mx) - mx
		y2 := rand.Intn(2*mx) - mx
		detail = append(detail, []int{x1, y1, x2, y2})
	}
	b := ComputeCake(n, detail)
	score := ComputeScore(a, b)
	now := time.Now().UnixNano()
	for now-start < 5*int64(1e8) {
		nk := rand.Intn(ans)
		var newDetail [][]int
		for i := 0; i < nk; i++ {
			x1 := rand.Intn(2*mx) - mx
			y1 := rand.Intn(2*mx) - mx
			x2 := rand.Intn(2*mx) - mx
			y2 := rand.Intn(2*mx) - mx
			newDetail = append(newDetail, []int{x1, y1, x2, y2})
		}
		nb := ComputeCake(n, newDetail)
		newScore := ComputeScore(a, nb)
		if newScore > score {
			ans = nk
			detail = newDetail
			score = newScore
		}
		now = time.Now().UnixNano()
	}
	//PrintHorizonaly(b)
	//score := ComputeScore(a, b)
	//fmt.Println(score)
	cnt := 0
	dir := []int{1, -1}
	for now-start < 2*int64(1e9)+6*int64(1e8) {
		//fmt.Println("cnt = ", cnt)
		cnt++
		//fmt.Println("cnt = ", cnt, " time = ", now-start)
		if cnt%2 == 0 {

			var newDetail [][]int
			w := rand.Intn(2)
			if w%2 == 0 {
				if ans == 100 {
					continue
				}
				updateRemain(n, detail)
				if len(rs) < 2 {
					continue
				}
				newDetail = make([][]int, len(detail))
				copy(newDetail, detail)

				x1 := rs[0].x
				y1 := rs[0].y
				x2 := rs[1].x
				y2 := rs[1].y
				newDetail = append(newDetail, []int{x1, y1, x2, y2})
			} else {
				if len(ls) == 0 {
					continue
				}
				//ComputeSimilarLine(detail)
				//del := ls[0].i
				del := rand.Intn(ans)

				for j := range detail {
					if j == del {
						continue
					}
					newDetail = append(newDetail, detail[j])
				}
			}
			nb := ComputeCake(n, newDetail)
			newScore := ComputeScore(a, nb)
			if newScore > score {
				//fmt.Printf("score:%d, new:%d\n", score, newScore)
				score = newScore
				ans += dir[w%2]
				detail = newDetail
			}
		} else {
			if len(ls) == 0 {
				continue
			}
			//ComputeSimilarLine(detail)
			idx := rand.Intn(ans)

			bx1, by1, bx2, by2 := detail[idx][0], detail[idx][1], detail[idx][2], detail[idx][3]
			rad := math.Pi / 180.0
			rad *= float64(rand.Intn(180))
			detail[idx][0], detail[idx][1], detail[idx][2], detail[idx][3] = Rotate(bx1, by1, bx2, by2, rad)
			nb := ComputeCake(n, detail)
			newScore := ComputeScore(a, nb)
			if newScore <= score {
				detail[idx][0], detail[idx][1], detail[idx][2], detail[idx][3] = bx1, by1, bx2, by2
			}
		}
		now = time.Now().UnixNano()
	}
	//fmt.Println("cnt = ", cnt)
	return ans, detail
}

func main() {
	start = time.Now().UnixNano()
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a := nextIntSlice(10)
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	k, p := solve02(n, k, a, x, y)
	PrintInt(k)
	PrintVertically(p)
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

func PrintVertically(x [][]int) {
	defer out.Flush()
	for _, v := range x {
		PrintHorizonaly(v)
		//fmt.Fprintln(out, v)
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

type IntQueue struct {
	q []int
}

func NewIntQueue() *IntQueue {

	return new(IntQueue)
}
func (this *IntQueue) Push(v int) {
	this.q = append(this.q, v)
}

func (this *IntQueue) Pop() (int, error) {
	if this.Size() == 0 {
		return -1, errors.New("")
	}
	ret := this.q[0]
	this.q = this.q[1:]
	return ret, nil
}

func (this *IntQueue) Size() int {
	return len(this.q)
}

func (this *IntQueue) PrintQueue() {
	fmt.Println(this.q)
}

type Pos struct {
	X int
	Y int
	D int
}

type Queue struct {
	ps []Pos
}

func NewQueue() *Queue {
	return new(Queue)
}

func (this *Queue) Push(p Pos) {
	this.ps = append(this.ps, p)
}

func (this *Queue) Pop() *Pos {
	if len(this.ps) == 0 {
		return nil
	}
	p := this.ps[0]
	this.ps = this.ps[1:]
	return &p
}

func (this *Queue) Find(x, y int) bool {
	for _, v := range this.ps {
		if x == v.X && y == v.Y {
			return true
		}
	}
	return false
}

func (this *Queue) Size() int {
	return len(this.ps)
}

type UnionFind struct {
	par  []int // parent numbers
	rank []int // height of tree
	size []int
}

func NewUnionFind(n int) *UnionFind {
	if n <= 0 {
		return nil
	}
	u := new(UnionFind)
	// for accessing index without minus 1
	u.par = make([]int, n+1)
	u.rank = make([]int, n+1)
	u.size = make([]int, n+1)
	for i := 0; i <= n; i++ {
		u.par[i] = i
		u.rank[i] = 0
		u.size[i] = 1
	}
	return u
}

func (this *UnionFind) Find(x int) int {
	if this.par[x] == x {
		return x
	} else {
		// compress path
		// ex. Find(4)
		// 1 - 2 - 3 - 4
		// 1 - 2
		//  L-3
		//  L 4
		this.par[x] = this.Find(this.par[x])
		return this.par[x]
	}
}

func (this *UnionFind) Size(x int) int {
	return this.size[this.Find(x)]
}

func (this *UnionFind) ExistSameUnion(x, y int) bool {
	return this.Find(x) == this.Find(y)
}

func (this *UnionFind) Unite(x, y int) {
	x = this.Find(x)
	y = this.Find(y)
	if x == y {
		return
	}
	// rank
	if this.rank[x] < this.rank[y] {
		//yがrootの木にxがrootの木を結合する
		this.par[x] = y
		this.size[y] += this.size[x]
	} else {
		// this.rank[x] >= this.rank[y]
		//xがrootの木にyがrootの木を結合する
		this.par[y] = x
		this.size[x] += this.size[y]
		if this.rank[x] == this.rank[y] {
			this.rank[x]++
		}
	}
}

func PrintUnionFind(u *UnionFind) {
	// for debuging. not optimize.
	fmt.Println(u.par)
	fmt.Println(u.rank)
	fmt.Println(u.size)
}

type BinaryIndexedTree struct {
	n     int
	nodes []int
	eval  func(x1, x2 int) int
}

func NewBinaryIndexTree(n int, f func(x1, x2 int) int) *BinaryIndexedTree {
	bt := new(BinaryIndexedTree)
	// 1-indexed
	bt.n = n + 1
	bt.nodes = make([]int, bt.n)
	bt.eval = f
	return bt
}

//i(0-indexed)をvに更新する
func (bt *BinaryIndexedTree) Update(i, v int) {
	//bt内部では1-indexedなのでここでインクリメントする
	i++
	for i < bt.n {
		bt.nodes[i] = bt.eval(bt.nodes[i], v)
		i += i & -1
	}
}

//i(0-indexed)の値を取得する
func (bt *BinaryIndexedTree) Query(i int) int {
	i++
	res := 0
	for i > 0 {
		res = bt.eval(bt.nodes[i], res)
		i -= i & -i
	}
	return res
}
