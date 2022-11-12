package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextString()
	ans := solve(n)
	PrintString(ans)
}

func solve(n string) string {
	//ビットが立っている数
	var cnt int
	for _, v := range n {
		if v == '1' {
			cnt++
		}
	}
	if cnt == 2 {
		return "Yes"
	} else if cnt <= 1 {
		return "No"
	}
	for i := 0; i < len(n)-1; i++ {
		if n[i] == '0' && n[i+1] == '1' {
			return "No"
		}
	}
	return "Yes"
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
