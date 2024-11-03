//（1）
// package main
// import "fmt"
//	func main() {
//	    fmt.Println("Hello, World!")
//	}

// （2）
// 小知识点，看他的中括号方的位置，必须方成这样
// package main
// import "fmt"
//	func main(){
//		var age int=25
//		fmt.Println("my age is",age)
//	}

// 3
// 可推断类型
// 感觉要求不太多呀，都可以加不加分号都无所谓
// package main
// import "fmt"
//	func main(){
//		var age=3
//		fmt.Println("hi",age)
//	}

// 4
// 声明多个变量
// package main
// import "fmt"
//
//	func main(){
//		var year,month int=2,3
//		fmt.Println("haha",year,"huhu",month)
//	}

// package main
// import "fmt"
//	func main(){
//		var year,month int
//		year=1
//		month=3
//		fmt.Println("jj",year,"kkk",month)
//	}

// 简写方法
// package main
// import "fmt"
//	func main(){
//	    var(
//			name="vfrdjfv"
//			age=10
//			height int
//		)
//		height=2
//		fmt.Println("hjfdsv",name,"vwjs",age,"fseghj",height)
//	}

// 简短声明的语法要求 := 操作符的左边至少有一个变量是尚未声明的。
// package main
// import "fmt"
//	func main(){
//		name,age :="hjfgrs",1
//		fmt.Println("efwsj",name,"hjfsr",age)
//	}

// package main
// import(
//	"fmt"
//	"math"
// )

//	func main(){
//		a,b :=111.1,222.2
//		c:=math.Min(a,b)
//		fmt.Println("hjdefssgiesdg",c)
//	}

// Go 是强类型（Strongly Typed）语言，因此不允许某一类型的变量赋值为其他类型的值
// package main
//	func main(){
//		age:=29
//		age="fjruwegu"
//	}

// 4....类型
// package main
// import "fmt"
//	func main(){
//		a:=true
//		b:=false
//		fmt.Println("a=",a,"b=",b)
//		c:=a&&b
//	    fmt.Println("c=",c)
//		d:=a||b
//	    fmt.Println("d=",d)
//	}

// package main
// import(
//	"fmt"
//	"unsafe"
// )
//	func main(){
//		//var a int=22;不是把，定义了没用都不行吗
//		b:=44
//		fmt.Println("type of is %T,size of b is %d",b,unsafe.Sizeof(b))
//	}

// 创建复数的两种方式
// package main
// import(
//	"fmt"
// )
//	func main(){
//		c1:=complex(5,7)
//		c2:=8+11i
//		cadd:=c1+c2
//		fmt.Println("sum:",cadd)
//		cmul:=c1*c2
//		fmt.Println("product",cmul)
//	}

// string类型
// package main
// import (
//	"fmt"
// )
//	func main(){
//		first:="jhefrwg"
//		last:="rfshk"
//		name:=first+" "+last
//		fmt.Println("111",name)
//	}

// package main
// import (
// 	"fmt"
// )
//	func main() {
//		i := 55
//		j := 11.1
//		sum := i + int(j)
//		fmt.Println(sum)
//	}
// package main

// import (
// 	"fmt"
// )
// func main() {
// 	i := 10
// 	var j float64 = float64(i) // 若没有显式转换，该语句会报错
// 	fmt.Println("j", j)
// }

// 这块还挺新奇的哈
// package main
// import (
// 	"fmt"
// 	"math"
// )
// func main() {
// 	var a = math.Sqrt(4) //可以被允许
// 	fmt.Println(a)
// 	//const b = math.Sqrt(4)//不能被允许
// }

// package main
// import (
// 	"fmt"
// )
// func main() {
// 	var name = "sam"
// 	fmt.Println("hello", "%T", "%v", name, name, name)
// 	fmt.Printf("%T----%v", name, name)//可以可以，高级高级
// }

// 这都不被允许，太强制爱了
// package main
// func main() {
// 	var defaultname = "sam"
// 	type mystring string
// 	var customname mystring = "sam"
// 	customname = defaultname
// }

