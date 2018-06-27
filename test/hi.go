// 程序的入口是包main
package main // 告诉我们当前文件属于哪个包, 除了main包之外,其他包都会生成*.a文件
// 同一个package可以直接使用,不用import

// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md
// http://www.vaikan.com/go/a-tour-of-go/#1
// 导入包fmt, os
import (
	// 点操作(.省略包名),可以直接使用Println(fmt.Println)
	// 别名操作(命名成容易记忆的名字),f.Println(fmt.Println)
	// _操作(引入该包,不适用该包函数,只是只是调用init函数)
	"fmt"
	// "os"
	"../util" // 一般不推荐这么做
	"math"
	"strconv"
)

// go get动态获取远程代码包,实际上是,1.git clone,2.go install
// go run hi.go执行,编译一个或多个,生成最终可执行文件,只能用于main包
// 或者会报错go run: cannot run non-main package
// go install,1.会生成可执行文件,2.会把编译的结果移到$GOPATH/bin,不可以直接编译一个go file
// 使用go fmt hi.go格式化代码
// gofmt -w src格式化整个项目
// go build -o filminfo 生成filminfo可执行文件

// go基本类型
// _占位符(python中pass)
// 变量,常量
// 字符串类型,string,utf-8编码
// 数字类型,int,float
// bool类型, true,false
// 派生类型: 指针,结构体,函数,array(list),slice,map(dict),interface,channel
// 算术运算符: +,-,*,/,%,++,--
// 关系运算符: ==,!=,>,<,>=,<=
// 逻辑运算符: && and,|| or,! not
// &(返回变量内存地址),*指针变量
// &生成一个指向变量对象的指针
// *可以访问指针指向的对象(取对象值)(取指针的值)

// new,type,make

// const定义常量
// (枚举)常量,关键字iota从0按行自增枚举值
// 如果枚举被打断,需显性恢复
// 被打断的枚举,其计数仍会增加
// 常量一般被省略
const (
	Sunday = iota // iota = 0, 相当于重置为0
	Monday        // 省略常量时,表示以前一个常量相同iota += 1
	Thesday
	Wednesday // 当iota被打断时,需显性恢复
	Thursday
	Friday
	Saturday
)

type User struct {
	id   int
	name string
}

// 结构体(struct)就是一个字段的集合
// 面向对象?(继承)
type Sex struct {
	User   // struct匿名字段(嵌入字段),可以继承User下字段名
	male   int
	female int
}

func (s *Sex) say() {
	fmt.Println(s)
}

type learner interface {
	say()
}

type Printer interface {
	Print()
}

