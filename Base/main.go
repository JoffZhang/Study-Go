package main

import (
	"fmt"
	"math"
	"reflect"
	"sync"
	"time"
	"unsafe"
)

/**
break    default      func    interface    select
case     defer        go      map          struct
chan     else         goto    package      switch
const    fallthrough  if      range        type
continue for          import  return       var
	var和const参考2.2Go语言基础里面的变量和常量申明
	package和import已经有过短暂的接触
	func 用于定义函数和方法
	return 用于从函数返回
	defer 用于类似析构函数
	go 用于并发
	select 用于选择不同类型的通讯
	interface 用于定义接口，参考2.6小节
	struct 用于定义抽象数据类型，参考2.5小节
	break、case、continue、for、fallthrough、else、if、switch、goto、default这些参考2.3流程介绍里面
	chan用于channel通讯
	type用于声明自定义类型
	map用于声明map类型数据
	range用于读取slice、map、channel数据
https://www.w3cschool.cn/yqbmht/d2a1cozt.html
*/

func main() {
	//1.变量
	variable()
	//2.常量
	constant()
	//3.运算
	operator()
	//4.条件语句
	conditionalStatement()
	//5.循环语句
	loopStatement()

	//结构体
	structural()
	//切片
	sliceDemo()
	//范围Range
	rangeDemo()
	//集合map
	mapDemo()
	//递归函数
	recurrenceDemo()
	//类型转换
	typeConversion()
	//接口
	interfaceDemo()
	//错误处理
	errorDemo()
	//反射
	reflectDemo()
	//并发
	concurrencyDemo()

}

/*******************************************************变量声明*******************************************************
变量名由字母、数字、下划线组成，其中首个字母不能为数字
声明变量的一般形式是使用 var 关键字
var identifier type
第一种，指定变量类型，声明后若不赋值，使用默认值。
第二种，根据值自行判定变量类型。
第三种，省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误。

*/

/**变量声明*******************/
var a = "变量声明"
var b string = "变量"
var c bool

/**多变量声明*****************/
var x, y int //类型相同多个变量, 非全局变量
var d, e int = 1, 2
var f, g = 1, "a" //自动推断类型

var ( //类型不同多个变量, 全局变量, 局部变量不能使用这种方式
	m int
	n int
)

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

/*
注意事项
如果在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明，例如：a := 20 就是不被允许的（因为在前文中已经声明了a:=50），编译器会提示错误 no new variables on left side of :=，但是 a = 20 是可以的，因为这是给相同的变量赋予一个新的值。
如果你在定义变量 a 之前使用它，则会得到编译错误 undefined: a。
如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误
但是全局变量是允许声明但不使用
同一类型的多个变量可以声明在同一行
多变量可以在同一行进行赋值
如果你想要交换两个变量的值，则可以简单地使用 a, b = b, a。
空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃
_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)。
*/
func variable() {
	g, h := 123, "hello" //声明后需使用，否则标红，编译declared but not used
	println(x, y, a, b, c, d, e, f, g, h, m, n)
}

/*******************************************************常量声明*******************************************************
常量是一个简单值的标识符，在程序运行时，不会被修改的量。
常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
常量的定义格式：
const identifier [type] = value
你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。

显式类型定义： const b string = "abc"
隐式类型定义： const b = "abc"

iota
iota，特殊常量，可以认为是一个可以被编译器修改的常量。
在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1
*/

//常量还可以用作枚举
const ( //数字 0、1 和 2 分别代表未知性别、女性和男性。
	Unknown = 0
	Female  = 1
	Male    = 2
)

//常量可以用len(), cap(), unsafe.Sizeof()常量计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不通过
const (
	a2 = "abc"
	b2 = len(a2)
	c2 = unsafe.Sizeof(a2)
)

//iota 可以被用作枚举值
const (
	a3 = iota
	b3 = iota
	c3 = iota
)

//第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2    可以简写为如下形式
/*const(
	a3 = iota
	b3
	c3
)*/

func constant() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a1, b1, c1 = 1, false, "str"
	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area)
	println()
	println(a1, b1, c1)

	println(a2, b2, c2)

	//iota 用法
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	const (
		m = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
	//iota表示从0开始自动加1，所以i=1<<0,j=3<<1（<<表示左移的意思），即：i=1,j=6，这没问题，关键在k和l，从输出结果看，k=3<<2，l=3<<3。

}

