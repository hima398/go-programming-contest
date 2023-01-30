package main

import (
	"bufio"
	"fmt"
	"math/big"
	"math/bits"
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

	n, m := nextInt(), nextInt()
	ans := solve(n, m)
	PrintInt(ans)
}

func solve(n, m int) int {
	ps := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73}
	mi := sort.Search(len(ps), func(i int) bool {
		return m < ps[i]
	})

	b := 1<<mi - 1
	ans := n
	bn := big.NewInt(int64(n))
	for pat := 1; pat <= b; pat++ {
		//72以下の素数の累積がオーバーフローする場合があるが制限時間内に計算できないほど大きくはならないのでbig.Intを使用
		w := big.NewInt(1)
		for i := 0; i < mi; i++ {
			if (pat>>i)&1 == 1 {
				p := big.NewInt(int64(ps[i]))
				w = w.Mul(w, p)
			}
		}
		if w.Cmp(bn) == 1 {
			continue
		}
		//fmt.Printf("%b, %d\n", pat, w)
		if bits.OnesCount(uint(pat))%2 == 0 {
			ans += n / int(w.Int64())
		} else {
			ans -= n / int(w.Int64())
		}
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
