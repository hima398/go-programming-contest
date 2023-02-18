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
	var ss []int
	for s1 := 1; s1 <= 9; s1++ {
		for s3 := 0; s3 <= 9; s3++ {
			for s4 := 0; s4 <= 9; s4++ {
				for s5 := 0; s5 <= 9; s5++ {
					for s7 := 0; s7 <= 9; s7++ {
						for s8 := 0; s8 <= 9; s8++ {
							v := s7 + 10*s8 + 100*s7 + 1000*s5 + 10000*s5 + 100000*s4 + 1000000*s3 + int(1e7)*s1 + int(1e8)*s1
							ss = append(ss, v)
						}
					}
				}
			}
		}
	}
	ans := ss[n-1]
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