/*******************************************************运算符*******************************************************
算术 关系 逻辑 位 赋值 其他

&	返回变量存储地址	&a; 将给出变量的实际地址。
*	指针变量。		*a; 是一个指针变量
*/
func operator() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

	/*  & 和 * 运算符实例 */
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Println("ptr 为 ", ptr)
	fmt.Printf("*ptr 为 %d\n", *ptr)
}

/*******************************************************循环语句*******************************************************
for init; condition; post { }
for condition { }
for { }	//for(;;)
	init： 一般为赋值表达式，给控制变量赋初值；
	condition： 关系表达式或逻辑表达式，循环控制条件；
	post： 一般为赋值表达式，给控制变量增量或减量。
for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
for key, value := range oldMap {
    newMap[key] = value
}

GO 语言支持以下几种循环控制语句：
	break 语句	经常用于中断当前 for 循环或跳出 switch 语句
	continue 语句	跳过当前循环的剩余语句，然后继续进行下一轮循环。
	goto 语句	将控制转移到被标记的语句。
		goto语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
		但是，在结构化程序设计中一般不主张使用goto语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。
无限循环
for true  {
        fmt.Printf("这是无限循环。\n");
}
*/

func loopStatement() {
	var b int = 15
	var a int
	numbers := [6]int{1, 2, 3, 5}
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值： %d \n", a)
	}
	for a < b {
		a++
		fmt.Printf("a 的值为：%d \n", a)
	}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	var i, j int
	for i = 2; i < 10; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 跳转到标签
				goto breakHere
			}
		}
	}
	// 手动返回, 避免执行进入标签
	return
	// 标签
breakHere:
	fmt.Println("done")

	//使用goto 集中相同错误异常处理

}

/*******************************************************函数*******************************************************
函数是基本的代码块，用于执行一个任务。

	Go 语言最少有1个 main() 函数。
	你可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务。
	函数声明告诉了编译器函数的名称，返回类型和参数。

程序的初始化和执行都起始于main包。如果main包还导入了其它的包，那么就会在编译时将它们依次导入。有时一个包会被多个包同时导入，那么它只会被导入一次（例如很多包可能都会用到fmt包，但它只会被导入一次，因为没有必要导入多次）。当一个包被导入时，如果该包还导入了其它的包，那么会先将其它包导入进来，然后再对这些包中的包级常量和变量进行初始化，接着执行init函数（如果有的话），依次类推。等所有被导入的包都加载完毕了，就会开始对main包中的包级常量和变量进行初始化，然后执行main包中的init函数（如果存在的话），最后执行main函数。

func function_name( [parameter list] ) [return_types]{
   函数体
}
函数定义解析：

func：函数由 func 开始声明
function_name：函数名称，函数名和参数列表一起构成了函数签名。
parameter list]：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序及参数个数。参数是可选的，也就是说函数也可以不包含参数。
return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
函数体：函数定义的代码集合。

函数参数
函数如果使用参数，该变量可称为函数的形参。
形参就像定义在函数体内的局部变量。

调用函数，可以通过两种方式来传递参数：
值传递	值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
引用传递	引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
	 调用 swap() 函数
		var a int = 100
		var b int= 200

	   * &a 指向 a 指针，a 变量的地址
	   * &b 指向 b 指针，b 变量的地址
		swap2(&a, &b)

函数用法
函数作为值	函数定义后可作为值来使用
闭包			闭包是匿名函数，可在动态编程中使用
				Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
				nextNumber 为一个函数，函数 i 为 0
				nextNumber := getSequence()

				调用 nextNumber 函数，i 变量自增 1 并返回
				fmt.Println(nextNumber())
				fmt.Println(nextNumber())

方法			方法就是一个包含了接受者的函数
				Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集
				func (variable_name variable_data_type) function_name() [return_type]{
				    函数体
				}
				  var c1 Circle
				  c1.radius = 10.00
				  fmt.Println("Area of Circle(c1) = ", c1.getArea())

defer语句
Go语言中的defer语句会将其后面跟随的语句进行延迟处理
在defer所属的函数即将返回时，将延迟处理的语句按照defer定义的顺序逆序执行，即先进后出


main函数和init函数
Go里面有两个保留的函数：init函数（能够应用于所有的package）和main函数（只能应用于package main）。这两个函数在定义时不能有任何的参数和返回值。虽然一个package里面可以写任意多个init函数，但这无论是对于可读性还是以后的可维护性来说，我们都强烈建议用户在一个package中每个文件只写一个init函数。

Go程序会自动调用init()和main()，所以你不需要在任何地方调用这两个函数。每个package中的init函数都是可选的，但package main就必须包含一个main函数。

*/
func swap(x, y string) (string, string) {
	return y, x
}
func swap2(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

//函数定义后可作为值来使用
func getSquareRoot() {
	squareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(squareRoot(9))
}

//创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量
func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//定义一个结构体类型和该类型的一个方法
/* 定义函数 */
type Circle struct {
	radius float64
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}
func deferDemo() {
	fmt.Println("开始")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("结束")
}

/*******************************************************变量作用域*******************************************************
Go 语言中变量可以在三个地方声明：

	函数内定义的变量称为局部变量
	函数外定义的变量称为全局变量
	函数定义中的变量称为形式参数

局部变量
在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。
全局变量
在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。
	Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。
形式参数
形式参数会作为函数的局部变量来使用

初始化局部和全局变量
不同类型的局部和全局变量默认值为：
	int				0
	float32			0
	pointer			nil
*/
/*******************************************************数组*******************************************************
数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整型、字符串或者自定义类型。
	数组元素可以通过索引（位置）来读取（或者修改），索引从0开始，第一个元素索引为 0，第二个索引为 1，以此类推。

var variable_name [SIZE] variable_type
一维数组的定义方式。数组长度必须是整数且大于 0
var balance [10] float32

初始化数组
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
 var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}

访问数组元素
数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。
float32 salary = balance[9]

多维数组			Go 语言支持多维数组，最简单的多维数组是二维数组
向函数传递数组	你可以像函数传递数组参数
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
多维数组可通过大括号来初始值
a = [3][4]int{
 {0, 1, 2, 3} ,   第一行索引为 0
{4, 5, 6, 7} ,    第二行索引为 1
{8, 9, 10, 11}    第三行索引为 2
}

如果你想向函数传递数组参数，你需要在函数定义时，声明形参为数组，我们可以通过以下两种方式来声明：
方式一形参设定数组大小：
void myFunction(param [10]int){}
方式二形参未设定数组大小：
void myFunction(param []int){}



*/
func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}

