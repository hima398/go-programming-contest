package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	q := make([]string, 4)
	q[0] = "AE"
	q[1] = "B"
	q[2] = "C"
	q[3] = "D"

	k := nextInt()
	for i := 0; i < k; i++ {
		c := q[i%4][0]
		q[i%4] = q[i%4][1:]
		q[(i+1)%4] += string(c)
	}
	PrintVertically(q)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x []string) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
