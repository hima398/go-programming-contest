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

	q := nextInt()
	t, x, k := make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t[i], x[i] = nextInt(), nextInt()
		if t[i] == 2 || t[i] == 3 {
			k[i] = nextInt()
		}
	}
	ans := solve(q, t, x, k)
	PrintVertically(ans)
}

func solve(q int, t, x, k []int) []int {
	avlTree := avltree.NewWithIntComparator()
	var ans []int
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			v, found := avlTree.Get(x[i])
			if found {
				avlTree.Put(x[i], v.(int)+1)
			} else {
				avlTree.Put(x[i], 1)
			}
		case 2:
			node, found := avlTree.Floor(x[i])
			if !found {
				ans = append(ans, -1)
				continue
			}
			var s int
			found = false
			for node != nil {
				s += node.Value.(int)
				found = s >= k[i]
				if found {
					break
				}
				node = node.Prev()
			}
			if found {
				ans = append(ans, node.Key.(int))
			} else {
				ans = append(ans, -1)
			}
		case 3:
			node, found := avlTree.Ceiling(x[i])
			if !found {
				ans = append(ans, -1)
				continue
			}
			var s int
			found = false
			for node != nil {
				s += node.Value.(int)
				found = s >= k[i]
				if found {
					break
				}
				node = node.Next()
			}
			if found {
				ans = append(ans, node.Key.(int))
			} else {
				ans = append(ans, -1)
			}
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
