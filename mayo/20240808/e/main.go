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

	n := nextInt()
	var m []int
	p, e := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		m = append(m, nextInt())
		for j := 0; j < m[i]; j++ {
			p[i] = append(p[i], nextInt())
			e[i] = append(e[i], nextInt())
		}
	}

	ans := solve(n, m, p, e)

	Print(ans)
}

func solve(n int, m []int, p, e [][]int) int {
	maxE := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < m[i]; j++ {
			maxE[p[i][j]] = Max(maxE[p[i][j]], e[i][j])
		}
	}
	primes := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < m[i]; j++ {
			if e[i][j] == maxE[p[i][j]] {
				primes[p[i][j]]++
			}
		}
	}
	var ans int
	var offset int
	for i := 0; i < n; i++ {
		var isDifferentLcm bool
		for j := 0; j < m[i]; j++ {
			isDifferentLcm = isDifferentLcm || (e[i][j] == maxE[p[i][j]] && primes[p[i][j]] == 1)
		}
		if isDifferentLcm {
			ans++
		} else if offset == 0 {
			//値を書き換えないaのLCMと同じケース
			offset++
		}
	}
	ans += offset
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
