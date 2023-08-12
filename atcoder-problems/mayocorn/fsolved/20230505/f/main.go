package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	ans := solve(n, k, s)
	PrintInt(ans)
}

func solve(n, k int, s []string) int {
	type cell struct {
		i, j int
	}

	var pat uint
	var ans int
	investigated := make(map[uint]struct{})
	di := []int{-1, 0, 0, 1}
	dj := []int{0, -1, 1, 0}
	var dfs func(rem int)
	dfs = func(rem int) {
		if _, found := investigated[pat]; found {
			return
		}
		investigated[pat] = struct{}{}

		if rem == 0 {
			ans++
			return
		}
		var next []cell

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				shift := n*i + j
				if s[i][j] == '#' || (pat>>shift)&1 > 0 {
					continue
				}
				var isNext bool
				//s[i][j]=='.'
				for k := 0; k < 4; k++ {
					ni, nj := i+di[k], j+dj[k]
					//nextIdx := n*ni + nj
					shift := n*ni + nj
					isNext = isNext || ((ni >= 0 && ni < n && nj >= 0 && nj < n) && ((pat >> uint(shift) & 1) > 0))
				}
				if isNext {
					next = append(next, cell{i, j})
				}
			}
		}
		for _, c := range next {
			shift := n*c.i + c.j
			pat |= 1 << shift
			dfs(rem - 1)
			pat ^= 1 << shift
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			shift := n*i + j
			if s[i][j] == '#' {
				continue
			}
			pat |= 1 << shift
			dfs(k - 1)
			pat ^= 1 << shift
		}
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
