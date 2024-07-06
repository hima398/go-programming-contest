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

	a, b := nextInt(), nextInt()
	if a == b {
		Print(-1)
	} else {
		IsCulprit := []bool{true, true, true}
		IsCulprit[a-1] = false
		IsCulprit[b-1] = false
		for i, ok := range IsCulprit {
			if ok {
				Print(i + 1)
				return
			}
		}
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
