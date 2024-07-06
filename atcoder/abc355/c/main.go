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

	n, t := nextInt(), nextInt()
	a := nextIntSlice(t)

	ans, err := solve(n, t, a)

	if err != nil {
		Print(-1)
	} else {
		Print(ans)
	}
}

func solve(n, t int, a []int) (int, error) {
	row := make([]int, n)
	col := make([]int, n)
	var c1, c2 int

	for k, ai := range a {
		i, j := (ai-1)/n, (ai-1)%n
		//fmt.Println(ai, i, j)
		row[i]++
		col[j]++
		if i == j {
			c1++
		}
		if i+j == n-1 {
			c2++
		}
		if row[i] == n || col[j] == n || c1 == n || c2 == n {
			return k + 1, nil
		}
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
