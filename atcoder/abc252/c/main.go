package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func solve(n int, s []string) int {
	// m[i][j]=数字iがリールの先頭からj(0<=j<10)番目ある個数
	var m [10][10]int

	for i := 0; i < n; i++ {
		for j := 0; j < 10; j++ {
			x := int(s[i][j] - '0')
			m[x][j]++
		}
	}

	//十分に大きな数を初期値にしておく
	ans := 10 * n
	//数字kで揃える場合にかかる時間を求めていく
	for i := 0; i < 10; i++ {
		t := 0
		for j := 0; j < 10; j++ {
			if m[i][j] > 0 {
				t = Max(t, j+(m[i][j]-1)*10)
			}
		}
		//fmt.Println(k, t)
		ans = Min(ans, t)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	// 入力
	n := nextInt()
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}

	// 解答
	ans := solve(n, s)

	// 出力
	fmt.Println(ans)
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
