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

// インタラクティブに関する入出力
type Gateway interface {
	PlayCard(c OutputCard)
	SubscribeProject(m int) []Project
	SubscribeMoney() int
	//TODO:Dealはこのアプリ視点じゃないかもしれないので命名ちょっと考える
	DealCard(k int) []DealedCard
	DeclareCard(r int, c DealedCard)
}

type StandardIO struct{}

func (gate *StandardIO) PlayCard(c OutputCard) {
	fmt.Println(c.c, c.m)
}

// 最新のプロジェクトの情報を受け取る
func (gate *StandardIO) SubscribeProject(m int) []Project {
	var res []Project
	for i := 0; i < m; i++ {
		h, v := nextInt(), nextInt()
		res = append(res, Project{h, v, false})
	}
	return res
}

// 最新の所持金の情報を受け取る
func (gate *StandardIO) SubscribeMoney() int {
	return nextInt()
}

// 方針カードの候補を受け取る
func (gate *StandardIO) DealCard(k int) []DealedCard {
	var res []DealedCard
	for i := 0; i < k; i++ {
		t, w, p := nextInt(), nextInt(), nextInt()
		res = append(res, DealedCard{t, w, p})
	}
	return res
}

// 候補から引く方針カードを宣言する
func (gate *StandardIO) DeclareCard(r int, c DealedCard) {
	fmt.Println(r)
}

// 独自ドメイン定義
type CardType int

const (
	//通常労働カード
	RegularWork CardType = iota
	//全力労働カード
	HardWork
	//キャンセルカード
	Cancel
	//業務転換カード
	Restructuring
	//増資カード
	Investment
)

// 方針カード
type PolicyCard struct {
	t       int
	w       int
	trashed bool
}

// 手札
type Hands struct {
	n     int
	cs    []PolicyCard
	kinds []int
}

func (h *Hands) Init(n int, t, w []int) {
	h.n = n
	//TODO:カードが5種類をここに持ってしまっているので要リファクタリング
	h.kinds = make([]int, 5)
	for i := 0; i < h.n; i++ {
		h.cs = append(h.cs, PolicyCard{t[i], w[i], false})
		h.kinds[h.cs[i].t]++
	}
}

func (h *Hands) hasRegularWork() bool {
	return h.kinds[RegularWork] > 0
}

func (h *Hands) hasHardWork() bool {
	return h.kinds[HardWork] > 0
}

func (h *Hands) hasCancel() bool {
	return h.kinds[Cancel] > 0
}

func (h *Hands) hasRestructuring() bool {
	return h.kinds[Restructuring] > 0
}

func (h *Hands) hasInvestment() bool {
	return h.kinds[Investment] > 0
}

func (h *Hands) PlayCard(idx int) PolicyCard {
	h.cs[idx].trashed = true
	h.kinds[h.cs[idx].t]--
	return h.cs[idx]
}

func (h *Hands) TakeCard(idx int, c PolicyCard) {
	h.cs[idx] = c
	h.kinds[c.t]++
}

// ターン毎に配られるカード
type DealedCard struct {
	t, w int
	cost int
}

func createDeck(t, w, p []int) []DealedCard {
	var res []DealedCard
	for i := range t {
		res = append(res, DealedCard{t[i], w[i], p[i]})
	}
	return res
}

// 出力に渡すカード
type OutputCard struct {
	c, m int
	//エラーのトレース用
	turn int
	t    CardType
	w    int
}

// プロジェクト
type Project struct {
	remain   int
	value    int
	finished bool
}

func createProjects(h, v []int) []Project {
	//len(h) != len(v)のときおかしな結果になるが
	//入力である程度揃っているのでチェックは一旦省略しておく

	var res []Project
	for i := range h {
		res = append(res, Project{h[i], v[i], false})
	}
	return res
}

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

var gateway Gateway

var debug bool

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, k, turn := nextInt(), nextInt(), nextInt(), nextInt()
	var t, w []int
	for i := 0; i < n; i++ {
		t = append(t, nextInt())
		w = append(w, nextInt())
	}
	var h, v []int
	for i := 0; i < m; i++ {
		h = append(h, nextInt())
		v = append(v, nextInt())
	}

	gateway = new(StandardIO)

	//solve00(n, m, k, turn, t, w, h, v)
	solve01(n, m, k, turn, t, w, h, v)

}

func SolveRegularWork(hands *Hands, projects []Project) OutputCard {
	res := OutputCard{c: -1, m: -1, t: RegularWork}
	diff := 1 << 60
	for i, u := range hands.cs {
		if u.t != int(RegularWork) {
			continue
		}
		for j, v := range projects {
			d := Abs(v.remain - u.w)
			if d < diff {
				res.c, res.m = i, j
				res.w = u.w
				diff = d
			}
		}
	}
	return res
}

