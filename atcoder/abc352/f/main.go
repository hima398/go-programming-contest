package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
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

	//計算量の見積もり
	//全探索しようとすると最大N! = 20,922,789,888,000 で2sでは間に合わない

	n, m := nextInt(), nextInt()
	var a, b, c []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		c = append(c, nextInt())
	}
	ans := solveHonestly(n, m, a, b, c)
	//ans := solve(n, m, a, b, c)
	PrintHorizonaly(ans)
}

func computeHash(s []int) [32]byte {
	var text []byte
	for _, v := range s {
		text = append(text, byte(v))
	}
	return sha256.Sum256(text)
}

func solve(n, m int, a, b, c []int) []int {
	//a, bを0-indexedで扱えるようにする
	/*
		for i := 0; i < m; i++ {
			a[i]--
			b[i]--
		}
	*/

	template := make([]int, n)
	for i := 0; i < n; i++ {
		template[i] = -1
	}

	type candidate struct {
		used  int   //順位が決まっている人を管理するフラグ
		ranks []int //順位
	}
	var candidates []candidate
	//candidates := make(map[[32]byte]candidate)
	var dfs func(d int, cs []candidate)
	dfs = func(d int, cs []candidate) {
		//テスト出力
		//fmt.Println(d, len(cs))
		//fmt.Println(d)
		//for _, v := range cs {
		//	fmt.Printf("%b, %v\n", v.used, v.ranks)
		//}

		if d == m {
			for _, v := range cs {
				if bits.OnesCount(uint(v.used)) == n-1 {
					idx := -1
					for i := 0; i < n; i++ {
						if (v.used>>i)&1 == 0 {
							idx = i
							break
						}
					}
					for i := range v.ranks {
						if v.ranks[i] == -1 {
							v.ranks[i] = idx
							v.used |= 1 << idx
						}
					}
				}
				//hash := computeHash(v.ranks)
				//candidates[hash] = v

				candidates = append(candidates, v)
			}
			return
		}

		var updated bool
		var nextCandidates []candidate
		for _, v := range cs {
			if (v.used>>b[d])&1 == 0 && (v.used>>a[d])&1 == 0 { //候補にbiもaiも一度も使用されていない
				for j := 0; j < n-c[d]; j++ {
					if v.ranks[j] >= 0 || v.ranks[j+c[d]] >= 0 {
						continue
					}
					nextRanks := make([]int, n)
					copy(nextRanks, v.ranks)
					nextRanks[j], nextRanks[j+c[d]] = b[d], a[d]
					/*
						if _, found := candidates[computeHash(nextRanks)]; found {
							continue
						}
					*/

					nextUsed := v.used | (1 << b[d]) | (1 << a[d])
					nextCandidates = append(nextCandidates, candidate{nextUsed, nextRanks})
					updated = true
				}
			} else if (v.used>>b[d])&1 > 0 { //候補の中で人b[d]の順位が決まっている場合
				for j := 0; j < n-c[d]; j++ {
					if v.ranks[j] != b[d] {
						continue
					}
					if v.ranks[j+c[d]] >= 0 {
						continue
					}
					nextRanks := make([]int, n)
					copy(nextRanks, v.ranks)

					nextRanks[j+c[d]] = a[d]
					/*
						if _, found := candidates[computeHash(nextRanks)]; found {
							continue
						}
					*/

					nextUsed := v.used | (1 << a[d])
					nextCandidates = append(nextCandidates, candidate{nextUsed, nextRanks})
					updated = true

				}
			} else if (v.used>>a[d])&1 > 0 { //候補の中で人a[d]の順位が決まっている場合
				for j := c[d]; j < n; j++ {
					if v.ranks[j] != a[d] {
						continue
					}
					if v.ranks[j-c[d]] >= 0 {
						continue
					}
					nextRanks := make([]int, n)
					copy(nextRanks, v.ranks)
					nextRanks[j-c[d]] = b[d]
					/*
						if _, found := candidates[computeHash(nextRanks)]; found {
							continue
						}
					*/

					nextUsed := v.used | (1 << b[d])
					nextCandidates = append(nextCandidates, candidate{nextUsed, nextRanks})
					updated = true
				}
			}
		}
		//Ai, Biの両方とも利用中の場合、順位の情報はすでに処理済みのため
		//else の処理は省略しておく
		if updated {
			dfs(d+1, nextCandidates)
		} else {
			dfs(d+1, cs)
		}
	}

	dfs(0, []candidate{{0, template}})

	//デバッグ用出力
	//for _, v := range candidates {
	//	fmt.Printf("%b, %v\n", v.used, v.ranks)
	//}

	idxes := make([]map[int]struct{}, n)
	for i := range idxes {
		idxes[i] = make(map[int]struct{})
	}
	for _, cs := range candidates {
		if cs.used != (1<<n)-1 {
			continue
		}
		for i, v := range cs.ranks {
			idxes[v][i] = struct{}{}
		}
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	for i := 0; i < n; i++ {
		if len(idxes[i]) == 1 {
			for k := range idxes[i] {
				ans[i] = k + 1
			}
		}
	}
	return ans
}

// 愚直解O(N**2 * N!)
func solveHonestly(n, m int, a, b, c []int) []int {
	//a, bを0-indexedで扱えるようにする
	/*
		for i := 0; i < m; i++ {
			a[i]--
			b[i]--
		}
	*/

	var p []int
	for i := 0; i < n; i++ {
		p = append(p, i)
	}
	var candidates [][]int
	for {
		ok := true
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if p[j] != b[i] {
					continue
				}
				if j >= n-c[i] {
					ok = false
					continue
				}
				//p[j]==b[i]
				ok = ok && p[j+c[i]] == a[i]
			}
		}
		//fmt.Println(p, ok)
		if ok {
			v := make([]int, n)
			copy(v, p)
			candidates = append(candidates, v)
		}
		if !NextPermutation(sort.IntSlice(p)) {
			break
		}
	}

	//デバッグ出力
	//fmt.Println(len(candidates), candidates)

	idxes := make([]map[int]struct{}, n)
	for i := range idxes {
		idxes[i] = make(map[int]struct{})
	}
	for _, cs := range candidates {
		for i, v := range cs {
			idxes[v][i] = struct{}{}
		}
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	for i := 0; i < n; i++ {
		if len(idxes[i]) == 1 {
			for k := range idxes[i] {
				ans[i] = k + 1
			}
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
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

// 愚直解用関数
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
