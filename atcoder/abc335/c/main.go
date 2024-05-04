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
	t := make([]int, q)
	c := make([]string, q)
	p := make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		switch t[i] {
		case 1:
			c[i] = nextString()
		case 2:
			p[i] = nextInt()
		}
	}
	ans := solve(n, q, t, c, p)
	for _, v := range ans {
		PrintHorizonaly(v)
	}
}

func solve(n, q int, t []int, c []string, p []int) [][]int {
	var dragon [][]int
	for i := n; i > 0; i-- {
		dragon = append(dragon, []int{i, 0})
	}

	var ans [][]int
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			next := make([]int, 2)
			copy(next, dragon[n-1])
			switch c[i] {
			case "R":
				next[0]++
			case "L":
				next[0]--
			case "U":
				next[1]++
			case "D":
				next[1]--
			}
			dragon = append(dragon, next)
			dragon = dragon[1:]
		case 2:
			idx := n - (p[i] - 1) - 1
			ans = append(ans, dragon[idx])
		}
		//fmt.Println(i, dragon)
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
