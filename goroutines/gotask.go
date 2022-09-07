package goroutines

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

func TaskFunc() {
	size := 5
	arr := [][]int{}
	fillArr(&arr, size)
	fmt.Println(arr)
	n := 8

	for i := 0; i < len(arr); i++ {
		wg2.Add(1)
		go multiplySliceToNum(arr[i], n)
	}
	wg2.Wait()

	fmt.Println(arr)
}

func fillArr(arr *[][]int, size int) {
	mmsl := [][]int{}
	for i := 1; i < size; i++ {
		newsl := []int{}
		for j := 1; j < size; j++ {
			newsl = append(newsl, j)
		}
		mmsl = append(mmsl, newsl)
	}
	*arr = mmsl
}

func multiplySliceToNum(slice []int, n int) {
	for i := range slice {
		slice[i] *= n
	}
	defer wg2.Done()
}
