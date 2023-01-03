package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	ans := solve(s)
	PrintInt(ans)
}

func solve(s string) int {
	m := make(map[string]struct{})
	for k := 1; k <= 3; k++ {
		for i := 0; i < len(s)-k+1; i++ {
			t := strings.Split(s[i:i+k], "")
			//fmt.Println(t)
			for pat := 0; pat < 1<<k; pat++ {
				ct := make([]string, len(t))
				copy(ct, t)
				for j := 0; j < k; j++ {
					//fmt.Println(k, i, pat, j)
					if pat>>j&1 == 1 {
						ct[j] = "."
					}
				}
				m[strings.Join(ct, "")] = struct{}{}
			}
		}
	}
	//fmt.Println(m)
	ans := len(m)
	return ans
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