/*******************************************************指针*******************************************************
变量是一种使用方便的占位符，用于引用计算机内存地址。

Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
什么是指针
一个指针变量可以指向任何一个值的内存地址它指向那个值的内存地址
var var_name *var-type
var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。
如何使用指针
指针使用流程：

	定义指针变量。
	为指针变量赋值。
	访问指针变量中指向地址的值。
	在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。

空指针
当一个指针被定义后没有分配到任何变量时，它的值为 nil。
nil 指针也称为空指针。
nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
一个指针变量通常缩写为 ptr。

Go 指针数组				你可以定义一个指针数组来存储地址
Go 指向指针的指针			Go 支持指向指针的指针
Go 向函数传递指针参数		通过引用或地址传参，在函数调用时可以改变其值

var ptr [MAX]*int
ptr[i] = &a[i] // 整数地址赋值给指针数组

如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址
var ptr **int
访问指向指针的指针变量值需要使用两个 * 号

Go 语言允许向函数传递指针，只需要在函数定义的参数上设置为指针类型即可
swap2()
*/

/*******************************************************结构体*******************************************************
结构体中我们可以为不同项定义不同的数据类型
结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
结构体表示一项记录

定义结构体
结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
一旦定义了结构体类型，它就能用于变量的声明
variable_name := structure_variable_type {value1, value2...valuen}
访问结构体成员
如果要访问结构体成员，需要使用点号 (.) 操作符，格式为："结构体.成员名"。
结构体类型变量使用struct关键字定义

结构体作为函数参数
结构体指针

*/
type Books struct {
	title  string
	author string
}

func structural() {
	var Book1 Books
	Book1.author = "我"
	Book1.title = "Go 语言"

	printBook(Book1, &Book1)
}
func printBook(book Books, book2 *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("*Books title : %s\n", book2.title)
	fmt.Printf("*Books author : %d\n", book2.author)
}

/*******************************************************切片Slice*******************************************************
切片是对数组的抽象
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

定义切片
声明一个未指定大小的数组来定义切片
var identifier []type
切片不需要说明长度。
或使用make()函数来创建切片:
var slice1 []type = make([]type, len)     简写  slice1 := make([]type, len)

也可以指定容量，其中capacity为可选参数。
make([]T, length, capacity)
这里 len 是数组的长度并且也是切片的初始长度。

切片初始化
s :=[] int {1,2,3 }
直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
s := arr[:]
初始化切片s,是数组arr的引用
s := arr[startIndex:endIndex]
将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
s := arr[startIndex:]
缺省endIndex时将表示一直到arr的最后一个元素
s := arr[:endIndex]
缺省startIndex时将表示从arr的第一个元素开始
s1 := s[startIndex:endIndex]
通过切片s初始化切片s1

s :=make([]int,len,cap)
通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片

len() 和 cap() 函数
切片是可索引的，并且可以由 len() 方法获取长度。
切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

空(nil)切片
一个切片在未初始化之前默认为 nil，长度为 0

切片截取
可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]
append() 和 copy() 函数
如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。

*/

