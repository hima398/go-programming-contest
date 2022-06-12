package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, r, n := nextInt(), nextInt(), nextInt()
	//if r == 1 {
	//	fmt.Println(a)
	//}
	ans := a
	for i := 1; i < n; i++ {
		ans *= r
		if ans > int(1e9) {
			fmt.Println("large")
			return
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
