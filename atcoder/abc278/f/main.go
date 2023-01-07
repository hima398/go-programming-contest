package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	ans := solve(n, s)
	PrintString(ans)
}

func solve(n int, s []string) string {
	mask := (1 << n) - 1
	//
	//使っていない文字の集合sと最後の文字cのとき先手必勝かどうかのフラグの集合t
	dp := make([]int, mask+1)

	var patterns []int
	for pat := 1; pat <= mask; pat++ {
		patterns = append(patterns, pat)
	}
	sort.Slice(patterns, func(i, j int) bool {
		return bits.OnesCount(uint(patterns[i])) < bits.OnesCount(uint(patterns[j]))
	})
	var words [][2]int
	for _, si := range s {
		words = append(words, [2]int{int(si[0] - 'a'), int(si[len(si)-1] - 'a')})
	}
	for _, pat := range patterns {
		for i := 0; i < n; i++ {
			//すでに文字s[i]を使用済
			if pat>>i&1 == 0 {
				continue
			}
			nextS := pat ^ 1<<i
			c := int(s[pat][len(s)-1] - 'a')
			ok = ok || dp[nextS][c]
		}
	}
	if bits.OnesCount(uint(dp[mask])) > 0 {
		return "First"
	} else {
		return "Second"
	}

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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
