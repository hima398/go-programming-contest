package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	ans := solve(n)
	Print(ans)
}

func solve(n int) int {
	var rep []int
	cur := 1
	for i := 0; i < 18; i++ {
		rep = append(rep, cur)
		cur = 10*cur + 1
	}
	var s []int
	for i := 0; i < len(rep); i++ {
		for j := i; j < len(rep); j++ {
			for k := j; k < len(rep); k++ {
				s = append(s, rep[i]+rep[j]+rep[k])
			}
		}
	}
	sort.Ints(s)
	ans := s[n-1]
	return ans
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
