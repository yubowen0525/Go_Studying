# Context
golang 的 Context包，是专门用来简化对于处理单个请求的多个goroutine之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用

例子1：
- 比如有一个网络请求Request，每个Request都需要开启一个goroutine做一些事情，这些goroutine又可能会开启其他的goroutine。
- 这样的话， 我们就可以通过Context，来跟踪这些goroutine，并且通过Context来控制他们的目的，这就是Go语言为我们提供的Context，中文可以称之为“上下文”。

例子2：
- 另外一个实际例子是，在Go服务器程序中，每个请求都会有一个goroutine去处理。
- 然而，处理程序往往还需要创建额外的goroutine去访问后端资源，比如数据库、RPC服务等。
- 由于这些goroutine都是在处理同一个请求，所以它们往往需要访问一些共享的资源，比如用户身份信息、认证token、请求截止时间等。
- 而且如果请求超时或者被取消后，所有的goroutine都应该马上退出并且释放相关的资源。**这种情况也需要用Context来为我们取消掉所有goroutine**

# 结构
```go
type Context interface {

Deadline() (deadline time.Time, ok bool)

Done() <-chan struct{}

Err() error

Value(key interface{}) interface{}

}
```
可以看到Context是一个interface，在golang里面，interface是一个使用非常广泛的结构，它可以接纳任何类型。Context定义很简单，一共4个方法，我们需要能够很好的理解这几个方法
1. `Deadline`方法是获取设置的截止时间的意思，第一个返回式是截止时间，到了这个时间点，Context会自动发起取消请求；第二个返回值ok==false时表示没有设置截止时间，如果需要取消的话，需要调用取消函数进行取消。
2. `Done`方法返回一个只读的chan，类型为struct{}，我们在goroutine中，如果该方法返回的chan可以读取，则意味着parent context已经发起了取消请求，我们通过Done方法收到这个信号后，就应该做清理操作，然后退出goroutine，释放资源。之后，Err 方法会返回一个错误，告知为什么 Context 被取消。
3. `Err`方法返回取消的错误原因，因为什么Context被取消。
4. `Value`方法获取该Context上绑定的值，是一个键值对，所以要通过一个Key才可以获取对应的值，这个值一般是线程安全的。