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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var a, b, c []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		c = append(c, nextInt()-1)
	}
	ans := solve(n, m, a, b, c)
	PrintInt(ans)
}

func solve(n, m int, a, b, c []int) int {
	var masks []int
	for i := 0; i < m; i++ {
		var mask int
		mask |= 1 << a[i]
		mask |= 1 << b[i]
		mask |= 1 << c[i]
		masks = append(masks, mask)
	}

	var ans int
	for pat := 0; pat < 1<<n; pat++ {
		e1 := false
		for _, mask := range masks {
			e1 = e1 || (pat&mask == mask)
		}
		if e1 {
			continue
		}
		s := 0
		for i := 0; i < n; i++ {
			substance := pat | (1 << i)
			if substance == pat {
				continue
			}
			e2 := false
			for _, mask := range masks {
				e2 = e2 || (substance&mask == mask)
			}
			if e2 {
				s++
			}
		}
		//fmt.Println(pat, s)

		ans = Max(ans, s)

	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
