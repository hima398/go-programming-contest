package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(t int, h, w, d []int) []string {
	var ans []string
	for i := 0; i < t; i++ {
		var dd int
		h2, w2 := h[i]*h[i], w[i]*w[i]
		if h[i]%2 == 0 {
			if w[i]%2 == 0 {
				dd = Min(h2, w2)
			} else {
				dd = Min(h2+1, w2)
			}
		} else {
			if w[i]%2 == 0 {
				dd = Min(h2, w2+1)
			} else {
				dd = Min(h2+1, w2+1)
			}
		}
		if dd <= d[i]*d[i] {
			ans = append(ans, "N")
		} else {
			ans = append(ans, "S")
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := nextInt()
	var h, w, d []int
	for i := 0; i < t; i++ {
		h = append(h, nextInt())
		w = append(w, nextInt())
		d = append(d, nextInt())
	}
	ans := solve(t, h, w, d)
	PrintVertically(ans)

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x []string) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
