package main

import (
	"bufio"
	"fmt"
	"os"
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

	n, m := nextInt(), nextInt()
	var u, v []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}
	k := nextInt()
	var p, d []int
	for i := 0; i < k; i++ {
		p = append(p, nextInt()-1)
		d = append(d, nextInt())
	}
	ok, s := solve03(n, m, u, v, k, p, d)
	if ok {
		PrintString("Yes")
		PrintString(s)
	} else {
		PrintString("No")
	}
}

func solve03(n, m int, u, v []int, k int, p, d []int) (bool, string) {
	const INF = 1 << 60

	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}

	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = INF
		}
	}
	bfs := func(x int) {
		var q []int
		q = append(q, x)
		dist[x][x] = 0

		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, next := range e[cur] {
				if dist[x][next] != INF {
					continue
				}
				q = append(q, next)
				dist[x][next] = dist[x][cur] + 1
			}
		}
	}
	for i := 0; i < n; i++ {
		bfs(i)
	}

	//点piから距離di未満の頂点を白に塗りつぶす
	ans := strings.Split(strings.Repeat("1", n), "")
	for i := 0; i < k; i++ {
		for j := range dist[p[i]] {
			if dist[p[i]][j] < d[i] {
				ans[j] = "0"
			}
		}
	}

	//1個以上の頂点が黒で塗られている
	var containsBlack bool
	for _, c := range ans {
		containsBlack = containsBlack || c == "1"
	}
	if !containsBlack {
		return false, ""
	}
	//点piから距離di未満の頂点に黒が存在しない
	for i := 0; i < k; i++ {
		for j, v := range dist[p[i]] {
			if v < d[i] {
				if ans[j] == "1" {
					return false, ""
				}
			}
		}
	}

	//点piから距離diの頂点が全て白ではない
	for i := 0; i < k; i++ {
		isAllWhite := true
		for j, v := range dist[p[i]] {
			if v == d[i] {
				isAllWhite = isAllWhite && ans[j] == "0"
			}
		}
		if isAllWhite {
			return false, ""
		}
	}

	return true, strings.Join(ans, "")
}

func solve02(n, m int, u, v []int, k int, p, d []int) (bool, string) {
	e1 := make([][]int, n)
	for i := 0; i < m; i++ {
		e1[u[i]] = append(e1[u[i]], v[i])
		e1[v[i]] = append(e1[v[i]], u[i])
	}
	ans := strings.Split(strings.Repeat("?", n), "")

	conditions := make([][]int, n)
	bfs := func(x, d int) bool {
		var q []int
		//visited := make([]bool, n)
		dist := make([]int, n)
		for i := range dist {
			dist[i] = 1 << 60
		}
		q = append(q, x)
		dist[x] = 0
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, next := range e1[cur] {
				if dist[next] != 1<<60 {
					continue
				}
				q = append(q, next)
				dist[next] = dist[cur] + 1
			}
		}
		for i := range dist {
			if dist[i] == d {
				conditions[x] = append(conditions[x], i)
			} else if dist[i] < d {
				if ans[i] == "1" {
					return false
				}
				ans[i] = "0"
			}
		}
		return true
	}
	//矛盾を判定する
	ok := true
	for i := 0; i < k; i++ {
		//条件を構築する
		ok = ok && bfs(p[i], d[i])
	}
	//fmt.Println(conditions, ans)

	for i := range conditions {
		if len(conditions[i]) == 0 {
			continue
		}
		//この中のどれかが黒でなければいけない(全部白で確定しているとだめ)
		var canPaintBlack bool
		for _, idx := range conditions[i] {
			canPaintBlack = canPaintBlack || (ans[idx] != "0")
			//if ans[idx] == "?" {
			//	ans[idx] = "1"
			//}
		}
		ok = ok && canPaintBlack
	}
	if !ok {
		return ok, ""
	}
	//fmt.Println(ans)
	//この先はOKなので文字列を構成する
	for i := 0; i < n; i++ {
		if ans[i] == "?" {
			ans[i] = "1"
		}
	}
	return ok, strings.Join(ans, "")
}

func solve01(n, m int, u, v []int, k int, p, d []int) (bool, string) {
	e1 := make([][]int, n)
	for i := 0; i < m; i++ {
		e1[u[i]] = append(e1[u[i]], v[i])
		e1[v[i]] = append(e1[v[i]], u[i])
	}
	e2 := make([][]int, n)
	degree := make([]int, n)
	bfs := func(x, d int) {
		var q []int
		//visited := make([]bool, n)
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		q = append(q, x)
		dist[x] = 0
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, next := range e1[cur] {
				if dist[next] >= 0 {
					continue
				}
				q = append(q, next)
				dist[next] = dist[cur] + 1
			}
		}
		for i := range dist {
			if dist[i] == d {
				e2[i] = append(e2[i], x)
				degree[x]++
			}
		}
	}
	for i := 0; i < k; i++ {
		bfs(p[i], d[i])
	}
	ans := make([]string, n)
	for i := range ans {
		ans[i] = "0"
	}
	var q2 []int
	for i := range degree {
		if degree[i] == 1 {
			q2 = append(q2, i)
		}
	}
	for len(q2) > 0 {
		cur := q2[0]
		q2 = q2[1:]
		for _, next := range e2[cur] {
			degree[next]--
			if degree[next] == 1 {
				q2 = append(q2, next)
			}
		}
	}
	ok := true
	for _, v := range degree {
		ok = ok && v == 0
	}
	return ok, strings.Join(ans, "")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
