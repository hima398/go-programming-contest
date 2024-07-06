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

	n := nextInt()
	var a, c []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		c = append(c, nextInt())
	}
	ans := solve(n, a, c)
	Print(len(ans))
	PrintHorizonaly(ans)
}

func solve(n int, a, c []int) []int {
	type card struct {
		id, power, cost int
	}
	var cards []card
	for i := 0; i < n; i++ {
		cards = append(cards, card{i, a[i], c[i]})
	}
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].power < cards[j].power
	})
	min := make([]int, n)
	min[n-1] = cards[n-1].cost
	for i := n - 2; i >= 0; i-- {
		min[i] = Min(min[i+1], cards[i].cost)
	}
	//fmt.Println(cards)
	//fmt.Println(min)
	var ans []int
	for i := 0; i < n-1; i++ {
		if cards[i].cost > min[i+1] {
			continue
		}
		ans = append(ans, cards[i].id+1)
	}
	ans = append(ans, cards[n-1].id+1)
	sort.Ints(ans)
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
