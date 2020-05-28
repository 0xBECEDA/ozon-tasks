package main
import (
	"fmt"
	"os"
	"math/rand"

)

// количесто элементов в изначальных слайсах (в финальном в 2 раза больше)

var amount_elts = 100

// перемещаем элементы с n по m из одного слайса в другой
func fill_the_slice (slice_to []int, slice_from []int, begin_to int,
	begin_from int, end_to int) {

	var j = begin_from

	fmt.Println("fill_the_slice\n")

	for i := begin_to; i < end_to; i++ {
		slice_to[i] = slice_from[j]
		j++
	}
}

// перемешивает элементы в слайсе
func mix_slice (slice []int) {

	for i := 0; i < amount_elts * 2 + 1; i++ {

		var indx1 int = rand.Intn(amount_elts * 2 + 1)
		var indx2 int = rand.Intn(amount_elts * 2 + 1)

		var elt1 int = slice[indx1]
		var elt2 int = slice[indx2]

		slice[indx1] = elt2
		slice[indx2] = elt1
	}
}

// создает финальный слайс с рандомными числами
func make_rand_numbers(slice []int) []int {

	for i := 0; i < amount_elts; i++ {
		rnd1 := rand.Intn(100000)
		slice[i] = rnd1
	}

	//fmt.Println("финальный слайс\n")

	// создаем финальный слайс
	var final_slice []int
    final_slice = make([]int, amount_elts * 2 + 1, amount_elts * 2 +1)

	// заполняем его элементами из исходного слайса (дважды)
	fill_the_slice(final_slice, slice, 0, 0, amount_elts)
	fill_the_slice(final_slice, slice, amount_elts, 0, amount_elts * 2)

	// добавляем один уникальный элемент
	final_slice[amount_elts * 2] = rand.Intn(1000000000)

	fmt.Println(final_slice[amount_elts * 2])
	// fmt.Println(final_slice[amount_elts * 2], "\n")

	// перемешиваем финальный слайс
	mix_slice(final_slice)
	return final_slice
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
	for i := 0; i < amount_elts * 2; i++ {
		fmt.Fprintln(file, finale_slice[i], "\n")

	}
}
