package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/test", test)

	a := r.Group("/oauth")
	a.Use(Logger())

	{
		a.GET("/test", test)
	}

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

func test(c *gin.Context) {
	example := c.MustGet("example").(string)
	// 打印："12345"
	log.Println(example)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 设置 example 变量
		c.Set("example", "12345")
		// 请求前
		//c.Next()
		// 请求后
		latency := time.Since(t)
		log.Print(latency)
		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}
