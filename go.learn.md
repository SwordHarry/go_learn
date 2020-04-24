# Go

## 第一章 入门
### Goroutine
### tcmalloc
### 垃圾回收：gc 堪用，离好用还有一段距离

### 用 go 开发的项目
- Docker
- Kubernetes
- etcd
- beego
- martini
- codis
- delve

### go 环境搭建

gocode: 代码补全
godef: 跳转到代码对应定义
guru: guru原来叫做oracle，能获得全部代码对应引用
gorename: 重命名Go源码文件
goimports: 自动帮你格式化import语句的，它来帮你决定import顺序

golangci-lint: 旧的用法是gometalinter，现在应该选择快5倍的golangci-lint，这个工具用来做静态检查工具，它会反馈代码风格并给出建议，这样就可以在源码保存是运行它能告诉你现有代码中有什么问题
go-outline: 当前文件做符号(Symbol)搜索，相当于文件大纲
go-symbols: 工作区符号搜索
gogetdoc: 鼠标悬停时可以显示文档(方法和类的签名等帮助信息)
gomodifytags: 提供tags管理，可以对struct的tag增删改
gotests: 用例测试
impl: 自动生成接口实现代码(stubs)
gopkgs: 列出Go包。

 【待下载】
fillstruct。根据结构体定义在使用时自动填充默认值


### go 命令
#### go build
如果不存在 main 包，则只对当前目录下的程序源码进行语法检查，不会生成可执行文件。
#### go run



***注意：Go 没有三目运算符，所以不支持* ?: *形式的条件判断。***

### 常量
常量必须是数字、字符串、布尔值

## 第二章 程序结构

## 可见性
实体在函数中声明，则只在函数局部有效
声明在函数外，对包里的所有源文件可见
实体第一个字母的大小写决定其可见性是否可跨包
若为大写，则是导出的

包名总是由小写字母组成

### 2.2 声明
- var
- const 
- type
- func

```go
func 函数名(参数列表) 返回值列表? {
    函数体
}
```

#### 短变量声明
`name := expression`
`:=` 表示声明；
`=` 表示赋值

短声明最少声明一个新变量


### 2.3.2 指针

### 2.3.3 new函数

new(T) 创建一个未命名的T类型变量，初始化为T类型的零值，并返回其地址（地址类型为*T）

### 2.3.4 变量的生命周期

包级别变量的生命周期是整个程序的执行时间

局部变量是动态的生命周期：执行语句时创建，直到不可访问

编译器选择使用堆或栈的空间分配，但这个选择不是基于使用 var 或 new 关键字来声明变量

#### 变量逃逸
```go
var global *int
func f() {
    var x int
    x = 1
    global = &x
}
```
x 变量使用了堆空间

避免在长生命周期对象中保持短生命周期对象不必要的指针

### 2.4.1多重赋值

```go
// 交换两个变量的值
x, y = y, x

// 最大公约数
func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}

// 斐波那契数列的第 n 个数
func fib(n int) int {
    x, y := 0, 1
    for i := 0;i < n;i++ {
        x, y = y, x+y
    }
    return x
}
```

### 2.5 类型声明
type 声明定义一个新的命名类型，它和某个已有类型使用同样的底层类型。
类型的声明通常出现在包级别，这里的命名类型在整个包中可见，如果名字是导出的（开头使用大写字母），其他的包也可以访问它

不同的类型之间需要 **显式类型转换**
对于每个类型T，都有一个对应的类型转换操作`T(x)`将值x转换为类型T
    当且仅当两个类型具有相同的底层类型或二者都是指向相同底层类型变量的未命名指针类型
数字类型间的转换，字符串和一些slice类型间的转换会改变值的表达方式

#### 改变类型的显示方式
```go
type Celsius float64
func (c Celsius) String() string {
    return fmt.Springtf("%gC", c)
}
```


