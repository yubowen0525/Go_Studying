# 复合数据类型
## 一 数组(栈区)
```go
//... 代表任意
list := [...]int{1,2,3}
```
## 二 切片(堆区)
### 2.1 创建数组切片
1. make和切片字面量
    - 注意有长度，容量的概念
```go
//长度和容量都是5
silce := make([]string,5)

//长度3，容量5
silce := make([]string,3,5)

```
2. nil和空切片
```go
//空切片
var slice []int

```

### 2.2 切片的使用
[demo](./Demo/slice.go)
1. 赋值和分割
    - 切片是[i,j) 左闭右开
    - 共享同一段底层数组，通过不同的切片会看到底层数组不同的部分，修改1个，会影响其他的切片
````go
slice := []int{10,20,30,40,50}
newslice := slice[1:3]
newslice2 := slice[2:3:4]
````
2. 切片扩容:append可增加容量，具体看是否需要扩容
    - 如果切片底层数组没有足够的可用容量，append()函数会创建一个新的底层数组
    - 将被引用的现有的值复制到新的数组里，在追加新的值
    - newSlice拥有一个全新的底层数组，且容量是原来的两倍
    - 容量小于1000，成倍增加。容量个数超过1000，增长因子设为1.25，每次增加25%的容量
```go
slice := []int{10,20,30,40}
newslice = append(slice,60)
```
3. 遍历:range
    - range 会返回两个值：(索引位置，对应元素值的**副本**)
    - range迭代从头部开始，想有更多控制使用for len([]int)
4. 限制容量：切片的第三个参数slice:=source[i:j:k]
    - 长度大小：j - i
    - 容量大小：k - i
    - 例如 source[0:1:3] 长度1，容量3
### 2.3 切片的传递
64位机器，切片需要24B内存，3x8B
- 指针字段：8B
- 长度：8B
- 容量：8B

## 三 映射(hashmap)
### 3.1 映射的创建
[demo](./Demo/map.go)
- key 需要有==语义，需要注意切片，函数，包含切片的数据结构由于具有引用语义，均不能作为映射的键

### 3.2 映射的使用
1. 赋值
    - =  
    ```go
        colors := map[string]string{}
        colors["Red"] = '#DA1337'
    ```
2. 查找与遍历
[demo](./Demo/map.go)
    - key不存在，会给key赋值相应的零值
3. 元素删除 delete
[demo](./Demo/map.go)
    -  delete()，如果删除key不存在，什么副作用没有。但如果传入map是nil，抛出异常panic

### 3.3 映射传递给函数
函数间传递映射**并不会制造出该映射的一个副本**。实际上，当传递映射给一个函数，对这个映射做了修改，所有对这个映射的引用都会察觉到这个修改。

# 四 链表
## list
Go语言的链表实现在标准库的container/list代码包中。这个代码包中有两个公开的程序实体——List和Element，List实现了一个双向链表（以下简称链表），而Element则代表了链表中元素的结构。
- List：代表链表
- Element：代表node
[demo](./Demo/list.go)

## ring
只有一层概念, prev,next,value

```
type Ring struct {
	next, prev *Ring
	Value      interface{} // for use by client; untouched by this library
}
```


[demo](./Demo/ring.go)