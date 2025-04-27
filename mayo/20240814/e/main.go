package main

import (
	"bufio"
	"errors"
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
	r := nextString()
	c := nextString()

	ans, err := solve(n, r, c)

	if err != nil {
		Print("No")
	} else {
		Print("Yes")
		for _, v := range ans {
			Print(v)
		}
	}
}

func check(n int, s []string, r, c string) bool {
	for i := 0; i < n; i++ {
		var m [256]int
		for j := 0; j < n; j++ {
			m[s[i][j]]++
		}
		if m['A'] != 1 || m['B'] != 1 || m['C'] != 1 || m['.'] != n-3 {
			return false
		}
	}
	for j := 0; j < n; j++ {
		var m [256]int
		for i := 0; i < n; i++ {
			m[s[i][j]]++
		}
		if m['A'] != 1 || m['B'] != 1 || m['C'] != 1 || m['.'] != n-3 {
			return false
		}
	}

	var sr, sc string
	//var sc string
	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			if s[i][j] != '.' {
				sc += string(s[i][j])
				break
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] != '.' {
				sr += string(s[i][j])
				break
			}
		}
	}

	return sr == r && sc == c
}

func solve(n int, r, c string) ([]string, error) {
	var masks []int
	for mask := 0; mask < 1<<n; mask++ {
		if bits.OnesCount(uint(mask)) == 3 {
			masks = append(masks, mask)
		}
	}
	sort.Ints(masks)
	var patterns [][]int
	const dot = '.' - 'A'
	pat := []int{0, 1, 2}
	m := len(pat)
	for i := 0; i < n-m; i++ {
		pat = append(pat, dot)
	}
	sort.Ints(pat)
	for {
		var v []int
		v = append(v, pat...)
		patterns = append(patterns, v)
		if !NextPermutation(sort.IntSlice(pat)) {
			break
		}
	}

	var dfs func(cur int, s []string) ([]string, error)
	dfs = func(cur int, s []string) ([]string, error) {
		if cur == n {
			if check(n, s, r, c) {
				return s, nil
			} else {
				return nil, errors.New("Impossible")
			}
		}
		for _, pat := range patterns {
			var ri byte
			for j := 0; j < n; j++ {
				if pat[j] != dot {
					ri = 'A' + byte(pat[j])
					break
				}
			}
			if ri != r[cur] {
				continue
			}

			var t string
			for i := 0; i < n; i++ {
				t += string('A' + byte(pat[i]))
			}
			var next []string
			next = append(next, s...)
			next = append(next, t)
			res, err := dfs(cur+1, next)
			if err == nil {
				return res, nil
			}
		}
		//}
		return nil, errors.New("Impossible")
	}

	return dfs(0, []string{})
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

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
