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

	s, t := nextString(), nextString()
	if s[0] == t[0] {
		//SとTどちらも地下
		ns, nt := int(s[1]-'0'), int(t[1]-'0')
		PrintInt(Abs(nt - ns))
	} else if s[1] == t[1] {
		//SとTどちらも地上
		ns, nt := int(s[0]-'0'), int(t[0]-'0')
		PrintInt(Abs(nt - ns))
	} else {
		//地下と地上の間を行き来する
		//sが地上、tが地下になるようにする
		if s > t {
			s, t = t, s
		}
		ans := int(s[0]-'0') + int(t[1]-'0') - 1
		fmt.Println(ans)
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
