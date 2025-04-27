package main

import (
	"bufio"
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

	n, x, k := nextInt(), nextInt(), nextInt()
	var p, u, c []int
	for i := 0; i < n; i++ {
		p = append(p, nextInt())
		u = append(u, nextInt())
		c = append(c, nextInt())
	}

	ans := solve(n, x, k, p, u, c)

	Print(ans)
}

func solve(n, x, k int, p, u, c []int) int {
	type product struct {
		p, u, c int
	}
	var products []product
	for i := 0; i < n; i++ {
		products = append(products, product{p[i], u[i], c[i]})
	}
	sort.Slice(products, func(i, j int) bool {
		return products[i].c < products[j].c
	})

	//i番目の商品まで見て、合計価格がj円、今の商品の色を購入すみかどうかの満足度
	dp := make([][][2]int, n+1)
	for i := range dp {
		dp[i] = make([][2]int, x+1)
	}

	prevColor := -1
	for i, product := range products {
		if product.c != prevColor {
			for j := 0; j <= x; j++ {
				dp[i][j][0] = Max(dp[i][j][0], dp[i][j][1])
			}
			prevColor = product.c
		}
		for j := 0; j <= x; j++ {
			for k := 0; k <= 1; k++ {
				dp[i+1][j][k] = Max(dp[i+1][j][k], dp[i][j][k])
			}
			nextTotalPrice := j + product.p
			if nextTotalPrice > x {
				continue
			}
			dp[i+1][nextTotalPrice][1] = Max(dp[i+1][nextTotalPrice][1], dp[i][j][1]+product.u)
			dp[i+1][nextTotalPrice][1] = Max(dp[i+1][nextTotalPrice][1], dp[i][j][0]+product.u+k)
		}
	}

	var ans int
	for j := 0; j <= x; j++ {
		for k := 0; k <= 1; k++ {
			ans = Max(ans, dp[n][j][k])
		}
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
