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
x 变量使用了堆空间，x从f中逃逸了

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

## 第三章 基本数据

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

rune 指unicode码

byte 指ascii 码

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
    c = 2
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

注意：map中的元素并不是一个变量，所以我们不能对map的元素进行取址操作

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



### 4.6 文本和HTML模板

双大括号指定输出，支持循环，管道

Text/template和html/template包，后者可以自动转义



## 第5章 函数

```go
func name(parameter-list) (result-list) {
  body
}
```

注意是值传递，则若传递slice进行递归，它并不会修改调用者原来传递的元素，所以当被调函数返回时，调用者的栈依旧保持原样。

Go语言针对递归使用了可变长度的栈，可达到1G左右

### 5.3 多返回值

```go
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}
```

保证resp.Body正确关闭使得网络资源正常释放。即使在发生错误的情况下也必须释放资源。Go语言的垃圾回收机制将回收未使用的内存，但不能指望它会释放未使用的操作系统资源，比如打开的文件和网络连接，必须显式关闭。

可以给返回值分别命名，尤其在一个函数返回多个结果且类型相同时，名字的选择很重要，如：

```go
func Size(rect image.Rectangle) (width, height int)
```

#### 裸返回

应保守使用

### 5.4 错误

如果错误只有一种情况，通常作为最后一个结果，返回布尔类型

对于IO操作，如果错误原因多样，则错误结果是error

error 是内置的接口类型

尽管Go有异常机制，但Go语言的异常只是针对程序bug导致的预期外的结果。

#### 5.4.1 错误处理策略

1. 将错误传递下去
   - 直接传递
   - 构建一个新的错误信息
   - fmt.Errorf 使用 fmt.Sprintf 函数格式化一条错误消息并且返回一个新的错误值
2. 对于不固定或不可预测的错误，进行重试，超出一定时间或者次数后再报错退出
3. 主程序部分：输出错误，停止程序。库函数应当将错误传递给调用者。log.Fatalf 将时间和日期作为前缀添加到错误信息前；或自定义log包的前缀，并将日期和时间略去
4. 只记录下错误信息然后程序继续运行，用log包添加日志前缀
5. 忽略

```go
func mayBeError() {
  data, error := handle()
  if error != nil {
    //...
  }
  
  // func-body
  return
}
```

#### 5.4.2 文件结束标识

io.EOF

### 5.5 函数变量

函数类型的零值是 nil，且本身不可比较，不可作为键出现在map中。

函数可以作为参数或者返回值使用

### 5.6 匿名函数

```go
strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
```

闭包

变量的生命周期不是由它的作用域所决定的

当匿名函数进行递归时，必须先声明一个变量对匿名函数进行赋值。如果两个步骤合并成一个声明，函数字面量将不能存在于变量的作用域中，就不能递归调用自己了

```go
visitAll := func(items []string) {
  // ...
  visitAll(m[item]) // undefined: visitAll
  // ...
}
```

#### 捕获迭代变量

类比js的循环闭包，go也有同样的问题

### 5.7 变长函数

可变的参数个数，用省略号表示，并且只能是最后一个参数

```go
func sum(vals ...int) int {
  total := 0
  for _, val := range vals {
    total += val
  }
  return total
}
```

Vals 是一个int类型的 slice

调用：

```go
values := []int {1, 2, 3, 4}
sum(values...)
```

**注意：**尽管 ...int 参数就像函数体内的 slice，但变长函数的类型和一个带有普通 slice 参数的函数类型不相同

`interface {}`意味着参数可以接受任何值

#### 总结 变长函数

- 省略号表示
- 会被转换成 slice

### 5.8 延迟函数调用

#### defer 机制

Defer 语句通常用于成对的操作，如打开关闭，链接断开，加锁解锁，即使是再复杂的控制流，资源在任何情况下都能正确释放。

实际调用推迟到包含defer 语句的函数结束后才执行，defer 没有限制使用次数，执行的时候调用defer 语句顺序为**倒序**进行。

延迟执行的匿名函数甚至能改变外层函数返回给调用者的结果：

```go
func triple(x int) (result int) {
  defer func() { result += x }
  return x + x
}

fmt.Println(triple(4)) // 12
```

#### 文件描述符

一个 Linux 进程可以打开成百上千个文件，为了表示和区分已经打开的文件，Linux 会给每个文件分配一个编号（一个 ID），这个编号就是一个整数，被称为文件描述符（File Descriptor）。

循环打开文件但是循环defer，有可能会耗尽文件描述符

