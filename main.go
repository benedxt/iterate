package main
import (
	"fmt"
	"math"
)

func GetData() (float64, float64, int) {
	var initializer float64 
	fmt.Print("initializer = ")
	fmt.Scanln(&initializer)

	var multiplier float64 
	fmt.Print("multiplier = ")
	fmt.Scanln(&multiplier)

	var iterations int
	fmt.Print("max rolls = ")
	fmt.Scanln(&iterations)

	return initializer, multiplier, iterations
}

func Iterate(in, mu float64, it int) []map[int]float64 { // may also return float64
	var Ws []map[int]float64

	for i := range it {
		in = math.Floor(in * mu)
		Ws = append(Ws, map[int]float64 {
			i+1: in,
		})
	}
	// return Ws[len(Ws)-1][len(Ws)] //value of last key of slice.
	return Ws
}

func main() {
	amount, odd, iterations := GetData()
	Ws := Iterate(amount, odd, iterations)
	fmt.Println(Ws)
}
 

