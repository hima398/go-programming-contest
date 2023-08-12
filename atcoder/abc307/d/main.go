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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := nextString()
	ans := solve(n, s)
	PrintString(ans)
}

func solve(n int, s string) string {
	var stack [][]string
	stack = append(stack, []string{})
	cur := 0
	for _, si := range s {
		//fmt.Println(stack)
		if si == '(' {
			stack = append(stack, []string{string(si)})
			cur++
		} else if si == ')' && cur > 0 {
			stack = stack[:cur]
			cur--
		} else {
			stack[cur] = append(stack[cur], string(si))
		}
	}
	var ans []string
	for i := 0; i < len(stack); i++ {
		for _, si := range stack[i] {
			ans = append(ans, si)
		}
	}
	return strings.Join(ans, "")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