一种解决方式是将循环体（包括defer语句）放到另一个函数里，在循环迭代后都会调用文件关闭函数。

#### 总结 defer

- 当defer被声明时，其参数就会被实时解析

  ```go
  func a() {
    i := 0
    defer fmt.Println(i) //输出0，因为i此时就是0
    i++
    defer fmt.Println(i) //输出1，因为i此时就是1
    return
  }
  ```

- defer 的顺序为先进后出

- defer即`在return执行完，函数退出前执行`。所以defer执行时可修改函数返回值（不推荐这么做）。

### 5.9 宕机

内置函数 panic 可以主动发生宕机

宕机发生时，正常程序执行会终止， goroutine 中的所有延迟函数会执行，程序异常退出并留下一条日志消息。

日志消息包括宕机的值，代表某种错误，goroutine 则在宕机时显示一个函数调用的栈跟踪消息。

只在发生严重错误时才会使用宕机

Defer 会在执行的函数在栈清理前调用

### 5.10 恢复

recover 函数在延迟函数内部调用，且包含该defer 语句的函数发生宕机，recover会终止当前的宕机状态并返回宕机的值。函数会从之前宕机的地方继续运行并正常返回。

如果 recover 在其他任何情况下运行则它没有任何效果且返回 nil

不应该随意恢复宕机

当一个 panic 被恢复后，调度并因此中断，会重新进入调度循环，进而继续执行 recover 后面的代码， 包括比 recover 更早的 defer（因为已经执行过得 defer 已经被释放，而尚未执行的 defer 仍在 goroutine 的 defer 链表中）， 或者 recover 所在函数的调用方。

```go
package main

import "fmt"

func main() {
	a := returnN()
	fmt.Println(a)
}

func returnN() (result int) {
	defer func() {
		if p := recover(); p != nil {
			result = p.(int)
			fmt.Println(p)
		}
	}()
	panic(3)
}
```

Panic 的值会成为 recover 的返回值

## 第 6 章 方法

封装和组合

方法声明在函数名字前多了一个参数，该参数把这个方法绑定到该参数对应的类型上

```go
func (p Point) Distance(q Point) float64 {
  return math.Hypot(q.X - p.X, q.Y - p.Y)
}

p.Distance(q)
```

附加的参数p为方法的接收者

p.Distance 成为**选择子**，它为接收者p 选择合适的 Distance 方法

Go 可以将方法绑定到任何类型上，如

```go
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
```

只要其类型不是指针和接口，则任何类型都可以声明方法

方法命名一般比函数更简短

### 6.2 指针接收者的方法

主调函数会复制每一个实参变量，如果函数需要更新变量，就必须使用指针来传递变量的地址

```go
func (p *Point) ScaleBy(factor float64) {
  p.X *= factor
  p.Y *= factor
}
```

习惯上都使用指针接收者。为了防止混淆，不允许本身是指针类型进行方法声明

```go
p := Point{1, 2}
(&p).ScaleBy(2)
fmt.Println(p) // {2, 4}
```

实际上方法调用可以简写成

```go
p.ScaleBy(2)
```

编译器会对变量进行 &p 的隐式转换。

故不能够对一个不能取地址的Point 接收者参数调用 *Point 方法，因为无法获取临时变量的地址

#### 接收者隐式转换机制

实际上，Go 编译器有隐式转换机制：

1. 实参和形参都是 *T 或者 T，不发生转换
2. 实参是T，形参是 *T，编译器会隐式获取变量地址`p -> (&p)`
3. 实参是*T，形参是 T，编译器会隐式解引用求值`p -> *p`

但无法获取临时变量的地址

```go
func main() {
	p := Point{1,2}
	p.ScaleBy(2)
	//等价于
	//(&p).ScaleBy(2)
	fmt.Println(p)
	pptr := &Point{3,3}
	pptr.ScaleBy(2)
	// 等价于
	//(*pptr).ScaleBy(2)
	fmt.Println(*pptr)

	//Point{1,2}.ScaleBy(2) // 临时变量T 无法 变成 *T
	//pptr.ScaleBy2(2) // *T 可以变成 T
}
```

### 6.3 结构体内嵌组成类型

结构体内嵌不仅能复用类型，还能复用方法

```go
type ColoredPoint struct {
  Point
  Color color.RGBA
}
```

Point 的方法都被纳入到 ColoredPoint类型中

ColoredPoint并不是Point，但是它包含一个Point

内嵌的字段会告诉编译器生成额外的包装方法来调用Point声明的方法

