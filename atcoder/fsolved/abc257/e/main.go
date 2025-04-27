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

func solve(n int, c []int) string {
	//使うお金はmのみ
	m := make(map[int]int)
	for i := 0; i < 9; i++ {
		m[c[i]] = i
	}
	type item struct {
		idx, cost int
	}
	uc := make([]item, 9)
	for k, v := range m {
		uc[v] = item{v, k}
	}
	sort.Slice(uc, func(i, j int) bool {
		return uc[i].cost < uc[j].cost
	})

	var mx int
	use := make([]int, 9)
	for _, v := range uc {
		if v.cost > 0 {
			mx = v.idx
			use[v.idx] = n / v.cost
			n = n - v.cost*use[v.idx]
			break
		}
	}
	offset := c[mx]
	for i := 0; i < 9; i++ {
		c[i] -= offset
	}
	//fmt.Println(n)
	//fmt.Println(c)
	for i := 8; i >= 0; i-- {
		if i > mx && c[i] > 0 {
			diff := n / c[i]
			use[i] += diff
			n -= diff * c[i]
			use[mx] -= diff
		}
		if n == 0 {
			break
		}
	}
	var ans []string
	for i := 8; i >= 0; i-- {
		if use[i] > 0 {
			ans = append(ans, strings.Repeat(strconv.Itoa(i+1), use[i]))
		}
	}
	return strings.Join(ans, "")
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	c := nextIntSlice(9)
	ans := solve(n, c)
	PrintString(ans)
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