### 2.6 包和文件
建议使用 go mod 方式构建包和导入包
https://segmentfault.com/q/1010000020696591/
此处的 replace 稍微介绍下，之所以要是 github.com/article 的格式，是因为在 go1.13 中， go module 名称规范要求路径的第一部分必须满足域名规范，否则可能汇报类似 malformed module path "article": missing dot in first path element 这样的错误。当然，在 go1.12 不会有报这个错误。建议的话，如果是公司内部使用，可以替换成公司内部域名。

#### 2.6.2 包初始化

##### 变量的初始化
包初始化从包级别的变量开始，先根据声明顺序初始化，但也根据依赖顺序初始化
若包由多个.go文件组成，初始化按照编译器收到的文件顺序进行，go工具会在调用编译器前将.go进行排序

除了简单赋值，还有`init`函数进行初始化，任何文件可以包含任意数量的声明如下的函数：
```go
func init() { /* ... */ }
```
init 函数不能被调用和引用

##### 包初始化
按照依赖顺序优先

预计算

### 2.7 作用域

声明的作用域指用到声明时名字的源代码段

词法块
包含了全部源代码的词法块，叫全局块

## 第三章 基本数据*

Go的数据类型有：
- 基础类型
    - 数字
    - 字符串
    - 布尔
- 聚合类型
    - 数组
    - 结构体
- 引用类型：间接指向程序变量或状态
    - 指针
    - slice
    - map
    - 函数
    - channel通道
- 接口类型


### 3.1 整数
按位分类：int8、int16、int32(rune)、int64
指最大值的范围
无符号：uint8(byte)、uint16、uint32、uint64
int ,uint
大小与原生相同，或等于平台上运行效率最高的值

uintptr ，大小不明确，但可完整存放指针

### 3.2 浮点数
float32
float64

### 3.3 复数
complex64，complex128，二者分别由float32和float64构成
相关函数有：
complex 函数根据给定的实部和虚部创建复数
real 提取复数的实部
imag 提取复数的虚部


math/cmplx 包提供了复数运算的库函数，如复数的平方根函数和幂函数


### 3.4 布尔值
布尔值无法隐式转换成数值

### 3.5 字符串
子串生成操作`s[i:j]`[i, j)，结果的大小是j-i个字节
i和j的默认值分别是0和`len(s)`

虽然可以将新值赋值给字符串变量，但是字符串值无法改变，故字符串内部数据无法修改
#### 为什么这样设计
不可变意味着两个字符串能安全地共用同一段底层内存，使得赋值任何长度字符串的开销都低廉

#### 3.5.4 字符串和字节slice
4个标准包用于操作字符串：bytes strings strconv unicode

##### strings
用于搜索、替换、比较、修整、切分与连接字符串
##### bytes
用于操作字节slice([]byte 类型)
由于字符串不可变，因此按照增量方式构建字符串会导致多次内存分配和复制。因此，使用 byte.Buffer 类型会更高效
##### strconv
用于在 布尔值、整数、浮点数 和 与之对应的字符串 之间的互相转换



##### unicode
判别文字符号值特性的函数，如是否是字母，大小写等

##### path 和 path/filepath
path 处理 url
filepath 根据宿主平台的规则处理文件名

##### bytes包和strings包
避免转换和不必要的内存分配，有对应的实用函数
- Contains
- Count
- Fields
- HasPrefix
- Index
- Join

##### bytes.Buffer
Buffer 起初为空，其大小随着各种类型数据的写入而增长

##### 字符串反转
现将string 变为 rune[]
```go
// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
```

### 3.6 常量
所有常量本质上都属于基本类型：布尔、字符串或数字
其值在编译时确定

若同时声明一组常量，除了第一项之外，其他项在等号右侧的表达式都可以省略，这会服用前一项的表达式及其类型
```go
const (
    a = 1
    b
    c = 1
    d
)

fmt.Println(a, b, c, d) // "1 1 2 2"
```

