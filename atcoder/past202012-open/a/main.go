package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()

	o := regexp.MustCompile("ooo")
	x := regexp.MustCompile("xxx")
	if o.MatchString(s) {
		PrintString("o")
	} else if x.MatchString(s) {
		PrintString("x")
	} else {
		PrintString("draw")
	}

}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
