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
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	reverse := func(s string) string {
		r := strings.Split(s, "")
		n := len(r)
		for i := 0; i < n/2; i++ {
			j := n - 1 - i
			r[i], r[j] = r[j], r[i]
		}
		return strings.Join(r, "")
	}
	m := make(map[string]struct{})
	var ans int
	for _, si := range s {
		if _, found := m[si]; found {
			continue
		}
		if _, found := m[reverse(si)]; found {
			continue
		}
		m[si] = struct{}{}
		ans++
	}
	PrintInt(ans)
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
