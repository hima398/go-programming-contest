package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

var h judge

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, w, k, s := nextInt(), nextInt(), nextInt(), nextInt()
	var a, b []int
	for i := 0; i < w; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	var c, d []int
	for i := 0; i < k; i++ {
		c = append(c, nextInt())
		d = append(d, nextInt())
	}
	var h2 hoge
	h = h2
	//solveHonestly(n, w, k, s, a, b, c, d)
	//solveGreedy(n, w, k, s, a, b, c, d)
	//solve01(n, w, k, s, a, b, c, d)
	solve02(n, w, k, s, a, b, c, d)
}

type judge interface {
	output(y, x, p int)
	input() int
}

type hoge struct {
}

func (h hoge) output(y, x, p int) {
	fmt.Println(y, x, p)
}

func (h hoge) input() int {
	return nextInt()
}

type excavation struct {
	y, x, p int
}

type output struct {
	y, x, p int
}
type cell struct {
	i, j int
}

type house struct {
	id   int
	i, j int
}

type waterSource struct {
	id   int
	i, j int
}

//柔らかい岩盤優先で採掘する2
func solve02(n, w, k, s int, a, b, c, d []int) ([]excavation, error) {
	const INF = 1 << 60
	const p = 100

	var wss []waterSource
	for i := range a {
		wss = append(wss, waterSource{i, a[i], b[i]})
	}
	var hs []house
	for i := range c {
		hs = append(hs, house{i, c[i], d[i]})
	}

	//各家に水が通ったかの管理
	completed := make([]bool, k)

	tq := make([]*TimesPriorityQueueV2, k)
	dq := make([]*DistPriorityQueueV2, k)
	for i := 0; i < k; i++ {
		tq[i] = &TimesPriorityQueueV2{}
		heap.Init(tq[i])
		dq[i] = &DistPriorityQueueV2{}
		heap.Init(dq[i])
	}

	uf := NewUnionFind(n * n)

	//プランニング
	plan := func() {
		fmt.Println("call plan()")
		nextTq := make([]*TimesPriorityQueueV2, k)
		nextDq := make([]*DistPriorityQueueV2, k)
		for i := 0; i < k; i++ {
			nextTq[i] = &TimesPriorityQueueV2{}
			heap.Init(nextTq[i])
			nextDq[i] = &DistPriorityQueueV2{}
			heap.Init(nextDq[i])
		}
		for _, h := range hs {
			if completed[h.id] {
				continue
			}
			for tq[h.id].Len() > 0 {
				cur := heap.Pop(tq[h.id]).(nodeV2)
				//水源iから水が通っている場所
				targets := make([][2]int, w)
				for i := range targets {
					targets[i] = [2]int{wss[i].i, wss[i].j}
				}
				dists := make([]int, w)
				for i := range dists {
					dists[i] = computeManhattanDist(cur.i, cur.j, wss[i].i, wss[i].j)
				}
				for i := 0; i < n; i++ {
					for j := 0; j < n; j++ {
						for _, ws := range wss {
							//水が通っていないマスは除く
							if !uf.ExistSameUnion(i*n+j, ws.i*n+j) {
								continue
							}
							dist := computeManhattanDist(i, j, cur.i, cur.j)
							if dists[ws.id] > dist {
								dists[ws.id] = dist
								targets[ws.id] = [2]int{i, j}
							}
						}
					}
				}
				for i := range targets {
					heap.Push(nextTq[h.id], nodeV2{h.id, i, cur.i, cur.j, targets[i][0], targets[i][1], 0, dists[i]})
				}
			}
			tq[h.id] = nextTq[h.id]

			for dq[h.id].Len() > 0 {
				cur := heap.Pop(dq[h.id]).(nodeV2)
				//水源iから水が通っている場所
				targets := make([][2]int, w)
				for i := range targets {
					targets[i] = [2]int{wss[i].i, wss[i].j}
				}
				dists := make([]int, w)
				for i := range dists {
					dists[i] = computeManhattanDist(cur.i, cur.j, wss[i].i, wss[i].j)
				}
				for i := 0; i < n; i++ {
					for j := 0; j < n; j++ {
						for _, ws := range wss {
							//水が通っていないマスは除く
							if !uf.ExistSameUnion(i*n+j, ws.i*n+j) {
								continue
							}
							dist := computeManhattanDist(i, j, cur.i, cur.j)
							if dists[ws.id] > dist {
								dists[ws.id] = dist
								targets[ws.id] = [2]int{i, j}
							}
						}
					}
				}
				for i := range targets {
					heap.Push(nextDq[h.id], nodeV2{h.id, i, cur.i, cur.j, targets[i][0], targets[i][1], 0, dists[i]})
				}
			}
			dq[h.id] = nextDq[h.id]
		}
	}

	visited := make([][][]bool, k)
	for kk := 0; kk < k; kk++ {
		visited[kk] = make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[kk][i] = make([]bool, n)
		}
	}
	//キューの準備
	//plan()
	for _, h := range hs {
		for _, ws := range wss {
			dist := computeManhattanDist(h.i, h.j, ws.i, ws.j)
			heap.Push(tq[h.id], nodeV2{h.id, ws.id, h.i, h.j, ws.i, ws.j, 0, dist})
			heap.Push(dq[h.id], nodeV2{h.id, ws.id, h.i, h.j, ws.i, ws.j, 0, dist})
			visited[h.id][h.i][h.j] = true
		}
	}
	var tlens, dlens []int
	for idx := range hs {
		tlens = append(tlens, tq[idx].Len())
		dlens = append(dlens, dq[idx].Len())
	}
	fmt.Println("h = ", len(hs))
	PrintHorizonaly(tlens)
	PrintHorizonaly(dlens)

	//fmt.Println(q)
	gdi := []int{-1, 0, 1, 0}
	gdj := []int{0, -1, 0, 1}

	crushed := make([][]bool, n)
	for i := 0; i < n; i++ {
		crushed[i] = make([]bool, n)
	}
	var ans []excavation
	var requiredPlan bool
	var totalLen int
	for i := 0; i < k; i++ {
		totalLen += tq[i].Len()
		totalLen += dq[i].Len()
	}
	//探索の処理
	for totalLen > 0 {
		//fmt.Println("totalLen = ", totalLen)
		for idx := range hs {
			fmt.Println(idx, tq[idx].Len(), dq[idx].Len())
			//すでに水が通っている家に関する探索であれば探索をしない
			if completed[idx] {
				continue
			}
			var cur nodeV2
			v := rand.Intn(2)
			if tq[idx].Len() > 0 && dq[idx].Len() > 0 {
				if v == 0 || dq[idx].Len() == 0 {
					cur = heap.Pop(tq[idx]).(nodeV2)
				} else if v == 1 || tq[idx].Len() == 0 {
					cur = heap.Pop(dq[idx]).(nodeV2)
				}
			} else if tq[idx].Len() > 0 {
				cur = heap.Pop(tq[idx]).(nodeV2)
				v = 0
			} else if dq[idx].Len() > 0 {
				cur = heap.Pop(dq[idx]).(nodeV2)
				v = 1
			} else {
				continue
			}

			//ここで採掘してみる
			if !crushed[cur.i][cur.j] {
				status := excavate(cur.i, cur.j, p)
				ans = append(ans, excavation{cur.i, cur.j, p})
				if status == -1 {
					return nil, errors.New("Already Crushed")
				} else if status == 0 {
					//取り出したキューに戻す
					if v == 0 {
						heap.Push(tq[idx], nodeV2{cur.houseId, cur.waterSourceId, cur.i, cur.j, cur.ti, cur.tj, cur.times + 1, cur.dist})
					} else if v == 1 {
						heap.Push(dq[idx], nodeV2{cur.houseId, cur.waterSourceId, cur.i, cur.j, cur.ti, cur.tj, cur.times + 1, cur.dist})
					}
					continue
				} else if status == 1 {
					crushed[cur.i][cur.j] = true
					for i := range gdi {
						pi, pj := cur.i+gdi[i], cur.j+gdj[i]
						if pi < 0 || pi >= n || pj < 0 || pj >= n {
							continue
						}
						if !crushed[pi][pj] {
							continue
						}
						if uf.ExistSameUnion(cur.i*n+cur.j, pi*n+pj) {
							continue
						}
						uf.Unite(cur.i*n+cur.j, pi*n+pj)
						//水源まで到達したかどうかをチェックする(要高速化)
						for _, ws := range wss {
							if uf.ExistSameUnion(cur.i*n+cur.j, ws.i*n+ws.j) {
								completed[cur.houseId] = true
								//目的地の再計画
								//plan()
								requiredPlan = true
							}
						}
					}
				} else if status == 2 {
					return ans, nil
				}
			}
			if requiredPlan {
				plan()
				var tlens, dlens []int
				for idx := range hs {
					tlens = append(tlens, tq[idx].Len())
					dlens = append(dlens, dq[idx].Len())
				}
				fmt.Println("h = ", len(hs))
				PrintHorizonaly(tlens)
				PrintHorizonaly(dlens)
				requiredPlan = false
				totalLen = 0
				for i := 0; i < k; i++ {
					totalLen += tq[i].Len()
					totalLen += dq[i].Len()
				}

				break
			}

			//これ以降は注目しているセルが採掘完了した時、次の目的地候補をキューに詰む
			var di []int
			var dj []int
			//次の経路の候補を2点ピックアップ
			dirI := cur.ti - cur.i
			dirJ := cur.tj - cur.j
			if dirI > 0 {
				di = append(di, 1)
				dj = append(dj, 0)
			} else if dirI < 0 {
				di = append(di, -1)
				dj = append(dj, 0)
			}
			if dirJ > 0 {
				di = append(di, 0)
				dj = append(dj, 1)
			} else if dirJ < 0 {
				di = append(di, 0)
				dj = append(dj, -1)
			}

			//キューに詰む
			for idx := range di {
				ni, nj := cur.i+di[idx], cur.j+dj[idx]
				if ni < 0 || ni >= n || nj < 0 || nj >= n {
					continue
				}
				if visited[cur.houseId][ni][nj] {
					continue
				}

				parents := make(map[int]int)
				for idx2 := range gdi {
					pi, pj := ni+gdi[idx2], nj+gdj[idx2]
					if pi < 0 || pi >= n || pj < 0 || pj >= n {
						continue
					}
					//TODO 角の処理
					parents[uf.Find(pi*n+pj)]++
				}
				dist := computeManhattanDist(ni, nj, wss[cur.waterSourceId].i, wss[cur.waterSourceId].j)
				heap.Push(tq[cur.houseId], nodeV2{cur.houseId, cur.waterSourceId, ni, nj, cur.ti, cur.tj, 0, dist})
				heap.Push(dq[cur.houseId], nodeV2{cur.houseId, cur.waterSourceId, ni, nj, cur.ti, cur.tj, 0, dist})
				visited[cur.houseId][ni][nj] = true
				//uf.Unite(cur.i*n+cur.j, ni*n+nj)
			}
		}
		totalLen = 0
		for i := 0; i < k; i++ {
			totalLen += tq[i].Len()
			totalLen += dq[i].Len()
		}
	}
	//ここまでくるとすべての家に水が通せなかった
	return ans, errors.New("Can't Complete.")
}