#### 3.6.1 常量生成器 iota
常量声明中，iota 从 0 开始取值，逐项加1
```go
type Weekday int
const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```
iota 也有局限，因为不存在指数运算符，所以无从生成1000的幂
#### 3.62 无类型常量
没有指明类型的常量并不从属于某一具体类型
借助推迟确定从属类型，无类型常量不仅能暂时维持更高的精度，并能写入更多表达式而无需转换类型
从属类型待定的常量共有6种，分别是：
- 无类型布尔
- 无类型整数
- 无类型文字符号
- 无类型浮点数
- 无类型复数
- 无类型字符串

变量声明，包括短变量声明中，假如没有显式指定类型，无类型常量会隐式转换成变量的默认类型。
```go
i := 0 // 无类型整数；隐式
r := '\000' // 无类型文字字符；隐式 rune('\000)
f := 0.0 // 无类型浮点数；隐式 float64(0.0)
c := 0i // 无类型复数；隐式 complex128(0i)
```

只有大小不确定的int类型，不存在大小不确定的 float 类型和 complex 类型
原因：如果浮点数据的大小不明，就很难写出正确的数值算法

## 第四章 复合数据类型
四种复合数据类型：
- 数组
- slice
- map
- 结构体

其中，数组和结构体都是聚合类型，其长度固定
slice 和 map 是动态数据结构，长度可动态增长

### 4.1 数组
省略号...出现在数组长度的位置，则长度由初始化数组的元素个数决定。
```go
q := [...]int {1,2,3}
```
除了按顺序给出一组值，还可以按索引位置给值，没有指定值的将补上零值
```go
r := [...]int{99: -1}
```
定义了一个拥有 100 个元素的数组r，除了最后一个元素是 -1 之外，其余都是0

#### 数组的比较
数组是可比较的，比较的结果是两边元素的值是否完全相同，比较方式有 == 和 !=
```go
a := [2]int{1,2}
b := [...]int{1,2}
c := [2]int{1,3}
fmt.Println(a == b, a == c, b == c) // true false false
d := [3]int{1,2}
fmt.Println(a == d) // 编译错误，[2]int 不能与 [3]int 比较
```

#### 对于数组的参数传递
当调用一个函数时，每个传入的参数都会创建一个副本，然后赋值给对应的函数变量，所以函数接受的是一个副本，而不是原始参数。
Go 把数组和其他类型都看成值传递，如果要使用引用传递，则要显式传递一个数组指针。
```go
func zero(ptr *[32]byte) {
    *ptr = [32]byte{}
}
```

### 4.2 slice
slice表示有相同类型元素的可变长度的序列。
slice写成 `[]T`，其中元素的类型都是T
slice可以用来访问数组的部分或全部元素，即底层为数组，有三个属性，指针、长度和容量
指针指向数组的第一个可以从slice中访问的元素，这个元素并不一定是数组的第一个元素
长度是指 slice 中的元素个数，它不能超过 slice 的容量。
容量大小通常是从 slice 的起始元素到底层数组的最后一个元素间元素的个数。

Go内置函数 len 和 cap 用来返回 slice 的长度和容量。

一个底层数组可以对应多个 slice，这些 slice 可以引用数组的任何位置，彼此可重叠



书 P 64 图 4-1

slice 操作符 `s[i:j]`（其中 0 <= i <= j <= cap(s)）创建一个新的 slice，从 i 到 j-1。这里的 s 可以是数组或指向数组的指针或 slice。语法和 python 类似

如果 slice 的引用超过了被引用对象的容量，即 `cap(s)`，则会宕机；但是如果超出的是被引用对象的长度，即 `len(s)`，则最终 slice 会比原 slice 长。

```go
	arr := []string{0: "a", 1: "b", 2: "c", 3: "d", 4: "e", 5: "f", 6: "g", 7: "h", 8: "i", 9: "j", 10: "k", 11: "l", 12: "m", 13: "n"}
	s := arr[4:7]
	s2 := arr[6:9]
	fmt.Println(s) // [e f g]
	fmt.Println((s2)) // [g h i]
	s3 := s2[:5]
	fmt.Println(s3) // [g h i j k] 6 + 5 = 11，[6:11]
```



##### slice 和 子串操作

两者语法相似，底层引用方式相同，均消耗常量时间



