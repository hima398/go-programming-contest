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

	s := nextIntSlice(8)
	for i := 0; i < 7; i++ {
		if s[i] > s[i+1] {
			PrintString("No")
			return
		}
	}
	for _, si := range s {
		if si < 100 || si > 675 {
			PrintString("No")
			return
		}
	}
	for _, si := range s {
		if si%25 != 0 {
			PrintString("No")
			return
		}
	}
	PrintString("Yes")
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
