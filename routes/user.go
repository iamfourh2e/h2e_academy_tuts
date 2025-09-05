package routes

import (
	"go_tuts/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoute(route *gin.Engine, client *mongo.Client, dbName string) {
	userService := models.NewUserService(client, dbName)
	group := route.Group("/users")
	group.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"data": id,
		})
	})
	group.POST("", func(c *gin.Context) {
		var user *models.UserModel
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, bson.M{
				"error": "Could not bind json. please make sure to give us the body",
			})
			return
		}
		user.BeforeInsert()
		user.HashPassword()
		res, err := userService.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, bson.M{
			"data": res,
		})

	})

}
