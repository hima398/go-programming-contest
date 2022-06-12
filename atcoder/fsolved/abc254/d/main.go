package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func solve(n int) int {
	sq := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sq[i] = i * i
	}
	//fmt.Println(sq)
	var ans int
	for i := 1; i <= n; i++ {
		// i = k * (x**2)に分解する
		k := i
		for x := 2; x*x <= k; x++ {
			for k%(x*x) == 0 {
				k /= x * x
			}
		}

		// j = k * (y**2)とすると
		// i * j = (k * x * y) ** 2

		var s int
		for y := 1; k*y*y <= n; y++ {
			s++
		}
		//s := sort.Search(n+1, func(idx int) bool {
		//	return k*sq[idx] > n
		//})
		//fmt.Println(i, k, s-1)
		ans += s
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	ans := solve(n)

	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
