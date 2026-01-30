package main
import (
	"fmt"
	"math"
)

func main() {
	initializer, multiplier, iterations := GetData()
	ms := Iterate(initializer, multiplier, iterations)
	current := ms[len(ms)-1][iterations]
	fmt.Println(current) 

	// find the iteration value at a specific index.
	fmt.Print("find value at index: ")
	var idx int
	fmt.Scan(&idx)
	favIdx := ms[idx-1]
	fmt.Printf("value at index %v: %v\n", idx, favIdx[idx])

}
 

func GetData() (float64, float64, int) {
	var initializer, multiplier float64 
	fmt.Print("initializer: ")
	fmt.Scanln(&initializer)
	fmt.Print("multiplier: ")
	fmt.Scanln(&multiplier)
	var iterations int
	fmt.Print("iterations: ")
	fmt.Scanln(&iterations)
	return initializer, multiplier, iterations
}

func Iterate(init, mtpl float64, itxn int) []map[int]float64 {
	var ms []map[int]float64
	for i := range itxn {
		init = math.Floor(init * mtpl)
		ms = append(ms, map[int]float64 {
			i+1: init,
		})
	}
	return ms
}
