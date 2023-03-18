package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

//インタラクティブに関する入出力
type Judge interface {
	send(y, x, p int)
	receive() int
}

var stdio Judge

func Excavate(y, x, p int) int {
	stdio.send(y, x, p)
	return stdio.receive()
}

//標準入出力への処理
type StandardIO struct {
}

func (judge StandardIO) send(y, x, p int) {
	fmt.Println(y, x, p)
}

func (judge StandardIO) receive() int {
	return nextInt()
}

//問題独自ドメインの定義
type Excavation struct {
	y, x, p int
}

type House struct {
	id   int
	i, j int
}

type WaterSource struct {
	id   int
	i, j int
}

type Cell struct {
	i, j int
}

var di = []int{-1, 0, 1, 0}
var dj = []int{0, -1, 0, 1}

func ExportMainImage(dir, name string, n int, field [][]int) {
	dest := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{n, n}})
	mx := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mx = Max(mx, field[i][j])
		}
	}
	//fmt.Println("mx = ", mx)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			v := field[y][x] * 255 / mx
			v = Abs(v - 255)
			//fmt.Println(y, x, field[y][x], mx)
			dest.Set(x, y, color.Gray{uint8(v)})
		}
	}
	path := filepath.Join(dir, name)
	file, err := os.Create(path)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	png.Encode(file, dest)
}

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
	var st StandardIO
	stdio = st

	solve07(n, w, k, s, 100, 2, 40, a, b, c, d)
	//solve06(n, w, k, s, 100, 2, 40, a, b, c, d)
	//solve05(n, w, k, s, 100, 2, 40, a, b, c, d)
	//solve04(n, w, k, s, 100, 200, 40, a, b, c, d)
	//solve035(n, w, k, s, 40, a, b, c, d)
	//solve03(n, w, k, s, a, b, c, d)
	//solve02(n, w, k, s, a, b, c, d)
}

func InField(n, i, j int) bool {
	return 0 <= i && i < n && 0 <= j && j < n
}

type Tree struct {
	ns []Cell
	st [][]bool
}

//岩盤調査を同じマンハッタン距離の領域で行う
func solve07(n, w, k, s, investPower, investTimes, defaultPower int, a, b, c, d []int) ([]Excavation, [][]int, int, int, error) {
	const INF = 1 << 60
	const investWidth = 20
	const weightWidth = investWidth / 2

	template := make([][]int, investWidth)
	for i := range template {
		template[i] = make([]int, investWidth)
		for j := range template[i] {
			template[i][j] = 1
		}
	}
	base := [][]int{{0, 0}, {0, investWidth - 1}, {investWidth - 1, 0}, {investWidth - 1, investWidth - 1}}
	for _, b := range base {
		for i := 0; i < len(template); i++ {
			for j := 0; j < len(template[0]); j++ {
				d := ComputeManhattanDist(i, j, b[0], b[1])
				if d < investWidth/2-1 {
					template[i][j] = 0
				}
			}
		}
	}

	var templateArea int
	for i := range template {
		for j := range template[i] {
			if template[i][j] > 0 {
				templateArea++
			}
		}
	}
	//ExportMainImage("./", "hillTemplate.png", len(hillTemplate), hillTemplate)
	//fmt.Println("templateArea = ", templateArea)
	//ExportMainImage("./", "template.png", len(template), template)
	uf := NewUnionFind(n * n)

	//水源、家を独自データ構造にしておく
	var wss []WaterSource
	for i := range a {
		wss = append(wss, WaterSource{i, a[i], b[i]})
	}
	var hs []House
	for i := range c {
		hs = append(hs, House{i, c[i], d[i]})
	}

	field := make([][]int, n)
	crushed := make([][]bool, n)
	openedWater := make([][]bool, n)
	for i := 0; i < n; i++ {
		field[i] = make([]int, n)
		crushed[i] = make([]bool, n)
		openedWater[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			field[i][j] = 5000*investWidth*investWidth + 1
		}
	}
	var ans []Excavation

	//白が岩盤調査で掘れたエリア、黒が掘れなかったエリア
	var whiteArea int
	blackArea := n * n
	fillField := func(i, j, status, w int) {
		bi, bj := len(template)/2, len(template[0])/2
		hw := investWidth / 2
		for ii := i - hw; ii < i+hw; ii++ {
			for jj := j - hw; jj < j+hw; jj++ {
				if !InField(n, ii, jj) {
					continue
				}
				//テンプレート上の位置
				ti, tj := bi+(ii-i), bj+(jj-j)
				//fmt.Println("(i, j) = ", i, j, "(bi, bj) = ", bi, bj, "(ti, tj) = ", ti, tj)
				if status == 1 && template[ti][tj] == 1 {
					field[ii][jj] = (w+1)*investPower*templateArea + 1
				} else if status == 0 && template[ti][tj] == 1 {
					field[ii][jj] = 5000*templateArea + 1
				}
			}
		}
		if status == 1 {
			whiteArea += templateArea
			blackArea -= templateArea
		}
	}
	//岩盤調査
	//fmt.Println("岩盤調査")
	for k := 0; k < investTimes; k++ {
		cnt := 0
		for i := 0; i < n; i += investWidth/2 + 1 {
			var w int
			if cnt%2 == 1 {
				w = investWidth / 2
			}
			for j := 0; j < n+investWidth; j += investWidth {
				oj := j
				if j+w >= n {
					j--
				}
				if !InField(n, i, j+w) {
					continue
				}

				if crushed[i][j+w] {
					continue
				}
				power := investPower
				status := Excavate(i, j+w, power)
				ans = append(ans, Excavation{i, j + w, power})

				if status == 1 {
					crushed[i][j+w] = true
				}
				fillField(i, oj+w, status, k)
			}
			cnt++
		}
	}
	//ExportMainImage("./", "0000.png", n, field)

	//平滑化
	for l := 0; l < investWidth; l++ {
		di := []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}
		dj := []int{-1, 0, 1, -1, 0, 1, -1, 0, 1}
		next := make([][]int, n)
		for i := 0; i < n; i++ {
			next[i] = make([]int, n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				num := 1
				sum := field[i][j]
				for k := 0; k < 8; k++ {
					ni, nj := i+di[k], j+dj[k]
					if !InField(n, ni, nj) {
						continue
					}
					sum += field[ni][nj]
					num++
				}
				next[i][j] = sum / num
				//if next[i][j] < 0 {
				//	fmt.Println(l, i, j, sum, num, next[i][j])
				//	panic("field has negative value..")
				//}
			}
		}
		field = next
	}

	//水源と家の岩盤をあらかじめ砕いておく
	//fmt.Println("水源と家の岩盤採掘")
	for _, ws := range wss {
		var cnt int
		var status int
		if crushed[ws.i][ws.j] {
			status = 1
		}
		for status != 1 {
			power := defaultPower + cnt*s
			status = Excavate(ws.i, ws.j, power)
			ans = append(ans, Excavation{ws.i, ws.j, power})
			if status == -1 {
				return nil, nil, whiteArea, blackArea, errors.New("Already Crushed.")
			}
			if status == 2 {
				return ans, field, whiteArea, blackArea, nil
			}
			cnt++
		}
		crushed[ws.i][ws.j] = true
		field[ws.i][ws.j] = 0
		openedWater[ws.i][ws.j] = true
	}
	for _, h := range hs {
		var cnt int
		var status int
		if crushed[h.i][h.j] {
			status = 1
		}
		for status != 1 {
			power := defaultPower + cnt*s
			status = Excavate(h.i, h.j, power)
			ans = append(ans, Excavation{h.i, h.j, power})
			if status == -1 {
				return nil, nil, whiteArea, blackArea, errors.New("Already Crushed.")
			}
			if status == 2 {
				return ans, field, whiteArea, blackArea, nil
			}
			cnt++
		}
		crushed[h.i][h.j] = true
		field[h.i][h.j] = 0
	}

	dijkstra := func(si, sj int) ([][]int, []int) {
		q := &PriorityQueue{}
		heap.Init(q)

		dist := make([][]int, n)
		for i := range dist {
			dist[i] = make([]int, n)
			for j := 0; j < n; j++ {
				dist[i][j] = INF
			}
		}
		path := make([]int, n*n)
		for i := 0; i < n*n; i++ {
			path[i] = i
		}
		push := func(ci, cj, ti, tj, cost int) {
			if dist[ti][tj] <= cost {
				return
			}
			dist[ti][tj] = cost
			heap.Push(q, NodeV3{ti, tj, cost})
			path[ti*n+tj] = ci*n + cj
		}
		push(si, sj, si, sj, 0)
		//di := []int{-1, 0, 1, 0}
		//dj := []int{0, -1, 0, 1}
		for q.Len() > 0 {
			cur := heap.Pop(q).(NodeV3)
			if dist[cur.i][cur.j] < cur.w {
				continue
			}
			for k := 0; k < 4; k++ {
				ni, nj := cur.i+di[k], cur.j+dj[k]
				if !InField(n, ni, nj) {
					continue
				}
				push(cur.i, cur.j, ni, nj, cur.w+field[ni][nj])
			}
		}
		return dist, path
	}

	//目指す水源の選択
	//水源が通った家の管理
	var areOpend int
	var path []int
	for areOpend < 1<<k-1 {
		//fmt.Printf("areOpened, mask = %b, %b\n", areOpend, 1<<k-1)
		houseId := -1
		dist := INF
		ti, tj := -1, -1
		//var dists [][]int
		//各家ごとの処理
		for _, h := range hs {
			//すでに水が通っている家は対象外とする
			if (areOpend>>h.id)&1 == 1 {
				continue
			}

			//一番近い水源を計算
			tdists, tpath := dijkstra(h.i, h.j)

			//家hから最も近い水源までの距離、座標
			d := INF
			tti, ttj := -1, -1
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if !openedWater[i][j] {
						continue
					}
					if d > tdists[i][j] {
						d = tdists[i][j]
						tti, ttj = i, j
					}
				}
			}
			if dist > d {
				houseId = h.id
				dist = d
				ti, tj = tti, ttj
				path = tpath
				//dists = tdists
			}
		}
		//fmt.Printf("id = %d, (ti, tj) = (%d, %d), (hi, hj) = (%d, %d)\n", houseId, ti, tj, hs[houseId].i, hs[houseId].j)

		restructRoute := func(ti, tj int) []Cell {
			res := []Cell{{ti, tj}}
			idx := ti*n + tj
			for path[idx] != idx {
				idx = path[idx]
				res = append(res, Cell{idx / n, idx % n})
			}
			return res
		}
		routes := restructRoute(ti, tj) //dfs(ti, tj, hs[houseId].i, hs[houseId].j)
		m := len(routes)
		//fmt.Println(routes[m-1].i, routes[m-1].j, routes[0].i, routes[0].j, len(routes))
		//fmt.Println(routes)
		pi, pj := routes[m-1].i, routes[m-1].j
		for i := m - 2; i >= 0; i-- {
			if !uf.ExistSameUnion(routes[i].i*n+routes[i].j, pi*n+pj) {
				uf.Unite(routes[i].i*n+routes[i].j, pi*n+pj)
			}
			var isOpenedWater bool
			for k := 0; k < 4; k++ {
				ni, nj := routes[i].i+di[k], routes[i].j+dj[k]
				if !InField(n, ni, nj) {
					continue
				}

				if openedWater[ni][nj] {
					uf.Unite(ni*n+nj, routes[i].i*n+routes[i].j)
					isOpenedWater = true
				}

			}
			if !uf.ExistSameUnion(ti*n+tj, routes[i].i*n+routes[i].j) {
				uf.Unite(ti*n+tj, routes[i].i*n+routes[i].j)
			}

			if !crushed[routes[i].i][routes[i].j] {
				//	continue
				//}
				var status int
				var cnt int
				for status != 1 {
					power := defaultPower + cnt*s
					status = Excavate(routes[i].i, routes[i].j, power)
					ans = append(ans, Excavation{routes[i].i, routes[i].j, power})
					if status == -1 {
						return nil, nil, whiteArea, blackArea, errors.New("Already Crushed.")
					}
					if status == 2 {
						return ans, field, whiteArea, blackArea, nil
					}
					cnt++
				}
				crushed[routes[i].i][routes[i].j] = true
				//openedWater[routes[i].i][routes[i].j] = true
				field[routes[i].i][routes[i].j] = 0
			}
			if isOpenedWater {
				for j := i; j >= 0; j-- {
					openedWater[routes[j].i][routes[j].j] = true
				}
				//fmt.Println(houseId, " route breaked.")
				break
			}
			pi, pj = routes[i].i, routes[i].j
		}
		openedWater[hs[houseId].i][hs[houseId].j] = true
		areOpend |= 1 << houseId
	}
	return ans, field, whiteArea, blackArea, nil
}

