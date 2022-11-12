package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, s string) int {
	b := []byte(s)
	m := make(map[[32]byte]struct{})
	for i := 0; i < n-1; i++ {
		b1, b2 := b[i], b[i+1]
		b[i], b[i+1] = 'B', 'B'
		k := sha256.Sum256(b)
		m[k] = struct{}{}
		b[i], b[i+1] = b1, b2
	}
	//fmt.Println(m)
	return len(m)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := nextString()
	ans := solve(n, s)
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
