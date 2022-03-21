package main

import (
	"concat_string/concat"
	"concat_string/util"
	"fmt"
	"time"
)

type AssertConcatFunc func(int, bool)

func main() {
	assertConcat(10000)
}

func assertConcat(n int) {
	list := []AssertConcatFunc{concat.UseConcat, concat.UseStringBuilder, concat.UseBytesBuffer}

	for _, f := range list {
		since := time.Now()
		f(n, false)
		fmt.Printf("Computation took: %v\n", time.Since(since))
		util.Brk()
	}

}
