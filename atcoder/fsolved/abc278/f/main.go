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
	set := (1 << n) - 1
	//
	//使っていない文字の集合iと最初文字cのとき先手必勝かどうかのフラグの集合j
	dp := make([]int, set+1)

	//残っているもじが小さい順に並び替え
	var patterns []int
	for pat := 1; pat <= set; pat++ {
		patterns = append(patterns, pat)
	}
	sort.Slice(patterns, func(i, j int) bool {
		return bits.OnesCount(uint(patterns[i])) < bits.OnesCount(uint(patterns[j]))
	})

	var words [][2]int
	for _, si := range s {
		words = append(words, [2]int{int(si[0] - 'a'), int(si[len(si)-1] - 'a')})
	}
	for _, subSet := range patterns {
		ok := 0
		for j := 0; j < n; j++ {
			//すでに文字s[j]を使用済
			if subSet>>j&1 == 0 {
				continue
			}
			nextSet := subSet ^ 1<<j
			//s[j]を使うと、相手が負ける場合
			if dp[nextSet]>>words[j][1]&1 == 0 {
				ok |= 1 << words[j][0]
			}
		}
		dp[subSet] = ok
	}
	if bits.OnesCount(uint(dp[set])) > 0 {
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
