package main

import (
	"bufio"
	"errors"
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

	t := nextInt()
	var ans []int
	for i := 0; i < t; i++ {
		n := nextInt()
		s := nextString()
		//v, err := solveHonestly(n, s)
		v, err := solve(n, s)
		if err != nil {
			ans = append(ans, -1)
		} else {
			ans = append(ans, v)
		}
	}
	PrintVertically(ans)
}

func solve(n int, s string) (int, error) {
	var no int
	for _, v := range s {
		if v == '1' {
			no++
		}
	}
	if no%2 == 1 {
		return -1, errors.New("Impossible")
	}
	if no == 2 && strings.Contains(s, "11") {
		if n == 2 || n == 3 {
			return -1, errors.New("Impossible")
		} else if n == 4 {
			if s == "0110" {
				return 3, nil
			} else {
				return 2, nil
			}
		} else {
			return 2, nil
		}
	}
	ans := no / 2
	return ans, nil
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
