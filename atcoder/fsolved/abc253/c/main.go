package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/emirpasic/gods/trees/redblacktree"
)

var sc = bufio.NewScanner(os.Stdin)

type query struct {
	t, x, c int
}

func solve(n int, q []query) (ans []int) {
	s := redblacktree.NewWithIntComparator()

	insert := func(x int) {
		v, found := s.Get(x)
		if found {
			s.Put(x, v.(int)+1)
		} else {
			s.Put(x, 1)
		}
	}
	erase := func(x, c int) {
		v, found := s.Get(x)
		if found {
			if v.(int) > c {
				s.Put(x, v.(int)-c)
			} else {
				//v <= c
				s.Remove(x)
			}
		}
	}

	for _, v := range q {
		switch v.t {
		case 1:
			insert(v.x)
		case 2:
			erase(v.x, v.c)
		case 3:
			max := s.Right().Key.(int)
			min := s.Left().Key.(int)
			ans = append(ans, max-min)
		}
	}

	return ans
}

func solveByMap(n int, q []query) (ans []int) {
	s := make(map[int]int)
	for _, v := range q {
		switch v.t {
		case 1:
			s[v.x]++
		case 2:
			if s[v.x] > v.c {
				s[v.x] -= v.c
			} else {
				delete(s, v.x)
			}
		case 3:
			//ここでmap[int]intの中から
			//max、minを効率よく求める必要があります。。
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	//入力、データ型を整理
	n := nextInt()
	var q []query
	for i := 0; i < n; i++ {
		t := nextInt()
		switch t {
		case 1:
			q = append(q, query{t, nextInt(), 0})
		case 2:
			q = append(q, query{t, nextInt(), nextInt()})
		case 3:
			q = append(q, query{t, 0, 0})
		}
	}

	//解答
	ans := solveByMap(n, q)
	//ans := solve(n, q)

	//出力
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
