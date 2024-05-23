package main

import (
	"fmt"
)

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

type PairWithLength struct {
	Pair
	Length int
}

func main() {
	p := Pair{"aPath", "aHash"}
	pl := PairWithLength{p, 100}
	fmt.Println(pl.Path, pl.Length)


	fmt.Println(pl) //this calls the String() function of the Pair struct
}