//岩盤調査後に重みをなだらかにする
func solve06(n, w, k, s, investPower, investTimes, defaultPower int, a, b, c, d []int) ([]Excavation, [][]int, error) {
	const INF = 1 << 60
	const investWidth = 20
	const weightWidth = investWidth / 2

	uf := NewUnionFind(n * n)

	//水源、家を独自データ構造にしておく
	var wss []WaterSource
	for i := range a {
		wss = append(wss, WaterSource{i, a[i], b[i]})
	}
	var hs []House
	for i := range c {
		hs = append(hs, House{i, c[i], d[i]})
	}

	field := make([][]int, n)
	crushed := make([][]bool, n)
	openedWater := make([][]bool, n)
	for i := 0; i < n; i++ {
		field[i] = make([]int, n)
		crushed[i] = make([]bool, n)
		openedWater[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			field[i][j] = 5000*investWidth*investWidth + 1
		}
	}
	var ans []Excavation
	//岩盤調査
	//fmt.Println("岩盤調査")
	for k := 0; k < investTimes; k++ {
		for i := weightWidth - 1; i < n; i += investWidth {
			for j := weightWidth - 1; j < n; j += investWidth {
				if crushed[i][j] {
					continue
				}
				status := Excavate(i, j, investPower)
				ans = append(ans, Excavation{i, j, investPower})

				if status == 1 {
					crushed[i][j] = true
					//field[i][j] = 0
					//割れた岩盤付近に重みをつける
					for ii := Max(i-weightWidth+1, 0); ii < Min(i+weightWidth, n); ii++ {
						for jj := Max(j-weightWidth+1, 0); jj < Min(j+weightWidth, n); jj++ {
							field[ii][jj] = k*investPower*investWidth*investWidth + 1
						}
					}
				}
			}
		}
	}
	//平滑化
	for l := 0; l < investWidth; l++ {
		next := make([][]int, n)
		for i := 0; i < n; i++ {
			next[i] = make([]int, n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				num := 1
				sum := field[i][j]
				for k := 0; k < 4; k++ {
					ni, nj := i+di[k], j+dj[k]
					if !InField(n, ni, nj) {
						continue
					}
					sum += field[ni][nj]
					num++
				}
				next[i][j] = sum / num
			}
		}
		field = next
	}

	//ExportImage(n, field)

	//水源と家の岩盤をあらかじめ砕いておく
	//fmt.Println("水源と家の岩盤採掘")
	for _, ws := range wss {
		var cnt int
		var status int
		if crushed[ws.i][ws.j] {
			status = 1
		}
		for status != 1 {
			power := defaultPower + cnt*s
			status = Excavate(ws.i, ws.j, power)
			ans = append(ans, Excavation{ws.i, ws.j, power})
			if status == -1 {
				return nil, nil, errors.New("Already Crushed.")
			}
			if status == 2 {
				return ans, field, nil
			}
			cnt++
		}
		crushed[ws.i][ws.j] = true
		field[ws.i][ws.j] = 0
		openedWater[ws.i][ws.j] = true
	}
	for _, h := range hs {
		var cnt int
		var status int
		if crushed[h.i][h.j] {
			status = 1
		}
		for status != 1 {
			power := defaultPower + cnt*s
			status = Excavate(h.i, h.j, power)
			ans = append(ans, Excavation{h.i, h.j, power})
			if status == -1 {
				return nil, nil, errors.New("Already Crushed.")
			}
			if status == 2 {
				return ans, field, nil
			}
			cnt++
		}
		crushed[h.i][h.j] = true
		field[h.i][h.j] = 0
	}

	dijkstra := func(si, sj int) ([][]int, []int) {
		q := &PriorityQueue{}
		heap.Init(q)

		dist := make([][]int, n)
		for i := range dist {
			dist[i] = make([]int, n)
			for j := 0; j < n; j++ {
				dist[i][j] = INF
			}
		}
		path := make([]int, n*n)
		for i := 0; i < n*n; i++ {
			path[i] = i
		}
		push := func(ci, cj, ti, tj, cost int) {
			if dist[ti][tj] <= cost {
				return
			}
			dist[ti][tj] = cost
			heap.Push(q, NodeV3{ti, tj, cost})
			path[ti*n+tj] = ci*n + cj
		}
		push(si, sj, si, sj, 0)
		//di := []int{-1, 0, 1, 0}
		//dj := []int{0, -1, 0, 1}
		for q.Len() > 0 {
			cur := heap.Pop(q).(NodeV3)
			if dist[cur.i][cur.j] < cur.w {
				continue
			}
			for k := 0; k < 4; k++ {
				ni, nj := cur.i+di[k], cur.j+dj[k]
				if !InField(n, ni, nj) {
					continue
				}
				push(cur.i, cur.j, ni, nj, cur.w+field[ni][nj])
			}
		}
		return dist, path
	}

	//const startPower = 50
	//目指す水源の選択
	//水源が通った家の管理
	var areOpend int
	var path []int
	for areOpend < 1<<k-1 {
		//fmt.Printf("areOpened, mask = %b, %b\n", areOpend, 1<<k-1)
		houseId := -1
		dist := INF
		ti, tj := -1, -1
		//var dists [][]int
		//各家ごとの処理
		for _, h := range hs {
			//すでに水が通っている家は対象外とする
			if (areOpend>>h.id)&1 == 1 {
				continue
			}

			//一番近い水源を計算
			tdists, tpath := dijkstra(h.i, h.j)

			//家hから最も近い水源までの距離、座標
			d := INF
			tti, ttj := -1, -1
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if !openedWater[i][j] {
						continue
					}
					if d > tdists[i][j] {
						d = tdists[i][j]
						tti, ttj = i, j
					}
				}
			}
			if dist > d {
				houseId = h.id
				dist = d
				ti, tj = tti, ttj
				path = tpath
				//dists = tdists
			}
		}
		//fmt.Printf("id = %d, (ti, tj) = (%d, %d), (hi, hj) = (%d, %d)\n", houseId, ti, tj, hs[houseId].i, hs[houseId].j)

		restructRoute := func(ti, tj int) []Cell {
			res := []Cell{{ti, tj}}
			idx := ti*n + tj
			for path[idx] != idx {
				idx = path[idx]
				res = append(res, Cell{idx / n, idx % n})
			}
			return res
		}
		routes := restructRoute(ti, tj) //dfs(ti, tj, hs[houseId].i, hs[houseId].j)
		m := len(routes)
		//fmt.Println(routes[m-1].i, routes[m-1].j, routes[0].i, routes[0].j, len(routes))
		//fmt.Println(routes)
		pi, pj := routes[m-1].i, routes[m-1].j
		for i := m - 2; i >= 0; i-- {
			if !uf.ExistSameUnion(routes[i].i*n+routes[i].j, pi*n+pj) {
				uf.Unite(routes[i].i*n+routes[i].j, pi*n+pj)
			}
			var isOpenedWater bool
			for k := 0; k < 4; k++ {
				ni, nj := routes[i].i+di[k], routes[i].j+dj[k]
				if !InField(n, ni, nj) {
					continue
				}

				if openedWater[ni][nj] {
					uf.Unite(ni*n+nj, routes[i].i*n+routes[i].j)
					isOpenedWater = true
				}

			}
			if !uf.ExistSameUnion(ti*n+tj, routes[i].i*n+routes[i].j) {
				uf.Unite(ti*n+tj, routes[i].i*n+routes[i].j)
			}

			if !crushed[routes[i].i][routes[i].j] {
				//	continue
				//}
				var status int
				var cnt int
				for status != 1 {
					power := defaultPower + cnt*s
					status = Excavate(routes[i].i, routes[i].j, power)
					ans = append(ans, Excavation{routes[i].i, routes[i].j, power})
					if status == -1 {
						return nil, nil, errors.New("Already Crushed.")
					}
					if status == 2 {
						return ans, field, nil
					}
					cnt++
				}
				crushed[routes[i].i][routes[i].j] = true
				//openedWater[routes[i].i][routes[i].j] = true
				field[routes[i].i][routes[i].j] = 0
			}
			if isOpenedWater {
				for j := i; j >= 0; j-- {
					openedWater[routes[j].i][routes[j].j] = true
				}
				//fmt.Println(houseId, " route breaked.")
				break
			}
			pi, pj = routes[i].i, routes[i].j
		}
		openedWater[hs[houseId].i][hs[houseId].j] = true
		areOpend |= 1 << houseId
	}
	return ans, field, nil
}

