package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/liyue201/gostl/ds/stack"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	ans := solve(s)
	Print(ans)
}

func flip(c byte) string {
	if 'A' <= c && c <= 'Z' {
		return string(c - 'A' + 'a')
	}
	return string(c - 'a' + 'A')
}

func solve(s string) string {
	ss := strings.Split(s, "")

	m := make(map[int]int)
	stk := stack.New[int]()
	var depth int
	for i, si := range s {
		switch si {
		case '(':
			stk.Push(i)
			depth++
		case ')':
			//fmt.Printf("stk: %v\n", stk)
			j := stk.Pop()
			m[i] = j
			m[j] = i
			depth--
		default:
			if depth%2 == 1 {
				ss[i] = flip(s[i])
			}
		}
	}
	s = strings.Join(ss, "")

	var ans []string
	var dfs func(l, r, d int)
	dfs = func(l, r, d int) {
		if d == 0 {
			for l <= r {
				if s[l] == '(' {
					//res = append(res, dfs(l+1, m[l]-1, 1))
					dfs(l+1, m[l]-1, 1)
					l = m[l]
				} else {
					ans = append(ans, string(s[l]))
				}
				l++
			}
		} else { // d == 1
			for l <= r {
				if s[r] == ')' {
					//res = append(res, dfs(m[r]+1, r-1, 0))
					dfs(m[r]+1, r-1, 0)
					r = m[r]
				} else {
					ans = append(ans, string(s[r]))
				}
				r--
			}
		}
	}
	dfs(0, len(s)-1, 0)
	return strings.Join(ans, "")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
