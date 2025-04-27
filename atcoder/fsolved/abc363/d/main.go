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

	for {
		//if n < cur+w { //ここのイコールがなくてもテストケースが通るがn=19, 199などの時に嘘解法になる
		if n <= cur+w {
			break
		}
		cur += w
		w *= 10

		digit += 2
	}
	idx := n - cur

	if idx > w/2 {
		idx -= w / 2
		digit++
	}
	w2 := 1
	for i := 0; i < Floor(digit-1, 2); i++ {
		w2 *= 10
	}
	//idxでdigit2桁の回文を作る
	t := strconv.Itoa(w2 + idx - 1)

	ans := "" + t
	for i := 0; i < len(t); i++ {
		if i == 0 && digit%2 == 1 {
			continue
		}
		//fmt.Println(i, string(t[i]))
		ans += string(t[len(t)-1-i])
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Floor(x, y int) int {
	/*
		return x / y
	*/
	r := (x%y + y) % y
	return (x - r) / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