//岩盤調査を複数回行う
func solve05(n, w, k, s, investPower, investTimes, defaultPower int, a, b, c, d []int) ([]Excavation, [][]int, error) {
	const INF = 1 << 60
	const investWidth = 20
	const weightWidth = investWidth / 2

	uf := NewUnionFind(n * n)

	//水源、家を独自データ構造にしておく
	var wss []WaterSource
	for i := range a {
		wss = append(wss, WaterSource{i, a[i], b[i]})
	}
	var hs []House
	for i := range c {
		hs = append(hs, House{i, c[i], d[i]})
	}

	field := make([][]int, n)
	crushed := make([][]bool, n)
	openedWater := make([][]bool, n)
	for i := 0; i < n; i++ {
		field[i] = make([]int, n)
		crushed[i] = make([]bool, n)
		openedWater[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			field[i][j] = 5000*investWidth*investWidth + 1
		}
	}
	var ans []Excavation
	//岩盤調査
	//fmt.Println("岩盤調査")
	for k := 0; k < investTimes; k++ {
		for i := weightWidth - 1; i < n; i += investWidth {
			for j := weightWidth - 1; j < n; j += investWidth {
				if crushed[i][j] {
					continue
				}
				status := Excavate(i, j, investPower)
				ans = append(ans, Excavation{i, j, investPower})

				if status == 1 {
					crushed[i][j] = true
					//field[i][j] = 0
					//割れた岩盤付近に重みをつける
					for ii := Max(i-weightWidth+1, 0); ii < Min(i+weightWidth, n); ii++ {
						for jj := Max(j-weightWidth+1, 0); jj < Min(j+weightWidth, n); jj++ {
							field[ii][jj] = k*investPower*investWidth*investWidth + 1
						}
					}
				}
			}
		}
	}

	//ExportImage(n, field)

	//水源と家の岩盤をあらかじめ砕いておく
	//fmt.Println("水源と家の岩盤採掘")
	for _, ws := range wss {
		var cnt int
		var status int
		if crushed[ws.i][ws.j] {
			status = 1
		}
		for status != 1 {
			power := defaultPower + cnt*s
			status = Excavate(ws.i, ws.j, power)
			ans = append(ans, Excavation{ws.i, ws.j, power})
			if status == -1 {
				return nil, nil, errors.New("Already Crushed.")
			}
			if status == 2 {
				return ans, field, nil
			}
			cnt++
		}
		crushed[ws.i][ws.j] = true
		field[ws.i][ws.j] = 0
		openedWater[ws.i][ws.j] = true
	}
	for _, h := range hs {
		var cnt int
		var status int
		if crushed[h.i][h.j] {
			status = 1
		}
		for status != 1 {
			power := defaultPower + cnt*s
			status = Excavate(h.i, h.j, power)
			ans = append(ans, Excavation{h.i, h.j, power})
			if status == -1 {
				return nil, nil, errors.New("Already Crushed.")
			}
			if status == 2 {
				return ans, field, nil
			}
			cnt++
		}
		crushed[h.i][h.j] = true
		field[h.i][h.j] = 0
	}

	dijkstra := func(si, sj int) ([][]int, []int) {
		q := &PriorityQueue{}
		heap.Init(q)

		dist := make([][]int, n)
		for i := range dist {
			dist[i] = make([]int, n)
			for j := 0; j < n; j++ {
				dist[i][j] = INF
			}
		}
		path := make([]int, n*n)
		for i := 0; i < n*n; i++ {
			path[i] = i
		}
		push := func(ci, cj, ti, tj, cost int) {
			if dist[ti][tj] <= cost {
				return
			}
			dist[ti][tj] = cost
			heap.Push(q, NodeV3{ti, tj, cost})
			path[ti*n+tj] = ci*n + cj
		}
		push(si, sj, si, sj, 0)
		//di := []int{-1, 0, 1, 0}
		//dj := []int{0, -1, 0, 1}
		for q.Len() > 0 {
			cur := heap.Pop(q).(NodeV3)
			if dist[cur.i][cur.j] < cur.w {
				continue
			}
			for k := 0; k < 4; k++ {
				ni, nj := cur.i+di[k], cur.j+dj[k]
				if !InField(n, ni, nj) {
					continue
				}
				push(cur.i, cur.j, ni, nj, cur.w+field[ni][nj])
			}
		}
		return dist, path
	}

	//const startPower = 50
	//目指す水源の選択
	//水源が通った家の管理
	var areOpend int
	var path []int
	for areOpend < 1<<k-1 {
		//fmt.Printf("areOpened, mask = %b, %b\n", areOpend, 1<<k-1)
		houseId := -1
		dist := INF
		ti, tj := -1, -1
		//var dists [][]int
		//各家ごとの処理
		for _, h := range hs {
			//すでに水が通っている家は対象外とする
			if (areOpend>>h.id)&1 == 1 {
				continue
			}

			//一番近い水源を計算
			tdists, tpath := dijkstra(h.i, h.j)

			//家hから最も近い水源までの距離、座標
			d := INF
			tti, ttj := -1, -1
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if !openedWater[i][j] {
						continue
					}
					if d > tdists[i][j] {
						d = tdists[i][j]
						tti, ttj = i, j
					}
				}
			}
			if dist > d {
				houseId = h.id
				dist = d
				ti, tj = tti, ttj
				path = tpath
				//dists = tdists
			}
		}
		//fmt.Printf("id = %d, (ti, tj) = (%d, %d), (hi, hj) = (%d, %d)\n", houseId, ti, tj, hs[houseId].i, hs[houseId].j)

		restructRoute := func(ti, tj int) []Cell {
			res := []Cell{{ti, tj}}
			idx := ti*n + tj
			for path[idx] != idx {
				idx = path[idx]
				res = append(res, Cell{idx / n, idx % n})
			}
			return res
		}
		routes := restructRoute(ti, tj) //dfs(ti, tj, hs[houseId].i, hs[houseId].j)
		m := len(routes)
		//fmt.Println(routes[m-1].i, routes[m-1].j, routes[0].i, routes[0].j, len(routes))
		//fmt.Println(routes)
		pi, pj := routes[m-1].i, routes[m-1].j
		for i := m - 2; i >= 0; i-- {
			if !uf.ExistSameUnion(routes[i].i*n+routes[i].j, pi*n+pj) {
				uf.Unite(routes[i].i*n+routes[i].j, pi*n+pj)
			}
			var isOpenedWater bool
			for k := 0; k < 4; k++ {
				ni, nj := routes[i].i+di[k], routes[i].j+dj[k]
				if !InField(n, ni, nj) {
					continue
				}

				if openedWater[ni][nj] {
					uf.Unite(ni*n+nj, routes[i].i*n+routes[i].j)
					isOpenedWater = true
				}

			}
			if !uf.ExistSameUnion(ti*n+tj, routes[i].i*n+routes[i].j) {
				uf.Unite(ti*n+tj, routes[i].i*n+routes[i].j)
			}

			if !crushed[routes[i].i][routes[i].j] {
				//	continue
				//}
				var status int
				var cnt int
				for status != 1 {
					power := defaultPower + cnt*s
					status = Excavate(routes[i].i, routes[i].j, power)
					ans = append(ans, Excavation{routes[i].i, routes[i].j, power})
					if status == -1 {
						return nil, nil, errors.New("Already Crushed.")
					}
					if status == 2 {
						return ans, field, nil
					}
					cnt++
				}
				crushed[routes[i].i][routes[i].j] = true
				//openedWater[routes[i].i][routes[i].j] = true
				field[routes[i].i][routes[i].j] = 0
			}
			if isOpenedWater {
				for j := i; j >= 0; j-- {
					openedWater[routes[j].i][routes[j].j] = true
				}
				//fmt.Println(houseId, " route breaked.")
				break
			}
			pi, pj = routes[i].i, routes[i].j
		}
		openedWater[hs[houseId].i][hs[houseId].j] = true
		areOpend |= 1 << houseId
	}
	return ans, field, nil
}

