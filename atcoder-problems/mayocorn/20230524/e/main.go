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
	a := nextIntSlice(n)
	b, c, err := solve(n, a)
	if err != nil {
		PrintString("No")
		return
	}
	PrintString("Yes")
	fmt.Fprintf(out, "%d ", len(b))
	PrintHorizonaly(b)
	fmt.Fprintf(out, "%d ", len(c))
	PrintHorizonaly(c)
}

func solve(n int, a []int) ([]int, []int, error) {
	const p = 200
	nn := Min(n, 8)
	mask := (1 << nn) - 1

	b := make([][]int, p)
	for pat := 1; pat <= mask; pat++ {
		sum := 0
		var c []int
		for j := 0; j < nn; j++ {
			if (pat>>j)&1 > 0 {
				c = append(c, j+1)
				sum += a[j]
				sum %= p
			}
		}
		if len(b[sum]) > 0 {
			return b[sum], c, nil
		} else {
			b[sum] = c
		}
	}
	return nil, nil, errors.New("Impossible")
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

func PrintString(x string) {
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
