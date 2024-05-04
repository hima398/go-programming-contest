package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	x := nextIntSlice(q)

	ans := solve(n, q, x)

	PrintHorizonaly(ans)
}
func solve(n, q int, x []int) []int {
	//数列A
	a := make([]int, n)
	//集合S
	s := make(map[int]struct{})
	//xiが集合Sから削除されたクエリを記録する
	deleted := make([]int, n+1)
	//i番目(1-indexed)のクエリを処理した時点での|S|累積和
	sum := make([]int, q+1)
	for i, xi := range x {
		if _, found := s[xi]; found { //集合Sにxiが含まれる
			a[xi-1] += sum[i] - sum[deleted[xi]]
			delete(s, xi)
			deleted[xi] = i
		} else { //集合Sにxiが含まれない
			deleted[xi] = i
			s[xi] = struct{}{}
		}
		sum[i+1] = len(s) + sum[i]
	}
	for j := range s {
		a[j-1] += sum[q] - sum[deleted[j]]
	}
	return a
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
