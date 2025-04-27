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

	x, y, err := solve(n)

	if err != nil {
		Print(-1)
	} else {
		fmt.Println(x, y)
	}
}

func solve(n int) (int, int, error) {
	for d := 1; d*d*d <= n; d++ {
		if n%d != 0 {
			continue
		}
		m := n / d
		y, err := quadraticEquation(3, 3*d, d*d-m)
		if err == nil && y > 0 {
			return y + d, y, nil
		}
	}
	return -1, -1, errors.New("not found")
}

func quadraticEquation(a, b, c int) (int, error) {
	l, r := 0, int(1e9)+1
	for r-l > 1 {
		//fmt.Println(l, r)
		mid := (l + r) / 2
		if a*mid*mid+b*mid+c <= 0 {
			l = mid
		} else {
			r = mid
		}
	}
	if a*l*l+b*l+c == 0 {
		return l, nil
	} else {
		return -1, errors.New("not found")
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