// go中有两个保留函数init,main
// 在定义时不能有任何参数和返回值
// go会自动调用init和main
// 每个package中init函数是可选,但package main必须包含main函数
func main() {
	// 结构体
	// 初始化
	v := Sex{User{1, "august"}, 0, 1}
	// 结构体可以通过.访问
	// 结构体可以通过指针访问,并修改
	v.male = 22
	// 当一个指针被定义后没有分配到任何变量,她的值为nil
	// var ptr *int
	// &可以取得内存地址
	q := &v // 是一个内存地址,相当于python的引用
	q.male = 33
	t := *q        // equal t = v
	fmt.Println(q) // q是获取到的变量地址
	fmt.Println(t) // t是通过指针获取的变量值
	name := v.User.name
	fmt.Println(name)   // 访问嵌入字段
	v.User.name = "tom" // 嵌入字段修改
	fmt.Println(v.User.name)
	v.say() // 结构体会继承方法
	util.GetMembers(v)

	// 接口(interface)
	var a interface{}
	// 空interface可以存储任意值
	// 匿名interface(嵌入interface)
	a = "hi"
	fmt.Println(a, 222)
	a = 5
	fmt.Println(a, 2222)
	// tt := a.(type) .(type)只能使用在switch语句中
	// fmt.Println(tt, 33333)
	var is learner // is的类型是learner接口
	is = &v        // ?具体什么意思
	is.say()

	// var定义变量
	var str string = "hi, 世界"
	// 字符串可以使用索引访问具体字节str[2]
	// 不能用序号获取元素指针&str[2]
	// 字符串是不可变类型
	// 字节数组不包括NULL
	//字符串一般由双引号("")或者反引号(``)定义
	cc := []rune(str)
	cc[1] = 'a' // 字符串转化为rune来修改
	str5 := string(cc)
	fmt.Println(str5, 233333)
	fmt.Println("Pre", str, math.Pi)
	fmt.Println(add(22, 33))
	// 用:=在明确类型的地方,用于替代var的定义,:=不能使用在函数外
	// 函数参数,和返回值也是已经明确定义类型
	// 不用重复进行:=,否则会报错
	a, b := swap("hi", "word!")
	fmt.Println(a, b)
	var x, y, z int // 多个变量赋值,零值为0
	// var c, python, ruby = true, false, "no"
	// 2.
	// 未使用的局部变量,编译器会报错,可以使用占位符来规避
	// 连接字符串时,+必须在上一行末尾,否则会报错
	var str2 = "good" +
		"work"
	// 指定不转义字符串
	str3 := `
    line1
    line2
    line3
    `
	_ = str3
	// 要修改字符串,必须转换成[]rune,[]byte
	// 使用中文的话用rune进行转化
	ustr := []rune(str)
	// go中双引号代表一个字符串,单引号代表一个字节(char)
	// 一般可以省略;
	for _, r := range str {
		// 打印出格式化的字符串
		fmt.Printf("char format %c,", r) // 格式化任意类型%v,注意可能会出现不希望得到的值
	}
	ustr[1] = 'e'
	fmt.Println(string(ustr))
	_ = str2
	c, python, ruby := true, false, "no"
	fmt.Println(x, y, z, c, python, ruby)
	// 常量Pi声明,常量可以为字符串,数字,布尔
	// 常量组
	// 未使用的常量不会引发编译错误
	const Pi = 3.14
	fmt.Println(Pi)
	fmt.Println(Sunday, "2333")
	// 声明这是一个指针
	var ptr *int
	var ss int = 123
	ptr = &ss // 可以通过指针读取
	// 分别为变量值,变量内存地址,*透过指针访问到目标对象,可以用.访问目标成员
	// 指针对象不能运算,可以转化uintptr变相进行运算
	fmt.Println(ss, ptr, *ptr)
	sum := 0
	// 循环
	for i := 0; i < 5; i++ {
		sum += i
		fmt.Println(sum)
	}
	sum = 1
	fmt.Println(sum, "check sum value")
	// 可以省去分号
	for sum <= 10 {
		sum += sum
		if sum > 10 {
			fmt.Println("sum big 10 is", sum)
		} else {
			fmt.Println("sum less 10 is", sum)
		}
	}
	// 数组,长度和类型都不可缺少
	// go中的数组(list)是按值传递,开销大,我们一般使用切片,传递切片成本低
	// 或者使用数组的指针
	var li = [5]int{0, 1, 2, 3, 4}
	var li2 = [3]string{"a", "b", "c"}
	// slice是对列表(array)的封装,比数组更加灵活,强大且方便
	// 用len取slice的长度,用cap取slice容量
	// 数组的len和cap永远都是相等的
	// slice是引用类型,不管append操作或者赋值操作都会影响引用的源数组,
	// slice都是指向其引用地址
	// slice引用传递会发生意外,但超过其容量时,系统会开辟新的内存空间,内存地址会发生变化
	// slice零值为nil,其长度和容量都为0
	// 创建slice,1.从数组中构建slice
	s := li[1:]
	s2 := [...]string{"august", "tom", "jack"} // 省略长度,go会自动计算长度
	fmt.Print(s2)
	// make和new都是用于分配内存,
	// new只分配内存,返回指针
	// make用于slic,map,chan的初始化
	// new将其置零,对于bool类型为false,int零值为0,string零值为空字符串,rune空值为0,byte为0x0
	// make返回初始化后的值
	// slice,map,chan,interface,func为nil
	// 可以用make创建slice,
	// 2. 使用make函数定义
	// s3创建了一个长度为3,容量为5的slice,第三参数可选,默认与长度相等
	s3 := make([]int, 3, 5)
	s3 = append(s3, 22) // 向slice添加元素
	fmt.Println(s3)
	// 遍历列表和slice
	for index, value := range s2 {
		fmt.Println(index, value)
	}
	fmt.Println(li)
	fmt.Println(li2)
	fmt.Println(s)
	fmt.Println(s2)
	// 词典(map)(dict)
	// map没有初始化,go会自动创建一个nil map,int类似,没有初始化会给自动赋值0
	// 包含slice的不可以作为map的键,否则会编译错误
	dict := make(map[string]int)
	dict["name"] = 233 // 创建dict,如果已经存在则更新
	dict["age"] = 25
	delete(dict, "name")      // 可以用delete(dict, "name")删除
	age, exist := dict["age"] // 取出dict中age,赋值给age,exist表示键值是否存在
	fmt.Println(len(dict))    // 返回dict个数, 也可以对slice,list,string个数计算
	fmt.Println(len(s2), len(str))
	// 不能将其他类型当bool值使用
	if exist {
		fmt.Println(age)
	}
	fmt.Println(dict)
	// 遍历词典
	for key, value := range dict {
		fmt.Println(key, value)
	}
	fmt.Println(dict)
	// 接口(interface),相当于python的类,是一个或者多个方法的集合
	// interface可以存储类型的值,任意的函数参数
	var t2 Printer = &User{1, "august"}
	t2.Print()
	fmt.Println(max(3, 8))
	var username string = "august"
	// 错误处理
	// 转化为int
	// 没有err会报错 multiple-value
	i, err := strconv.Atoi(username)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	// 转为字符串
	bb := strconv.Itoa(ss)
	fmt.Println(bb)

	// todo go, chan
	// 网络和并发是go语言的两大feature
	// 在调用函数前添加go就可以创建并发单元(和yield类似)
	// goroutine是轻量级实现的并发
}

// 该String函数属于User type类型对象的方法
func (self *User) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}
func (self *User) Print() {
	fmt.Println(self.String())
}

// 函数
// 匿名函数, 任意类型的不定参数
// 判断较大值
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// 变量类型在变量名之后
// x int, y int, 定义返回值的类型
func add(x, y int) int {
	return x + y
}

// 前面是变量名类型声明,后面是返回变量名声明
func swap(x, y string) (string, string) {
	return y, x
}
