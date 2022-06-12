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

	n, k := nextInt(), nextInt()
	a, b := nextIntSlice(n), nextIntSlice(k)

	max := Max(a)

	favorites := make(map[int]struct{})
	for i, v := range a {
		if v == max {
			favorites[i+1] = struct{}{}
		}
	}
	for _, v := range b {
		if _, found := favorites[v]; found {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Max(x []int) int {
	res := x[0]
	for i := 1; i < len(x); i++ {
		if res < x[i] {
			res = x[i]
		}
	}
	return res
}
