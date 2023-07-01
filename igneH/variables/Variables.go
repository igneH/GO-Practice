package main

import (
	"fmt"
	"strconv"
)

var (
	I int = 68  //globally visible
	j int = 700 //only visible in the package
	k int = 419
)

func main() {
	var i int //the shorter the lifespan of the variable the shorter the name
	i = 69
	var j int = 699 //only visible in the block
	k := 420
	fmt.Println(k, i, j)
	fmt.Printf("%v, %T\n", k, k)

	var theURL string = "https://rule34.xxx" //Acronyms always upper case (for readability)
	fmt.Println(theURL)

	var u int = 18
	fmt.Printf("%v %T\n", u, u)
	var v float32
	v = float32(u) //type-conversion
	var w string
	w = strconv.Itoa(u)
	fmt.Printf("%v %T\n", v, v)
	fmt.Printf("%v, %T\n", w, w)
}
