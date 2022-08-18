package adsgraphs

import (
	"ads-golang/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputGraph() {
	args := os.Args[1:]
	if len(args) != 1 {
		panic("needs input file")
	}
	filePath := args[0]
	input := utils.ReadFile(filePath)
	rows := strings.Split(string(input), "\r\n")
	n1, n2 := getMeta(rows[0])
	arr := convertToIntArray(rows[1:])
	if !isValidGraph(arr) {
		panic("this graph couldn't exist")
	} else {
		fmt.Println("valid")
	}
	fmt.Print(n1)
	fmt.Println("//")
	fmt.Print(n2)
	fmt.Println("//")
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
