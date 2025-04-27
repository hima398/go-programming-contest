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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	x0, y0 := nextInt(), nextInt()
	xn, yn := nextInt(), nextInt()

	ans := solve(n, x0, y0, xn, yn)

	fmt.Println(real(ans), imag(ans))
}

func solve(n, x0, y0, xn, yn int) complex128 {
	c := complex(float64(x0+xn)/2.0, float64(y0+yn)/2.0)

	ans := complex(float64(x0), float64(y0)) - c
	theta := 2 * math.Pi / float64(n)
	rotate := complex(math.Cos(theta), math.Sin(theta))
	ans *= rotate
	ans += c

	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
