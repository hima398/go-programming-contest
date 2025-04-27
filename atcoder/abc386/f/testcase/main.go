package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/rand"
)

func main() {
	k := 20
	n := 5 * int(1e5)
	var s []string
	for i := 0; i < n; i++ {
		v := string('a' + rand.Intn(26))
		s = append(s, v)
	}
	var t []string
	for i := 0; i < n; i++ {
		v := string('a' + rand.Intn(26))
		t = append(t, v)
	}

	fmt.Println(k)
	fmt.Println(strings.Join(s, ""))
	fmt.Println(strings.Join(t, ""))
}
