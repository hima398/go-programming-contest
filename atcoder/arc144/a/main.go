package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int) (int, string) {
	memo := make([]int, 5)
	m := 2 * n
	for i := 4; n > 0 && i >= 1; i-- {
		d := n / i
		n -= i * d
		memo[i] = d
	}
	//fmt.Println(memo)
	var x string
	for i := 1; i <= 4; i++ {
		if memo[i] > 0 {
			x += strings.Repeat(string(byte(i)+'0'), memo[i])
		}
	}
	return m, x
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	m, x := solve(n)
	PrintInt(m)
	PrintString(x)
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
