// Алгоритм "Двух сумм" Оптимизированный алгоритм
//
// https://practicum.yandex.ru/trainer/algorithms/lesson/29fdc9f9-9476-491e-9297-596ad7d03d4b/
// https://contest.yandex.ru/contest/26365/problems/E/
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoSum(array []int, targetSum int) []int {
	var result []int
	store := make(map[int]bool)

	for _, v := range array {
		y := targetSum - v
		if ok := store[y]; ok {
			result = append(result, v, y)

			return result
		} else {
			store[v] = true
		}
	}

	return result
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	array := readArray(scanner)
	targetSum := readInt(scanner)
	result := twoSum(array, targetSum)
	if len(result) == 0 {
		fmt.Println("None")
	} else {
		fmt.Print(result[0])
		fmt.Print(" ")
		fmt.Print(result[1])
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
