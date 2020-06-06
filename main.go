package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	var r bool
	flag.BoolVar(&r, "r", false, "reverse")
	var n int
	flag.IntVar(&n, "n", 1, "number of shift")
	flag.Parse()

	if flag.NArg() < 1 {
		scn := bufio.NewScanner(os.Stdin)
		for scn.Scan() {
			if err := scn.Err(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
		run(scn.Text(), n, r)
	} else {
		run((flag.Args())[0], n, r)
	}

	os.Exit(0)
}

func run(chars string, num int, reverse bool) {
	cs := shift(chars, num, reverse)
	fmt.Println(cs)
}