但是，当形参是Point的时候，必须显式使用它

```go
p.Distance(q.Point)
```

### 6.4 方法变量与表达式

可以将选择子`p.Distance`赋值给方法变量，这样函数只需要提供实参而不需要提供接收者就能调用，这样的变量为方法变量

#### 方法表达式

方法表达式写成 T.f 或者 (*T).f，其中 T 是类型，把原来方法的接收者替换成函数的第一个形参了，使得其可以像平常函数一样调用

```go
distance := Point.Distance // 方法表达式
distance(p, q) // 相当于 p.Distance(q)
```

### 6.5 位向量

数据结构用 slice：[]uint 64 实现 IntSet

位运算

### 6.6 封装

在Go 语言中封装的单元是包而不是类型

## 第7章 接口

抽象类型，隐式实现

隐式实现指某类型实现这个接口所指定的内容，即为该接口的实现。不需要关键词进行说明

```go
type Writer interface {
  Write(p []byte) (n int, err error)
}

type ByteCounter int
func (c *ByteCounter) Write(p []byte) (int, error) {
  *c += ByteCounter(len(p))
  return len(p), nil
}
```

则 ByteCounter 隐式实现了 Writer

可取代性

可以通过组合已有接口得到新的接口

```go
type ReadWriter interface {
  Reader
  Writer
}
```

如上语法为嵌入式接口

空接口类型 interface{} 能接受任意类型的值

可以用接口实现类型共性

### 7.5 接口值

接口类型的值其实有两部分：具体类型和该类型的值。二者称为接口的动态类型和动态值

 ```go
var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil
 ```

类型描述符

动态分发：编译器必须生成一段代码来从类型描述符拿到名为Write的方法地址，再间接调用该方法

第二行类型为指针类型 *os.File ，值指向了代表进程标准输出的 os.File 类型的指针

第三行类型为 *bytes.Buffer，值是一个指向新分配缓冲区的指针 data []byte

#### 接口的比较

如果两个接口都是 nil 或 二者的动态类型一致且动态值相等（使用动态类型的==操作符进行比较），则两个接口值相等

若对应的动态值不可比较（如 slice），则会崩溃

调试的时候，拿到接口值的动态类型很有帮助，可以使用 fmt 包的 %T 来实现这个需求

fmt 内部使用反射来获取动态类型的名字

#### 含有空指针的非空接口

空的接口值（其中不含任何信息）和动态值为 nil 的接口值不同，这是陷阱

如下将会报错

```go
const debug = false

func main() {
	var buf *bytes.Buffer
	// var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
}

func f(out io.Writer) {
	fmt.Println(out != nil)
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
```

out 的动态值为空，但它的动态类型是 *bytes.Buffer，表示 out 是一个包含空指针的非空接口，所以 out != nil 为 true

解决方案是将 buf 声明为 io.Writer

### 7.6 sort.Interface

sort包提供了针对任意序列根据任意排序函数原地排序的功能

满足sort.Interface 接口

```go
type Interface interface {
  Len() int
  Less(i, j int) bool // i, j 是序列的下标
  Swap(i, j int)
}
```

`sort.Strings` 用于 slice []string 的排序

sort.Reverse

sort.IsSorted

### 7.7 http.Handler

```go
package http

type Handler interface {
  ServeHTTP(w ResponseWriter, r *Request)
}
func ListenAndServe(address string, h Handler) error
```

req *http.Request

req.URL.Path 路径

#### 请求多工转发器 ServeMux

http.NewServeMux()

```go
mux := http.NewServeMux()
mux.Handle("/index", http.HandlerFunc(db))
```

http.HandleFunc 其实是类型转换，而不是函数调用

```go
package http

type HandlerFunc func(w ResponseWriter, r *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
  f(w, r)
}
```

同时也是一个函数类型，采用了适配器模式

### 7.8 error 接口

Error 实际上是一个接口类型，包含返回错误消息的方法

```go
type error interace {
  Error() string
}
```

完整error 包只有4行代码

```go
package errors

func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```

没有直接用字符串，为了避免将来无意间的布局变更

fmt.Errorf 封装了 errors，额外提供字符串格式化功能

### 7.10 类型断言

类型断言是作用在接口值上的操作，x.(T)

- 若T是具体类型：检查动态类型是否就是T
  - 是，类型 为T，结果为x的动态值；类型从接口类型变为具体类型了
- T是接口类型：检查动态类型是否满足T
  - 是，从一个接口变为另一个拥有更多方法的接口；但保留接口值中的动态类型和动态值部分
