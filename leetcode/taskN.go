package leetcode

import "fmt"

func TaskN(s []string) []string {
	fmt.Println(s)

	return nil
}

func ReverseString(s []byte, c int, e int) {
	if c > len(s)/2 {
		s[c], s[e] = s[e], s[c]
		c++
		e--

		fmt.Println(string(s))
		ReverseString(s, c, e)
	}
	// fmt.Println(string(s))
}

// func reverseN(input []string, output []string) []string {

// }
