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

	n := nextInt()
	a := nextIntSlice(n)

	ans := solve(n, a)

	Print(ans)
}

func solve(n int, a []int) int {
	var d []int
	for i := 0; i < n-1; i++ {
		d = append(d, a[i+1]-a[i])
	}
	e := RunLengthEncoding(d)

	ans := n
	for _, ei := range e {
		ans += (ei.length * (ei.length + 1)) / 2
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

type RunLengthNode struct {
	value  int
	length int
}

func RunLengthEncoding(arr []int) []RunLengthNode {
	if len(arr) == 0 {
		return nil
	}

	var res []RunLengthNode
	for _, v := range arr {
		if len(res) == 0 || res[len(res)-1].value != v {
			res = append(res, RunLengthNode{v, 1})
		} else {
			res[len(res)-1].length++
		}
	}
	return res
}
