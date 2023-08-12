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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	var c [][]int
	for i := 0; i < 3; i++ {
		c = append(c, nextIntSlice(3))
	}
	ok := solve(c)
	if ok {
		PrintString("Yes")
	} else {
		PrintString("No")
	}
}

func solve(c [][]int) bool {
	a := []int{0}
	b := []int{c[0][0]}

	a = append(a, c[1][0]-b[0])
	a = append(a, c[2][0]-b[0])
	b = append(b, c[0][1]-a[0])
	b = append(b, c[0][2]-a[0])

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if a[i]+b[j] != c[i][j] {
				return false
			}
		}
	}
	return true
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