func sliceDemo() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)

	/* 初始化切片 */
	numbers = []int{0, 1, 2, 3, 4}
	/* 打印原始切片 */
	printSlice(numbers)
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)
	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)
	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers3 := make([]int, len(numbers), (cap(numbers))*2)
	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers3, numbers)
	printSlice(numbers3)
}

func printSlice(x []int) {
	fmt.Printf("x is nil %b , len=%d cap=%d slice=%v\n", x == nil, len(x), cap(x), x)
}

/*******************************************************范围Range*******************************************************
Go 语言中 range 关键字用于for循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值。
对于映射，它返回下一个键值对的键。Range返回一个值或两个值。如果在Range表达式的左侧只使用了一个值，则该值是下表中的第一个值。

Range表达式					第一个值	 				第二个值[可选的]
Array 或者 slice a [n]E	 	索引 i int	 			a[i] E
String s string type	 	索引 i int	 			rune int
map m map[K]V	 			键 k K	 				值 m[k] V
channel c chan E	 		元素 e E	 				none

*/
func rangeDemo() {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}

/*******************************************************Map(集合)*******************************************************
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

定义 Map
可以使用内建函数 make 也可以使用 map 关键字来定义 Map:
	声明变量，默认 map 是 nil
var map_variable map[key_data_type]value_data_type
	使用 make 函数
map_variable = make(map[key_data_type]value_data_type)

如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对


delete() 函数
delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。实例如下：





*/

func mapDemo() {
	var countryCapitalMap map[string]string
	//创建集合
	countryCapitalMap = make(map[string]string)
	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"
	/* 使用 key 输出 map 值 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if ok {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}
	/* 删除元素 */
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")
	fmt.Println("删除元素后 map")
	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
}

/*******************************************************递归函数*******************************************************
递归，就是在运行的过程中调用自己。
func recursion() {
   recursion() // 函数调用自身
}
Go 语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。
递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。
*/
func recurrenceDemo() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(i))

	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}

}

//阶乘
func Factorial(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * Factorial(x-1)
	}
	return
}

//斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

/*******************************************************类型转换*******************************************************
类型转换用于将一种数据类型的变量转换为另外一种类型的变量。Go 语言类型转换基本格式如下：
type_name(expression)
type_name 为类型，expression 为表达式。

go 不支持隐式转换类型

*/
func typeConversion() {
	var sum int = 18
	var count int = 4
	var mean float32
	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)
	//go 不支持隐式转换类型
	//var a int64 = 3
	//var b int32
	//b = a		//编辑器会提示报错		如果改成 ​b = int32(a) ​就不会报错了
	//fmt.Printf("b 为 : %d", b)

}

/*******************************************************接口*******************************************************
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
定义接口
type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
}
定义结构体
type struct_name struct {
    variables
}
实现接口方法
func (struct_name_variable struct_name) method_name1() [return_type] {
}
func (struct_name_variable struct_name) method_namen() [return_type] {
}
*/

type Phone interface {
	call()
}
type NokiaPhone struct{}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct{}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

//定义了一个接口Phone，接口里面有一个方法call()。然后我们在main函数里面定义了一个Phone类型变量，并分别为之赋值为NokiaPhone和IPhone。然后调用call()方法
func interfaceDemo() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}

/*******************************************************错误处理*******************************************************
通过内置的错误接口提供了非常简单的错误处理机制。
error类型是一个接口类型，这是它的定义：
type error interface {
    Error() string
}
可以在编码中通过实现 error 接口类型来生成错误信息。
函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息
Panic和Recover
Go没有像Java那样的异常机制，它不能抛出异常，而是使用了panic和recover机制。一定要记住，你应当把它作为最后的手段来使用，也就是说，你的代码中应当没有，或者很少有panic的东西。这是个强大的工具，请明智地使用它。那么，我们应该如何使用它呢？

Panic

是一个内建函数，可以中断原有的控制流程，进入一个令人恐慌的流程中。当函数F调用panic，函数F的执行被中断，但是F中的延迟函数会正常执行，然后F返回到调用它的地方。在调用的地方，F的行为就像调用了panic。这一过程继续向上，直到发生panic的goroutine中所有调用的函数返回，此时程序退出。恐慌可以直接调用panic产生。也可以由运行时错误产生，例如访问越界的数组。

Recover

是一个内建的函数，可以让进入令人恐慌的流程中的goroutine恢复过来。recover仅在延迟函数中有效。在正常的执行过程中，调用recover会返回nil，并且没有其它任何效果。如果当前的goroutine陷入恐慌，调用recover可以捕获到panic的输入值，并且恢复正常的执行。
*/
//定义一个DivideError结构
type DivideError struct {
	dividee int
	divider int
}

