package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Kilogram float64
type Pound float64

const c = 2.20462

func (k Kilogram) KtoP() Pound {
	return Pound(k * c)
}
func (p Pound) PtoK() Kilogram {
	return Kilogram(p / c)
}

func main() {
	ReadInput()
}

func ReadInput() {
	var w float64
	i := flag.Float64("w", 0, "weight")
	flag.Parse()
	if *i == 0 {
		fmt.Print("Input weight: ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		temp := sc.Text()
		w, _ = strconv.ParseFloat(temp, 64)
	} else {
		w = *i
	}
	k := Kilogram(w)
	p := Pound(w)
	fmt.Println(k.KtoP(), p.PtoK())
}
