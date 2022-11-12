package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, a [][]int) (int, error) {
	isConnected := make([]map[int]struct{}, n)
	q := make([][]int, n)
	isPlayed := make([][]bool, n)
	for i := 0; i < n; i++ {
		isConnected[i] = make(map[int]struct{})
		isPlayed[i] = make([]bool, n)
		isPlayed[i][i] = true
	}
	var ans int
	var cnt int
	for {
		//ans日目に行われる試合数
		s := 0
		//ans日目に試合予定があるかどうか
		isPlayToday := make([]bool, n)
		if ans < n-1 {
			for i := 0; i < n; i++ {
				isConnected[i][a[i][ans]-1] = struct{}{}
				q[i] = append(q[i], a[i][ans]-1)
			}
		}
		for i := 0; i < n; i++ {
			cnt++
			if len(q[i]) == 0 {
				continue
			}
			if len(q[q[i][0]]) == 0 {
				continue
			}
			k1 := q[i][0]
			k2 := q[k1][0]
			if i == k2 && !isPlayed[k1][k2] && !isPlayed[k2][k1] && !isPlayToday[k1] && !isPlayToday[k2] {
				s++
				q[i] = q[i][1:]
				q[k1] = q[k1][1:]
				isPlayed[k1][k2] = true
				isPlayed[k2][k1] = true
				isPlayToday[k1] = true
				isPlayToday[k2] = true
			}
		}
		/*
			for k1 := range isConnected {
				for k2 := range isConnected[k1] {
					_, f1 := isConnected[k1][k2]
					_, f2 := isConnected[k2][k1]
					if f1 && f2 && !isPlayed[k1][k2] && !isPlayed[k2][k1] && !isPlayToday[k1] && !isPlayToday[k2] {
						s++
						delete(isConnected[k1], k2)
						delete(isConnected[k2], k1)
						isPlayed[k1][k2] = true
						isPlayed[k2][k1] = true
						isPlayToday[k1] = true
						isPlayToday[k2] = true
					}
				}
			}
		*/
		//fmt.Println("s = ", s)
		//fmt.Println(isConnected)
		//PrintVertically(isPlayed)

		if s == 0 {
			break
		}
		ans++
	}
	fmt.Println("cnt = ", cnt)
	ok := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ok = ok && isPlayed[i][j]
		}
	}
	if ok {
		return ans, nil
	} else {
		return 0, errors.New("Impossible")
	}
}

type Pair struct {
	i, a, b int
}
type PriorityQueue []Pair

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].i < pq[j].i
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(Pair))
}

func (pq *PriorityQueue) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

func solve02(n int, a [][]int) (int, error) {
	isConnected := make([]map[int]struct{}, n)
	isPlayed := make([][]bool, n)
	for i := 0; i < n; i++ {
		isConnected[i] = make(map[int]struct{})
		isPlayed[i] = make([]bool, n)
		isPlayed[i][i] = true
	}
	q := &PriorityQueue{}
	heap.Init(q)
	m := make([]map[[2]int]struct{}, n*n)
	for i := 0; i < n*n; i++ {
		m[i] = make(map[[2]int]struct{})
	}
	for j := 0; j < n-1; j++ {
		for i := 0; i < n; i++ {
			x, y := i, a[i][j]-1
			if x > y {
				x, y = y, x
			}
			heap.Push(q, Pair{j, x, y})
		}
	}
	var ans int
	for q.Len() > 0 && ans <= n*n {
		p := heap.Pop(q).(Pair)
		ans = Max(ans, p.i)
		if _, found := m[p.i][[2]int{p.a, p.b}]; found {
			p.i++
			heap.Push(q, p)
		} else {
			m[p.i][[2]int{p.a, p.b}] = struct{}{}
			isPlayed[p.a][p.b] = true
			isPlayed[p.b][p.a] = true
		}
	}
	ok := true
	if ok {
		return ans, nil
	} else {
		return 0, errors.New("Impossible")
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextIntSlice(n - 1)
	}
	ans, err := solve(n, a)
	if err != nil {
		PrintInt(-1)
	} else {
		PrintInt(ans)
	}
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

func PrintHorizonaly(x []bool) {
	defer out.Flush()
	fmt.Fprintf(out, "%v", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %v", x[i])
	}
	fmt.Fprintln(out)
}

func PrintVertically(x [][]bool) {
	defer out.Flush()
	for _, v := range x {
		//fmt.Fprintln(out, v)
		PrintHorizonaly(v)
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
