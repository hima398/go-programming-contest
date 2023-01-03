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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	t := make([]int, q)
	a := make([]int, q)
	b := make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		switch t[i] {
		case 1:
			a[i] = nextInt() - 1
			b[i] = nextInt() - 1
		case 2, 3:
			a[i] = nextInt() - 1
		}
	}
	ans := solve(n, q, t, a, b)
	PrintVertically(ans)
}

func solve(n, q int, t, a, b []int) []string {
	e := make([][]bool, n)
	for i := 0; i < n; i++ {
		e[i] = make([]bool, n)
	}
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			//fmt.Println(a[i], b[i])
			e[a[i]][b[i]] = true
		case 2:
			next := make([]bool, n)
			copy(next, e[a[i]])
			for j := 0; j < n; j++ {
				if e[j][a[i]] && j != a[i] {
					//fmt.Println(a[i], j)
					next[j] = true
				}
			}
			e[a[i]] = next
		case 3:
			next := make([]bool, n)
			copy(next, e[a[i]])
			for j := 0; j < n; j++ {
				if e[a[i]][j] {
					for k := 0; k < n; k++ {
						if e[j][k] && k != a[i] {
							next[k] = true
						}
					}
				}
			}
			e[a[i]] = next
		}
	}
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		var buf []string
		for j := 0; j < n; j++ {
			if e[i][j] {
				buf = append(buf, "Y")
			} else {
				buf = append(buf, "N")
			}
		}
		ans[i] = strings.Join(buf, "")
	}
	return ans
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

func PrintVertically(x []string) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
