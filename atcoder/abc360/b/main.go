package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s, t := nextString(), nextString()
	for i := 1; i < len(s); i++ {
		var u []string
		for j := 0; j < len(s); j += i {
			u = append(u, s[j:Min(j+i, len(s))])
		}
		//fmt.Println(u)
		for j := 0; j < i; j++ {
			var v []string
			for k := 0; k < len(u); k++ {
				if j < len(u[k]) {
					v = append(v, string(u[k][j]))
				}
			}
			//fmt.Println(v)
			if strings.Join(v, "") == t {
				Print("Yes")
				return
			}
		}

	}
	Print("No")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
