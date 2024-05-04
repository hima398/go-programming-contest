package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	n, m := nextInt(), nextInt()
	p := nextIntSlice(n)
	l := nextIntSlice(m)
	d := nextIntSlice(m)

	ans := solve(n, m, p, l, d)

	Print(ans)
}

func solve(n, m int, p, l, d []int) int {
	type coupon struct {
		l, d int
	}

	tree := avltree.NewWithIntComparator()
	for _, pi := range p {
		v, found := tree.Get(pi)
		if found {
			tree.Put(pi, v.(int)+1)
		} else {
			tree.Put(pi, 1)
		}
	}

	sort.Ints(p)

	var cs []coupon
	for i := 0; i < m; i++ {
		cs = append(cs, coupon{l[i], d[i]})
	}
	sort.Slice(cs, func(i, j int) bool {
		if cs[i].d == cs[j].d {
			return cs[i].l > cs[j].l
		}
		return cs[i].d > cs[j].d
	})
	var ans int
	for _, pi := range p {
		ans += pi
	}
	for _, c := range cs {
		node, found := tree.Ceiling(c.l)
		if !found {
			continue
		}
		ans -= c.d
		if node.Value.(int) == 1 {
			tree.Remove(node.Key)
		} else {
			tree.Put(node.Key, node.Value.(int)-1)
		}
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
