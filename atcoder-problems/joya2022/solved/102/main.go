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
	buf := make([]byte, 1024*1024)
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
	PrintInt(ans)
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
	ans := 0
	isSame := false
	for i := 0; i < n; i++ {
		isDifferent := false
		for j := 0; j < m[i]; j++ {
			isDifferent = isDifferent || e[i][j] == maxE[p[i][j]] && primes[p[i][j]] == 1
		}
		if isDifferent {
			ans++
		} else if !isSame {
			ans++
			isSame = true
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