//实现error接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func errorDemo() {
	//正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println(" 100/10 = ", result)
	}
	//当被除数为0返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is :", errorMsg)
	}
}

/*******************************************************反射（Reflect）*******************************************************
Go语言提供了一种机制，在不知道具体类型的情况下，可以用反射来更新变量值，查看变量类型
Typeof
Typeof返回接口中保存的值得类型，Typeof(nil)会返回nil
ValueOf
ValueOf返回一个初始化为interface接口保管的具体值得Value，ValueOf(nil)返回Value零值

使用建议
1、大量使用反射的代码通常会变得难以理解
2、反射的性能低下，基于反射的代码会比正常的代码运行速度慢一到两个数量级
*/
func reflectsetvalue1(x interface{}) {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.String {
		value.SetString("欢迎来到W3Cschool")
	}
}
func reflectsetvalue2(x interface{}) {
	value := reflect.ValueOf(x)
	//反射中使用Elem()方法获取指针所指向的值
	if value.Elem().Kind() == reflect.String {
		value.Elem().SetString("欢迎来到W3Cschool")
	}
}
func reflectDemo() {
	var booknum float32 = 6
	var isbook bool = true
	bookauthor := "www.w3cschool.cn"
	bookdetail := make(map[string]string)
	bookdetail["Go语言教程"] = "www.w3cschool.cn"
	fmt.Println(reflect.TypeOf(booknum))
	fmt.Println(reflect.TypeOf(isbook))
	fmt.Println(reflect.TypeOf(bookauthor))
	fmt.Println(reflect.TypeOf(bookdetail))

	fmt.Println(reflect.ValueOf(booknum))
	fmt.Println(reflect.ValueOf(isbook))
	fmt.Println(reflect.ValueOf(bookauthor))
	fmt.Println(reflect.ValueOf(bookdetail))

	//通过反射设置值
	address := "www.w3cschool.cn"
	reflectsetvalue1(address)
	fmt.Println(address)

	// 反射修改值必须通过传递变量地址来修改。若函数传递的参数是值拷贝，则会发生下述错误。
	// panic: reflect: reflect.Value.SetString using unaddressable value
	reflectsetvalue2(&address)
	fmt.Println(address)
}

