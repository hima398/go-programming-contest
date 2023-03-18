package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/emirpasic/gods/trees/avltree"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	var t, x, y []int
	for i := 0; i < q; i++ {
		t = append(t, nextInt())
		x = append(x, nextInt()-1)
		y = append(y, nextInt()-1)
	}
	ans := solve(n, q, t, x, y)
	PrintHorizonaly(ans)
}

func solve(n, q int, t, x, y []int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i + 1
	}

	idxes := avltree.NewWithIntComparator()
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			a[x[i]], a[x[i]+1] = a[x[i]+1], a[x[i]]
			for j := -1; j <= 1; j++ {
				if x[i]+j < 0 || x[i]+j+1 >= n {
					continue
				}
				if a[x[i]+j] < a[x[i]+j+1] {
					idxes.Remove(x[i] + j)
				} else {
					idxes.Put(x[i]+j, 1)
				}
			}
		case 2:
			for {
				node, found := idxes.Ceiling(x[i])
				if !found {
					break
				}
				if node.Key.(int) >= y[i] {
					break
				}
				cur := node.Key.(int)
				a[cur], a[cur+1] = a[cur+1], a[cur]
				for j := -1; j <= 1; j++ {
					if cur+j < 0 || cur+j+1 >= n {
						continue
					}
					if a[cur+j] < a[cur+j+1] {
						idxes.Remove(cur + j)
					} else {
						idxes.Put(cur+j, 1)
					}
				}
			}
		}
	}
	return a
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