Slice 包含了指向数组元素的指针，所以将一个 slice 传递给函数时，可以在函数内部修改底层数组的元素。

创建一个 slice 等于为数组创建了一个别名



和数组不同的是，slice 无法做比较，因此不能用==来测试两个slice是否拥有相同元素。

标准库中提供高度优化的函数 bytes.Equal 来比较两个字节 slice([]byte)。但是其他 slice，必须自己写函数来比较。



##### 关于 slice 不能使用 == 操作符

首先，slice的元素不是直接的，如果底层数组元素改变，同一个slice在不同时间会拥有不同的元素。

slice需要深度比较，所以不能用slice作map的键。

对于引用类型，==检查的是引用相等性。若==操作符对slice和数组的行为不一致，会带来困扰。

Slice 唯一允许的比较操作是和 nil 做比较

slice的零值是 nil。值为 nil 的 slice 没有对应的底层数组，长度和容量都是0

但也有非 nil 的 slice 长度和容量是零，如 `[]int{}`



**检查一个 slice 是否为空，使用 `len(s) == 0`**

内置 make 函数可以创建一个无名数组并返回了它的一个 slice

```go
make([]T, len [,cap]) // 和 make([]T, cap)[:len] 功能相同
```

Cap 省略的情况下，len 和 cap 相同

#### 4.2.1 append 函数

内置 append 函数用来将元素追加到 slice 的后面

内置的 append 函数使用了复杂的增长策略。

- 不清楚会不会导致一次新的内存分配
- 不能假设原始 slice 和 append 后的结果指向同一个底层数组
- 不能证明它们就指向不同的底层数组
- 无法假设对旧 slice 的操作是否会影响新的 slice 元素

所以，通常将 append 的调用再次赋值给传入 append 的函数的 slice

Slice 并不是纯引用类型，更像是聚合类型：

```go
type IntSlice struct {
  ptr *int
  len, cap int
}
```

任何有可能改变 slice 的长度或容量，或是使得 slice 指向不同的底层数组，都应该更新 slice 变量

Append 函数可以同时给 slice 添加多个元素，甚至添加另一个 slice

`func(slice []Type, elems ...Type) []Type`

### 4.3 map

map[k]v， 内置函数make 可以创建map，也可以用字面量创建带初始化键值对元素的字典

```go
ages := make(map[string]int)
ages := map[string]int {}

// 访问
ages["alice"] = 32
fmt.Println(ages["alice"]) // 32
```

可以用内置函数 delete 从字典中根据键移除一个元素：

```go
delete(ages, "alice") // 移除元素 ages["alice"]
```

若键不在map中，求值操作也是安全地，对应元素不存在，就返回值类型的零值

所以通过下标的方式访问map中的元素总是会有值

```go
ages["bob"] = ages["bob"] + 1 // 0 + 1
```

但是我们无法获取map 元素的地址，因为 map 的增长可能会导致已有元素被重新散列到新的存储位置，这样就可能使得获取的地址无效。

for range 循环会遍历键值对，但迭代顺序不固定

向零值map中设置元素会导致错误。

```go
var ages map[string]int
ages["carol"] = 21 // 宕机
```

故设置元素之前，必须初始化 map



其实通过下标方式访问 map 中的元素输出两个值，第二个值是一个布尔值，用来报告该元素是否存在，消除二义性。

布尔值一般叫作 ok 

```go
if age, ok := ages["bob"]; !ok {
  
}
```

和 slice 一样， map 不可比较，唯一合法的比较也是与 nil 比较。

注意如何使用 `!ok` 来区分“元素不存在”和“元素存在但值为零” 的情况



##### 将 slice 用作键

若一定要用 slice 用作键，可以编写一个转换函数k，将 slice 转变为可比较的如 string

```go

func k(list []string) string {
	return fmt.Sprintf("q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}

```



### 4.4 结构体

```go
type Employee struct {
  ID int
  Name string
  DoB time.Time
  Salary int
  ManagerID int
}
```

成员通过点号访问