// package main
// import (
// 	"fmt"
// )
// func main() {
// 	const a = 5
// 	var intvar int = a
// 	var int32var int32 = a
// 	var float64var float64 = a
// 	var complex64var complex64 = a
// 	fmt.Println("intvar", intvar, "\nint32var", int32var, "\nfloat64var", float64var, "\ncomplex64var", complex64var)
// }

// package main
// import (
// 	"fmt"
// )
// func main() {
// 	var i = 5
// 	var f = 5.6
// 	var c = 5 + 6i
// 	fmt.Printf("i's type %T,f' type %T,c's type %T", i, f, c)
// }

// 函数
// package main
// import (
// 	"fmt"
// )
// func A(price int, no int) int {
// 	var B = price * no
// 	return B
// }
// func main() {
// 	price, no := 90, 6
// 	total := A(price, no)
// 	fmt.Println("total", total)
// }
//两个参数类型一样，可以写成这样func A(price,no int) int

// 多返回值
// package main
// import (
// 	"fmt"
// )
//	func A(price, no int) (int, int) {
//		var B = price * no
//		var C = B + 10
//		return B, C
//	}
//
// 也可以写成
// func A(price, no int)(B,C int) {
// 	var B = price * no
// 	var C = B + 10
// 	return
// }
// func main() {
// 	price, no := 90, 6
// 	total, final := A(price, no)
// 	fmt.Println("total", total, "final", final)
// }

//空白夫
// package main
// import (
// 	"fmt"
// )
// func A(price, no int) (int, int) {
// 	var B = price * no
// 	var C = B + 10
// 	return B, C
// }
// func main() {
// 	price, no := 90, 6
// 	total, _ := A(price, no)
// 	fmt.Println("total", total)
// }

// if-else
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	num := 10
// 	if num%2 == 0 {
// 		fmt.Println("hahaha")
// 	} else {
// 		fmt.Println("wuwuwu")
// 	}
// }

// package main
// import (
// 	"fmt"
// )
// func main() {
// 	if num := 10; num%2 == 0 {
// 		fmt.Println("hahaha")
// 	} else {
// 		fmt.Println("wuwuwu")
// 	}
// }

