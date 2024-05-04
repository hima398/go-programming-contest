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

	k := nextInt()
	t := nextString()
	a := nextString()

	ans := solve(k, t, a)

	Print(ans)
}

func solve(k int, t, a string) float64 {
	//持っているカード
	var takahashi, aoki [10]int
	for i := 0; i < 4; i++ {
		takahashi[t[i]-'0']++
		aoki[a[i]-'0']++
	}

	//残っているカード
	var rem [10]int
	for i := range rem {
		rem[i] = k
	}
	for i := 0; i < 4; i++ {
		rem[int(t[i]-'0')]--
		rem[int(a[i]-'0')]--
	}

	//スコア計算で使う関数
	pow := func(x, y int) int {
		res := 1
		for i := 0; i < y; i++ {
			res *= x
		}
		return res
	}
	//スコアの計算
	computeScore := func(hand [10]int, x int) int {
		var t [10]int
		for i := range hand {
			t[i] = hand[i]
		}
		t[x]++
		//fmt.Println(t)
		var res int
		for i := 1; i <= 9; i++ {
			res += i * pow(10, t[i])
		}
		return res
	}

	var cnt int
	for i := 1; i <= 9; i++ {
		if rem[i] == 0 {
			continue
		}
		for j := 1; j <= 9; j++ {
			if i == j {
				if rem[i] < 2 {
					continue
				}
				if computeScore(takahashi, i) <= computeScore(aoki, i) {
					continue
				}
				cnt += rem[i] * (rem[i] - 1)
			} else {
				if rem[j] == 0 {
					continue
				}
				if computeScore(takahashi, i) <= computeScore(aoki, j) {
					continue
				}
				cnt += rem[i] * rem[j]
			}
		}
	}

	ans := float64(cnt) / float64(9*k-8) / float64(9*k-9)
	return ans
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