```go
var dilbert Employee
dilbert.Salary = 25000
```

可以获取成员变量的地址，然后通过指针访问

点号也可以用在结构体指针上：

```go
var employeeOfTheMonth *Employee = &dilbert
employeeOfTheMonth.Position += " (proactive team player)"
```



##### 结构体与结构体指针

来看一个函数例子

```go
func EmployeeByID(id int) *Employee { /* ... */ }
EmployeeByID(0).Salary = 0
```

若函数 `EmployeeByID` 返回的是 `Employee` 而不是 `*Employee`，会编译错误，因为复制表达式的左侧无法识别出一个变量。原因是每次调用这个函数都新建了一个对象。



结构体成员通常一行一个，名称在前，类型在后，但相同类型的连续成员变量可以写在一行上。

成员变量的顺序对于结构体同一性很重要，交换属性的顺序，就定义了一个不同的结构体类型

成员变量名称大写，则可导出，否则不能访问和赋值

匿名结构体类型

命名结构体类型 s 不可以定义一个拥有相同结构体类型s的成员变量，即聚合类型不可以包含自己。但是s中可以定义一个s的指针类型，即*s。如此就可以定义递归数据类型，如链表和树。

结构体的零值由结构体成员的零值组成



### 4.4.1 结构体字面量

结构体字面量

- 按照正确顺序，依次指定值，仅用于变量顺序约定的小结构体
- 指定部分或全部成员变量的名称和值来初始化

两种初始化方式不能混用

结构体类型的值可以作为参数传递给函数或者作为函数的返回值。如下函数将Point缩放了一个比率：

```go
func Scale(p Point, factor int) Point {
  return Point{p.X * factor, p.Y * factor}
}
```

出于效率的考虑，大型结构体通常都使用结构体指针的方式直接传递给函数或者从函数中返回。

通常结构体都通过指针的方式使用，因此可以使用一种简单的方式来创建、初始化一个 struct 类型的变量并获取其地址：

```go
pp := &Point{1,2}
// 等价于
pp := new(Point)
*pp := Point{1,2}
```

### 4.4.2 结构体比较

若结构体的所有成员变量都可以比较，则该结构体可比较，可以使用 == 或 !=

可比较的结构体可以作为map的键类型

### 4.4.3 结构体嵌套和匿名成员

Go允许我们定义不带名称的结构体成员，只需要指定类型即可；称为匿名成员。这个结构体类型必须是一个命名类型或者指向命名类型的指针

```go
type Point struct {
  X,Y int
}

type Circle struct {
  Point
  Radius int
}

type Wheel struct {
  Circle
  Spokes int
}

var w Wheel
w.X = 8 // 等价于 w.Circle.Point.X = 8

// 初始化
w = Wheel{
  Circle: Circle{
    Point: Point{X: 8, Y:8},
    Radius: 5,
  },
  Spokes: 20, // 注意尾部的逗号是必须的
}
```

结构体嵌套的使用，为语法糖

因为匿名成员拥有隐式名字，所以不能在一个结构体里面定义两个相同类型的匿名成员

匿名成员的可导出性由类型决定

组合是Go面向对象编程方式的核心

### 4.5 JSON

标准库 encoding/json

#### 成员标签

```go
type Movie struct {
  Title string
  Year int `json:"released"`
  Color bool `json:"color,omitempty"`
  Actors []string
}
```

#### Marshal

数据结构转换为JSON为 Marshal

MarshalIndent 整齐格式化，定义每行的前缀字符串，定义缩进字符串

```go
data, err := json.MarshalIndent(movies, "", "    ")
```



**注意：**只有可导出的成员可以转换为JSON字段

成员标签相当于JSON别名

键json控制encoding/json 的行为。

Omitempty 表示成员的值为零值或空，则不输出到JSON



#### Unmarshal

Json到Go数据结构

通过合理定义Go的数据结构，可以选择哪部分JSON数据解码到结构体对象中，哪些数据可以丢弃。

在 unMarshal 过程中是忽略大小写的