/*******************************************************并发*******************************************************
并发与并行
并发：同一时间段内执行多个任务（你早上在编程狮学习Java和Python）

并行：同一时刻执行多个任务（你和你的网友早上都在使用编程狮学习Go）
Go语言中的并发程序主要是通过基于CSP（communicating sequential processes）的goroutine和channel来实现，当然也支持使用传统的多线程共享内存的并发方式

goroutine
Go语言中使用goroutine非常简单，只需要在函数或者方法前面加上go关键字就可以创建一个goroutine，从而让该函数或者方法在新的goroutine中执行
匿名函数同样也支持使用go关键字来创建goroutine去执行
一个goroutine必定对应一个函数或者方法，可以创建多个goroutine去执行相同的函数或者方法


sync.WaitGroup
Go语言中的sync包为我们提供了一些常用的并发原语
当你并不关心并发操作的结果或者有其它方式收集并发操作的结果时，WaitGroup是实现等待一组并发操作完成的好方法
动态栈
操作系统的线程一般都有固定的栈内存（通常为2MB），而 Go 语言中的 goroutine 非常轻量级，一个 goroutine 的初始栈空间很小（一般为2KB），所以在 Go 语言中一次创建数万个 goroutine 也是可能的。并且 goroutine 的栈不是固定的，可以根据需要动态地增大或缩小， Go 的 runtime 会自动为 goroutine 分配合适的栈空间。

goroutine调度
在经过数个版本迭代之后，目前Go语言的调度器采用的是GPM调度模型

G: 表示goroutine，存储了goroutine的执行stack信息、goroutine状态以及goroutine的任务函数等；另外G对象是可以重用的。
P: 表示逻辑processor，P的数量决定了系统内最大可并行的G的数量（前提：系统的物理cpu核数>=P的数量）；P的最大作用还是其拥有的各种G对象队列、链表、一些cache和状态。
M: M代表着真正的执行计算资源。在绑定有效的p后，进入schedule循环；而schedule循环的机制大致是从各种队列、p的本地队列中获取G，切换到G的执行栈上并执行G的函数，调用goexit做清理工作并回到m，如此反复。M并不保留G状态，这是G可以跨M调度的基础。

GOMAXPROCS
Go运行时，调度器使用GOMAXPROCS的参数来决定需要使用多少个OS线程来同时执行Go代码。默认值是当前计算机的CPU核心数。例如在一个8核处理器的电脑上，GOMAXPROCS默认值为8。Go语言中可以使用runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU核心数

channel
单纯地将函数并发执行是没有意义的，函数与函数间需要交换数据才能体现并发执行函数的意义
虽然可以使用共享内存进行数据交换，但是共享内存在不同的 goroutine 中容易发生竞态问题。为了保证数据交换的正确性，很多并发模型中必须使用互斥锁对内存进行加锁，这种做法势必造成性能问题

Go语言采用的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存，而不是通过共享内存而实现通信
Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

channel类型
声明通道类型变量方法如下
var 变量名 chan 元素类型
其中chan是关键字，元素类型指通道中传递的元素的类型
	var a chan int //声明一个传递int类型的通道
	var b chan string // 声明一个传递string类型的通道
	var c chan bool //声明一个传递bool类型的通道

channel零值
未经初始化的通道默认值为nil

初始化channel
声明的通道类型变量需要使用内置的make函数初始化之后才能使用，具体格式如下
	make(chan 元素类型,[缓冲大小])
channel的缓冲大小是可选的
a:=make(chan int)
b:=make(chan int,10)//声明一个缓冲大小为10的通道

channel操作
通道共有发送，接收，关闭三种操作，而发送和接收操作均用​<-​符号，举几个例子
声明通道并初始化
a := make(chan int) //声明一个通道并初始化
给一个通道发送值
a <- 10  //把10发送给a通道
从一个通道中取值
x := <-a //x从a通道中取值
<-a      //从a通道中取值，忽略结果
关闭通道
close(a) //关闭通道
一个通道值是可以被垃圾回收掉的。通道通常由发送方执行关闭操作，并且只有在接收方明确等待通道关闭的信号时才需要执行关闭操作。它和关闭文件不一样，通常在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。

关闭后的通道有以下特点

对一个关闭的通道再发送值就会导致 panic。
对一个关闭的通道进行接收会一直获取值直到通道为空。
对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
关闭一个已经关闭的通道会导致 panic。

无缓冲的通道
无缓冲的通道又称为阻塞的通道

deadlock表示我们程序中所有的goroutine都被挂起导致程序死锁了，为什么会出现这种情况呢？
这是因为我们创建的是一个无缓冲区的通道，无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段。同理，如果对一个无缓冲通道执行接收操作时，没有任何向通道中发送值的操作那么也会导致接收操作阻塞。

有缓冲区的通道
另外还有一种方法解决上述死锁的问题，那就是使用有缓冲区的通道。我们可以在使用make函数初始化通道时，为其指定缓冲区大小
只要通道的容量大于零，那么该通道就属于有缓冲的通道，通道的容量表示通道中最大能存放的元素数量。当通道内已有元素数达到最大容量后，再向通道执行发送操作就会阻塞，除非有从通道执行接收操作。
我们可以使用内置的len函数获取通道的长度，使用cap函数获取通道的容量

判断通道关闭
当向通道中发送完数据时，我们可以通过close函数来关闭通道。当一个通道被关闭后，再往该通道发送值会引发panic。从该通道取值的操作会先取完通道中的值。通道内的值被接收完后再对通道执行接收操作得到的值会一直都是对应元素类型的零值。
value, ok := <-ch

value：表示从通道中所取得的值
ok：若通道已关闭，返回false，否则返回true

for range接收值
通常我们会使用for range循环来从通道中接收值，当通道关闭后，会在通道内所有值被取完之后退出循环


单向通道
在某些场景下我们可能会将通道作为参数在多个任务函数间进行传递，通常我们会选择在不同的任务函数中对通道的使用进行限制，比如限制通道在某个函数中只能执行发送或只能执行接收操作

<- chan int // 只接收通道，只能接收不能发送
chan <- int // 只发送通道，只能发送不能接收

select多路复用
在某些场景下我们可能需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以被接收那么当前 goroutine 将会发生阻塞。Go语言内置了select关键字，使用它可以同时响应多个通道的操作，具体格式如下
select {
case <-ch1:
	//...
case data := <-ch2:
	//...
case ch3 <- 10:
	//...
default:
	//默认操作
}
select语句具有以下特点

可处理一个或多个channel的发送/接收操作
如果多个case同时满足，select会随机选择一个执行
对于没有case的select会一直阻塞，可用于阻塞 main 函数，防止退出

并发安全和互斥锁
有时候我们的代码中可能会存在多个 goroutine 同时操作一个资源的情况，这种情况下就会发生数据读写错乱的问题

互斥锁
互斥锁是一种常用的控制共享资源访问的方法，它能够保证同一时间只有一个 goroutine 可以访问共享资源。Go语言中使用sync包中提供的Mutex类型来实现互斥锁
使用互斥锁能够保证同一时间有且只有一个 goroutine 进入临界区，其他的 goroutine 则在等待锁；当互斥锁释放后，等待的 goroutine 才可以获取锁进入临界区，多个 goroutine 同时等待一个锁时，唤醒的策略是随机的

读写互斥锁
互斥锁是完全互斥的，但是实际上有很多场景是读多写少的，当我们并发的去读取一个资源而不涉及资源修改的时候是没有必要加互斥锁的，这种场景下使用读写锁是更好的一种选择。在Go语言中使用sync包中的RWMutex类型来实现读写互斥锁
读写锁分为两种：读锁和写锁。当一个 goroutine 获取到读锁之后，其他的 goroutine 如果是获取读锁会继续获得锁，如果是获取写锁就会等待；而当一个 goroutine 获取写锁之后，其他的 goroutine 无论是获取读锁还是写锁都会等待


*/
func hello1() {
	fmt.Println("hello")
}

