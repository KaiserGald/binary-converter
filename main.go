package main

import (
	"bytes"
	"flag"
	"fmt"
	"math"
	"strconv"

	"github.com/GoesToEleven/golang-web-dev/000_temp/36_packages/stringutil"
)

func main() {
	binFlag := flag.Bool("b", false, "a boolean value")
	decFlag := flag.Bool("d", false, "a boolean value")
	flag.Parse()

	args := flag.Args()

	if *binFlag {
		binarytodecimal(args[0])
	} else if *decFlag {
		decimaltobase(func(s string) int {
			n, _ := strconv.Atoi(s)
			return n
		}(args[0]))

	}
}

// Algorithms found and adapted from
// http://www.cs.trincoll.edu/~ram/cpsc110/inclass/conversions.html

func binarytodecimal(num string) {
	var s float64
	n := len(num) - 1
	for i := n; i >= 0; i-- {
		s += float64(num[n-i]-'0') * math.Pow(2, float64(i))
	}
	fmt.Printf("%s in decimal is %v\n", num, s)
}

func decimaltobase(num int) {
	var r int
	tmp := num
	var buffer bytes.Buffer
	for num > 0 {
		r = num % 2
		t := strconv.Itoa(r)
		buffer.WriteString(t)
		num /= 2
	}
	fmt.Printf("%v in binary is %s\n", tmp, stringutil.Reverse(buffer.String()))
}
