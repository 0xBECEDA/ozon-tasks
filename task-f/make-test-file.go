package main
import (
	"fmt"
	"os"
	"math/rand"

)

// количесто элементов в изначальных слайсах (в финальном в 2 раза больше)

var amount_elts = 20000


// перемешивает элементы в слайсе
func mix_slice (slice []int) {

	for i := 0; i < len(slice); i++ {

		var indx1 int = rand.Intn(len(slice))
		var indx2 int = rand.Intn(len(slice))

		var elt1 int = slice[indx1]
		var elt2 int = slice[indx2]

		slice[indx1] = elt2
		slice[indx2] = elt1
	}
}

// создает финальный слайс с рандомными числами
func make_rand_numbers(slice []int) []int {

	for i := 0; i < len(slice); i++ {
		rnd1 := rand.Intn(999999)
		slice[i] = rnd1
	}

	mix_slice(slice)
	return slice
}

func main () {

	file, err := os.Create("input.txt")// создаем файл

	fmt.Println("file", file)
	if err != nil {                          // если возникла ошибка
		fmt.Println("Unable to create file:", err)
		os.Exit(1)                          // выходим из программы
	}

	// создаем слайс
	var slice []int
	slice = make([]int, amount_elts, amount_elts)

	var finale_slice []int = make_rand_numbers(slice)

	// печатаем все элементы из финального слайса в файл
	for i := 0; i < len(slice); i++ {
		fmt.Fprint(file, finale_slice[i], " ")

	}
}