//事前に岩盤調査を2段階行う
func solve04(n, w, k, s, investPower1, investPower2, defaultPower int, a, b, c, d []int) ([]Excavation, [][]int, error) {
	const INF = 1 << 60
	const investWidth = 20
	const weightWidth = investWidth / 2

	uf := NewUnionFind(n * n)

	//水源、家を独自データ構造にしておく
	var wss []WaterSource
	for i := range a {
		wss = append(wss, WaterSource{i, a[i], b[i]})
	}
	var hs []House
	for i := range c {
		hs = append(hs, House{i, c[i], d[i]})
	}

	field := make([][]int, n)
	crushed := make([][]bool, n)
	openedWater := make([][]bool, n)
	for i := 0; i < n; i++ {
		field[i] = make([]int, n)
		crushed[i] = make([]bool, n)
		openedWater[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			field[i][j] = 5000*investWidth*investWidth + 1
		}
	}
	var ans []Excavation
	//岩盤調査
	//fmt.Println("岩盤調査")
	for i := weightWidth - 1; i < n; i += investWidth {
		for j := weightWidth - 1; j < n; j += investWidth {
			status := Excavate(i, j, investPower1)
			ans = append(ans, Excavation{i, j, investPower1})

			if status == 1 {
				crushed[i][j] = true
				//field[i][j] = 0
				//割れた岩盤付近に重みをつける
				for ii := Max(i-weightWidth+1, 0); ii < Min(i+weightWidth, n); ii++ {
					for jj := Max(j-weightWidth+1, 0); jj < Min(j+weightWidth, n); jj++ {
						field[ii][jj] = 1
					}
				}
			}
		}
	}
	//2回目の岩盤調査
	for i := weightWidth - 1; i < n; i += investWidth {
		for j := weightWidth - 1; j < n; j += investWidth {
			if crushed[i][j] {
				continue
			}
			power := investPower2 - investPower1
			status := Excavate(i, j, power)
			ans = append(ans, Excavation{i, j, power})

			if status == 1 {
				crushed[i][j] = true
				//field[i][j] = 0
				//割れた岩盤付近に重みをつける
				for ii := Max(i-weightWidth+1, 0); ii < Min(i+weightWidth, n); ii++ {
					for jj := Max(j-weightWidth+1, 0); jj < Min(j+weightWidth, n); jj++ {
						field[ii][jj] = power*investWidth*investWidth + 1
					}
				}
			}
		}
	}

	//ExportImage(n, field)

	//水源と家の岩盤をあらかじめ砕いておく
	//fmt.Println("水源と家の岩盤採掘")
	for _, ws := range wss {
		var cnt int
		var status int
		if crushed[ws.i][ws.j] {
			status = 1
		}
		for status != 1 {
			power := defaultPower + cnt*s
			status = Excavate(ws.i, ws.j, power)
			ans = append(ans, Excavation{ws.i, ws.j, power})
			if status == -1 {
				return nil, nil, errors.New("Already Crushed.")
			}
			if status == 2 {
				return ans, field, nil
			}
			cnt++
		}
		crushed[ws.i][ws.j] = true
		field[ws.i][ws.j] = 0
		openedWater[ws.i][ws.j] = true
	}
	for _, h := range hs {
		var cnt int
		var status int
		if crushed[h.i][h.j] {
			status = 1
		}
		for status != 1 {
			power := defaultPower + cnt*s
			status = Excavate(h.i, h.j, power)
			ans = append(ans, Excavation{h.i, h.j, power})
			if status == -1 {
				return nil, nil, errors.New("Already Crushed.")
			}
			if status == 2 {
				return ans, field, nil
			}
			cnt++
		}
		crushed[h.i][h.j] = true
		field[h.i][h.j] = 0
	}

	dijkstra := func(si, sj int) ([][]int, []int) {
		q := &PriorityQueue{}
		heap.Init(q)

		dist := make([][]int, n)
		for i := range dist {
			dist[i] = make([]int, n)
			for j := 0; j < n; j++ {
				dist[i][j] = INF
			}
		}
		path := make([]int, n*n)
		for i := 0; i < n*n; i++ {
			path[i] = i
		}
		push := func(ci, cj, ti, tj, cost int) {
			if dist[ti][tj] <= cost {
				return
			}
			dist[ti][tj] = cost
			heap.Push(q, NodeV3{ti, tj, cost})
			path[ti*n+tj] = ci*n + cj
		}
		push(si, sj, si, sj, 0)
		//di := []int{-1, 0, 1, 0}
		//dj := []int{0, -1, 0, 1}
		for q.Len() > 0 {
			cur := heap.Pop(q).(NodeV3)
			if dist[cur.i][cur.j] < cur.w {
				continue
			}
			for k := 0; k < 4; k++ {
				ni, nj := cur.i+di[k], cur.j+dj[k]
				if !InField(n, ni, nj) {
					continue
				}
				push(cur.i, cur.j, ni, nj, cur.w+field[ni][nj])
			}
		}
		return dist, path
	}

	//const startPower = 50
	//目指す水源の選択
	//水源が通った家の管理
	var areOpend int
	var path []int
	for areOpend < 1<<k-1 {
		//fmt.Printf("areOpened, mask = %b, %b\n", areOpend, 1<<k-1)
		houseId := -1
		dist := INF
		ti, tj := -1, -1
		//var dists [][]int
		//各家ごとの処理
		for _, h := range hs {
			//すでに水が通っている家は対象外とする
			if (areOpend>>h.id)&1 == 1 {
				continue
			}

			//一番近い水源を計算
			tdists, tpath := dijkstra(h.i, h.j)

			//家hから最も近い水源までの距離、座標
			d := INF
			tti, ttj := -1, -1
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if !openedWater[i][j] {
						continue
					}
					if d > tdists[i][j] {
						d = tdists[i][j]
						tti, ttj = i, j
					}
				}
			}
			if dist > d {
				houseId = h.id
				dist = d
				ti, tj = tti, ttj
				path = tpath
				//dists = tdists
			}
		}
		//fmt.Printf("id = %d, (ti, tj) = (%d, %d), (hi, hj) = (%d, %d)\n", houseId, ti, tj, hs[houseId].i, hs[houseId].j)

		restructRoute := func(ti, tj int) []Cell {
			res := []Cell{{ti, tj}}
			idx := ti*n + tj
			for path[idx] != idx {
				idx = path[idx]
				res = append(res, Cell{idx / n, idx % n})
			}
			return res
		}
		routes := restructRoute(ti, tj) //dfs(ti, tj, hs[houseId].i, hs[houseId].j)
		m := len(routes)
		//fmt.Println(routes[m-1].i, routes[m-1].j, routes[0].i, routes[0].j, len(routes))
		//fmt.Println(routes)
		pi, pj := routes[m-1].i, routes[m-1].j
		for i := m - 2; i >= 0; i-- {
			if !uf.ExistSameUnion(routes[i].i*n+routes[i].j, pi*n+pj) {
				uf.Unite(routes[i].i*n+routes[i].j, pi*n+pj)
			}
			var isOpenedWater bool
			for k := 0; k < 4; k++ {
				ni, nj := routes[i].i+di[k], routes[i].j+dj[k]
				if !InField(n, ni, nj) {
					continue
				}

				if openedWater[ni][nj] {
					uf.Unite(ni*n+nj, routes[i].i*n+routes[i].j)
					isOpenedWater = true
				}

			}
			if !uf.ExistSameUnion(ti*n+tj, routes[i].i*n+routes[i].j) {
				uf.Unite(ti*n+tj, routes[i].i*n+routes[i].j)
			}

			if !crushed[routes[i].i][routes[i].j] {
				//	continue
				//}
				var status int
				var cnt int
				for status != 1 {
					power := defaultPower + cnt*s
					status = Excavate(routes[i].i, routes[i].j, power)
					ans = append(ans, Excavation{routes[i].i, routes[i].j, power})
					if status == -1 {
						return nil, nil, errors.New("Already Crushed.")
					}
					if status == 2 {
						return ans, field, nil
					}
					cnt++
				}
				crushed[routes[i].i][routes[i].j] = true
				//openedWater[routes[i].i][routes[i].j] = true
				field[routes[i].i][routes[i].j] = 0
			}
			if isOpenedWater {
				for j := i; j >= 0; j-- {
					openedWater[routes[j].i][routes[j].j] = true
				}
				//fmt.Println(houseId, " route breaked.")
				break
			}
			pi, pj = routes[i].i, routes[i].j
		}
		openedWater[hs[houseId].i][hs[houseId].j] = true
		areOpend |= 1 << houseId
	}
	return ans, field, nil
}

