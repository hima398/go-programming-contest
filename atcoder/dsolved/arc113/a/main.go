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

	k := nextInt()

	e := make([]int, 2*int(1e5)+1)
	s := make([]int, 2*int(1e5)+1)
	for i := 1; i <= 2*int(1e5); i++ {
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				if j == i/j {
					e[i]++
				} else {
					e[i] += 2
				}
			}
		}
	}
	for i := 1; i <= 2*int(1e5); i++ {
		s[i] = e[i] + s[i-1]
	}
	var ans int
	for i := 1; i <= k; i++ {
		bc := k / i
		ans += s[bc]
		//fmt.Println(bc, s[bc])
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
