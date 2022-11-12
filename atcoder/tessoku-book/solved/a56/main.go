package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, q int, s string, a, b, c, d []int) []string {
	const p = 2147483647
	t := make([]int, n+1)
	for i := 1; i <= n; i++ {
		t[i] = int(s[i-1]-'a') + 1
	}
	pow := make([]int, n+1)
	pow[0] = 1
	for i := 1; i <= n; i++ {
		pow[i] = 100 * pow[i-1] % p
	}
	h := make([]int, n+1)
	for i := 1; i <= n; i++ {
		h[i] = (100*h[i-1] + t[i]) % p
	}
	computeHash := func(l, r int) int {
		res := h[r+1] - (h[l] * pow[r-l+1] % p)
		if res < 0 {
			res += p
		}
		return res
	}
	var ans []string
	for i := 0; i < q; i++ {
		h1 := computeHash(a[i], b[i])
		h2 := computeHash(c[i], d[i])
		if h1 == h2 {
			ans = append(ans, "Yes")
		} else {
			ans = append(ans, "No")
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	s := nextString()
	var a, b, c, d []int
	for i := 0; i < q; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		c = append(c, nextInt()-1)
		d = append(d, nextInt()-1)
	}
	ans := solve(n, q, s, a, b, c, d)
	PrintVertically(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintVertically(x []string) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
