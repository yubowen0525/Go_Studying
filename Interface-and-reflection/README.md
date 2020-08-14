# 接口与反射
## 一 接口
### 1.1 接口的举例
[demo](./Demo/interface.go)
 
### 1.2 接口类型与约束
接口之间可以匿名继承，或组合继承

#### 1.2.1 动态类型
类型断言测试[demo](./Demo/assert.go)
```go
if v,ok := varI.(T); ok{
    Process(v)
    return
}
```

### 1.2.2 类型判断
利用case根据不同类型，执行不同分支 [demo](./Demo/interface.go)
```go
func clasifier(items ...interface{})  {
	for i,item:= range items{
		switch  item.(type){
		case int:
			fmt.Printf("参数 #%d 类型是 int\n",i)
		case bool:
			fmt.Printf("参数 #%d 类型是 bool\n",i)
		case string:
			fmt.Printf("参数 #%d 类型是 string\n",i)
		default:
			fmt.Printf("参数 #%d 类型未知\n",i)
		}
	}
}
```

### 1.2.3 接口实现
没太看明白
### 1.2.4 嵌套接口
匿名嵌套
```go
type ReadWrite interface{
    Read(b Buffer) bool
    Write(b Buffer) bool
}

type Lock interface{
    Lock()
    Unlock()
}

type File interface{
    ReadWrite
    Lock
    Close()
}
```
### 1.2.5 接口赋值
1. 方法的接口者不是指针，但赋值是指针时，会自动成功一个以指针为接口者的方法：定义方法， `func(a Integer) Less(b Interger) bool` 如果使用的时候是地址赋值给接口 
`var a LessAdder = &a` 
2. 将一个接口赋值给另一个接口：
    - 接口的方法完全相同，或者时子集的关系，都可以赋值

### 1.2.6 接口查询
检查file1接口指向的对象实例是否实现了two.IStream接口。
```go
// two.IStream类型是interface
var file1 Writer = ...
if file5,ok := file1.(two.IStream);ok{
    ...
}
```

### 1.2.7 接口组合
```go
type ReadWriter interface{
    Reader
    Writer
}

type ReadWriter interface{
    Reade(p []byte)(n int,err error)
    Writer(p []byte)(n int,err error)
}
```

## 二 反射
反射是程序检查代码中所拥有的结构尤其是类型的一种能力，这是元编程的一种形式。尽量避免使用或小心使用。

反射可以从接口值反射到对象，也可以从对象反射到接口值
### 2.1.1 方法和类型的反射
[demo](./Demo/reflect.go)

### 2.1.2 通过反射修改设置值
CanSet查看是否可以反射，首先传送指针，在Elem()，才可修改值
[demo](./Demo/reflect.go)

### 2.1.3 反射结构
利用反射，解刨数据类型
[demo](./Demo/reflect_struct.go)

### 2.1.4 Printf 和 反射
[demo](./Demo/printf_and_reflect.go)