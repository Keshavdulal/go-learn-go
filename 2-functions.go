package main

// avoid this way of importing multiple packages
// import "fmt"
// import "math"

// multiple packages import
// import ("fmt"
// 		"math")

// importing packages and specific subpackages
import ("fmt"
		"math"
		"math/rand")

func main() {
	fmt.Println("Using math package.\nThe square root of 4 is ",math.Sqrt(4),"\n")
	foo()
}

func foo(){
	fmt.Println("Using specific subpackage.\nA random no: ",rand.Intn(100))
}
