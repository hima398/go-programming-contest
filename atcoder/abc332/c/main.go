package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	s := nextString()

	ans := solve(n, m, s)

	Print(ans)
}

func solve(n, m int, s string) int {
	//何も予定が無い日までに必要なロゴ入りTシャツの数を求める
	f := func(s string) int {
		var c1, c2 int
		for _, r := range s {
			switch r {
			case '1':
				c1++
			case '2':
				c2++
			}
		}
		res := c2
		//食事に行く予定に足りないTシャツ分ロゴ入りTシャツを買う
		res += Max(c1-m, 0)
		return res
	}
	schedules := strings.Split(s, "0")
	var ans int
	for _, v := range schedules {
		ans = Max(ans, f(v))
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
