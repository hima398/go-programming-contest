package main

import (
	"bufio"
	"errors"
	"fmt"
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

	n := nextInt()

	ans, err := solve(n)
	if err != nil {
		Print(-1)
		return
	}
	Print(ans)
}

func dividePrime(x int) map[int]int {
	res := make(map[int]int)
	for i := 2; i*i <= x; i++ {
		for x%i == 0 {
			res[i]++
			x /= i
		}
		if x == 1 {
			break
		}
	}
	if x != 1 {
		res[x]++
	}
	return res
}

func isPalindrome(s string) bool {
	n := len(s)
	ok := true
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		ok = ok && s[i] == s[j]
	}
	return ok
}

func reverseString(s string) string {
	res := strings.Split(s, "")
	n := len(res)
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		res[i], res[j] = res[j], res[i]
	}
	return strings.Join(res, "")
}

// 整数xを問題文の条件を満たす式で表現する
func buildPalindromicExpression(x int) (string, error) {
	res := strconv.Itoa(x)
	if isPalindrome(res) && !strings.Contains(res, "0") {
		return res, nil
	}
	for i := 2; i*i <= x; i++ {
		if x%i != 0 {
			continue
		}
		if strings.Contains(strconv.Itoa(i), "0") {
			continue
		}

		j, err := strconv.Atoi(reverseString(strconv.Itoa(i)))
		//問題の制約ではほぼ発生しない見込みだけど、エラーハンドリングを書いておく
		if err != nil {
			return "", errors.New("can't make reversed i")
		}
		if (x/i)%j != 0 {
			continue
		}
		mid, err := buildPalindromicExpression(x / i / j)
		if err != nil {
			continue
		}
		return strconv.Itoa(i) + "*" + mid + "*" + strconv.Itoa(j), nil
	}
	return "", errors.New("impossible")
}

func solve(n int) (string, error) {
	return buildPalindromicExpression(n)
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