//以下main NUM  都对应的是main函数  因一个文件下只能有一个主函数，所以暂代
func main1() {
	go hello1()
	fmt.Println("欢迎来到编程狮")
}

/*
只在终端控制台输出了“欢迎来到编程狮”，并没有打印“hello”
其实在Go程序中，会默认为main函数创建一个goroutine，而在上述代码中我们使用go关键字创建了一个新的goroutine去调用hello函数。而此时main的goroutine还在往下执行中，我们的程序中存在两个并发执行的goroutine。当main函数结束时，整个程序也结束了，所有由main函数创建的子goroutine也会跟着退出，也就是说我们的main函数执行过快退出导致另一个goroutine内容还未执行就退出了，导致未能打印出hello
所以我们这边要想办法让main函数等一等，让另一个goroutine的内容执行完。其中最简单的方法就是在main函数中使用time.sleep睡眠一秒钟

*/

func main2() {
	go hello1()
	fmt.Println("欢迎来到编程狮")
	time.Sleep(time.Second)
}

/*
为什么会先打印欢迎来到编程狮呢？

这是因为在程序中创建 goroutine 执行函数需要一定的开销，而与此同时 main 函数所在的 goroutine 是继续执行的。
*/

var wg sync.WaitGroup

func hello2() {
	fmt.Println("hello2")
	defer wg.Done() //把计数器-1

}

//sync.WaitGroup
func main3() {
	wg.Add(1) //把计数器+1
	go hello2()
	fmt.Println("欢迎来到编程狮")
	wg.Wait() //阻塞代码的运行，直到计算器为0
}

//启动多个goroutine
var wg1 sync.WaitGroup

func hello3(i int) {
	fmt.Printf("hello,欢迎来到编程狮%v\n", i)
	defer wg.Done() //goroutine结束计数器-1
}

//执行多次上述代码你会发现输出顺序并不一致，这是因为10个goroutine都是并发执行的，而goroutine的调度是随机的
func main4() {
	for i := 0; i < 10; i++ {
		go hello3(i)
		wg.Add(1) //启动一个goroutine计数器+1
	}
	wg.Wait() //等待所有的goroutine执行结束
}

func receive(x chan int) {
	ret := <-x
	fmt.Println("接收成功", ret)
}
func receive2(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}
func receive2_forRange(ch chan int) {
	for i := range ch {
		fmt.Printf("v:%v", i)
	}
}

/*并发安全和互斥锁*/
var (
	x2  int64
	wg2 sync.WaitGroup //等待组
)

// add 对全局变量x执行5000次加1操作
func add() {
	for i := 0; i < 5000; i++ {
		x2 = x2 + 1
	}
	wg2.Done()
}

