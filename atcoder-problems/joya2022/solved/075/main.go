package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/emirpasic/gods/trees/redblacktree"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	q := nextInt()
	t, x, c := make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		switch t[i] {
		case 1:
			x[i] = nextInt()
		case 2:
			x[i], c[i] = nextInt(), nextInt()
		}
	}
	ans := solve(q, t, x, c)
	PrintVertically(ans)
}

func solve(q int, t, x, c []int) []int {
	var ans []int
	rbTree := redblacktree.NewWithIntComparator()
	for i := range t {
		switch t[i] {
		case 1:
			if v, found := rbTree.Get(x[i]); found {
				rbTree.Put(x[i], v.(int)+1)
			} else {
				rbTree.Put(x[i], 1)
			}
		case 2:
			if v, found := rbTree.Get(x[i]); found {
				if c[i] < v.(int) {
					rbTree.Put(x[i], v.(int)-c[i])
				} else {
					rbTree.Remove(x[i])
				}
			}
		case 3:
			ans = append(ans, rbTree.Right().Key.(int)-rbTree.Left().Key.(int))
		}
		//fmt.Printf("rbTree: %v\n", rbTree)
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
