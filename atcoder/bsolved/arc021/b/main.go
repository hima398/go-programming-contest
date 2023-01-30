package main

import (
	"bufio"
	"errors"
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

	l := nextInt()
	b := nextIntSlice(l)
	ans, err := solve(l, b)
	if err != nil {
		PrintInt(-1)
		return
	}
	PrintVertically(ans)
}

func solve(l int, b []int) ([]int, error) {
	ans := []int{0}
	for i := 0; i < l-1; i++ {
		ans = append(ans, ans[len(ans)-1]^b[i])
	}
	a0 := ans[len(ans)-1] ^ b[l-1]
	if ans[0] != a0 {
		return []int{-1}, errors.New("Impossible")
	}
	return ans, nil
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
