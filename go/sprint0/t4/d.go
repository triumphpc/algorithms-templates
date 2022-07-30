// Алгоритм "Двух сумм" - наивный алгоритм
// https://practicum.yandex.ru/trainer/algorithms/lesson/99be9f47-9445-4d66-97e8-74d28a2f809d/
// https://contest.yandex.ru/contest/26365/problems/D/

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

	l := len(array)

	// Наивный алгоритм - реализация
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if array[i]+array[j] == targetSum {
				result = append(result, array[i])
				result = append(result, array[j])

				return result
			}
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
