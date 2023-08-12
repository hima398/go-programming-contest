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

	n, m := nextInt(), nextInt()
	var c []int
	var a [][]int
	for i := 0; i < m; i++ {
		c = append(c, nextInt())
		a = append(a, nextIntSlice(c[i]))
	}
	//fmt.Println(a)
	var ans int
	for pat := 1; pat < 1<<m; pat++ {
		x := make(map[int]struct{})
		for i := 0; i < m; i++ {
			if (pat>>i)&1 == 0 {
				continue
			}
			for _, aij := range a[i] {
				x[aij] = struct{}{}
			}
		}
		if len(x) == n {
			ans++
		}
	}
	PrintInt(ans)
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
