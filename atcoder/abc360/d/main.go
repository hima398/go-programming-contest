package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, t := nextInt(), nextInt()
	s := nextString()
	x := nextIntSlice(n)

	ans := solve(n, t, s, x)

	Print(ans)
}

func solve(n, t int, s string, x []int) int {
	var neg []int //負の方を向いているアリ
	var pos []int //正の方を向いているアリ
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			neg = append(neg, x[i])
		} else {
			pos = append(pos, x[i])
		}
	}
	sort.Ints(neg)
	sort.Ints(pos)
	//fmt.Println("pos = ", pos)
	//fmt.Println("neg = ", neg)

	var ans int
	for _, ant := range pos {
		for len(neg) > 0 && ant > neg[0] {
			neg = neg[1:]
		}
		idx := sort.Search(len(neg), func(i int) bool {
			return 2*t < Abs(neg[i]-ant)
		})
		ans += idx
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
