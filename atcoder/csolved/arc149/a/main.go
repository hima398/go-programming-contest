package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func firstSolve(n, m int) string {
	loop := make([]map[int]struct{}, 10)
	for i := 0; i < 10; i++ {
		loop[i] = make(map[int]struct{})
	}
	for i := 1; i < 10; i++ {
		len := 1
		for {
			s := strings.Repeat(strconv.Itoa(i), len)
			x, _ := strconv.Atoi(s)
			x %= m
			if _, found := loop[i][x]; found {
				break
			}
			loop[i][x] = struct{}{}
			len++
		}
	}
	type node struct {
		idx, len int
	}
	//fmt.Println(loop)
	var ns []node
	for i := 1; i < 10; i++ {
		if _, found := loop[i][0]; found {
			ns = append(ns, node{i, len(loop[i])})
		}
	}
	if len(ns) == 0 {
		return "-1"
	}
	sort.Slice(ns, func(i, j int) bool {
		if n%ns[i].len == n%ns[j].len {
			return ns[i].idx > ns[j].idx
		}
		return n%ns[i].len < n%ns[j].len
	})
	lenAnswer := n - n%ns[0].len
	ans := strings.Repeat(strconv.Itoa(ns[0].idx), lenAnswer)
	return ans
}

func solve(n, m int) string {
	type node struct {
		idx, len int
	}
	var ns []node
	for i := 1; i <= 9; i++ {
		x := i
		for j := 1; j <= n; j++ {
			if x%m == 0 {
				ns = append(ns, node{i, j})
			}
			x = (10*x + i) % m
		}
	}
	if len(ns) == 0 {
		return "-1"
	}
	sort.Slice(ns, func(i, j int) bool {
		if ns[i].len == ns[j].len {
			return ns[i].idx > ns[j].idx
		}
		return ns[i].len > ns[j].len
	})
	ans := strings.Repeat(strconv.Itoa(ns[0].idx), ns[0].len)
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	ans := solve(n, m)
	PrintString(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
