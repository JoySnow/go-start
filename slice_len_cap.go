
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	fmt.Println(s, &s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)
	fmt.Println(s, &s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)
	fmt.Println(s, &s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
	fmt.Println(s, &s)
	
	s = s[:4]
	printSlice(s)
	fmt.Println(s, &s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


// Output: 
// ============================
//len=6 cap=6 [2 3 5 7 11 13]
//[2 3 5 7 11 13] &[2 3 5 7 11 13]
//len=0 cap=6 []
//[] &[]
//len=4 cap=6 [2 3 5 7]
//[2 3 5 7] &[2 3 5 7]
//len=2 cap=4 [5 7]
//[5 7] &[5 7]
//len=4 cap=4 [5 7 11 13]
//[5 7 11 13] &[5 7 11 13]
// ============================
