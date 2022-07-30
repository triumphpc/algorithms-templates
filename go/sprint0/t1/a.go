// https://practicum.yandex.ru/trainer/algorithms/lesson/29fdc9f9-9476-491e-9297-596ad7d03d4b/
// Алгоритм A + B
//
package main

import (
	"fmt"
	"strconv"
)

func getSum(a int, b int) int {
	return a + b
}

func main() {
	a := readInt()
	b := readInt()
	fmt.Println(getSum(a, b))

}

func readInt() int {
	var aString string
	fmt.Scan(&aString)
	a, _ := strconv.Atoi(aString)

	return a
}
