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

	n := nextInt()
	ans := solve(n)
	Print(ans)
}

func solve(n int) string {
	if n == 1 {
		return "0"
	}
	w := 9 * 2
	cur := 1
	digit := 1

	for cur+w < n {
		cur += w
		w *= 10
		digit += 2
	}

	offset := n - cur
	if offset > w/2 {
		offset -= w / 2
		digit++
	}
	base := 1
	for i := 0; i < Floor(digit-1, 2); i++ {
		base *= 10
	}
	half := strconv.Itoa(base + offset - 1)

	//halfでdigit桁の回文を作る
	ans := half
	for i := 0; i < len(half); i++ {
		if i == 0 && digit%2 == 1 {
			continue
		}
		j := len(half) - 1 - i
		ans += string(half[j])
	}
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

func Floor(x, y int) int {
	/*
		return x / y
	*/
	r := (x%y + y) % y
	return (x - r) / y
}
