package routes

import (
	"fmt"
	"go_tuts/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ProductRoute(route *gin.Engine) {
	g := route.Group("/products")
	///products have protector (middleware)
	//use middleware to check before enter the route
	g.Use(middleware.AuthMiddleWare())
	g.GET("", func(ctx *gin.Context) {
		user, ok := ctx.Get("user")
		if !ok {
			fmt.Println("No context found")
		}
		ctx.JSON(http.StatusOK, bson.M{
			"data": user,
		})
	})
}
