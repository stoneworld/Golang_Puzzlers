### 入门

Go是一门编译型语言，Go语言的工具链将源代码及其依赖转换成计算机的机器指令。
Go语言提供的工具都通过一个单独的命令go调用，go命令有一系列子命令。
最简单的一个子命令就是run，这个命令编译一个或多个以.go结尾的源文件，链接库文件，并运行最终生成的可执行文件。
go build 命令生成可执行的二进制文件，之后你可以随时运行它，不需任何处理。

变量声明：
变量会在声明时直接初始化。如果变量没有显式初始化，则被隐式地赋予其类型的零值（zero value），数值类型是0，字符串类型是空字符串 ""。 
类似 `s := ""`  短变量声明只能用户函数内部，不能用于包变量

map: 
它是一个无序的key/value对的集合，其中所有的key都是不同的。
一个map就是一个哈希表的引用，map类型可以写为map[K]V，其中K和V分别对应key和value。map中所有的key都有相同的类型，所有的value也有着相同的类型，但是key和value之间可以是不同的数据类型。
其中K对应的key必须是支持==比较运算符的数据类型，所以map可以通过测试key是否相等来判断是否已经存在。

> counts := make(map[string]int) // make 函数创建空 map

map 是一个由 make 函数创建的数据结构的引用。map 作为参数传递给某函数时，该函数接收这个引用的一份拷贝（copy，或译为副本），被调用函数对map底层数据结构的任何修改，调用者函数都可以通过持有的map引用看到。


// todo
- [ ] 字符串和字节切片


// fatal error: all goroutines are asleep - deadlock!
所有的协程都已经 asleep，但此时还在等待 <-ch 数据接收，自然就死锁了。

```go
ch := make(chan int)
<-ch
```


### 程序结构

函数内部定义的变量只能在函数内部调用，包一级声明语句声明的名字可在整个包对应的每个源文件中访问，而不是仅仅在其声明语句所在的源文件中访问。名字的开头字母的大小写决定了名字在包外的可见性。

#### 声明

变量声明：

> var 变量 类型 = 表达式

如果表达式被省略，使用零值初始化该变量，因此在Go语言中不存在未初始化的变量，一组变量也可以通过调用一个函数，由函数返回的多个返回值初始化：

> var f, err = os.Open(name) // os.Open returns a file and an error

引用类型：slice 指针、chan、map、函数

|  变量类似   | 零值  |
|  ----  | ----  |
| 数值类型  | 0 |
| 字符串  | 空字符串 |
| 接口  | nil |
| 引用类型  | nil |


简短变量声明：
请记住“:=”是一个变量声明语句，而“=”是一个变量赋值操作。和普通var形式的变量声明语句一样，简短变量声明语句也可以用函数的返回值来声明和初始化变量。简短变量声明会屏蔽函数（代码块）外部的声明

> f, err := os.Open(name)


简短变量声明语句中必须至少要声明一个新的变量，下面的代码将不能编译通过：

```go

f, err := os.Open(infile)
// ...
f, err := os.Create(outfile) // compile error: no new variables

```

#### 指针
如果用“var x int”声明语句声明一个x变量，那么&x表达式（取x变量的内存地址）将产生一个指向该整数变量的指针，指针对应的数据类型是*int，指针被称之为“指向int类型的指针”。
如果指针名字为p，那么可以说“p指针指向变量x”，或者说“p指针保存了x变量的内存地址”。同时 *p 表达式对应 p 指针指向的变量的值。

```go

x := 1
p := &x         // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2          // equivalent to x = 2
fmt.Println(x)  // "2"

```

在Go语言中，返回函数中局部变量的地址也是安全的。例如下面的代码，调用f函数时创建局部变量v，在局部变量地址被返回之后依然有效，因为指针p依然引用这个变量。
每次调用f函数都将返回不同的结果。

```go

var p = f()

func f() *int {
    v := 1
    return &v
}

fmt.Println(f() == f()) // "false"

```

new 函数
表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T，且每次调用new函数都是返回一个新的变量的地址

```go

p := new(int)   // p, *int 类型, 指向匿名的 int 变量
fmt.Println(*p) // "0"
*p = 2          // 设置 int 匿名变量的值为 2
fmt.Println(*p) // "2"

p := new(int)
q := new(int)

fmt.Println(p == q) // "false"
```

#### 作用域

不要将作用域和生命周期混为一谈。声明语句的作用域对应的是一个源代码的文本区域；它是一个编译时的属性。
一个变量的生命周期是指程序运行时变量存在的有效时间段，在此时间区域内它可以被程序的其他部分引用；是一个运行时的概念。

#### 字符串
一个字符串是一个不可改变的字节序列（因为字符串是不可修改的，因此尝试修改字符串内部数据的操作也是被禁止的）。
字符串可以包含任意的数据，包括byte值0，但是通常是用来包含人类可读的文本。
文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列，我们稍后会详细讨论这个问题。
内置的len函数可以返回一个字符串中的字节数目（不是rune字符数目），索引操作s[i]返回第i个字节的字节值，i必须满足0 ≤ i< len(s)条件约束。

#### 数组

数组是一个由固定长度的特定类型元素组成的序列，和数组对应的类型是Slice（切片），它是可以增长和收缩的动态序列。
数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型。数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定。

#### Map

但是map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作：

> _ = &ages["bob"] // compile error: cannot take address of map element

Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。
在实践中，遍历的顺序是随机的，每一次遍历的顺序都不相同。这是故意的，每次都使用随机的遍历顺序可以强制要求程序不会依赖具体的哈希函数实现。
如果要按顺序遍历key/value对，我们必须显式地对key进行排序，可以使用sort包的Strings函数对字符串slice进行排序。下面是常见的处理方式：

```go

import "sort"

var names []string
for name := range ages {
    names = append(names, name)
}
sort.Strings(names)
for _, name := range names {
    fmt.Printf("%s\t%d\n", name, ages[name])
}

```

在向map存数据前必须先创建map。

通过key作为索引下标来访问map将产生一个value。如果key在map中是存在的，那么将得到与key对应的value；如果key不存在，那么将得到value对应类型的零值，正如我们前面看到的ages["bob"]那样。这个规则很实用，但是有时候可能需要知道对应的元素是否真的是在map之中。
例如，如果元素类型是一个数字，你可能需要区分一个已经存在的0，和不存在而返回零值的0，可以像下面这样测试：

```go

age, ok := ages["bob"]
if !ok { /* "bob" is not a key in this map; age == 0. */ }

if age, ok := ages["bob"]; !ok { /* ... */ }

```
