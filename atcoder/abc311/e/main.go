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

	h, w, n := nextInt(), nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	ans := solve(h, w, n, a, b)
	PrintInt(ans)
}

func solve(h, w, n int, a, b []int) int {
	s := make([][]int, h)
	for i := range s {
		s[i] = make([]int, w)
	}

	isHole := make([][]bool, h)
	for i := range isHole {
		isHole[i] = make([]bool, w)
	}
	for k := 0; k < n; k++ {
		i, j := a[k], b[k]
		isHole[i][j] = true
	}

	if !isHole[0][0] {
		s[0][0] = 1
	}
	for i := 1; i < h; i++ {
		if isHole[i][0] {
			continue
		}
		s[i][0] = 1
	}
	for j := 1; j < w; j++ {
		if isHole[0][j] {
			continue
		}
		s[0][j] = 1
	}
	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			if isHole[i][j] {
				continue
			}
			s[i][j] = Min(s[i-1][j], Min(s[i][j-1], s[i-1][j-1])) + 1
		}
	}
	//for _, si := range s {
	//	fmt.Println(si)
	//}
	var ans int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans += s[i][j]
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
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
