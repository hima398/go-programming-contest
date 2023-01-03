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

	x, k := nextInt(), nextInt()
	ans := x
	w1, w2 := 10, 1
	for i := 0; i < k; i++ {
		y := ans % w1
		next := -y
		if y/w2 >= 5 {
			next += w1
		}
		ans += next
		w1 *= 10
		w2 *= 10
	}
	PrintInt(ans)
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
