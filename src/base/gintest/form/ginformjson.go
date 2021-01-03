package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
	"time"
)

type Login struct {
	User     string `json:"user" form:"user" uri:"user" binding:"required"`
	Password string `json:"password" form:"password" uri:"password" binding:"required"`
}

func main() {
	g := gin.Default()
	v1 := g.Group("/")
	{
		v1.POST("postjson", func(ctx *gin.Context) {
			var jsondata Login
			ctx.ShouldBindJSON(&jsondata)
			fmt.Println(jsondata)
			ctx.JSON(http.StatusOK, jsondata)
		})
		v1.POST("postform", func(ctx *gin.Context) {
			var formdata Login
			ctx.Bind(&formdata)
			fmt.Println(formdata)
			ctx.JSON(http.StatusOK, formdata)
		})
		v1.POST("postyaml", func(ctx *gin.Context) {
			a := map[string]interface{}{
				"name":  "ralap",
				"aget":  20,
				"hobby": "devops coding",
			}
			ctx.YAML(http.StatusOK, a)
		})
		v1.POST("postprotobuf", func(ctx *gin.Context) {
			repo := []int64{int64(1), int64(2)}
			label := "label"
			data := &protoexample.Test{
				Label: &label,
				Reps:  repo,
			}
			ctx.JSON(http.StatusOK, data)
		})
		v1.Use(Middleware())
		v1.POST("middleware", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, nil)
		})
		v1.Use(gin.BasicAuth(gin.Accounts{
			"admin": "123456",
		}))
		v1.GET("/basicauth", func(c *gin.Context) {
			c.JSON(200, "首页")
		})
	}

	g.Run(":8080")
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("middler ware starting")
		c.Set("request", "中间件")
		c.Next()
		status := c.Writer.Status()
		time.Sleep(3 * time.Second)
		fmt.Println("中间件执行完毕", status)
		fmt.Println("time", time.Since(t))
	}
}
