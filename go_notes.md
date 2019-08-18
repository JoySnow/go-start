






## Go tour - reading note
https://tour.go-zh.org/flowcontrol/14

### Basic

#### 零值
没有明确初始值的变量声明会被赋予它们的 零值。

零值是：
 - 数值类型为 0，
 - 布尔类型为 false，
 - 字符串为 ""（空字符串）。
 - 切片的零值是 nil。 (nil 切片的长度和容量为 0 且没有底层数组。)
 - 映射的零值为 nil 。(nil 映射既没有键，也不能添加键。)

### FlowControl

#### defer

defer 语句会将函数推迟到外层函数返回之后执行。

推迟调用的函数其**参数会立即求值**，但直到外层函数返回前**该函数都不会被调用**。

推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照**后进先出(LIFO)**的顺序调用。

更多关于 defer 语句的信息，请阅读此[blog:defer-panic-and-recover](https://blog.go-zh.org/defer-panic-and-recover)。


### More Types

#### 数组

Require 创建时声明数组长度 like C.

表达式 `var a [10]int` , 会将变量 a 声明为拥有 10 个整数的数组。
数组的长度是其类型的一部分，因此数组不能改变大小。


#### Slice

切片就像数组的引用, **切片并不存储任何数据**，它只是描述了底层数组中的一段。


#### 切片的长度与容量

切片拥有 长度 和 容量。

切片的长度就是它所包含的元素个数。（是本次slice之后的结果，print出来的元素的个数。）

切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。(Base on 内存中数组实际的存储量。以切片后的第一个元素开始，到内存中数组的最后。)

切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。

切片下界的默认值为 0，上界则是该切片的长度, len, not cap。


```
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
len=6 cap=6 [2 3 5 7 11 13]
[2 3 5 7 11 13] &[2 3 5 7 11 13]
len=0 cap=6 []
[] &[]
len=4 cap=6 [2 3 5 7]
[2 3 5 7] &[2 3 5 7]
len=2 cap=4 [5 7]
[5 7] &[5 7]
len=4 cap=4 [5 7 11 13]
[5 7 11 13] &[5 7 11 13]

```


#### 用 make 创建切片
创建动态数组的方式。
分配一个元素为零值的数组并返回一个引用了它的切片：

第三个参数：可选，指定它的容量：
`b := make([]int, 0, 5) // len(b)=0, cap(b)=5`



#### slice append

When the old cap is not big enough, a new array space will be created in memory with larger cap.
And copy the values there, return the new address.
(Kind link Python's dict.) 

当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。

了解关于切片的更多内容，请阅读文章[Go 切片：用法和本质](https://blog.go-zh.org/go-slices-usage-and-internals) .
