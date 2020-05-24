package main
import (
	"fmt"
	"os"
	"time"
)

func main() {
	t0 := time.Now()
	var n1 float64
	var n2 float64
	fmt.Fscan(os.Stdin, &n1)
	fmt.Fscan(os.Stdin, &n2)
	var result float64 = n1 + n2
	fmt.Fprintln(os.Stdout, result)
	t1 := time.Now();
	fmt.Printf("Elapsed time: %v ", t1.Sub(t0))

	fmt.Fprintln(os.Stdout, result - n2)

}
