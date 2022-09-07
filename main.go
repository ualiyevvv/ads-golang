package main

import (
	"ads-golang/leetcode"
	"fmt"
)

func main() {
	// tree := &bt.TreeNode{Val: 8}
	// tree.Add(4)
	// tree.Add(2)
	// tree.Add(77)
	// tree.Add(3)
	// tree.Add(10)
	// tree.Add(6)
	// tree.Add(15)
	// tree.Add(7)

	// tree.InorderTraversal()
	// fmt.Println("\nmax = ", tree.FindMax())
	// fmt.Println("min = ", tree.FindMin())
	// // tree.Delete(20)
	// // tree.InorderTraversal()

	// tree.Delete(77)
	// fmt.Println()
	// tree.InorderTraversal()
	// fmt.Println("\nmax = ", tree.FindMax())

	// dfs := tree.Dfs(10)
	// for _, r := range dfs {
	// 	fmt.Println(r)
	// }
	// adsgraphs.InputGraph()
	// goroutines.Start()
	// goroutines.TaskFunc()
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	// fmt.Println(leetcode.TaskN(s))
	leetcode.ReverseString(s, 0, len(s)-1)
	fmt.Println(string(s))
	// b := s[len(s)-1:]
	// fmt.Println(string(b))
	// fmt.Println(string(s[:len(s)-1]))
}
