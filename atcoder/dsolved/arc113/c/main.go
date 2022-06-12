package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()

	n := len(s)
	var ans int
	m := make(map[byte]int)
	for i := n - 1; i > 0; i-- {
		if s[i-1] == s[i] {
			//変更する文字の数
			nr := n - i - 1
			ans += nr - m[s[i]]
			m = make(map[byte]int)
			m[s[i]] = nr + 2
			i--
		} else {
			m[s[i]]++
		}
	}
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