- 上述若检查失败，则崩溃，可以用ok结果防止崩溃，多返回一个布尔型的返回值指示断言是否成功；ok为false，则第一个返回值为断言类型的零值

若操作数为空接口值，类型断言都失败

```go
var w io.Writer = os.Stdout
f, ok := w.(*os.FIle) // 成功 ok, f == os.Stdout
b, ok := w.(*bytes.Buffer) // 失败 !ok, b == nil
```

### 7.11 使用类型断言来识别错误

IO失败原因有很多，但有三类原因必须单独处理：

- 文件已存储，创建操作
- 文件没找到，读取操作
- 权限不足

```go
package os
func IsExist(err error) bool
func IsNotExist(err error) bool
func IsPermission(err error) bool
```

```go
func IsNotExist(err error) bool {
  if pe, ok := err.(*PathError); ok {
    err = pe.Err
  }
  return err == syscall.ENOENT || err == ErrNotExist
}
```

错误识别通常必须在失败操作发生时马上处理，而不是等到错误消息返回给调用者之后

#### io.WriteString

避免多余的内存分配和内存复制，回收

### 7.13 类型分支

接口有两种不同风格

- 第一种风格，强调了方法，而不是具体类型，子类型多态
- 第二种风格，把接口作为这些类型的联合来使用，强调满足这个接口的具体类型，而不是这个接口的方法，不注重信息隐藏，称为可识别联合。为特设多态

类型分支：操作数为x.(type)，每个分支是一个或多个类型。类型分支的分支判定基于接口值的动态类型。其中 nil 分支需要 x==nil

```go
switch x.(type): {
  case nil:
  case int, uint:
  case bool:
  case string:
  default:
}
```

类型分支不允许使用 fallthrough

可以把分支中提取出来的原始值绑定到新变量

```go
switch x := x.(type) { /*...*/ }
```

类型断言隐式创建了一个词法块

## 第8章 goroutine和通道

### 8.1 goroutine

每个并发执行的活动为 goroutine

```go
go f()
```

### 8.4 通道

通道是goroutine 的链接。每个通道是一个具体类型的管道，叫做通道的元素类型

如int

```go
chan int
```

#### 通道的比较

同种类型可以使用==符号，当二者都是同一个通道数据的引用时，为true,也可以和nil 比较

#### 发送和接收

```go
ch <- x // 发送语句
x = <- ch // 赋值语句中的接收表达式
<- ch // 接收，并丢弃
```

#### 关闭

```go
close(ch)
```

设置一个标志位来指示值当前已经发送完毕，这个通道后面没有值了

关闭后的发送操作将导致宕机

接收操作将获取所有已发送的值，直到通道为空；这时任何接收都会立即完成，同时获取到通道元素类型对应的零值

#### 无缓冲通道

make函数接受第一个可选参数（go使用变长参数取代可选参数），表示通道容量，若为0，则创建一个无缓冲通道

无缓冲通道 的发送和接收都将会阻塞，直到另一个goroutine在对应的通道上执行接收和发送

如此，无缓冲通道也叫同步通道

通道除了发送值得时候，若没有携带额外信息，则为事件，目的是进行同步

#### 8.4.2 管道

通道用来连接goroutine，一个输出是另一个输入

判断通道是否关闭，可以用第二个返回值bool ok

```go
x, ok := <- myChannel
```

为 true 时表示接收成功，false 表示在一个关闭且读完的通道上

但由于该模式通用，所以提供了 for range 在通道上迭代，接收完最后一个值后关闭循环

```go
for x:= range myChannel {
  
}
```

结束时的关闭close，垃圾回收器根据其是否可达决定是否回收，而不是根据是否关闭

试图关闭一个已经关闭的通道将宕机，就像关闭空通道一样

通道的关闭可以作为广播机制

#### 8.4.3 单向通道类型

当一个通道用作函数的形参时，它几乎总是被有意地限制不能发送或不能接收

故Go提供单向通道类型。chan <- int 只能发送和 <- chan in 只能接收

close 操作说明通道上没有数据再发送，仅仅可以在发送方 chan<- int 上调用，所以试图关闭一个仅能接收的通道将报错

在任何赋值操作中将双向通道转换成单向通道都是允许的，但反过来不行



#### 8.4.4 缓冲通道

缓冲通道有一个元素队列

```go
ch = make(chan string, 3)
```

通道不满，则不会阻塞

cap 函数可以知道缓冲区的容量

len 函数获取当前通道内元素个数

