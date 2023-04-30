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

	n := nextInt()
	q := nextInt()
	t, x, y := make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		x[i] = nextInt()
		if t[i] == 1 {
			y[i] = nextInt()
		}
	}
	ans := solve(n, q, t, x, y)

	PrintVertically(ans)
}

func solve(n, q int, t, x, y []int) (ans [][]int) {
	boxes := make([][]int, 2*int(1e5)+1)
	cards := make([]map[int]struct{}, 2*int(1e5)+1)

	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			boxes[y[i]] = append(boxes[y[i]], x[i])
			if cards[x[i]] == nil {
				cards[x[i]] = make(map[int]struct{})
			}
			cards[x[i]][y[i]] = struct{}{}
		case 2:
			s := make([]int, len(boxes[x[i]]))
			copy(s, boxes[x[i]])
			ans = append(ans, s)
		case 3:
			var s []int
			for k := range cards[x[i]] {
				s = append(s, k)
			}
			ans = append(ans, s)
		}
	}
	for _, v := range ans {
		sort.Ints(v)
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func PrintVertically(x [][]int) {
	defer out.Flush()
	for _, v := range x {
		//fmt.Fprintln(out, v)
		PrintHorizonaly(v)
	}
}
