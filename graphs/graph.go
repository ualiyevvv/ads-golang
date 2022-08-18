package adsgraphs

import (
	"ads-golang/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	GlobalArr [][]int
)

func InputGraph() {
	args := os.Args[1:]
	if len(args) != 1 {
		panic("needs input file")
	}
	filePath := args[0]
	input := utils.ReadFile(filePath)
	rows := strings.Split(string(input), "\r\n")
	_, n2 := getMeta(rows[0])
	arr := convertToIntArray(rows[1:])
	GlobalArr = arr
	if !isValidGraph(arr) {
		panic("this graph couldn't exist")
	}
	countOf := OneGraph(arr[n2-1], n2-1, 0)
	fmt.Println("solve = ", countOf)
}

func getMeta(text string) (int, int) {
	s := strings.Split(text, " ")
	n1, err := strconv.Atoi(s[0])
	utils.CheckErr(err)
	n2, err := strconv.Atoi(s[1])
	utils.CheckErr(err)

	return n1, n2
}

func convertToIntArray(text []string) [][]int {
	arr := [][]int{}
	for _, t := range text {
		slice := goToInt(t)
		arr = append(arr, slice)
	}

	return arr
}

func goToInt(text string) []int {
	slice := []int{}
	s := strings.Split(text, " ")
	for _, r := range s {
		n, err := strconv.Atoi(r)
		utils.CheckErr(err)
		slice = append(slice, n)
	}

	return slice
}
