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

	lm, ld := nextInt(), nextInt()
	y, m, d := nextInt(), nextInt(), nextInt()

	if d == ld && m == lm {
		fmt.Println(y+1, 1, 1)
	} else if d == ld {
		fmt.Println(y, m+1, 1)
	} else {
		fmt.Println(y, m, d+1)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