//事前に岩盤調査をする
func solve035(n, w, k, s, defaultPower int, a, b, c, d []int) ([]Excavation, [][]int, error) {
	const INF = 1 << 60
	const p = 100
	const investWidth = 20
	const weightWidth = investWidth / 2

	uf := NewUnionFind(n * n)

	//水源、家を独自データ構造にしておく
	var wss []WaterSource
	for i := range a {
		wss = append(wss, WaterSource{i, a[i], b[i]})
	}
	var hs []House
	for i := range c {
		hs = append(hs, House{i, c[i], d[i]})
	}

	field := make([][]int, n)
	crushed := make([][]bool, n)
	openedWater := make([][]bool, n)
	for i := 0; i < n; i++ {
		field[i] = make([]int, n)
		crushed[i] = make([]bool, n)
		openedWater[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			field[i][j] = 100 * p
		}
	}
	var ans []Excavation
	//岩盤調査
	//fmt.Println("岩盤調査")
	for i := weightWidth - 1; i < n; i += investWidth {
		for j := weightWidth - 1; j < n; j += investWidth {
			status := Excavate(i, j, p)
			ans = append(ans, Excavation{i, j, p})

			if status == 1 {
				crushed[i][j] = true
				//field[i][j] = 0
				//割れた岩盤付近に重みをつける
				for ii := Max(i-weightWidth+1, 0); ii < Min(i+weightWidth, n); ii++ {
					for jj := Max(j-weightWidth+1, 0); jj < Min(j+weightWidth, n); jj++ {
						//w := ComputeManhattanDist(i, j, ii, jj)
						//if w == 0 {
						//	field[ii][jj] = Max(field[ii][jj]-100*p, 0)
						//} else {
						//	field[ii][jj] = Max(field[ii][jj]-100*p/w, 0)
						//}
					}
				}
			} else {
				//割れなかった岩盤付近に重みをつける
				for ii := Max(i-weightWidth+1, 0); ii <= Min(i+weightWidth, n); ii++ {
					for jj := Max(j-weightWidth+1, 0); jj <= Min(j+weightWidth, n); jj++ {
						field[ii][jj] += 100 * p
						//w := ComputeManhattanDist(i, j, ii, jj)
						//if w == 0 {
						//	field[ii][jj] += 100 * p
						//} else {
						//	field[ii][jj] += 100 * p / w
						//}
					}
				}
			}
		}
	}
	//ExportImage(n, field)

	//水源と家の岩盤をあらかじめ砕いておく
	//fmt.Println("水源と家の岩盤採掘")
	for _, ws := range wss {
		var cnt int
		var status int
		if crushed[ws.i][ws.j] {
			status = 1
		}
		for status != 1 {
			power := p + cnt*s
			status = Excavate(ws.i, ws.j, power)
			ans = append(ans, Excavation{ws.i, ws.j, power})
			if status == -1 {
				return nil, nil, errors.New("Already Crushed.")
			}
			if status == 2 {
				return ans, field, nil
			}
			cnt++
		}
		crushed[ws.i][ws.j] = true
		field[ws.i][ws.j] = 0
		openedWater[ws.i][ws.j] = true
	}
	for _, h := range hs {
		var cnt int
		var status int
		if crushed[h.i][h.j] {
			status = 1
		}
		for status != 1 {
			power := p + cnt*s
			status = Excavate(h.i, h.j, power)
			ans = append(ans, Excavation{h.i, h.j, power})
			if status == -1 {
				return nil, nil, errors.New("Already Crushed.")
			}
			if status == 2 {
				return ans, field, nil
			}
			cnt++
		}
		crushed[h.i][h.j] = true
		field[h.i][h.j] = 0
	}

	dijkstra := func(si, sj int) ([][]int, []int) {
		q := &PriorityQueue{}
		heap.Init(q)

		dist := make([][]int, n)
		for i := range dist {
			dist[i] = make([]int, n)
			for j := 0; j < n; j++ {
				dist[i][j] = INF
			}
		}
		path := make([]int, n*n)
		for i := 0; i < n*n; i++ {
			path[i] = i
		}
		push := func(ci, cj, ti, tj, cost int) {
			if dist[ti][tj] <= cost {
				return
			}
			dist[ti][tj] = cost
			heap.Push(q, NodeV3{ti, tj, cost})
			path[ti*n+tj] = ci*n + cj
		}
		push(si, sj, si, sj, 0)
		//di := []int{-1, 0, 1, 0}
		//dj := []int{0, -1, 0, 1}
		for q.Len() > 0 {
			cur := heap.Pop(q).(NodeV3)
			if dist[cur.i][cur.j] < cur.w {
				continue
			}
			for k := 0; k < 4; k++ {
				ni, nj := cur.i+di[k], cur.j+dj[k]
				if !InField(n, ni, nj) {
					continue
				}
				push(cur.i, cur.j, ni, nj, cur.w+field[ni][nj])
			}
		}
		return dist, path
	}

	//const startPower = 50
	//目指す水源の選択
	//水源が通った家の管理
	var areOpend int
	var path []int
	for areOpend < 1<<k-1 {
		//fmt.Printf("areOpened, mask = %b, %b\n", areOpend, 1<<k-1)
		houseId := -1
		dist := INF
		ti, tj := -1, -1
		//var dists [][]int
		//各家ごとの処理
		for _, h := range hs {
			//すでに水が通っている家は対象外とする
			if (areOpend>>h.id)&1 == 1 {
				continue
			}

			//一番近い水源を計算
			tdists, tpath := dijkstra(h.i, h.j)

			//家hから最も近い水源までの距離、座標
			d := INF
			tti, ttj := -1, -1
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if !openedWater[i][j] {
						continue
					}
					if d > tdists[i][j] {
						d = tdists[i][j]
						tti, ttj = i, j
					}
				}
			}
			if dist > d {
				houseId = h.id
				dist = d
				ti, tj = tti, ttj
				path = tpath
				//dists = tdists
			}
		}
		//fmt.Printf("id = %d, (ti, tj) = (%d, %d), (hi, hj) = (%d, %d)\n", houseId, ti, tj, hs[houseId].i, hs[houseId].j)

		restructRoute := func(ti, tj int) []Cell {
			res := []Cell{{ti, tj}}
			idx := ti*n + tj
			for path[idx] != idx {
				idx = path[idx]
				res = append(res, Cell{idx / n, idx % n})
			}
			return res
		}
		routes := restructRoute(ti, tj) //dfs(ti, tj, hs[houseId].i, hs[houseId].j)
		m := len(routes)
		//fmt.Println(routes[m-1].i, routes[m-1].j, routes[0].i, routes[0].j, len(routes))
		//fmt.Println(routes)
		pi, pj := routes[m-1].i, routes[m-1].j
		for i := m - 2; i >= 0; i-- {
			if !uf.ExistSameUnion(routes[i].i*n+routes[i].j, pi*n+pj) {
				uf.Unite(routes[i].i*n+routes[i].j, pi*n+pj)
			}
			var isOpenedWater bool
			for k := 0; k < 4; k++ {
				ni, nj := routes[i].i+di[k], routes[i].j+dj[k]
				if !InField(n, ni, nj) {
					continue
				}

				if openedWater[ni][nj] {
					uf.Unite(ni*n+nj, routes[i].i*n+routes[i].j)
					isOpenedWater = true
				}

			}
			if !uf.ExistSameUnion(ti*n+tj, routes[i].i*n+routes[i].j) {
				uf.Unite(ti*n+tj, routes[i].i*n+routes[i].j)
			}

			if !crushed[routes[i].i][routes[i].j] {
				//	continue
				//}
				var status int
				var cnt int
				for status != 1 {
					power := defaultPower + cnt*s
					status = Excavate(routes[i].i, routes[i].j, power)
					ans = append(ans, Excavation{routes[i].i, routes[i].j, power})
					if status == -1 {
						return nil, nil, errors.New("Already Crushed.")
					}
					if status == 2 {
						return ans, field, nil
					}
					cnt++
				}
				crushed[routes[i].i][routes[i].j] = true
				//openedWater[routes[i].i][routes[i].j] = true
				field[routes[i].i][routes[i].j] = 0
			}
			if isOpenedWater {
				for j := i; j >= 0; j-- {
					openedWater[routes[j].i][routes[j].j] = true
				}
				//fmt.Println(houseId, " route breaked.")
				break
			}
			pi, pj = routes[i].i, routes[i].j
		}
		openedWater[hs[houseId].i][hs[houseId].j] = true
		areOpend |= 1 << houseId
	}
	return ans, field, nil
}

