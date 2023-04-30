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

	n, d := nextInt(), nextInt()
	ans := solve(n, d)
	PrintInt(ans)
}

func solve(n, d int) int {
	const p = 998244353

	p2 := make([]int, Max(n, d)+1)
	p2[0] = 1
	for i := 1; i <= Max(n, d); i++ {
		p2[i] = 2 * p2[i-1]
		p2[i] %= p
	}
	//葉の方向に深さx進める点の数
	points := func(x int) int {
		if n-x < 0 {
			return 0
		} else {
			return (p2[n-x] - 1 + p) % p
		}
	}
	//葉の方向に距離d進むと到達できる点の数
	toLeaf := func(d int) int {
		if d < 0 {
			return 0
		} else if d == 0 {
			return 1
		} else {
			return p2[d-1]
		}
	}
	var ans int
	for i := 0; i <= d; i++ {
		//左方向に深さi、右方向に深さd-i進むと到達できる点の数の積
		v := toLeaf(i) * toLeaf(d-i) % p
		v *= points(Max(i, d-i))
		v %= p
		//fmt.Println(i, d-i, toLeaf(i), toLeaf(d-i), Max(i, d-i), points(Max(i, d-i)))
		ans += 2 * v
		ans %= p
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
