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

func solveRedBlackTree(q int, t, x []int) []int {
	rbTree := redblacktree.NewWithIntComparator()
	var ans []int
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			//下記のコメントアウトのように実装する手もあるが
			//すでに置かれているカードは追加されない制約があるのでシンプルに書いておく
			rbTree.Put(x[i], 1)
			//if v, found := rbTree.Get(x[i]); found {
			//	rbTree.Put(x[i], v.(int)+1)
			//} else {
			//	rbTree.Put(x[i], 1)
			//}
		case 2:
			//下記のコメントアウトのように実装する手もあるが
			//置かれていないカードが除去されることはない制約があるのでシンプルに書いておく
			rbTree.Remove(x[i])
			//if v, found := rbTree.Get(x[i]); found {
			//	rbTree.Put(x[i], v.(int)-1)
			//} else {
			//	rbTree.Remove(x[i])
			//}
		case 3:
			if rbTree.Empty() {
				ans = append(ans, -1)
			} else {
				node, found := rbTree.Ceiling(x[i])
				if found {
					ans = append(ans, node.Key.(int))
				} else {
					ans = append(ans, -1)
				}
			}
		}
	}
	return ans
}

func solve(q int, t, x []int) []int {
	var ans []int
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	q := nextInt()
	var t, x []int
	for i := 0; i < q; i++ {
		t = append(t, nextInt())
		x = append(x, nextInt())
	}
	ans := solveRedBlackTree(q, t, x)
	PrintVertically(ans)
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