//事前に岩盤調査をする
func solve03(n, w, k, s int, a, b, c, d []int) ([]Excavation, [][]int, error) {
	const INF = 1 << 60
	const p = 100
	const investWidth = 20
	const weightWidth = investWidth / 2

	uf := NewUnionFind(n * n)

	//水源、家を独自データ構造にしておく
	var wss []WaterSource
	for i := range a {
		wss = append(wss, WaterSource{i, a[i], b[i]})
	}
	var hs []House
	for i := range c {
		hs = append(hs, House{i, c[i], d[i]})
	}

	field := make([][]int, n)
	crushed := make([][]bool, n)
	openedWater := make([][]bool, n)
	for i := 0; i < n; i++ {
		field[i] = make([]int, n)
		crushed[i] = make([]bool, n)
		openedWater[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			field[i][j] = 100 * p
		}
	}
	var ans []Excavation
	//岩盤調査
	//fmt.Println("岩盤調査")
	for i := weightWidth - 1; i < n; i += investWidth {
		for j := weightWidth - 1; j < n; j += investWidth {
			status := Excavate(i, j, p)
			ans = append(ans, Excavation{i, j, p})

			if status == 1 {
				crushed[i][j] = true
				//field[i][j] = 0
				//割れた岩盤付近に重みをつける
				for ii := Max(i-weightWidth+1, 0); ii < Min(i+weightWidth, n); ii++ {
					for jj := Max(j-weightWidth+1, 0); jj < Min(j+weightWidth, n); jj++ {
						//w := ComputeManhattanDist(i, j, ii, jj)
						//if w == 0 {
						//	field[ii][jj] = Max(field[ii][jj]-100*p, 0)
						//} else {
						//	field[ii][jj] = Max(field[ii][jj]-100*p/w, 0)
						//}
					}
				}
			} else {
				//割れなかった岩盤付近に重みをつける
				for ii := Max(i-weightWidth+1, 0); ii <= Min(i+weightWidth, n); ii++ {
					for jj := Max(j-weightWidth+1, 0); jj <= Min(j+weightWidth, n); jj++ {
						field[ii][jj] += 100 * p
						//w := ComputeManhattanDist(i, j, ii, jj)
						//if w == 0 {
						//	field[ii][jj] += 100 * p
						//} else {
						//	field[ii][jj] += 100 * p / w
						//}
					}
				}
			}
		}
	}
	//ExportImage(n, field)

	//水源と家の岩盤をあらかじめ砕いておく
	//fmt.Println("水源と家の岩盤採掘")
	for _, ws := range wss {
		var cnt int
		var status int
		if crushed[ws.i][ws.j] {
			status = 1
		}
		for status != 1 {
			power := p + cnt*s
			status = Excavate(ws.i, ws.j, power)
			ans = append(ans, Excavation{ws.i, ws.j, power})
			if status == -1 {
				return nil, nil, errors.New("Already Crushed.")
			}
			if status == 2 {
				return ans, field, nil
			}
			cnt++
		}
		crushed[ws.i][ws.j] = true
		field[ws.i][ws.j] = 0
		openedWater[ws.i][ws.j] = true
	}
	for _, h := range hs {
		var cnt int
		var status int
		if crushed[h.i][h.j] {
			status = 1
		}
		for status != 1 {
			power := p + cnt*s
			status = Excavate(h.i, h.j, power)
			ans = append(ans, Excavation{h.i, h.j, power})
			if status == -1 {
				return nil, nil, errors.New("Already Crushed.")
			}
			if status == 2 {
				return ans, field, nil
			}
			cnt++
		}
		crushed[h.i][h.j] = true
		field[h.i][h.j] = 0
	}

	dijkstra := func(si, sj int) ([][]int, []int) {
		q := &PriorityQueue{}
		heap.Init(q)

		dist := make([][]int, n)
		for i := range dist {
			dist[i] = make([]int, n)
			for j := 0; j < n; j++ {
				dist[i][j] = INF
			}
		}
		path := make([]int, n*n)
		for i := 0; i < n*n; i++ {
			path[i] = i
		}
		push := func(ci, cj, ti, tj, cost int) {
			if dist[ti][tj] <= cost {
				return
			}
			dist[ti][tj] = cost
			heap.Push(q, NodeV3{ti, tj, cost})
			path[ti*n+tj] = ci*n + cj
		}
		push(si, sj, si, sj, 0)
		//di := []int{-1, 0, 1, 0}
		//dj := []int{0, -1, 0, 1}
		for q.Len() > 0 {
			cur := heap.Pop(q).(NodeV3)
			if dist[cur.i][cur.j] < cur.w {
				continue
			}
			for k := 0; k < 4; k++ {
				ni, nj := cur.i+di[k], cur.j+dj[k]
				if !InField(n, ni, nj) {
					continue
				}
				push(cur.i, cur.j, ni, nj, cur.w+field[ni][nj])
			}
		}
		return dist, path
	}

	//目指す水源の選択
	//水源が通った家の管理
	var areOpend int
	var path []int
	for areOpend < 1<<k-1 {
		//fmt.Printf("areOpened, mask = %b, %b\n", areOpend, 1<<k-1)
		houseId := -1
		dist := INF
		ti, tj := -1, -1
		//var dists [][]int
		//各家ごとの処理
		for _, h := range hs {
			//すでに水が通っている家は対象外とする
			if (areOpend>>h.id)&1 == 1 {
				continue
			}

			//一番近い水源を計算
			tdists, tpath := dijkstra(h.i, h.j)

			//家hから最も近い水源までの距離、座標
			d := INF
			tti, ttj := -1, -1
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if !openedWater[i][j] {
						continue
					}
					if d > tdists[i][j] {
						d = tdists[i][j]
						tti, ttj = i, j
					}
				}
			}
			if dist > d {
				houseId = h.id
				dist = d
				ti, tj = tti, ttj
				path = tpath
				//dists = tdists
			}
		}
		//fmt.Printf("id = %d, (ti, tj) = (%d, %d), (hi, hj) = (%d, %d)\n", houseId, ti, tj, hs[houseId].i, hs[houseId].j)

		restructRoute := func(ti, tj int) []Cell {
			res := []Cell{{ti, tj}}
			idx := ti*n + tj
			for path[idx] != idx {
				idx = path[idx]
				res = append(res, Cell{idx / n, idx % n})
			}
			return res
		}
		routes := restructRoute(ti, tj) //dfs(ti, tj, hs[houseId].i, hs[houseId].j)
		m := len(routes)
		//fmt.Println(routes[m-1].i, routes[m-1].j, routes[0].i, routes[0].j, len(routes))
		//fmt.Println(routes)
		pi, pj := routes[m-1].i, routes[m-1].j
		for i := m - 2; i >= 0; i-- {
			if !uf.ExistSameUnion(routes[i].i*n+routes[i].j, pi*n+pj) {
				uf.Unite(routes[i].i*n+routes[i].j, pi*n+pj)
			}
			var isOpenedWater bool
			for k := 0; k < 4; k++ {
				ni, nj := routes[i].i+di[k], routes[i].j+dj[k]
				if !InField(n, ni, nj) {
					continue
				}

				if openedWater[ni][nj] {
					uf.Unite(ni*n+nj, routes[i].i*n+routes[i].j)
					isOpenedWater = true
				}

			}
			if !uf.ExistSameUnion(ti*n+tj, routes[i].i*n+routes[i].j) {
				uf.Unite(ti*n+tj, routes[i].i*n+routes[i].j)
			}

			if !crushed[routes[i].i][routes[i].j] {
				//	continue
				//}
				var status int
				var cnt int
				for status != 1 {
					power := p + cnt*s
					status = Excavate(routes[i].i, routes[i].j, power)
					ans = append(ans, Excavation{routes[i].i, routes[i].j, power})
					if status == -1 {
						return nil, nil, errors.New("Already Crushed.")
					}
					if status == 2 {
						return ans, field, nil
					}
					cnt++
				}
				crushed[routes[i].i][routes[i].j] = true
				//openedWater[routes[i].i][routes[i].j] = true
				field[routes[i].i][routes[i].j] = 0
			}
			if isOpenedWater {
				for j := i; j >= 0; j-- {
					openedWater[routes[j].i][routes[j].j] = true
				}
				//fmt.Println(houseId, " route breaked.")
				break
			}
			pi, pj = routes[i].i, routes[i].j
		}
		openedWater[hs[houseId].i][hs[houseId].j] = true
		areOpend |= 1 << houseId
	}
	return ans, field, nil
}

