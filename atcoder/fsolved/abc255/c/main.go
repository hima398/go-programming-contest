package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(x, a, d, n int) int {
	if d == 0 {
		return Abs(x - a)
	}
	l := a + d*(n-1)
	if d < 0 {
		a, l = l, a
		d *= -1
	}
	if x <= a {
		return a - x
	} else if l <= x {
		return x - l
	}
	// a < x < l
	//il, ir := 0, n-1
	//for ir-il > 1 {
	//	mid := (il + ir) / 2
	//	if a+d*mid < x {
	//		il = mid
	//	} else {
	//		ir = mid
	//	}
	//}
	//fmt.Println(il, ir)
	il := (x - a) / d
	ir := il + 1

	return Min(Abs(a+il*d-x), Abs(a+ir*d-x))
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	x, a, d, n := nextInt(), nextInt(), nextInt(), nextInt()
	ans := solve(x, a, d, n)
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

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
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

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