type nodeV2 struct {
	houseId       int
	waterSourceId int
	i, j          int //現在地
	ti, tj        int //目的地
	times         int
	dist          int
}

type PriorityQueueV2 []nodeV2

func (pq PriorityQueueV2) Len() int {
	return len(pq)
}
func (pq PriorityQueueV2) Less(i, j int) bool {
	//採掘回数、距離
	if pq[i].times == pq[j].times {
		return pq[i].dist < pq[j].dist
	}
	return pq[i].times < pq[j].times
}
func (pq PriorityQueueV2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

type TimesPriorityQueueV2 []nodeV2

func (pq TimesPriorityQueueV2) Len() int {
	return len(pq)
}
func (pq TimesPriorityQueueV2) Less(i, j int) bool {
	//採掘回数、距離
	if pq[i].times == pq[j].times {
		return pq[i].dist < pq[j].dist
	}
	return pq[i].times < pq[j].times
}
func (pq TimesPriorityQueueV2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *TimesPriorityQueueV2) Push(item interface{}) {
	*pq = append(*pq, item.(nodeV2))
}

func (pq *TimesPriorityQueueV2) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

type DistPriorityQueueV2 []nodeV2

func (pq DistPriorityQueueV2) Len() int {
	return len(pq)
}
func (pq DistPriorityQueueV2) Less(i, j int) bool {
	//採掘回数、距離
	if pq[i].dist == pq[j].dist {
		return pq[i].times < pq[j].times
	}
	return pq[i].dist < pq[j].dist
}
func (pq DistPriorityQueueV2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *DistPriorityQueueV2) Push(item interface{}) {
	*pq = append(*pq, item.(nodeV2))
}

func (pq *DistPriorityQueueV2) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

//柔らかい岩盤優先で採掘する
func solve01(n, w, k, s int, a, b, c, d []int) ([]excavation, error) {
	const INF = 1 << 60
	const p = 100
	var wss []waterSource
	for i := range a {
		wss = append(wss, waterSource{i, a[i], b[i]})
	}
	var hs []house
	for i := range c {
		hs = append(hs, house{i, c[i], d[i]})
	}

	//各家に水が通ったかの管理
	completed := make([]bool, k)

	tq := &TimesPriorityQueue{}
	dq := &DistPriorityQueue{}
	heap.Init(tq)
	heap.Init(dq)

	uf := NewUnionFind(n * n)

	//プランニング
	/*
		plan := func() {
			tq = &TimesPriorityQueue{}
			dq = &DistPriorityQueue{}
			heap.Init(tq)
			heap.Init(dq)
			//キューの準備
			for _, h := range hs {
				//水が通った家はスキップする
				if completed[h.i] {
					continue
				}
				//水源iから水が通っている場所
				targets := make([][2]int, w)
				for i := range targets {
					targets[i] = [2]int{wss[i].i, wss[i].j}
				}
				dists := make([]int, w)
				for i := range dists {
					dists[i] = computeManhattanDist(wss[i].i, wss[i].j, h.i, h.j)
				}
				for i := 0; i < n; i++ {
					for j := 0; j < n; j++ {
						for _, ws := range wss {
							//水が通っていないマスは除く
							if !uf.ExistSameUnion(i*n+j, ws.i*n+j) {
								continue
							}
							dist := computeManhattanDist(i, j, h.i, h.j)
							if dists[ws.i] > dist {
								dists[ws.i] = dist
								targets[ws.i] = [2]int{i, j}
							}
						}
					}
				}
				for i := range targets {
					heap.Push(tq, node{h.id, i, h.i, h.j, 0, computeManhattanDist(h.i, h.j, targets[i][0], targets[i][1])})
					heap.Push(dq, node{h.id, i, h.i, h.j, 0, computeManhattanDist(h.i, h.j, targets[i][0], targets[i][1])})
				}
			}
		}
	*/

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}
	//キューの準備
	for _, h := range hs {
		for _, ws := range wss {
			heap.Push(tq, node{h.id, ws.id, h.i, h.j, 0, computeManhattanDist(h.i, h.j, ws.i, ws.j)})
			heap.Push(dq, node{h.id, ws.id, h.i, h.j, 0, computeManhattanDist(h.i, h.j, ws.i, ws.j)})
			visited[h.i][h.j] = true
		}
	}
	//fmt.Println(q)

	crushed := make([][]bool, n)
	for i := 0; i < n; i++ {
		crushed[i] = make([]bool, n)
	}
	var ans []excavation
	//探索の処理
	for tq.Len() > 0 || dq.Len() > 0 {
		var cur node
		v := rand.Intn(2)
		if v == 0 {
			cur = heap.Pop(tq).(node)
		} else if v == 1 {
			cur = heap.Pop(dq).(node)
		}
		//すでに水が通っている家に関する探索であれば探索をしない
		if completed[cur.houseId] {
			continue
		}

		//ここで採掘してみる
		if !crushed[cur.i][cur.j] {
			status := excavate(cur.i, cur.j, p)
			ans = append(ans, excavation{cur.i, cur.j, p})
			if status == -1 {
				return nil, errors.New("Already Crushed")
			} else if status == 0 {
				//取り出したキューに戻す
				if v == 0 {
					heap.Push(tq, node{cur.houseId, cur.waterSourceId, cur.i, cur.j, cur.times + 1, cur.dist})
				} else if v == 1 {
					heap.Push(dq, node{cur.houseId, cur.waterSourceId, cur.i, cur.j, cur.times + 1, cur.dist})
				}
				continue
			} else if status == 1 {
				crushed[cur.i][cur.j] = true
				di := []int{-1, 0, 1, 0}
				dj := []int{0, -1, 0, 1}
				for i := range di {
					pi, pj := cur.i+di[i], cur.j+dj[i]
					if pi < 0 || pi >= n || pj < 0 || pj >= n {
						continue
					}
					if !crushed[pi][pj] {
						continue
					}
					if uf.ExistSameUnion(cur.i*n+cur.j, pi*n+pj) {
						continue
					}
					uf.Unite(cur.i*n+cur.j, pi*n+pj)
					//水源まで到達したかどうかをチェックする(要高速化)
					for _, ws := range wss {
						//if cur.i == ws.i && cur.j == ws.j {
						//	completed[cur.houseId] = true
						//}
						if uf.ExistSameUnion(cur.i*n+cur.j, ws.i*n+ws.j) {
							completed[cur.houseId] = true
						}
					}
				}
			} else if status == 2 {
				return ans, nil
			}
		}

		//これ以降は注目しているセルが採掘完了した時、次の目的地候補をキューに詰む

		var di []int
		var dj []int
		//次の経路の候補を2点ピックアップ
		dirI := wss[cur.waterSourceId].i - cur.i
		dirJ := wss[cur.waterSourceId].j - cur.j
		if dirI > 0 {
			di = append(di, 1)
			dj = append(dj, 0)
		} else if dirI < 0 {
			di = append(di, -1)
			dj = append(dj, 0)
		}
		if dirJ > 0 {
			di = append(di, 0)
			dj = append(dj, 1)
		} else if dirJ < 0 {
			di = append(di, 0)
			dj = append(dj, -1)
		}

		//キューに詰む
		for idx := range di {
			ni, nj := cur.i+di[idx], cur.j+dj[idx]
			if ni < 0 || ni >= n || nj < 0 || nj >= n {
				continue
			}
			if visited[ni][nj] {
				continue
			}
			dist := computeManhattanDist(ni, nj, wss[cur.waterSourceId].i, wss[cur.waterSourceId].j)
			heap.Push(tq, node{cur.houseId, cur.waterSourceId, ni, nj, 0, dist})
			heap.Push(dq, node{cur.houseId, cur.waterSourceId, ni, nj, 0, dist})
			visited[ni][nj] = true
			//uf.Unite(cur.i*n+cur.j, ni*n+nj)
		}
	}
	//ここまでくるとすべての家に水が通せなかった
	return ans, errors.New("Can't Complete.")
}

