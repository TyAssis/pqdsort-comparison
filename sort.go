package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

const (
	runs = 10
	size = 10000
)

var inputTypes = []string{
	"random",
	"ascending",
	"descending",
	"asc-mid-desc",
	"desc-mid-asc",
}

func quickSortHoare(arr []int, low, high int) {
	if low < high {
		p := hoarePartition(arr, low, high)
		quickSortHoare(arr, low, p)
		quickSortHoare(arr, p+1, high)
	}
}

func hoarePartition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low - 1
	j := high + 1

	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}
		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func generateSlice(kind string, size int) []int {
	s := make([]int, size)
	switch kind {
	case "random":
		for i := range s {
			s[i] = rand.Intn(100000)
		}
	case "ascending":
		for i := range s {
			s[i] = i
		}
	case "descending":
		for i := range s {
			s[i] = size - i
		}
	case "asc-mid-desc":
		mid := size / 2
		for i := 0; i < mid; i++ {
			s[i] = i
		}
		for i := mid; i < size; i++ {
			s[i] = size - i
		}
	case "desc-mid-asc":
		mid := size / 2
		for i := 0; i < mid; i++ {
			s[i] = size - i
		}
		for i := mid; i < size; i++ {
			s[i] = i - mid
		}
	}
	return s
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Benchmark de QuickSort (Hoare) vs sort.Sort")
	fmt.Println(strings.Repeat("=", 60))

	for _, kind := range inputTypes {
		var quickTimes []time.Duration
		var stdlibTimes []time.Duration

		fmt.Printf("\nTipo de entrada: %s\n", kind)
		for i := 0; i < runs; i++ {
			original := generateSlice(kind, size)

			// QuickSort
			quickSlice := make([]int, len(original))
			copy(quickSlice, original)
			start := time.Now()
			quickSortHoare(quickSlice, 0, len(quickSlice)-1)
			quickDuration := time.Since(start)

			// sort.Sort
			stdlibSlice := make([]int, len(original))
			copy(stdlibSlice, original)
			start = time.Now()
			sort.Ints(stdlibSlice) // it calls sort.Sort after interace conversion
			stdlibDuration := time.Since(start)

			quickTimes = append(quickTimes, quickDuration)
			stdlibTimes = append(stdlibTimes, stdlibDuration)
		}

		fmt.Printf("%-10s | %-20s | %-20s\n", "Execução", "QuickSort", "sort.Sort")
		fmt.Println(strings.Repeat("-", 60))
		for i := 0; i < runs; i++ {
			fmt.Printf("%-10d | %-20v | %-20v\n", i+1, quickTimes[i], stdlibTimes[i])
		}
	}
}
