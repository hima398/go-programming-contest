package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
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

	h, w, k := nextInt(), nextInt(), nextInt()
	var s []string
	for i := 0; i < h; i++ {
		s = append(s, nextString())
	}
	ans, err := solve(h, w, k, s)
	if err != nil {
		Print(-1)
	} else {
		Print(ans)
	}
}

// "o"と"."の文字列から、k個連続する"o"を作るのに"."を"o"に変える最小の手数を求める
func computeMinimumOperation(s string, k int) (int, error) {
	if len(s) < k {
		return -1, errors.New("Impossible")
	}
	var cnt int
	for i := 0; i < k; i++ {
		if s[i] == '.' {
			cnt++
		}
	}
	res := cnt
	for r := k; r < len(s); r++ {
		l := r - k
		if s[l] == '.' {
			cnt--
		}
		if s[r] == '.' {
			cnt++
		}
		res = Min(res, cnt)
	}
	return res, nil
}

func solve(h, w, k int, s []string) (int, error) {
	const INF = math.MaxInt
	ans := INF
	for _, si := range s {
		for _, ss := range strings.Split(si, "x") {
			v, err := computeMinimumOperation(ss, k)
			if err != nil {
				continue
			}
			ans = Min(ans, v)
		}
	}
	//sを90度回転
	t := make([][]string, w)
	for j := range t {
		t[j] = make([]string, h)
	}
	for j := 0; j < w; j++ {
		for i := 0; i < h; i++ {
			ni, nj := w-1-j, i
			t[ni][nj] = string(s[i][j])
		}
	}
	var sd []string
	for _, row := range t {
		sd = append(sd, strings.Join(row, ""))
	}

	for _, sdi := range sd {
		for _, ss := range strings.Split(sdi, "x") {
			v, err := computeMinimumOperation(ss, k)
			if err != nil {
				continue
			}
			ans = Min(ans, v)
		}
	}

	if ans == INF {
		return -1, errors.New("Impossible")
	}

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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