//等間隔に、デフォルトのパワーで採掘できれば1、そうでなければ0

type NodeV3 struct {
	i, j, w int
}

type PriorityQueue []NodeV3

func (q PriorityQueue) Len() int {
	return len(q)
}

func (q PriorityQueue) Less(i, j int) bool {
	return q[i].w < q[j].w
}

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(NodeV3))
}

func (pq *PriorityQueue) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

//柔らかい岩盤優先で採掘する2
func solve02(n, w, k, s int, a, b, c, d []int) ([]Excavation, error) {
	const INF = 1 << 60
	const p = 100
	var wss []WaterSource
	for i := range a {
		wss = append(wss, WaterSource{i, a[i], b[i]})
	}
	var hs []House
	for i := range c {
		hs = append(hs, House{i, c[i], d[i]})
	}

	//各家に水が通ったかの管理
	completed := make([]bool, k)

	//柔らかい岩盤優先キュー
	tq := &TimesPriorityQueue{}
	//硬い岩盤優先キュー
	dq := &DistPriorityQueue{}
	heap.Init(tq)
	heap.Init(dq)

	//水源と家との連結を管理
	uf1 := NewUnionFind(w + k)
	uf := NewUnionFind(n * n)

	//プランニング
	//visited := make([][]bool, n)
	//for i := 0; i < n; i++ {
	//	visited[i] = make([]bool, n)
	//}

	//キューの準備
	for _, h := range hs {
		for _, ws := range wss {
			heap.Push(tq, Node{h.id, ws.id, h.i, h.j, ws.i, ws.j, 0, ComputeManhattanDist(h.i, h.j, ws.i, ws.j)})
			//visited[h.i][h.j] = true
		}
	}
	//fmt.Println(q)

	crushed := make([][]bool, n)
	for i := 0; i < n; i++ {
		crushed[i] = make([]bool, n)
	}
	var ans []Excavation

	//事前の地盤調査
	//for i := 0; i < n; i += 10 {
	//	for j := 0; j < n; j += 10 {
	//		status := Excavate(i, j, p)
	//		if status == 1 {
	//			crushed[i][j] = true
	//		}
	//		ans = append(ans, Excavation{i, j, p})
	//	}
	//}

	bfs := func(waterSourceId, ci, cj int) (int, int) {
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, n)
		}
		var q []Cell
		si, sj := wss[waterSourceId].i, wss[waterSourceId].j
		q = append(q, Cell{si, sj})
		visited[si][sj] = true
		dist := ComputeManhattanDist(si, sj, ci, cj)
		ti, tj := si, sj
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			di := []int{-1, 0, 1, 0}
			dj := []int{0, -1, 0, 1}
			for k := 0; k < 4; k++ {
				ni, nj := cur.i+di[k], cur.j+dj[k]
				if ni < 0 || ni >= n || nj < 0 || nj >= n {
					continue
				}
				if visited[ni][nj] {
					continue
				}
				if !crushed[ni][nj] {
					continue
				}
				q = append(q, Cell{ni, nj})
				visited[ni][nj] = true
				newDist := ComputeManhattanDist(ni, nj, ci, cj)
				if dist > newDist {
					dist = newDist
					ti, tj = ni, nj
				}
			}
		}
		return ti, tj
	}
	replan := func() {
		//fmt.Println("Called replan")
		nextTq := &TimesPriorityQueue{}
		heap.Init(nextTq)
		for tq.Len() > 0 {
			cur := heap.Pop(tq).(Node)

			ti, tj := bfs(cur.waterSourceId, cur.i, cur.j)
			//fmt.Printf("Change Target: houseId:%d waterSourceId:%d (%d, %d) -> (%d, %d)\n", cur.houseId, cur.waterSourceId, cur.ti, cur.tj, ti, tj)
			dist := ComputeManhattanDist(cur.i, cur.j, ti, tj)
			heap.Push(nextTq, Node{cur.houseId, cur.waterSourceId, cur.i, cur.j, ti, tj, 0, dist})

			//heap.Push(nextTq, cur)
		}
		tq = nextTq
		nextDq := &DistPriorityQueue{}
		heap.Init(nextDq)
		for dq.Len() > 0 {
			cur := heap.Pop(dq).(Node)

			ti, tj := bfs(cur.waterSourceId, cur.i, cur.j)
			dist := ComputeManhattanDist(cur.i, cur.j, ti, tj)
			heap.Push(nextTq, Node{cur.houseId, cur.waterSourceId, cur.i, cur.j, ti, tj, 0, dist})

			//heap.Push(nextTq, cur)
		}
		dq = nextDq
	}

	//探索の処理
	for tq.Len() > 0 || dq.Len() > 0 {
		var cur Node

		if tq.Len() > 0 {
			cur = heap.Pop(tq).(Node)
		} else {
			cur = heap.Pop(dq).(Node)
		}
		//すでに水が通っている家に関する探索であれば探索をしない
		if completed[cur.houseId] {
			continue
		}

		//ここで採掘してみる
		if !crushed[cur.i][cur.j] {
			power := p + cur.times*s
			status := Excavate(cur.i, cur.j, power)
			ans = append(ans, Excavation{cur.i, cur.j, power})
			if status == -1 {
				return nil, errors.New("Already Crushed")
			} else if status == 0 {
				//取り出したキューに戻す
				if cur.times < 10 {
					heap.Push(tq, Node{cur.houseId, cur.waterSourceId, cur.i, cur.j, cur.ti, cur.tj, cur.times + 1, cur.dist})
				} else {
					heap.Push(dq, Node{cur.houseId, cur.waterSourceId, cur.i, cur.j, cur.ti, cur.tj, cur.times + 1, cur.dist})
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
						if uf.ExistSameUnion(cur.i*n+cur.j, ws.i*n+ws.j) {
							uf1.Unite(cur.waterSourceId, w+cur.houseId)
							completed[cur.houseId] = true
							//rePlan = true
							replan()
						}
					}
				}
			} else if status == 2 {
				return ans, nil
			}
		}

		//これ以降は注目しているセルが採掘完了した時、次の目的地候補をキューに詰む

		//キューに詰む
		//for idx := range di {
		ni, nj := PlanRoute(cur.i, cur.j, cur.ti, cur.tj) //cur.i+di[idx], cur.j+dj[idx]
		if ni < 0 || ni >= n || nj < 0 || nj >= n {
			continue
		}
		if crushed[ni][nj] {
			continue
		}
		dist := ComputeManhattanDist(ni, nj, cur.ti, cur.tj)
		heap.Push(tq, Node{cur.houseId, cur.waterSourceId, ni, nj, cur.ti, cur.tj, 0, dist})
		//visited[ni][nj] = true
	}
	//ここまでくるとすべての家に水が通せなかった
	return ans, errors.New("Can't Complete.")

}

