// Алгоритм "Скользящее среднее" или оптимизированный "Метод двух указателей"
// https://practicum.yandex.ru/trainer/algorithms/lesson/9536d99f-2148-47c2-aa33-f9dcda7384e6/
// https://contest.yandex.ru/contest/26365/problems/C/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func movingAverage(array []int, windowSize int) []float64 {
	// optimization realization
	l := len(array)
	s := l - windowSize + 1
	result := make([]float64, s)

	current := array[0:windowSize]
	currentSum := sumSlice(current)

	result[0] = float64(currentSum) / float64(windowSize)

	for i := 0; i <= l-windowSize-1; i++ {
		currentSum = currentSum - array[i]
		currentSum += array[i+windowSize]
		avg := float64(currentSum) / float64(windowSize)
		result[i+1] = avg
	}

	return result
}

func sumSlice(array []int) int {
	res := 0
	for _, v := range array {
		res += v
	}

	return res
}

func movingAverageBad(array []int, windowSize int) []float64 {
	// bad realization
	l := len(array)
	s := l - windowSize + 1
	result := make([]float64, s)

	for i := 0; i <= l-windowSize; i++ {
		end := i + windowSize
		current := 0

		active := array[i:end]
		for _, v := range active {
			current += v
		}

		avg := float64(current) / float64(windowSize)
		result[i] = avg
	}

	return result
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	array := readArray(scanner)
	windowSize := readInt(scanner)
	printArray(movingAverage(array, windowSize))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(arr []float64) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.FormatFloat(arr[i], 'f', 8, 64))
		writer.WriteString(" ")
	}
	writer.Flush()
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
