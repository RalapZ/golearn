package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
			ctx.JSON(http.StatusOK, gin.H{
				"status": 200,
			})
		})
		v1.POST("postform", func(ctx *gin.Context) {
			var formdata Login
			ctx.Bind(&formdata)
			fmt.Println(formdata)
			ctx.JSON(http.StatusOK, formdata)
		})
	}
	g.Run(":8080")

}
