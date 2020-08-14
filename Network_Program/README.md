# 一 Socket
Socket编程的API都在net包中
## IP类型
[demo](./Demo/ip.go)
## Dial()函数
Dial函数连接服务器，Listen监听，Accept接收连接

无论什么协议建立什么连接，调用net.Dial()即可 `func Dial(net,addr string)(Conn,error)`

[demo](./Demo/icmp_demo.go)
```go
//TCP连接：
conn,err := net.Dial("tcp","192.168.0.10:2100")

//UDP
conn,err := net.Dial("udp","192.168.0.12:975")

//ICMP
conn,err := net.Dial("ip4:icmp","www.baidu.com")
conn,err := net.Dial("ip4:1","10.0.0.3")
```
## TCP socket
```go
func (c *conn) Write(b []byte) (int, error) 
func (c *conn) Read(b []byte) (int, error)

type TCPAddr struct{
    IP IP
    Port int

zone string
}

//解析tcpaddr
func ResolveTCPAddr(net, addr string) (*TCPAddr, error)
// net : TCP4,TCP6,TCP
```

### 控制TCP连接
```go
//设置连接超时时间
func DialTimeout(net,addr string,timeout time.Duration)(Conn,error)

//
func (c *conn) SetReadDeadline(t time.Time) error
func (c *conn) SetWriteDeadline(t time.Time) error
```

# 二 HTTP
## HTTP client
[demo](./Demo/http_client.go)

## HTTP server
[demo](./Demo/http_server.go)
`func ListenAndServe(addr string, handler Handler) error`