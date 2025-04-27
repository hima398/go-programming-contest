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

	n, t := nextInt(), nextString()
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}

	ans := solve(n, t, s)

	Print(ans)
}

func reverseString(s string) string {
	r := strings.Split(s, "")
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - 1 - i
		r[i], r[j] = r[j], r[i]
	}
	return strings.Join(r, "")
}

// 文字列sに文字列tのうち連続しない部分文字を左から何文字持つか計算する
func subString(s, t string) int {
	var idx int
	for i := range s {
		if s[i] == t[idx] {
			idx++
		}
		if idx == len(t) {
			return idx
		}
	}
	return idx
}

func subStringRev(s, t string) int {
	idx := len(t) - 1
	var res int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == t[idx] {
			idx--
			res++
		}
		if res == len(t) {
			return len(t)
		}
	}
	return res
}

func solve(n int, t string, s []string) int {
	l, r := make([]int, len(t)+1), make([]int, len(t)+1)
	//s[i]の中に、tの部分文字列を左から、右からそれぞれいくつ持つか計算する
	for _, si := range s {
		li, ri := subString(si, t), subStringRev(si, t)
		l[li]++
		r[ri]++
	}
	//後処理のために累積和
	for i := len(t); i > 0; i-- {
		r[i-1] += r[i]
	}
	var ans int
	for i := 0; i <= len(t); i++ {
		j := len(t) - i
		ans += l[i] * r[j]
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
