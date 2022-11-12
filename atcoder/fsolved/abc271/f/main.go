package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, a [][]int) int {
	max := 1<<(n-1) - 1
	var halfPatterns []int
	for pat := 0; pat <= max; pat++ {
		halfPatterns = append(halfPatterns, pat)
	}
	//fmt.Println(len(halfPatterns))
	//fmt.Println(max, halfPatterns)
	low := make([][]int, n)
	//fmt.Println(max, patterns)
	for _, pat := range halfPatterns {
		//low
		i, j := 0, 0
		x := a[i][j]
		for k := 0; k < n-1; k++ {
			if pat>>k&1 == 0 {
				j++
			} else {
				i++
			}
			x ^= a[i][j]
		}
		//後で残り半分と結合する時a[i][j]が2重にカウントされるので
		//ここで打ち消しておく
		x ^= a[i][j]
		idx := bits.OnesCount(uint(pat))
		low[idx] = append(low[idx], x)
	}
	high := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		high[i] = make(map[int]int)
	}
	for _, pat := range halfPatterns {
		i, j := n-1, n-1
		x := a[i][j]
		for k := 0; k < n-1; k++ {
			if pat>>k&1 == 0 {
				j--
			} else {
				i--
			}
			x ^= a[i][j]
		}
		idx := bits.OnesCount(uint(pat))
		high[idx][x]++
	}
	//fmt.Println(high)
	//fmt.Println(low)
	var ans int
	for i := 0; i < n; i++ {
		j := n - i - 1
		for _, vl := range low[i] {
			ans += high[j][vl]
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = append(a[i], nextIntSlice(n)...)
	}
	ans := solve(n, a)
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
