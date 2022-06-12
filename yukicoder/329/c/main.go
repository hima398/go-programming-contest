package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(t int, testCase []int) []int {
	var ans []int
	f := make([]int, int(1e6)+1)
	f[3] = 4
	for i := 3; i < int(1e6); i++ {
		f[i+1] += f[i] + 2*(i*i-i)
	}
	for _, n := range testCase {
		ans = append(ans, f[n])
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := nextInt()
	var testCase []int
	for i := 0; i < t; i++ {
		testCase = append(testCase, nextInt())
	}

	ans := solve(t, testCase)
	PrintVertically(ans)
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
