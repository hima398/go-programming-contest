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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var a []string
	for i := 0; i < n; i++ {
		a = append(a, nextString())
	}
	isCorrect := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if a[i][j] == 'D' {
				isCorrect = isCorrect && a[j][i] == 'D'
			} else if a[i][j] == 'W' {
				isCorrect = isCorrect && a[j][i] == 'L'
			} else {
				isCorrect = isCorrect && a[j][i] == 'W'
			}
		}
	}
	if isCorrect {
		PrintString("correct")
	} else {
		PrintString("incorrect")
	}
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