func printFirstHalf(turn int, hands *Hands, projects []Project, node []int, outputCard OutputCard, money int) {
	fmt.Println("# ", "turn:", turn, ".0")
	fmt.Println("# ", *hands)
	fmt.Println("# ", projects)
	fmt.Println("# ", node, outputCard)
	fmt.Println("# money = ", money)
}

func printSecondHalf(turn int, hands *Hands, projects []Project, node []int, dealedCard []DealedCard, money int) {
	fmt.Println("# turn:", turn, ".5")
	fmt.Println("# ", *hands)
	fmt.Println("# ", projects)
	fmt.Println("# ", dealedCard)
	fmt.Println("# ", node)
	fmt.Println("# money = ", money)
}

// 条件を改善
func solve02(n, m, k, turn int, t, w, h, v []int) ([][]int, int) {
	const INF = 1 << 60
	const canUseInvestment = 4

	hands := new(Hands)
	hands.Init(n, t, w)

	projects := createProjects(h, v)

	money := 0
	usedInvestmentCards := 0
	var ans [][]int

	for cnt := 0; cnt < turn; cnt++ {

		var playCandidate []OutputCard
		if hands.hasInvestment() && usedInvestmentCards < canUseInvestment {
			for i, u := range hands.cs {
				if u.t == int(Investment) {
					playCandidate = append(playCandidate, OutputCard{i, 0, cnt, HardWork, u.w})
					break
				}
			}
			usedInvestmentCards++
		} else {
			//手札から方針カードを選択して実行する
			//TODO:最適な選び方
			if hands.hasRegularWork() {
				candidate := SolveRegularWork(hands, projects)
				candidate.turn = cnt
				playCandidate = append(playCandidate, candidate)
			}
			if hands.hasHardWork() {
				for i, v := range hands.cs {
					if v.t == int(HardWork) {
						playCandidate = append(playCandidate, OutputCard{i, 0, cnt, HardWork, v.w})
						break
					}
				}
			}
			if hands.hasCancel() {
				candidate := OutputCard{turn: cnt, t: Cancel}
				for i, v := range hands.cs {
					if v.t == int(Cancel) {
						candidate.c = i
						break
					}
				}
				//使うプロジェクトを選ぶ
				//candidate.m = 0
				max := projects[0].remain
				for j, v := range projects {
					if v.remain > max {
						max = v.remain
						candidate.m = j
					}
				}
				playCandidate = append(playCandidate, candidate)
			}
			//TODO:最適な選び方
			if hands.hasRestructuring() {
				for i, v := range hands.cs {
					if v.t == int(Restructuring) {
						playCandidate = append(playCandidate, OutputCard{i, 0, cnt, Restructuring, v.w})
						break
					}
				}
			}
		}
		idx := rand.Intn(len(playCandidate))
		outputCard := playCandidate[idx]

		hands.PlayCard(playCandidate[idx].c)
		gateway.PlayCard(outputCard)
		node := []int{playCandidate[idx].c, playCandidate[idx].m}
		//ans = append(ans, []int{playCandidate[idx].c, playCandidate[idx].m})

		//turn.0のところでテスト出力
		if debug {
			printFirstHalf(cnt, hands, projects, node, outputCard, money)
		}

		//存在するプロジェクトの数がM個を下回った場合、プロジェクトの数がM個になるように与えられる
		//プロジェクトの更新はジャッジから知れる
		projects = gateway.SubscribeProject(m)
		//所持金も同様にジャッジから知れる
		money = gateway.SubscribeMoney()

		//K枚の方針カードとコストが提示される。1枚選んで手札に加える。
		dealedCard := gateway.DealCard(k)

		r := -1
		for i, dc := range dealedCard {
			if dc.t == int(Investment) && usedInvestmentCards < canUseInvestment && money >= dc.cost {
				r = i
			}
		}
		if r < 0 {
			r = 0
		}

		gateway.DeclareCard(r, dealedCard[r])

		node = append(node, r)
		ans = append(ans, node)
		//TODO:ここをtakeCardにした方が良いか考える
		hands.TakeCard(outputCard.c, PolicyCard{dealedCard[r].t, dealedCard[r].w, false})
		//テスト出力
		if debug {
			printSecondHalf(cnt, hands, projects, node, dealedCard, money)
		}
		//

	}
	return ans, money
}

