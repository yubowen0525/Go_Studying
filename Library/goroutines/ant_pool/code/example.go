// MIT License

// Copyright (c) 2018 Andy Pan

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	time.Sleep(10 * time.Millisecond)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func main() {
	// defer ants.Release()

	runTimes := 100000

	// // Use the common pool.
	var wg sync.WaitGroup
	// syncCalculateSum := func() {
	// 	demoFunc()
	// 	wg.Done()
	// }
	// for i := 0; i < runTimes; i++ {
	// 	wg.Add(1)
	// 	_ = ants.Submit(syncCalculateSum)
	// }
	// wg.Wait()
	// fmt.Printf("running goroutines: %d\n", ants.Running())
	// fmt.Printf("finish all tasks.\n")

	// Use the pool with a method,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.

	// p, _ := ants.NewPool(10000)
	type funchandler func(i interface{})
	var f funchandler = func(i interface{}) {
		myFunc(i)
		wg.Done()
	}

	// wg.Done 必须在所有 wg.Add 之后执行，所以要保证两个函数都在main协程中调用；
	// wg.Done 在 worker协程里调用，尤其要保证调用一次，不能因为 panic 或任何原因导致没有执行（建议使用 defer wg.Done()）；
	// wg.Done 和 wg.Wait 在时序上是没有先后。

	p, _ := ants.NewPoolWithFunc(10000, f)
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	// 阻塞所有woker，直到Done的次数 跟 add 次数相同
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
	// if sum != 499500 {
	// 	panic("the final result is wrong!!!")
	// }
}
