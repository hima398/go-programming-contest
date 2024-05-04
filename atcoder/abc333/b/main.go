package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	t := nextString()

	d1 := (int(s[0]-'A') - int(s[1]-'A') + 5) % 5
	d2 := (int(t[0]-'A') - int(t[1]-'A') + 5) % 5
	//fmt.Println(d1, d2)
	switch d1 {
	case 1:
	case 4:
		if d2 == 1 || d2 == 4 {
			Print("Yes")
			return
		}
	case 2:
	case 3:
		if d2 == 2 || d2 == 3 {
			Print("Yes")
			return
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
