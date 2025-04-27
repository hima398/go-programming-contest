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

	h, w, x, y := nextInt(), nextInt(), nextInt()-1, nextInt()-1
	var s []string
	for i := 0; i < h; i++ {
		s = append(s, nextString())
	}
	var ans int
	for cur := x + 1; cur < h && s[cur][y] == '.'; cur++ {
		ans++
	}
	for cur := x - 1; cur >= 0 && s[cur][y] == '.'; cur-- {
		ans++
	}
	for cur := y + 1; cur < w && s[x][cur] == '.'; cur++ {
		ans++
	}
	for cur := y - 1; cur >= 0 && s[x][cur] == '.'; cur-- {
		ans++
	}
	ans++
	Print(ans)
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
