package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	t, x, y := make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		if t[i] == 3 {
			x[i] = nextInt() - 1
		} else {
			x[i], y[i] = nextInt()-1, nextInt()-1
		}
	}
	ans := solve(n, q, t, x, y)
	PrintVertically(ans)
}

func solve(n, q int, t, x, y []int) [][]int {
	prev := make([]int, n)
	next := make([]int, n)
	for i := range next {
		prev[i] = i
		next[i] = i
	}
	var ans [][]int
	for i := range t {
		switch t[i] {
		case 1:
			next[x[i]] = y[i]
			prev[y[i]] = x[i]
		case 2:
			next[x[i]] = x[i]
			prev[y[i]] = y[i]
		case 3:
			cur := x[i]
			for prev[cur] != cur {
				cur = prev[cur]
			}

			v := []int{cur + 1}
			for next[cur] != cur {
				cur = next[cur]
				v = append(v, cur+1)
			}
			ans = append(ans, v)
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d %d", len(x), x[0])
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
