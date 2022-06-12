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

	n, m, q := nextInt(), nextInt(), nextInt()
	solver := make([][]int, m)
	score := make([]int, n)
	//for i := range score {
	//	score[i] = n
	//}
	var ans []int
	for k := 0; k < q; k++ {
		t := nextInt()
		switch t {
		case 1:
			nn := nextInt()
			nn--
			ans = append(ans, score[nn])
		case 2:
			nn, mm := nextInt(), nextInt()
			nn--
			mm--
			for _, v := range solver[mm] {
				score[v] = Max(score[v]-1, 0)
			}
			solver[mm] = append(solver[mm], nn)
			score[nn] += n - len(solver[mm])
		}
	}
	PrintSlice(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintSlice(a []int) {
	defer out.Flush()
	for _, v := range a {
		fmt.Fprintln(out, v)
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
