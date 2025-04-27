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
	h := nextIntSlice(n)

	ans := solve(n, h)

	PrintHorizonaly(ans)
}

func solve(n int, h []int) []int {
	var ans []int
	stack := stack.New[int]()
	for i := n - 1; i >= 0; i-- {
		ans = append(ans, stack.Size())
		for stack.Size() > 0 && h[i] > stack.Top() {
			stack.Pop()
		}
		stack.Push(h[i])
	}
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		ans[i], ans[j] = ans[j], ans[i]
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
