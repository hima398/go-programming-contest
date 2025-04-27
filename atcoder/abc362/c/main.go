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

	n := nextInt()
	var l, r []int
	for i := 0; i < n; i++ {
		l = append(l, nextInt())
		r = append(r, nextInt())
	}
	ans, err := solve(n, l, r)
	if err != nil {
		Print("No")
		return
	}
	Print("Yes")
	PrintHorizonaly(ans)
}

func solve(n int, l, r []int) ([]int, error) {
	var res []int
	var min int
	for _, li := range l {
		res = append(res, li)
		min += li
	}
	if min > 0 {
		return nil, errors.New("impossible")
	}
	for i := 0; i < n; i++ {
		diff := Min(-min, r[i]-l[i])
		res[i] += diff
		min += diff
	}
	if min == 0 {
		return res, nil
	} else {
		return nil, errors.New("impossible")
	}
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
