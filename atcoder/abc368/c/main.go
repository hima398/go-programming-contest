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
	h := nextIntSlice(n)

	ans := solve(n, h)

	Print(ans)
}

func solve(n int, h []int) int {
	var t int
	simulateAttack := func(t, h int) int {
		var res int
		for i := t % 3; i < 3; i++ {
			if (i+1)%3 == 0 {
				h -= 3
			} else {
				h--
			}
			res++
			if h <= 0 {
				return res
			}
		}
		res += (h / 5) * 3
		h %= 5
		if h <= 0 {
			return res
		}
		for i := 0; i < 3; i++ {
			if (i+1)%3 == 0 {
				h -= 3
			} else {
				h--
			}
			res++
			if h <= 0 {
				return res
			}
		}
		return res
	}

	for _, hi := range h {
		t += simulateAttack(t, hi)
	}
	return t
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