// 条件を列挙して操作
func solve01(n, m, k, turn int, t, w, h, v []int) ([][]int, int) {
	const INF = 1 << 60

	hands := new(Hands)
	hands.Init(n, t, w)
	//cards := createCards(t, w)
	projects := createProjects(h, v)

	money := 0
	usedInvestmentCards := 0
	var ans [][]int

	for cnt := 0; cnt < turn; cnt++ {
		//手札から方針カードを選択して実行する
		var playCandidate []OutputCard
		if hands.hasRegularWork() {
			candidate := SolveRegularWork(hands, projects)
			candidate.turn = cnt
			playCandidate = append(playCandidate, candidate)
		}
		//TODO:最適な選び方
		if hands.hasHardWork() {
			for i, v := range hands.cs {
				if v.t == int(HardWork) {
					playCandidate = append(playCandidate, OutputCard{i, 0, cnt, HardWork, v.w})
					break
				}
			}
		}
		if hands.hasCancel() {
			candidate := OutputCard{turn: cnt, t: Cancel}
			for i, v := range hands.cs {
				if v.t == int(Cancel) {
					candidate.c = i
					break
				}
			}
			//使うプロジェクトを選ぶ
			//candidate.m = 0
			max := projects[0].remain
			for j, v := range projects {
				if v.remain > max {
					max = v.remain
					candidate.m = j
				}
			}
			playCandidate = append(playCandidate, candidate)
		}
		//TODO:最適な選び方
		if hands.hasRestructuring() {
			for i, v := range hands.cs {
				if v.t == int(Restructuring) {
					playCandidate = append(playCandidate, OutputCard{i, 0, cnt, Restructuring, v.w})
					break
				}
			}
		}
		//TODO:最適な選び方
		if hands.hasInvestment() {
			for i, v := range hands.cs {
				if v.t == int(Investment) {
					playCandidate = append(playCandidate, OutputCard{i, 0, cnt, Investment, v.w})
					break
				}
			}
		}
		idx := rand.Intn(len(playCandidate))
		outputCard := playCandidate[idx]
		if outputCard.t == Investment {
			usedInvestmentCards++
		}

		//gateway.PlayCard(ci, mi, c)
		hands.PlayCard(playCandidate[idx].c)
		gateway.PlayCard(outputCard)
		node := []int{playCandidate[idx].c, playCandidate[idx].m}
		//ans = append(ans, []int{playCandidate[idx].c, playCandidate[idx].m})

		//turn.0のところでテスト出力
		//fmt.Println("# ", "turn:", cnt, ".0")
		//fmt.Println("# ", *hands)
		//fmt.Println("# ", projects)
		//fmt.Println("# ", node, outputCard)
		//fmt.Println("# money = ", money)

		//存在するプロジェクトの数がM個を下回った場合、プロジェクトの数がM個になるように与えられる
		//プロジェクトの更新はジャッジから知れる
		projects = gateway.SubscribeProject(m)
		//所持金も同様にジャッジから知れる
		money = gateway.SubscribeMoney()

		//K枚の方針カードとコストが提示される。1枚選んで手札に加える。
		dealedCard := gateway.DealCard(k)

		r := 0
		//r := rand.Intn(k)
		//for money < dealedCard[r].cost || dealedCard[r].t == 3 || (usedInvestmentCards >= 20 && dealedCard[r].t == 4) {
		//	r = rand.Intn(k)
		//}

		gateway.DeclareCard(r, dealedCard[r])

		node = append(node, r)
		ans = append(ans, node)
		//TODO:ここをtakeCardにした方が良いか考える
		hands.TakeCard(outputCard.c, PolicyCard{dealedCard[r].t, dealedCard[r].w, false})
		//テスト出力
		//fmt.Println("# turn:", cnt, ".5")
		//fmt.Println("# ", *hands)
		//fmt.Println("# ", projects)
		//fmt.Println("# ", dealedCard)
		//fmt.Println("# ", node)
		//fmt.Println("# money = ", money)
		//

	}
	return ans, money
}

// solve00, すべてランダムに操作する

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
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

func DividePolicyCard(A []PolicyCard, K int) ([]PolicyCard, []PolicyCard, error) {
	if K < 0 {
		return nil, nil, errors.New("Impossible")
	}
	if len(A) <= K {
		return nil, nil, errors.New("Impossible")
	}
	return A[:K], A[K+1:], nil
}

func DivideProject(s []Project, k int) ([]Project, []Project, error) {
	if k < 0 {
		return nil, nil, errors.New("Impossible")
	}
	if len(s) <= k {
		return nil, nil, errors.New("Impossible")
	}
	return s[:k], s[k+1:], nil

}
