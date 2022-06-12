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

	m, d := nextInt(), nextInt()

	var ans int
	for i := 1; i <= m; i++ {
		for j := 2; j <= d; j++ {
			d1, d2 := j%10, j/10
			if d1 <= 1 || d2 <= 1 {
				continue
			}
			if d1*d2 == i {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
