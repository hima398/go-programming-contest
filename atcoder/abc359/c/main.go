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

	sx, sy := nextInt(), nextInt()
	tx, ty := nextInt(), nextInt()

	ans := solve(sx, sy, tx, ty)

	Print(ans)
}

func solve(sx, sy, tx, ty int) int {
	//sx<=txになるようにする
	if sx > tx {
		sx, tx = tx, sx
		sy, ty = ty, sy
	}
	//y方向には必ず差分の通行料が発生
	y := Abs(ty - sy)

	sx = Min(sx+y, tx)

	//x方向
	var x int
	if ty%2 == 0 {
		x = Abs(tx/2 - sx/2)
	} else {
		x = Abs((tx-1)/2 - (sx-1)/2)
	}

	ans := x + y
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
