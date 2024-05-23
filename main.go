package main

import (
	"fmt"
)

type Pair struct {
	Path string
	Hash string
}

type PairWithLength struct {
	Pair
	Length int
}

func main() {
	pl := PairWithLength{Pair{"aPath", "aHash"}, 100}
	fmt.Println(pl.Path, pl.Length)
}
