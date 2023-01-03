package main

import (
	"bufio"
	"errors"
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

	n := nextInt()
	a := nextIntSlice(n)
	ans := solve(n, a)
	PrintInt(ans)
}

func solve(n int, a []int) int {
	offset := 2 * int(1e9)
	rbTree := redblacktree.NewWithIntComparator()
	push := func(x int) {
		if v, found := rbTree.Get(x); found {
			rbTree.Put(x, v.(int)+1)
		} else {
			rbTree.Put(x, 1)
		}
	}
	pop := func(x int) {
		if v, _ := rbTree.Get(x); v.(int) > 1 {
			rbTree.Put(x, v.(int)-1)
		} else {
			rbTree.Remove(x)
		}
	}
	read := func(x int) {
		if v, _ := rbTree.Get(x); v.(int) > 1 {
			//読み終わった巻が残る場合、十分に大きい巻数を持たせて
			//売るものとして扱う
			rbTree.Put(offset+x, v.(int)-1)
		}
		rbTree.Remove(x)
	}
	buy := func() error {
		if n <= 1 {
			return errors.New("impossible")
		}

		for i := 0; i < 2; i++ {
			k := rbTree.Right().Key.(int)
			pop(k)
		}
		return nil
	}
	find := func(x int) bool {
		_, found := rbTree.Get(x)
		return found
	}
	for _, ai := range a {
		if ai <= n {
			if find(ai) {
				push(offset + ai)
			} else {
				push(ai)
			}
		} else {
			push(ai)
		}
	}
	cur := 0
	for {
		if find(cur + 1) {
			//巻を持っているので読む
			read(cur + 1)
			n--
		} else {
			//2冊売って買う
			err := buy()
			if err != nil {
				//売れる本がない
				break
			}
			n -= 2
		}
		//fmt.Printf("rbTree: %v\n", rbTree)

		cur++
	}
	return cur
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
