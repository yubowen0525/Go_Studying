# 一 并发编程基础
## 1.1 线程，进程，协程
- 如果连时钟阻塞、 线程切换这些功能我们都不需要了，自己在进程里面写一个逻辑流调度的东西。那么我们即可以利用到并发优势，又可以避免反复系统调用，还有进程切换造成的开销，分分钟给你上几千个逻辑流不费力。这就是用户态线程。
- 从上面可以看到，实现一个用户态线程有两个必须要处理的问题：一是碰着阻塞式I\O会导致整个进程被挂起；二是由于缺乏时钟阻塞，进程需要自己拥有调度线程的能力。如果一种实现使得每个线程需要自己通过调用某个方法，主动交出控制权。那么我们就称这种用户态线程是协作式的，即是协程。

go协程调用跟切换比线程效率高。
- go协程之间的切换发生在用户态，因为用户态没有时钟中断，系统调用等机制。
- 占用内存更少，协程栈空间(4~5KB),根据相应的数据进行伸缩。默认情况，线程栈大小是1MB。

go的CSP并发模型实现:M,P,G[随笔](https://www.cnblogs.com/sunsky303/p/9115530.html)
## 1.2 指定使用核心数
[demo](./Demo/test.go)
```go
var numCores = flag.Int("n",2,"CPU核心数")

in main()
flag.Pars()
runtime.GOMAXPROCS(*numCores)
```

# 二 协程
- 其他库实现，不一定完整，不一定真的并发调度，可能同时阻塞于I/O
- Go实现goroutine,提供所有系统调用操作支持，进行协程调度（有CPU优先级，与同级协程竞争CPU资源），交给Go负责统一调度

## 2.1 协程基础
[demo](./Demo/goroutine_basic.go)

## 2.2 协程间通信
共享数据：多个并发单元分别保持对同一个数据的引用，实现对数据的共享。
1. 基于锁来解决问题[demo](./Demo/goroutine_basic.go)
    - 但是当逻辑复杂，业务庞大很容易出现问题，所以go基于消息机制而非共享内存作为通信方式
2. channel


# 三 通道 channel
语言级别提供的goroutine间的通信方式，使用channel在两个或多个goroutine之间传递消息。

channel类型相关，一个channel传递一种类型的值

## 3.1 基础语法
基础
- 声明 `var chanName chan ElementType`
- 初始化:  `ch := make(chan int)`
- 写： `ch <- value`
- 读: `value := <- ch`

### select 
用于处理异步I/O问题

随机读取I/O内容
[demo](./Demo/channel_select.go)

### 缓冲机制
简单的消息队列
```go
	c := make(chan int ,1024)

    //顺序读取c缓存队列的内容
	for i := range c {
		fmt.Println(i)
	}
```

### 超时和计时器
[demo](./Demo/channel_timeout.go)设置超时时间，防止死锁

### channel的传递
[demo](./Demo/channel_transfer.go)

### 单向channel
禁止权限问题
- 初始化 `ch4 := make(chan int)`
- 只能读： `ch5 := <-chan int(ch4)` 强制类型转换，用于函数参数可禁止函数的操作
- 只写: `ch6 := chan<- int(ch4)`

### 关闭channel
关闭ch
```go
close(ch)
```
判断channel是否关闭
```go
//返回false表示通道关闭
x,ok := <-ch
```

## 协程同步
[demo](./Demo/chan_sync.go)
if需要判断，for会自动判断通道的关闭

## 协程和恢复
```go
// 只读的通道workChan
func server(workChan <-chan *Work){
    for work := range workChan{
        go safelyDo(work)
    }
}

// 执行函数时，包一层safelydo
func safelyDo(work *Work){
    defer func{
        if err := recover(); err != nil{
            log.Printf()
        }
    }
    
    do(work)
}
```