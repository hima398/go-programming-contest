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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	firstCanWin := solve(n, s)
	if firstCanWin {
		PrintString("First")
	} else {
		PrintString("Second")
	}
}

func solve(n int, s []string) bool {
	dp := make([]int, 1<<n)
	var ps []int
	for pat := 1; pat < 1<<n; pat++ {
		ps = append(ps, pat)
	}
	sort.Slice(ps, func(i, j int) bool {
		if bits.OnesCount(uint(ps[i])) == bits.OnesCount(uint(ps[j])) {
			return ps[i] < ps[j]
		}
		return bits.OnesCount(uint(ps[i])) < bits.OnesCount(uint(ps[j]))
	})
	for _, pat := range ps {
		var flags int
		for i := 0; i < n; i++ {
			//すでに文字s[i]を使用済み
			if (pat>>i)&1 == 0 {
				continue
			}
			next := pat ^ (1 << i)
			if (dp[next]>>int(s[i][len(s[i])-1]-'a'))&1 == 0 {
				flags |= 1 << int(s[i][0]-'a')
			}
		}
		dp[pat] = flags
	}
	return dp[(1<<n)-1] > 0
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