#### Goroutine 泄露

发送响应结果给goroutine，若缓冲区已满，没有接收，则造成泄露

泄露的goroutine不会自动回收，所以确保goroutine在不再需要的时候可以自动结束

#### 组装流水线

是对于通道和goroutine合适的比喻

### 8.5 并行循环

对于完全独立的子问题组成的问题为高度并行

### 8.6 并发的web爬虫实例

通道作技术信号量使用

保持信号量操作离其所约束的IO操作越近越好

```go
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
  tokens <- struct{}{}
  list, err := links.Extract(url)
  <-tokens
  if err != nil {
    log.Print(err)
  }
  return list
}
```



### 8.7 使用 select 多路复用

多路复用，类似于 Promise.race；每种情况指定一次通信，发送或接收

```go
select {
  case <- ch1:
  case x := <- ch2:
  case ch3 <- y:
  default:
}
```

select一直等待，直到一次通信来告知有些情况可以执行。然后，进行这次通信，执行此情况对应的语句；其他通信将不会发生。对于没有对应情况的 select，select{} 将永远等待

如下例子在偶数时发送，在奇数时接收

```go
ch := make(chan int, 1)
for i:= 0;i < 10;i++ {
  select {
    case x := <-ch:
    	fmt.Println(x)
    case ch <-i:
  }
}
```

select的default语句默认情况，用来指定在没有其他通信发生时可以立即执行的动作

这将可以产生非阻塞通信，重复这个动作称为对通道轮询

### 8.8 取消goroutine机制

Select 中的case 有通道已被关闭的情况

对于不再使用的通道不必显示关闭。如果没有goroutine引用这个通道，这个通道就会被垃圾回收。

注意如果需要把关闭通道作为一个控制信号告知其他goroutine没有更多数据的情况下，需要显示关闭

无缓冲的通道，则传值后立马close，则会在close之前阻塞，有缓冲的通道则即使close了也会继续让接收后面的值chanNum

#### 关于close

close函数是一个内建函数， 用来关闭channel，这个channel要么是双向的， 要么是只写的（chan<- Type）。

这个方法应该只由发送者调用， 而不是接收者。

**当最后一个发送的值都被接收者从关闭的channel(下简称为c)中接收时，接下来所有接收的值都会非阻塞直接成功，返回channel元素的零值。**

如果c已经关闭（c中所有值都被接收）， x, ok := <- c， 读取ok将会得到false。

#### 第8章的练习待巩固，第89章均是并发难点

## 第9章使用共享变量实现并发

### 9.1竞态

对于绝大部分变量，如要回避并发访问，要么限制变量只存在于一个goroutine内，要么维护一个高层的*互斥不变量*

竞态指在多个goroutine按某些交错顺序执行时程序无法给出正确的结果。

#### 避免数据竞态的三种方法

1. 不修改变量；这种方法无法用在存在更新的场景中
2. 避免从多个goroutine访问同一个变量；必须使用通道来向受限goroutine发送查询请求或更新变量。使用通道请求代理一个受限变量的所有访问的goroutine称为该变量的监控。采用流水线的机制对变量进行*串行受限*。这是第八章的内容
3. 允许多个goroutine访问同一个变量，但在同一时刻只有一个goroutine可以访问，即互斥机制

### 9.2互斥锁：sync.Mutex

用通道可以模拟一个信号量

一个计数上限为1的信号量称为二进制信号量

sync包有一个单独的Mutex类型来支持这种模式

互斥锁不可再入，不能锁了又锁

通过封装函数，可以解决互斥锁再入的问题

### 9.3 读写互斥锁： sync.RWMutex

允许只读操作并发执行，但写操作需要获得完全独享的访问权限。这种锁为多读单写锁，即sync.RWMutex

RLock仅可用于在临界区域内对共享变量无写操作的情形

仅在绝大部分goroutine都在获取读锁并且锁竞争比较激烈时，RWMutex才有优势。因为其内部需要更复杂的内部记事簿工作，所以在竞争不激烈时它比普通的互斥锁慢

### 9.4 内存同步

需要互斥锁的原因有两个

- 防止该操作查到其他操作中间
- 同步不仅涉及多个 goroutine 的执行顺序问题，同步还会影响到内存，故锁起来则可以达到内存同步。

> 计算机一般有多个处理器，每个处理器都有内存的本地缓存，为了提高效率，对内存的写入是缓存在每个处理器中的，只在必要时才刷回内存。像通道通信和互斥锁操作这样的同步原语，都会导致处理器把累积的写操作刷回内存并提交，所以这个时刻goroutine的执行结果就保证了对运行在其他处理器的goroutine可见

