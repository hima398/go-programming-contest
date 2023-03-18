package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
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
		x1, x2, x3 := nextInt(), nextInt(), nextInt()
		//v, err := solveHonestly(x1, x2, x3)
		v, err := solve(x1, x2, x3)
		if err != nil {
			ans = append(ans, -1)
		} else {
			ans = append(ans, v)
		}
	}
	PrintVertically(ans)
}

func solve(x1, x2, x3 int) (int, error) {
	if x1%2 != x2%2 || x2%2 != x3%2 || x3%2 != x1%2 {
		return -1, errors.New("Impossible")
	}
	//xの値がすべて同じ
	if x1 == x2 && x2 == x3 && x3 == x1 {
		return 0, nil
	}
	x := []int{x1, x2, x3}
	sort.Ints(x)
	d1 := Min(Abs(x[1]-x[0]), Abs(x[2]-x[1]))
	x[0] += 7 * d1 / 2
	x[1] += 5 * d1 / 2
	x[2] += 3 * d1 / 2

	ans := d1 / 2

	d2 := Abs(x[2]-x[1]) + Abs(x[2]-x[0])
	if d2%6 != 0 {
		return -1, errors.New("Impossible")
	}
	//x[0]<=x[1]<=x[2]
	switch {
	case x[1] == x[2]:
		d2 := (Abs(x[1] - x[0]))
		ans += d2 / 3
		return ans, nil
	case x[0] == x[1]:
		d2 := Abs(x[0] - x[2])
		ans += d2 / 3
		return ans, nil
	}
	return -1, errors.New("Impossible")
}

func solveHonestly(x1, x2, x3 int) (int, error) {
	if x1%2 != x2%2 || x2%2 != x3%2 || x3%2 != x1%2 {
		return -1, errors.New("Impossible")
	}
	//xの値がすべて同じ
	if x1 == x2 && x2 == x3 && x3 == x1 {
		return 0, nil
	}
	x := []int{x1, x2, x3}
	sort.Ints(x)
	d := Abs(x[2]-x[1]) + Abs(x[2]-x[0])
	if d%6 != 0 {
		return -1, errors.New("Impossible")
	}
	var ans int
	for !(x[0] == x[1] && x[1] == x[2] && x[2] == x[0]) {
		sort.Ints(x)
		//fmt.Println(x1, x2, x3, x[0], x[1], x[2])
		x[0] += 7
		x[1] += 5
		x[2] += 3
		ans++
	}
	return ans, nil
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
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