//else不能放到下一行，这样会在}这个后面自动加一个分号
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	if num := 10; num%2 == 0 {
// 		fmt.Println("hahaha")
// 	} else if num%3 == 0 {
// 		fmt.Println("wuwuwu")
// 	} else {
// 		fmt.Println("hehehe")
// 	}
// }

// 循环
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	for i := 1; i < 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// package main
// import (
// 	"fmt"
// )
// func main() {
// 	for i := 1; i < 10; i++ {
// 		if i > 5 {
// 			break
// 		}
// 		fmt.Printf("%d ", i)
// 	}
// 	fmt.Printf("hahah")
// }

// package main
// import (
// 	"fmt"
// )
// func main() {
// 	for i := 1; i < 10; i++ {
// 		if i%2 == 0 {
// 			continue
// 		}
// 		fmt.Printf("%d ", i)
// 	}
// }

// package main
// import (
// 	"fmt"
// )
// func main() {
// 	i := 1
// 	for i < 10 {
// 		fmt.Printf("%d ", i)
// 		i++
// 	}
// }
//给for ;i<10;{}这种也可以，但是肯定没有不写简单

// package main
// import (
// 	"fmt"
// )
// func main() {
// 	for {
// 		fmt.Println("hahaha")
// 	}
// }
//无限循环

// switch
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	finger := 6
// 	//这里也可以省略finger
// 	switch finger{
// 	case 1:
// 		fmt.Println("Thumb")
// 	case 2:
// 		fmt.Println("Index")
// 	case 3:
// 		fmt.Println("Middle")
// 		//一个case可以包含多个表达式
// 	case 4, 8:
// 		fmt.Println("Ring")
// 	// case 4:
// 	// 	fmt.Println("Pinky")
// 	//go语言不允许出现重复向哦
// 	case 5:
// 		fmt.Println("Pinky")
// 	default:
// 		fmt.Println("incorrect finger number")
// 	}
// }

// package main
// import (
// 	"fmt"
// )
// func main() {
// 	num := 75
// 	switch { // 表达式被省略了
// 	case num >= 0 && num <= 50:
// 		fmt.Println("num is greater than 0 and less than 50")
// 	case num >= 51 && num <= 100:
// 		fmt.Println("num is greater than 51 and less than 100")
// 	case num >= 101:
// 		fmt.Println("num is greater than 100")
// 	}
// }

// package main
// import (
// 	"fmt"
// )
// func number() int {
// 	num := 15 * 5
// 	return num
// }
// func main() {
// 	switch num := number(); {
// 	case num < 50:
// 		fmt.Println("num is less than 50")
// 		fallthrough
// 	case num < 100:
// 		fmt.Println("num is less than 100")
// 		fallthrough
// 	case num < 200:
// 		fmt.Println("num is less than 200")
// 		//最后以不不需要哈
// 	}
// }

// 数组
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	var a [3]int
// 	a[0] = 12
// 	a[1] = 78
// 	a[2] = 50
// 	fmt.Println(a[0], a[1], a[2])
// 	fmt.Println(a)
// }

// package main
// import (
// 	"fmt"
// )
// func main() {
// 	a := [3]int{1, 2, 3}
// 	fmt.Println(a)
// }

// 数组是值类型,感觉就是没有指针哪种的
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	a := [...]string{"apple", "banana", "pear", "orange"}
// 	b := a
// 	b[0] = "grape"
// 	fmt.Println(a, b)
// }

// 计算数组长度
// package main

// import (
// 	"fmt"
// )
// //想一下两个打印函数的区别
// func main() {
// 	a := [...]float64{1.1, 3.3, 5.5}
// 	fmt.Printf("len(a)=%d\n", len(a))
// 	fmt.Println("len(a) is ", len(a))
// }

// //range函数的用法
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	a := [...]float64{1.1, 3.3, 5.5}
// 	sum := float64(0)
//     //如果这里你不想要i,可以for _, v := range a
// 	for i, v := range a {
// 		sum += v
// 		fmt.Printf("a[%d]=%v\n", i, v)
// 	}
// 	fmt.Println(sum)
// }

// 切片，这里咋根python好像
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	a := [5]int{1, 2, 3, 4, 5}
// 	var b []int = a[1:4]//4还要-1
// 	fmt.Println(b)
// }

// package main
// import (
// 	"fmt"
// )
// func main() {
// 	c := []int{1, 2, 3, 4, 5}
// 	fmt.Println(c)
// }

// 切片和数组的区别，就像引用一样，改了就是改了
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	a := [...]int{1, 2, 3, 4, 5}
// 	b := a[1:3]
// 	fmt.Println(a)
// 	b[0] = 100
// 	fmt.Println(a)
// }

// numa [:] 缺少开始和结束值。开始和结束的默认值分别为 0 和 len (numa)
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	a := [...]int{1, 2, 3, 4, 5}
// 	b := a[:]
// 	c := a[:]
// 	fmt.Println(a)
// 	b[1] = 9
// 	fmt.Println(a)
// 	c[2] = 8
// 	fmt.Println(a)
// }

// cap函数直观开始的值，切片可以重置其容量
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	a := [...]int{1, 2, 3, 4, 5, 6}
// 	b := a[2:4]
// 	fmt.Printf("length of b:=%d cap(b)=%d\n", len(b), cap(b))
// 	b = b[:cap(b)]
// 	fmt.Printf("length of b:=%d cap(b)=%d\n", len(b), cap(b))
// }

//func make（[]T，len，cap）[]T 通过传递类型，长度和容量来创建切片。容量是可选参数, 默认值为切片长度。make 函数创建一个数组，并返回引用该数组的切片。
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	i := make([]int, 5, 5)
// 	fmt.Println(i)
// }

// 追加切片元素,cap(a)为啥凡被呀，我感觉很奇怪
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	a := []string{"apple", "banana", "pear", "orange"}
// 	fmt.Printf("%d %d", len(a), cap(a))
// 	a = append(a, "grape")
// 	fmt.Printf("%d %d", len(a), cap(a))
// }

// 切片类型的零值为 nil。一个 nil 切片的长度和容量为 0。可以使用 append 函数将值追加到 nil 切片
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	var names []string
// 	if names == nil {
// 		names = append(names, "tom", "jerry", "mike")
// 		fmt.Printf("%d %d", len(names), cap(names))
// 	}
// }

// 使用 ... 运算符将一个切片添加到另一个切片,不是，太像python了
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	a := []string{"apple", "banana", "pear", "orange"}
// 	b := []string{"grape", "watermelon"}
// 	a = append(a, b...)
// 	fmt.Println(a)
// }

// 可变参函数//传入数组//传入切片
// package main
// import (
// 	"fmt"
// )
// func find(num int, nums ...int) {
// 	fmt.Printf("%T", nums)
// 	found := false
// 	for _, v := range nums {
// 		if v == num {
// 			found = true
// 			break
// 		}
// 	}
// 	if !found {
// 		fmt.Println(num, "not found in ", nums)
// 	}
// 	fmt.Printf("\n")
// }
// func main() {
// 	find(1, 1, 2, 3)
// 	find(2, 2, 3)
// 	nums := []int{1, 2, 3, 4, 5} //数组
// 	find(2, nums...)             //切片
// }

//说过切片想引用，但是append函数不会改变切片本身，而是返回一个包含原切片所有元素加上新添加元素的新切片
// package main
// import (
// 	"fmt"
// )
// func change(s ...string) {
// 	s[0] = "Go"
// 	s = append(s, "playground")
// 	fmt.Println(s)
// }
// func main() {
// 	welcome := []string{"hello", "world"}
// 	change(welcome...)
// 	fmt.Println(welcome)
// }

// // // map
// // //map 的零值是 nil,map 必须使用 make 函数初始化。
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	var a map[string]int
// 	if a == nil {
// 		fmt.Println("a is nil")
// 		a = make(map[string]int)
// 	}
// 	a["one"] = 1
// 	a["two"] = 2
// 	fmt.Println(a)
// }

// 这个初始化好像不用make
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	a := map[string]int{
// 		"one": 1,
// 		"two": 2,
// 	}
// 	a["three"] = 3
// 	fmt.Println(a)
// 	employee := "one"
// 	fmt.Println(a[employee])
// 	//如果我们获取一个不存在的元素，会返回 int 类型的零值 0
// 	fmt.Println(a["four"])
// 	//map 中到底是不是存在这个 key
// 	value, ok := a["four"]
// 	if ok == true {
// 		fmt.Println(value)
// 	} else {
// 		fmt.Println("key does not exist")
// 	}
// 	//获取map长度
// 	fmt.Println(len(a))
// 	delete(a, "one")
// 	fmt.Println(a)
// }

// //map是引用类型
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	a := map[string]int{
// 		"one": 1,
// 		"two": 2,
// 	}
// 	a["three"] = 3
// 	fmt.Println(a)
// 	b := a
// 	b["three"] = 4
// 	fmt.Println(a)
// }

// map之间不能使用==操作符判断，==只能用来检查map是否为nil
// package main
// import "fmt"
// func main() {
// 	map1 := map[string]int{
// 		"one": 1,
// 		"two": 2,
// 	}
// 	map2 := map1
// 	if map1 == map2 {
// 		fmt.Println("map1==map2")
// 	}
// }

// //看看分开打印有什么问题
// package main
// import (
// 	"fmt"
// )
// func A(s string) {
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%x", s[i])
// 	}
// 	fmt.Printf("\n")
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%c", s[i])
// 	}
// }
// func main() {
// 	A("Señor")
// }

// //b := []rune(s)这个可以解决这个问题
// 计算字符串的长度
// package main
// import (
// 	"fmt"
// 	"unicode/utf8"
// )
// func A(s string) {
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%x", s[i])
// 	}
// 	fmt.Printf("\n")
// 	b := []rune(s)
// 	for i := 0; i < len(b); i++ {
// 		fmt.Printf("%c", b[i])
// 	}
// 	fmt.Printf("字符串的长度为：%d", utf8.RuneCountInString(s))
// }
// func main() {
// 	A("Señor")
// }

//用range可以达到相同的效果
// package main
// import (
// 	"fmt"
// )
// func A(s string) {
// 	for _, b := range s {
// 		fmt.Printf("%c", b)
// 	}
// }
// func main() {
// 	A("Señor")
// }

// // 用字节切片构造函数
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	//a := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
// 	//若把16进制改成10进制会怎么杨,证明还是可以地
// 	a := []byte{67, 97, 102, 195, 169}
// 	str := string(a)
// 	fmt.Println(str)
// }

// // 用rune切片构造字符串,为啥上面那个输出和上面哪个不一样？？？？
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	a := []rune{0x0043, 0x0061, 0x0066, 0x00C3, 0x00A9}
// 	str := string(a)
// 	fmt.Println(str)
// 	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
// 	str1 := string(runeSlice)
// 	fmt.Println(str1)
// }

// 字符串的长度是不可编的,会报错
// package main
// import (
// 	"fmt"
// )
// func A(s string) string {
// 	s[0] = 'w'
// 	return s
// }
// func main() {
// 	k := "hello"
// 	fmt.Println(A(k))
// }

// // 可以先将他转成rune切片，修改后再转回字符串
// package main
// import (
// 	"fmt"
// )
// func A(s []rune) string {
// 	s[0] = 'a'
// 	return string(s)
// }
// func main() {
// 	k := "hello"
// 	fmt.Println(A([]rune(k)))
// }

// 指针
// package main
// import (
// 	"fmt"
// )
// func A(a *int) {
// 	*a = 10
// }
// func main() {
// 	b := 1
// 	var a *int//声明一个指针变量a,这样开始他就是nil
// 	if a == nil {
// 		a = &b
// 	}
// 	fmt.Println(*a)//界引用
// 	fmt.Printf("%T\n", a)//输出他的类型
// 	fmt.Println("address of b is", a)
// 	*a++
// 	fmt.Println(b)
// 	A(a)//将指针当作参数传入
// 	fmt.Println(b)
// }

// package main
// import (
// 	"fmt"
// )
// func A(arr *[3]int) {
// 	//(*arr)[0] = 100
// 	//也可以表示为
// 	arr[0] = 100
// }
// func main() {
// 	arr := [3]int{1, 2, 3}
// 	A(&arr)
// 	fmt.Println(arr)
// }

// // 切片传递
// package main
// import (
// 	"fmt"
// )
// func A(a []int) {
// 	a[0] = 100
// }
// func main() {
// 	a := [3]int{1, 2, 3}
// 	A(a[:])
// 	fmt.Println(a)
// }

// // go语言不支持指针运算，和c语言还是有一定区别的哈
// package main
// func main() {
// 	b := [3]int{1, 2, 3}
// 	p := &b
// 	p++
// }

// // 结构体
// package main
// import (
// 	"fmt"
// )
// type Person struct {
// 	f, l       string
// 	age, month int
// }
// func main() {
// 	var q Person
// 	fmt.Println(q)//初始化但是未复制那么都为0
// 	r := Person{
// 		f:     "zhang",
// 		month: 1,
// 		l:     "san",
// 		age:   18,
// 	}
// 	fmt.Println(r)
// 	p := Person{"zhang", "san", 18, 1}
// 	fmt.Println(p)
// }

// // 匿名结构体
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	a := struct {
// 		name string
// 		age  int
// 	}{
// 		name: "zhangsan",
// 		age:  18,
// 	}
// 	fmt.Println(a)
// }

// // 结构体指针
// package main
// import (
// 	"fmt"
// )
// type Person struct {
// 	f, l       string
// 	age, month int
// }
// func main() {
// 	a := &Person{"zhang", "san", 18, 1}
// 	fmt.Println("", (*a).l)
// 	fmt.Println("", a.l)
// }

// // 匿名字段
// package main
// import (
// 	"fmt"
// )
// type Person struct {
// 	string
// 	int
// }
// func main() {
// 	a := Person{"zhangsan", 18}
// 	fmt.Println(a)
// 	a.string = "lisi"
// 	a.int = 20
// 	fmt.Println(a)
// }

// // 结构体嵌套//提升字段，访问这两个字段就像在 Person 里直接声明的一样，因此我们称之为提升字段
// package main
// import (
// 	"fmt"
// )
// type Person struct {
// 	string
// 	int
// }
// type Student struct {
// 	Person
// 	int
// }
// func main() {
// 	a := Student{Person{"zhangsan", 18}, 20}
// 	fmt.Println(a.string, a.int)
// }

//结构体相等性
//结构体是值类型。如果它的每一个字段都是可比较的，则该结构体也是可比较的,map时不可以比较的

// // 方法，这个和c++比较像
// package main
// import (
// 	"fmt"
// )
// type Person struct {
// 	f, l       string
// 	age, month int
// }

// func (p Person) SayHi() {
// 	fmt.Println("Hi, my name is", p.f, p.l)
// }
// func main() {
// 	g := Person{
// 		f:     "zhang",
// 		l:     "san",
// 		age:   18,
// 		month: 1,
// 	}
// 	g.SayHi()
// }

// // 这个是只使用函数
// package main
// import (
// 	"fmt"
// )
// type Person struct {
// 	f, l       string
// 	age, month int
// }

// func SayHi(p Person) {
// 	fmt.Println("Hi, my name is", p.f, p.l)
// }
// func main() {
// 	g := Person{
// 		f:     "zhang",
// 		l:     "san",
// 		age:   18,
// 		month: 1,
// 	}
// 	SayHi(g)
// }

// // 相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的
// package main
// import (
// 	"fmt"
// 	"math"
// )
// type Rectangle struct {
// 	length int
// 	width  int
// }
// type Circle struct {
// 	radius float64
// }

// func (r Rectangle) Area() int {
// 	return r.length * r.width
// }
// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
// func main() {
// 	r := Rectangle{10, 20}
// 	c := Circle{10}
// 	fmt.Println(r.Area())
// 	fmt.Println(c.Area())
// }

// // 指针接收器与值接收器
// package main
// import (
// 	"fmt"
// )
// type Person struct {
// 	f, l       string
// 	age, month int
// }
// func (p Person) SayHi() {
// 	p.f = "wang"
// }
// func (p *Person) SayHi2() { //好奇怪呀，他这个指针传进来也可以用.
// 	p.f = "wang"
// }
// func main() {
// 	p := Person{"zhang", "san", 18, 1}
// 	p.SayHi()
// 	fmt.Println(p.f)
// 	//这两个效果一样
// 	//p.SayHi2()
// 	(&p).SayHi2()
// 	fmt.Println(p.f)
// }

// // 我们通过使用 p.fullAddress() 来访问 address 结构体的 fullAddress() 方法。明确的调用 p.address.fullAddress() 是没有必要的
// package main
// import (
// 	"fmt"
// )
// type Person struct {
// 	f, l       string
// 	age, month int
// }
// type Address struct {
// 	city, state string
// 	Person
// }
// func (a Person) fullAddress() {
// 	fmt.Printf("Person: %s %s\n", a.f, a.l)
// }
// func main() {
// 	a := Address{"Beijing", "China", Person{"zhang", "san", 18, 1}}
// 	a.fullAddress()
// }

// package main
// import(
// 	"fmt"
// )
// type rectangle struct{
// 	length int
// 	width int
// }
// func (r rectangle) area() int{
// 	return r.length*r.width
// }
// func area(r rectangle) int{
// 	return r.length*r.width
// }
// func main(){
// 	r := rectangle{10,20}
// 	fmt.Println(r.area())
// 	fmt.Println(area(r))
// 	p:=&r
// 	fmt.Println(p.area())
// 	fmt.Println(area(*p))
// 	//fmt.Println(area(p))这是错误的
// }

// 在非结构体类型上定义方法
package main

import (
	"fmt"
)

type myInt int

func (m myInt) add(i myInt) myInt {
	return m + i
}
func main() {
	m := myInt(10)
	i := myInt(20)
	fmt.Println(m.add(i))
}
