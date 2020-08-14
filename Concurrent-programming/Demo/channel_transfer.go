package main

type PipeData struct {
	value   int
	handler func(int) int
	next    chan int
}

//定义一个pipedata队列，传递给handle函数
func handle(queue chan *PipeData) {
	for data := range queue {
		//pipedata的handle处理完，返回值赋予data.next(chan)，此时阻塞于下一个读
		// 而for会去读取下一个chan
		data.next <- data.handler(data.value)
	}
}

func main() {

}
