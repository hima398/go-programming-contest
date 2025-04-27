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
	s := nextString()

	ans := solve(n, s)

	Print(ans)
}

func solve(n int, s string) int {
	//1の左に0がいくつかるか
	var d []int
	var t int
	for _, si := range s {
		switch si {
		case '0':
			t++
		case '1':
			d = append(d, t)
			t = 0
		}
	}
	d[0] = 0
	//fmt.Println(d)
	sum := make([]int, len(d))
	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] + d[i]
	}
	//fmt.Println(sum)
	var ans int
	for _, si := range sum {
		ans += si
	}
	candidate := ans
	for i := 1; i < len(d); i++ {
		l, r := i, len(d)-i
		candidate = candidate + l*d[i] - r*d[i]
		ans = Min(ans, candidate)
	}

	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