考虑如下代码：

```go
var x, y int
go func() {
  x = 1
  fmt.Print("y:", y, " ")
}()
go func() {
  y = 1
  fmt.Print("x:", x, " ")
}()
```

如果输出如下结果，则超乎我们的意料了：

x:0 y:0

y:0 x:0

原因和某些特定的编译器和CPU有关，详见P209

### 9.5 延迟初始化 sync.Once

不先释放一个共享锁，就无法直接把它升级为互斥锁

sync.Once 包含一个布尔变量和一个互斥量；布尔变量记录初始化是否已经完成，互斥量则保证这个布尔变量和客户端的数据结构

sync.Once 适合延迟初始化同步

```go
func loadIcons() {
  icons = map[string]image.Image{
    "spades.png": loadIcon("spades.png")
    ...
  }
}

var loadIconsOnce sync.Once
var icons map[string]image.Image
// 并发安全
func Icon(name string) image.Image {
  loadIconsOnce.Do(loadIcons)
  return icons[name]
}
```

Do函数内部就是先检查是否已置done为已生成，没生成则使用锁锁住初始化的过程，生成则直接return

### 9.7 并发非阻塞缓存

函数记忆

重复抑制：避免重复额外的处理

两种方案构建并发结构：

- 共享变量上锁
- 通信顺序进程



#### 共享变量上锁

和传统模型类似，易懂

```go
package memo

import "sync"

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

//!+
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}

//!-

```





#### 通信顺序进程

存在通道传值，阅读代码会跳跃

```go
package memo

//!+Func

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

//!-Func

//!+get

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

//!-get

//!+monitor

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}

//!-monitor

```

### 9.8 goroutine与线程

#### 可增长的栈

每个OS线程都有一个固定大小的栈内存，通常为2MB，固定栈大小在很多场景下不够灵活

作为对比，一个 goroutine 在生命周期开始时只有一个很小的栈，典型情况下为2KB，可以按需增大和缩小

#### OS 调度

OS调度线程的大致流程

OS线程由OS内核调度。每隔几毫秒，一个硬件时钟中断发送到CPU，CPU调用一个叫调度器的内核函数。该函数暂停当前正在运行的线程，把它的寄存器信息保存到内存，查看线程列表并决定接下来运行哪一个线程，再从内存恢复线程的注册表信息，最后继续执行选中的线程。

OS线程由内核调度，所以控制权限从一个线程到另一个线程需要一个完整的上下文切换：即保存一个线程状态到内存，再恢复另外一个线程的状态，最后更新调度器的数据结构。

上述操作涉及内存局域性以及涉及内存访问数量，还有访问内存所需的CPU周期数量的增加，该操作其实很慢

#### Goroutine 调度

Go 运行时 包含一个自己的调度器，该调度器使用一个称为 m:n 调度的技术，因为它可以复用/调度 m 个 goroutine 到 n 个OS线程。Go调度器与内核调度器的工作类似，但Go调度器只需关心单个 Go程序的 goroutine 调度问题

Go 调度器不是由硬件时钟来定期触发的，而是由特定的Go 语言结构触发的。如当一个 goroutine 调用 time.Sleep 或被通道阻塞或互斥量操作时，调度器会将这个 goroutine 设置为休眠模式，并运行其他 goroutine 直到前一个可重新唤醒为止。因为它不需要切换内核语境，所以调用一个 goroutine 比调度 一个线程成本低很多。

#### GOMAXPROCS

Go 调度器使用 GOMAXPROCS 参数来确定使用多少个 OS 线程来同时执行 Go 代码。默认值是机器上的 CPU 数量。

GOMAXPROCS 是m:n 调度中的 n。

正在休眠或正被通道通信阻塞的 goroutine 不需要占用线程。阻塞在 IO 和其他系统调用中或调用非Go 语言写的函数的 goroutine 需要一个独立的OS线程，但这个线程不计算在GOMAXPROCS 内

#### Goroutine 没有标识

当前线程都有一个独特标识，通常可以取一个整数或指针，这个特性可以构建一个 线程的局部存储，本质上是一个全局 map，这样线程可以独立地用这个 map 存储和获取值，并互不干扰

但这样会导致一种不健康的”超距作用“，即函数的行为不仅取决于其参数，还取决于运行它的线程标识

所以，goroutine 没有可供程序员访问的标识，能影响一个函数行为的参数应当是显式指定的。