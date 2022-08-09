package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(t int, n []int) []int {
	memo := map[int]bool{1: true, 3: true, 4: true, 5: true, 7: true, 8: true}
	var ans []int
	for i := 0; i < t; i++ {
		if n[i] <= 8 {
			if _, contains := memo[n[i]]; contains {
				ans = append(ans, 1)
			} else {
				ans = append(ans, 2)
			}
		} else {
			ans = append(ans, 2)
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := nextInt()
	var n []int
	for i := 0; i < t; i++ {
		n = append(n, nextInt())
	}
	ans := solve(t, n)
	PrintVertically(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
