package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/liyue201/gostl/ds/stack"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, d := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(m)

	ans, err := solve(n, m, d, a, b)
	if err != nil {
		PrintInt(-1)
	} else {
		PrintInt(ans)
	}
}

func solve(n, m, d int, a, b []int) (int, error) {
	sort.Ints(a)
	sa := stack.New[int]()
	for _, ai := range a {
		sa.Push(ai)
	}
	sort.Ints(b)
	sb := stack.New[int]()
	for _, bi := range b {
		sb.Push(bi)
	}

	for !sa.Empty() && !sb.Empty() {
		x, y := sa.Pop(), sb.Pop()
		if Abs(x-y) <= d {
			return x + y, nil
		}
		if x < y {
			sa.Push(x)
		} else {
			sb.Push(y)
		}
		//fmt.Println(sa, sb)
	}

	return -1, errors.New("Impossible")
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
