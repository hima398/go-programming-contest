package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	t := nextString()
	ans := solve(s, t)
	for _, ok := range ans {
		if ok {
			Print("Yes")
		} else {
			Print("No")
		}
	}
}

func solve(s, t string) []bool {
	//sの末尾len(t)文字とtを比較
	offset := len(s) - len(t)
	//右側に幾つ異なる文字があるか
	sum := make([]int, len(t))
	for j := 0; j < len(t); j++ {
		i := offset + j
		ok := s[i] == '?' || t[j] == '?' || s[i] == t[j]
		if !ok {
			sum[j]++
		}
	}
	for i := len(t) - 1; i >= 1; i-- {
		sum[i-1] += sum[i]
	}
	var ans []bool
	ans = append(ans, sum[0] == 0)
	sum = append(sum, 0)
	sum = sum[1:]
	ok := true
	for i := 0; i < len(t); i++ {
		ok = ok && (s[i] == '?' || t[i] == '?' || s[i] == t[i])
		ans = append(ans, ok && sum[i] == 0)
	}
	return ans
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
