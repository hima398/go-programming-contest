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
	var a, b []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	k := nextInt()
	var c, d []int
	for i := 0; i < k; i++ {
		c = append(c, nextInt()-1)
		d = append(d, nextInt()-1)
	}
	mask := (1 << k) - 1
	var ans int
	//0:ciにボールを置く、1:diにボールを置く
	for pat := 0; pat <= mask; pat++ {
		dishes := make([]int, n)
		for i := 0; i < k; i++ {
			if (pat>>i)&1 == 0 {
				dishes[c[i]]++
			} else {
				dishes[d[i]]++
			}
		}
		var s int
		for i := 0; i < m; i++ {
			if dishes[a[i]] > 0 && dishes[b[i]] > 0 {
				s++
			}
		}
		ans = Max(ans, s)
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
