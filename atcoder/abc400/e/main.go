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

	q := nextInt()
	var a []int
	for i := 0; i < q; i++ {
		a = append(a, nextInt())
	}

	ans := solve(q, a)
	for _, v := range ans {
		Print(v)
	}
	/*
		sieve := NewSieveOfEratosthenes(int(1e6))
		var ps []int
		for i := 2; i <= int(1e6); i++ {
			if sieve.IsPrime(i) && i*i <= int(1e6) {
				ps = append(ps, i*i)
			}
		}
		fmt.Println("ps = ", ps)
		ns := make(map[int]struct{})
		for i := 0; i < len(ps); i++ {
			for j := i + 1; j < len(ps); j++ {
				a := ps[i]
				for k := 1; a*ps[j] <= int(1e12); k++ {
					b := ps[j]
					for l := 1; a*b <= int(1e12); l++ {
						ns[a*b] = struct{}{}
						b *= ps[j]
					}
					a *= ps[i]
				}
			}
		}
		var numbers []int
		for k := range ns {
			numbers = append(numbers, k)
		}
		sort.Ints(numbers)
		fmt.Println("numbers = ", numbers)
		//var ans []int
		for _, ai := range a {
			idx := sort.Search(len(numbers), func(i int) bool {
				return ai < numbers[i]
			})
			//ans = append(ans, numbers[i])
			Print(numbers[idx-1])
		}

		//sieve.init(int(1e6))
	*/
}

func solve(q int, a []int) []int {
	sieve := NewSieveOfEratosthenes(int(1e6))
	//fmt.Println(sieve.kinds)
	var numbers []int
	for i := 0; i <= int(1e6); i++ {
		if sieve.kinds[i] == 2 {
			numbers = append(numbers, i*i)
		}
	}
	//fmt.Println(len(numbers))
	sort.Ints(numbers)
	var ans []int
	for _, ai := range a {
		idx := sort.Search(len(numbers), func(i int) bool {
			return ai < numbers[i]
		})
		ans = append(ans, numbers[idx-1])
	}

	return ans
}

// エラトステネスの篩
type SieveOfEratosthenes struct {
	n int
	//isNotPrime []bool
	kinds []int
}

func New(n int) *SieveOfEratosthenes {
	return NewSieveOfEratosthenes(n)
}

func NewSieveOfEratosthenes(n int) *SieveOfEratosthenes {
	sieve := new(SieveOfEratosthenes)
	sieve.init(n)
	return sieve
}

func (sieve *SieveOfEratosthenes) init(n int) {
	sieve.n = n + 1
	//sieve.isNotPrime = make([]bool, sieve.n)
	//sieve.isNotPrime[0] = true
	//sieve.isNotPrime[1] = true
	sieve.kinds = make([]int, sieve.n)
	sieve.kinds[0] = -1
	sieve.kinds[1] = -1
	//for j := 4; j < sieve.n; j += 2 {
	//	sieve.isNotPrime[j] = true
	//}
	for i := 2; i < sieve.n; i += 2 {
		sieve.kinds[i]++
	}
	for i := 3; i < sieve.n; i += 2 {
		if sieve.kinds[i] > 0 {
			continue
		}
		for j := i; j < sieve.n; j += i {
			sieve.kinds[j]++
		}
	}
}

/*
func (sieve *SieveOfEratosthenes) IsPrime(x int) bool {
	return sieve.kinds[x] == 0
	//return !sieve.isNotPrime[x]
}
*/

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
