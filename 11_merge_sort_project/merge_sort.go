package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func TakeUserInput() []int {
	fmt.Println("Please enter a series of integers to be sorted (example: 32 23562 425 2462 64277 79 2245 2352): ")
	fmt.Printf("> ")

	scanner := bufio.NewScanner(os.Stdin)
	numbers := make([]int, 0)

	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")

	for _, value := range inputs {
		number, err := strconv.Atoi(value)
		if err == nil {
			numbers = append(numbers, number)
		} else {
			fmt.Println(err)
		}
	}
	return numbers
}

func Merge(left_slice []int, right_slice []int, merged_slices *[]int) {
	i, j, k := 0, 0, 0

	for i < len(left_slice) && j < len(right_slice) {

		if left_slice[i] < right_slice[j] {
			(*merged_slices)[k] = left_slice[i]
			i++
		} else if right_slice[j] < left_slice[i] {
			(*merged_slices)[k] = right_slice[j]
			j++
		} else {
			fmt.Println("NOT DETECTED CASE!")
		}
		k++

	}
	for i < len(left_slice) {
		(*merged_slices)[k] = left_slice[i]
		i++
		k++
	}
	for j < len(right_slice) {
		(*merged_slices)[k] = right_slice[j]
		j++
		k++
	}
}

func Sort(slice *[]int) {
	if len(*slice) < 2 {
		return
	}
	half := int(len(*slice) / 2)
	left_slice := (*slice)[0:half]
	right_slice := (*slice)[half:]

	Sort(&left_slice)
	Sort(&right_slice)

	left_slice_copy := make([]int, len(left_slice))
	copy(left_slice_copy, left_slice)

	right_slice_copy := make([]int, len(right_slice))
	copy(right_slice_copy, right_slice)

	Merge(left_slice_copy, right_slice_copy, slice)
}

func MergeSort(slice []int, wg *sync.WaitGroup, c chan []int) {
	fmt.Printf("A subarray that will be sorted by a goroutine: %v\n", slice)
	Sort(&slice)
	c <- slice
	wg.Done()
}

func PartitionSlice(input_slice []int, number_of_partitions int) [][]int {
	slices := make([][]int, 0)
	n := int(len(input_slice) / number_of_partitions)

	i, j := 0, n
	for k := 0; k < number_of_partitions; k++ {
		if k < number_of_partitions-1 {
			slices = append(slices, input_slice[i:j])
			i += n
			j += n
			continue
		}
		slices = append(slices, input_slice[i:])
		break
	}

	return slices
}

func main() {
	numbers := TakeUserInput()

	slices := PartitionSlice(numbers, 4)

	c1 := make(chan []int, 1)
	c2 := make(chan []int, 1)
	c3 := make(chan []int, 1)
	c4 := make(chan []int, 1)

	wg := sync.WaitGroup{}
	wg.Add(4)

	fmt.Println()
	go MergeSort(slices[0], &wg, c1)
	go MergeSort(slices[1], &wg, c2)
	go MergeSort(slices[2], &wg, c3)
	go MergeSort(slices[3], &wg, c4)

	wg.Wait()

	sorted_1 := <-c1
	sorted_2 := <-c2
	sorted_3 := <-c3
	sorted_4 := <-c4

	left_sorted := make([]int, len(sorted_1)+len(sorted_2))
	right_sorted := make([]int, len(sorted_3)+len(sorted_4))

	Merge(sorted_1, sorted_2, &left_sorted)
	Merge(sorted_3, sorted_4, &right_sorted)
	Merge(left_sorted, right_sorted, &numbers)

	fmt.Printf("\nThe fully sorted array is: %v", numbers)

}
