package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, a, b, p []int) ([][2]int, error) {
	for i := 0; i < n; i++ {
		p[i]--
	}
	type person struct {
		i, w int
	}
	for i := 0; i < n; i++ {
		if a[i] <= b[p[i]] && i != p[i] {
			return nil, errors.New("Impossible")
		}
	}
	//i：荷物, v：人
	mb := make([]int, n)
	for i := 0; i < n; i++ {
		mb[p[i]] = i
	}
	//i:人、v：荷物
	mp := make([]int, n)
	for i, v := range mb {
		mp[v] = i
	}
	var people []person
	for i := 0; i < n; i++ {
		if i != p[i] {
			people = append(people, person{i, a[i]})
		}
	}
	sort.Slice(people, func(i, j int) bool {
		if people[i].w == people[j].w {
			return i < j
		}
		return people[i].w < people[j].w
	})
	var ans [][2]int
	for len(people) > 0 {
		pi, pj := people[0].i, mb[people[0].i]
		if pi != pj {
			ans = append(ans, [2]int{pi + 1, pj + 1})
		}
		bi, bj := mp[pi], mp[pj]
		mb[bi], mb[bj] = pj, pi
		mp[pi], mp[pj] = bj, bi
		people = people[1:]
	}
	return ans, nil
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(n)
	p := nextIntSlice(n)
	ans, err := solve(n, a, b, p)
	if err != nil {
		PrintInt(-1)
		return
	}
	PrintInt(len(ans))
	PrintVertically(ans)
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

func PrintVertically(x [][2]int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v[0], v[1])
	}
}
