package main

import (
	"bufio"
	"fmt"
	"os"
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

	n := nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	isExist := solve(n, a, b)
	if isExist {
		Print("Yes")
	} else {
		Print("No")
	}
}

func solve(n int, a, b []int) bool {
	t := make([]int, 2*n)
	for i := 0; i < n; i++ {
		t[a[i]-1] = i
		t[b[i]-1] = i
	}
	s := stack.New[int]()
	for _, ti := range t {
		if s.Empty() || s.Top() != ti {
			s.Push(ti)
		} else {
			s.Pop()
		}
	}
	return !s.Empty()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
