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

type task struct {
	i, a, b int
	w       float64
}

type taskSlice []task

func (t taskSlice) Len() int {
	return len(t)
}
func (t taskSlice) Less(i, j int) bool {
	//if t[i].b == t[j].b {
	//	return t[i].a < t[j].a
	//}
	//return t[i].b > t[j].b
	return t[i].i < t[j].i
}
func (t taskSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func solveHonestly(n int, a, b []int) int {
	const p = int(1e9) + 7

	var tasks taskSlice
	for i := range a {
		tasks = append(tasks, task{i: i, a: a[i], b: b[i]})
		tasks[i].w = float64(tasks[i].b-1) / float64(tasks[i].a)
	}

	ans := 0
	for {
		s := 0
		level := 1
		for _, t := range tasks {
			s = (s + level*t.a) % p
			level = (level * t.b) % p
		}
		if ans < s {
			ans = s
			fmt.Println(s, tasks)
		}
		if !NextPermutation(tasks) {
			break
		}
	}
	return ans
}

func solve(n int, a, b []int) int {
	const p = int(1e9) + 7
	var tasks taskSlice
	for i := range a {
		tasks = append(tasks, task{i: i, a: a[i], b: b[i]})
		tasks[i].w = float64(tasks[i].b-1) / float64(tasks[i].a)
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].w > tasks[j].w
	})
	level := 1
	ans := 0
	for _, t := range tasks {
		ans = (ans + level*t.a) % p
		level = (level * t.b) % p
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(n)

	//ans := solveHonestly(n, a, b)
	ans := solve(n, a, b)
	PrintInt(ans)
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
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
