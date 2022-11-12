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

	n, m := nextInt(), nextInt()
	s := nextString()
	ans, err := solveGreedy(n, m, s)
	if err != nil {
		PrintInt(-1)
		return
	}
	PrintHorizonaly(ans)
}

func solveGreedy(n, m int, s string) ([]int, error) {
	idx := n
	var rev []int
	for idx > 0 {
		next := Max(idx-m, 0)
		//fmt.Println(idx, next)
		for s[next] == '1' {
			next++
		}
		if next == idx {
			return nil, errors.New("Impossible")
		}
		rev = append(rev, idx-next)
		idx = next
	}
	var ans []int
	for i := len(rev) - 1; i >= 0; i-- {
		ans = append(ans, rev[i])
	}
	return ans, nil
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
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
