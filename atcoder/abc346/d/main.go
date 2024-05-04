package main

import (
	"bufio"
	"fmt"
	"math"
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
	c := nextIntSlice(n)

	ans := solve(n, s, c)

	Print(ans)
}

func solve(n int, s string, c []int) int {
	const INF = 1 << 60
	l, r := make([][2]int, n+1), make([][2]int, n+1)
	l[0][0] = 0
	l[0][1] = 0
	for i := 0; i < n; i++ {
		l[i+1][0] = l[i][1]
		l[i+1][1] = l[i][0]
		//siが'0'なら末尾を'1'にするために、
		//siが'1'ならば末尾を'0'にするためにコストを払う
		w := int(s[i] - '0')
		l[i+1][w^1] += c[i]
	}
	//fmt.Println(l)
	r[n][0] = 0
	r[n][1] = 0
	for i := n - 1; i >= 0; i-- {
		r[i][0] = r[i+1][1]
		r[i][1] = r[i+1][0]
		w := int(s[i] - '0')
		r[i][w^1] += c[i]
	}
	//fmt.Println(r)
	ans := math.MaxInt - 1
	for i := 1; i <= n-1; i++ {
		ans = Min(ans, Min(l[i][0]+r[i][0], l[i][1]+r[i][1]))
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
