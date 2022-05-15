package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(a, v, b, w, t int) string {
	//if v <= w {
	//	return "NO"
	//}
	if Abs(a-b) <= (v-w)*t {
		return "YES"
	} else {
		return "NO"
	}

}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, v := nextInt(), nextInt()
	b, w := nextInt(), nextInt()
	t := nextInt()
	ans := Solve(a, v, b, w, t)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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
