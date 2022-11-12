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

func solveHonestly(n int) int {
	//if n == 2 {
	//	return 1
	//}
	var ans int
	for k := 2; k <= n; k++ {
		nn := n
		for nn%k == 0 && nn > 1 {
			nn /= k
		}
		m := nn % k
		if m == 1 || m-k == 1 {
			//fmt.Println("k = ", k)
			ans++
		}
	}
	return ans
}

func divide(x int) []int {
	m := make(map[int]struct{})
	for i := 1; i*i <= x; i++ {
		if x%i == 0 {
			m[i] = struct{}{}
			m[x/i] = struct{}{}
		}
	}
	var res []int
	for k := range m {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}

func isContains(s []int, x int) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

func solve(n int) int {
	//if n == 2 {
	//	return 2
	//}
	d1 := divide(n)
	//1を除外
	d1 = d1[1:]
	//fmt.Println(d1)
	d2 := divide(n - 1)
	//1を除外
	d2 = d2[1:]

	ans := len(d2)
	f := func(n, k int) int {
		for n >= k && n%k == 0 {
			n = n / k
		}
		return n % k
	}
	for i := range d1 {
		if f(n, d1[i]) == 1 && !isContains(d2, d1[i]) {
			ans++
		}
	}
	//n自信分を加算
	//ans++
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	//ans := solveHonestly(n)
	ans := solve(n)
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
