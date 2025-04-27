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

	s := nextIntSlice(5)

	ans := solve(s)

	for _, v := range ans {
		Print(v)
	}
}

func solve(s []int) []string {
	t := "ABCDE"
	type participant struct {
		score int
		name  string
	}

	var ps []participant
	for pat := 1; pat < 1<<5; pat++ {
		var score int
		var name string
		for i := 0; i < 5; i++ {
			if (pat>>i)&1 > 0 {
				score += s[i]
				name += string(t[i])
			}
		}
		ps = append(ps, participant{score, name})
	}

	sort.Slice(ps, func(i, j int) bool {
		if ps[i].score == ps[j].score {
			return ps[i].name < ps[j].name
		}
		return ps[i].score > ps[j].score
	})

	var ans []string
	for _, p := range ps {
		ans = append(ans, p.name)
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
