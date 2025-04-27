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

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := nextInt()
	var ans []int
	for i := 0; i < t; i++ {
		//fmt.Println("# t = ", t)
		n := nextInt()
		a := nextIntSlice(2 * n)
		ans = append(ans, solve(n, a))
	}
	PrintVertically(ans)
}

func solve(n int, a []int) int {
	m := make([][2]int, n+1)
	for i := range m {
		m[i][0], m[i][1] = -1, -1
	}
	for i, ai := range a {
		if m[ai][0] >= 0 {
			m[ai][1] = i
		} else {
			m[ai][0] = i
		}
	}

	counted := make([]map[int]struct{}, n+1)
	for i := range counted {
		counted[i] = make(map[int]struct{})
	}

	var ans int
	for k := 0; k < 2*n-1; k++ {
		if a[k] == a[k+1] {
			continue
		}

		u, v := a[k], a[k+1]
		if u > v {
			u, v = v, u
		}
		if _, found := counted[u][v]; found {
			continue
		}
		if m[u][1]-m[u][0] == 1 || m[v][1]-m[v][0] == 1 {
			continue
		}
		idxes := []int{m[u][0], m[u][1], m[v][0], m[v][1]}
		sort.Ints(idxes)
		if idxes[1]-idxes[0] == 1 && idxes[3]-idxes[2] == 1 {
			//fmt.Println("# ", u, v, idxes)
			ans++
			counted[u][v] = struct{}{}
		}
	}
	return ans
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
