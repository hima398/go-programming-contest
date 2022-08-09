package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, p := nextInt(), nextInt()
	pp, qq := float64(p)/100.0, float64(100-p)/100.0
	var ans float64
	for i := 0; i <= n; i++ {
		v := float64(i) * float64(Combination(n, i)) * math.Pow(pp, float64(i)) * math.Pow(qq, float64(n-i))
		//fmt.Println(i, v)
		ans += v
	}
	PrintFloat64(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Combination(N, K int) int {
	if K == 0 {
		return 1
	}
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}
