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

	n := nextInt()
	var d1, d2 []int
	for i := 0; i < n; i++ {
		d1 = append(d1, nextInt())
		d2 = append(d2, nextInt())
	}
	var cnt int
	for i := 0; i < n; i++ {
		if d1[i] == d2[i] {
			cnt++
		} else {
			cnt = 0
		}
		if cnt >= 3 {
			PrintString("Yes")
			return
		}
	}
	PrintString("No")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
