package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	httpen := gin.Default()
	httpen.GET("/", func(c *gin.Context) {
		c.String(200, "test")
	})
	v1 := httpen.Group("/v1")
	v1.POST("/test", test)
	v1.POST("/testpost/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, name)
	})
	v1.GET("/testget/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, name)
	})
	v1.POST("/testpost/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		ctx.String(http.StatusOK, name+" is "+action)
	})
	v1.POST("/welcome", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("name", "ok")
		ctx.String(http.StatusOK, name)
	})
	v1.POST("/welcomeform", func(ctx *gin.Context) {
		name := ctx.QueryMap("name")
		key := ctx.PostFormMap("key")
		ctx.String(http.StatusOK, fmt.Sprintf("%v,%v", name, key["1"]))
	})
	v1.POST("/upload", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		ctx.SaveUploadedFile(file, file.Filename)
		ctx.JSON(http.StatusOK, gin.H{
			"code":     200,
			"result":   "success",
			"filename": file.Filename,
		})
	})
	httpen.Run(":8000")

}

func test(ctx *gin.Context) {
	name := ctx.Query("name")
	//name=ctx.GetString("name")
	fmt.Println("======", name)
	ctx.String(http.StatusOK, name)
}
