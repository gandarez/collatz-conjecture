package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/gandarez/collatz-conjecture/internal/algorithm"
)

func main() {
	flag.Parse()

	input, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Printf("faield to parse %q as integer", flag.Arg(0))
	}

	results := algorithm.Calculate(input)

	fmt.Printf("it took %d iterations to complete\n", len(results))
	fmt.Printf("INPUT: %d\n", input)
	fmt.Printf("OUTPUT: [%s]\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(results)), ","), "[]"))
}