/*互斥锁*/
var (
	x3  int64
	wg3 sync.WaitGroup
	m3  sync.Mutex //互斥锁
)

func add3() {
	for i := 0; i < 5000; i++ {
		m3.Lock() //修改x前加锁
		x = x + 1
		m3.Unlock() //改完解锁
	}
	wg.Done()
}

/*读写互斥锁*/
var (
	x4      = 0
	wg4     sync.WaitGroup
	rwLock4 sync.RWMutex
)

func read() {
	defer wg4.Done()
	rwLock4.RLock()
	fmt.Println(x4)
	time.Sleep(time.Microsecond)
	rwLock4.RUnlock()
}
func write() {
	defer wg4.Done()
	rwLock4.Lock()
	x += 1
	time.Sleep(time.Microsecond * 5)
	rwLock4.Unlock()
}

func main5() {
	//开启了2个goroutine去执行add函数，某个goroutine对全局变量x的修改可能会覆盖掉另外一个goroutine中的操作，所以导致结果与预期不符
	wg2.Add(2)
	go add()
	go add()
	wg2.Wait()
	fmt.Println(x2)

	/*互斥锁*/
	wg3.Add(2)
	go add3()
	go add3()
	wg3.Wait()
	fmt.Println(x3)

	/*读写互斥锁*/
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg4.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go read()
		wg4.Add(1)
	}
	wg4.Wait()
	fmt.Println(time.Since(start))

}

func concurrencyDemo() {

	/*无缓冲区的通道*/
	//以下代码编译通过，执行报错
	/*a := make(chan int)
	a <- 10
	fmt.Println("发送成功")*/
	//deadlock表示我们程序中所有的goroutine都被挂起导致程序死锁了

	a := make(chan int)
	go receive(a)
	a <- 10
	fmt.Println("发送成功", 10)

	/*有缓冲区的通道*/
	a = make(chan int, 1)
	a <- 1
	fmt.Println("发送成功", 1)

	/*判断通道关闭*/
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	receive2(ch)

	//select多路复用

	ch = make(chan int, 1) //创建一个类型为int，缓冲区大小为1的通道
	for i := 1; i < 10; i++ {
		select {
		case x := <-ch: //第一次循环由于没有值，所以该分支不满足
			fmt.Println(x)
		case ch <- i: //将i发送给通道(由于缓冲区大小为1，缓冲区已满，第二次不会走该分
		}
	}
}

/*******************************************************条件语句*******************************************************
if 语句			if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
if...else 语句	if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
if 嵌套语句		你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
switch 语句		switch 语句用于基于不同条件执行不同动作。
select 语句		select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
					select是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。
					select随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。
				以下描述了 select 语句的语法：
					每个case都必须是一个通信
					所有channel表达式都会被求值
					所有被发送的表达式都会被求值
					如果任意某个通信可以进行，它就执行；其他被忽略。
					如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
					否则：
					如果有default子句，则执行该语句。
					如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
				select的知识点小结如下：
					select语句只能用于信道的读写操作
					select中的case条件(非阻塞)是并发执行的，select会选择先操作成功的那个case条件去执行，如果多个同时返回，则随机选择一个执行，此时将无法保证执行顺序。对于阻塞的case语句会直到其中有信道可以操作，如果有多个信道可操作，会随机选择其中一个 case 执行
					对于case条件语句中，如果存在信道值为nil的读写操作，则该分支将被忽略，可以理解为从select语句中删除了这个case语句
					如果有超时条件语句，判断逻辑为如果在这个时间段内一直没有满足条件的case,则执行这个超时case。如果此段时间内出现了可操作的case,则直接执行这个case。一般用超时语句代替了default语句
					对于空的select{}，会引起死锁
					对于for中的select{}, 也有可能会引起cpu占用过高的问题
*/

func conditionalStatement() {
	/* 定义局部变量 */
	var a int = 4
	if a < 5 {
		fmt.Printf("a 小于 5 \n")
	}
	if a < 4 {
		fmt.Printf("a 小于 4 \n")
	} else {
		fmt.Printf("a 不小于 4 \n")
	}

	var b int = 5
	if a < 5 {
		if b == 5 {
			fmt.Printf("a 的值为 4 ， b 的值为 5\n")
		}
	}
	fmt.Printf("a 的值为 : %d\n", a)
	fmt.Printf("b 的值为 : %d\n", b)

	var grade string = "B"
	var marks int = 90
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)

	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}

	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
