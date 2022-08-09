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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	mask := 1<<n - 1
	w := 1
	for i := 0; i < n; i++ {
		w *= 2
	}
	ans := make([]float64, n+1)
	for pat := 0; pat <= mask; pat++ {
		idx := bits.OnesCount(uint(pat))
		ans[idx]++
	}
	for i := 0; i <= n; i++ {
		ans[i] /= float64(w)
	}
	PrintVertically(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x []float64) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