func PlanRoute(si, sj, ti, tj int) (int, int) {
	di, dj := ti-si, tj-sj
	v := rand.Intn(2)
	if di == 0 {
		if dj > 0 {
			return si, sj + 1
		} else if dj < 0 {
			return si, sj - 1
		}
	} else if di > 0 {
		if dj > 0 {
			if v == 0 {
				return si + 1, sj
			} else {
				return si, sj + 1
			}
		} else if dj < 0 {
			if v == 0 {
				return si + 1, sj
			} else {
				return si, sj - 1
			}
		} else {
			return si + 1, sj
		}
	} else {
		// di < 0
		if dj > 0 {
			if v == 0 {
				return si - 1, sj
			} else {
				return si, sj + 1
			}
		} else if dj < 0 {
			if v == 0 {
				return si - 1, sj
			} else {
				return si, sj - 1
			}
		} else {
			return si - 1, sj
		}
	}
	return si, sj
}

//問題特有の汎用ライブラリ
func ComputeManhattanDist(i1, j1, i2, j2 int) int {
	return Abs(i2-i1) + Abs(j2-j1)
}

type Node struct {
	houseId       int
	waterSourceId int
	i, j          int //現在地
	ti, tj        int //目的地
	times         int
	dist          int
}

type TimesPriorityQueue []Node

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
	*pq = append(*pq, item.(Node))
}

func (pq *TimesPriorityQueue) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

type DistPriorityQueue []Node

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
	*pq = append(*pq, item.(Node))
}

func (pq *DistPriorityQueue) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

//テンプレート外に定義している典型的なライブラリ類
type UnionFind struct {
	parent []int // parentent numbers
	rank   []int // height of tree
	size   []int
}

//func New(n int) *UnionFind {
//	return NewUnionFind(n)
//}

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

//テンプレートにあらかじめ用意している典型的なライブラリ類
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
