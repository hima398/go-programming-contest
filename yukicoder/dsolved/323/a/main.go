package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unsafe"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func isPrime(x int) bool {
	if x == 1 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func clone(s string) string {
	b := make([]byte, len(s))
	copy(b, s)
	return *(*string)(unsafe.Pointer(&b))
}

func solve(s string) int {
	n := len(s)
	var ans int
	//answers := make(map[int]struct{})
	for mask := 0; mask < 1<<(n-1); mask++ {
		//fmt.Printf("%b ", mask)
		cs := clone(s)
		cm := mask
		sum := 0
		for i := 0; i < len(cs)-1; i++ {
			if cm>>i&1 > 0 {
				//fmt.Printf("%s %s ", cs[:i+1], cs[i+1:])
				x, _ := strconv.Atoi(cs[:i+1])
				sum += x
				cs = cs[i+1:]
				cm >>= i + 1
				i = -1
			}
		}
		x, _ := strconv.Atoi(cs)
		sum += x
		//fmt.Println(mask, sum)
		if isPrime(sum) {
			ans++
		}
		//fmt.Println()
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	ans := solve(s)
	PrintInt(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
