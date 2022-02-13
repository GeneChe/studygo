package main

import "fmt"

// 变量
// var name string
// var age int

// 批量声明变量
// var (
// 	height float64
// 	isOk   bool
// )

// 常量
// const pi = 3.14
// 批量声明常量
// const (
// 	statusOk = 200
// 	notFound = 404
// )

// 批量声明常量时, 如果某一行没有赋值, 默认与`上一行` `一致`
const (
	n1 = 100
	n2 // 等效 n2 = 100
	n3
)

// iota 是go语言中的常量计数器
// 在const出现时 iota被重置为0; const中每`新增一行`常量声明iota值加一
const (
	a1 = iota // 0
	a2        // a2 = iota, iota此时为1
	a3
)

const (
	b1 = iota
	b2
	_  //2
	b3 //3
)

const (
	c1 = iota
	c2 = 100
	c3        //100
	c4 = iota //3
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 //1, 2
	d3, d4 = iota + 1, iota + 2 //2, 3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) // 1024
	MB
	GB
	TB
)

func main() {
	var s1 = "wjja"
	fmt.Println(s1)
	fmt.Println("hello world!")
	// 同一作用域不能声明同名变量
	{
		s1 := 22 // 简短变量声明只能在函数内使用
		fmt.Println(s1)
	}
	// _ 匿名变量
	fmt.Println(n1, n2, n3)
	fmt.Println(a1, a2, a3)
	fmt.Println(b1, b2, b3)
	fmt.Println(c1, c2, c3, c4)
	fmt.Println(d1, d2, d3, d4)
	fmt.Println(KB, MB, GB, TB)
}

// go build 	编译当前目录文件
// go run [.go] 编译并执行对应.go文件
// go install 	编译当前目录文件并将生成的可执行文件copy到gopath的bin目录下
