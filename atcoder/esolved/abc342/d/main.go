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
	const maxA = 2 * int(1e5)
	var d []int
	for i := 2; i*i <= maxA; i++ {
		d = append(d, i*i)
	}
	b := make([]int, maxA+1)
	for _, ai := range a {
		if ai == 0 {
			b[0]++
		} else {
			bi := ai
			for _, v := range d {
				for bi%v == 0 {
					bi /= v
				}
			}
			b[bi]++
		}
	}
	var ans int
	for i := 0; i < b[0]; i++ {
		ans += n - 1 - i
	}
	for i := 1; i <= maxA; i++ {
		ans += b[i] * (b[i] - 1) / 2
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