func excavate(y, x, p int) int {
	h.output(y, x, p)
	return h.input()
}

type node struct {
	houseId       int
	waterSourceId int
	i, j          int
	times         int
	dist          int
}
type TimesPriorityQueue []node

func (pq TimesPriorityQueue) Len() int {
	return len(pq)
}
func (pq TimesPriorityQueue) Less(i, j int) bool {
	//採掘回数、距離
	if pq[i].times == pq[j].times {
		return pq[i].dist < pq[j].dist
	}
	return pq[i].times < pq[j].times
}
func (pq TimesPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *TimesPriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(node))
}

func (pq *TimesPriorityQueue) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

type DistPriorityQueue []node

func (pq DistPriorityQueue) Len() int {
	return len(pq)
}
func (pq DistPriorityQueue) Less(i, j int) bool {
	//採掘回数、距離
	if pq[i].dist == pq[j].dist {
		return pq[i].times < pq[j].times
	}
	return pq[i].dist < pq[j].dist
}
func (pq DistPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *DistPriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(node))
}

func (pq *DistPriorityQueue) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
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

//各家から一番近い水源へ、L字のルートでcずつパワーを上げて掘削する
func solveGreedy(n, w, k, s int, a, b, c, d []int) ([]excavation, error) {
	const INF = 1 << 60
	crushed := make([][]bool, n)
	for i := range crushed {
		crushed[i] = make([]bool, n)
	}

	var res []excavation
	for i := 0; i < k; i++ {
		minDist := INF
		target := -1
		for j := 0; j < w; j++ {
			dist := computeManhattanDist(a[j], b[j], c[i], d[i])
			if minDist > dist {
				target = j
				minDist = dist
			}
		}
		//fmt.Println("i, target = ", i, target)
		y1, y2 := Min(a[target], c[i]), Max(a[target], c[i])
		x1, x2 := Min(b[target], d[i]), Max(b[target], d[i])
		var y int
		if (y1 == a[target] && x1 == b[target]) || (y1 == c[i] && x1 == d[i]) {
			//L字に掘削する
			y = y2
		} else {
			//「字に掘削する
			y = y1
		}
		for ii := y1; ii <= y2; ii++ {
			if crushed[ii][x1] {
				continue
			}
			var status int
			power := 100
			for status != 1 {
				e := excavation{ii, x1, power}
				power = Min(power+s, 5000)
				//fmt.Println(o.y, o.x, o.p)
				write(h, e.y, e.x, e.p)
				res = append(res, e)
				//status = nextInt()
				status = read(h)
				if status == -1 {
					return nil, errors.New("Already Crushed")
				}
				if status == 2 {
					return res, nil
				}
			}
			crushed[ii][x1] = true
		}
		for jj := x1; jj <= x2; jj++ {
			if crushed[y][jj] {
				continue
			}
			var status int
			power := 100
			for status != 1 {
				o := excavation{y, jj, power}
				power = Min(power+s, 5000)
				//fmt.Println(o.y, o.x, o.p)
				write(h, o.y, o.x, o.p)
				res = append(res, o)
				//status = nextInt()
				status = read(h)
				if status == -1 {
					return nil, errors.New("Already Crushed")
				}
				if status == 2 {
					return res, nil
				}
			}
			crushed[y][jj] = true
		}
	}
	return res, errors.New("Can't Completed")
}

//全部掘削する(TLE)
func solveHonestly(n, w, k, s int, a, b, c, d []int) {
	//defer out.Flush()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var res int
			for res != 1 {
				//fmt.Fprintln(out, i, j, c)
				//fmt.Println(i, j, s)
				res = nextInt()
				if res == -1 || res == 2 {
					return
				}
			}
		}
	}
}

func write(j judge, y, x, p int) {
	j.output(y, x, p)
}

func read(j judge) int {
	return j.input()
}

func computeManhattanDist(i1, j1, i2, j2 int) int {
	return Abs(i2-i1) + Abs(j2-j1)
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
