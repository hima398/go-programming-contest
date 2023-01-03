package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	k, n := nextInt(), nextInt()
	var v, w []string
	for i := 0; i < n; i++ {
		v = append(v, nextString())
		w = append(w, nextString())
	}
	solve(k, n, v, w)
}

func eval(s []int, v, w []string) ([]string, error) {
	k := len(s)
	n := len(v)
	m := make(map[int]string)
	wd := make([]string, len(w))
	for i := 0; i < n; i++ {
		wd[i] = w[i]
	}

	for i := 0; i < n; i++ {
		for _, vi := range v[i] {
			idx, _ := strconv.Atoi(string(vi))
			if s[idx-1] > len(wd[i]) {
				return nil, errors.New("Impossible")
			}
			si := wd[i][:s[idx-1]]
			wd[i] = wd[i][s[idx-1]:]
			if _, found := m[idx]; found {
				if m[idx] != si {
					return nil, errors.New("Impossible")
				}
			} else {
				m[idx] = si
			}
		}
		if len(wd[i]) > 0 {
			return nil, errors.New("Impossible")
		}
	}
	//ループを抜けてた=語呂合わせに矛盾がない
	ans := make([]string, k)
	for k, v := range m {
		ans[k-1] = v
	}
	return ans, nil
}

func solve(k, n int, v, w []string) {
	var dfs func(d int, s []int)
	dfs = func(d int, s []int) {
		if d == 0 {
			ans, err := eval(s, v, w)
			if err == nil {
				PrintVertically(ans)
				os.Exit(0)
			}
			return
		}
		for i := 1; i <= 3; i++ {
			ns := make([]int, len(s))
			copy(ns, s)
			ns = append(ns, i)
			dfs(d-1, ns)
		}
	}
	dfs(k, []int{})
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

func PrintVertically(x []string) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
