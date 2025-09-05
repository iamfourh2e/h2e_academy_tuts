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
	//username, password
	group.POST("/login", func(c *gin.Context) {
		var user *models.UserModel
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, bson.M{
				"error": "Username and password is rquried",
			})
			return
		}
		// 1. find exist user ✅,
		// 2. convert request password to hash✅
		// 3. request password  == exist password✅
		// 4. success return JWT(setup later)✅
		// 5. failed return error✅
		res, err := userService.FindUserByUsername(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{
				"error": err.Error(),
			})
			return
		}

		user.HashPassword()
		if user.Password != res.Password {
			c.JSON(http.StatusNotFound, bson.M{
				"message": "Username or password is incorrect",
			})
			return
		}
		//generate token key
		userClaim := new(models.UserClaimModel)
		token, err := userClaim.GenerateToken(res)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, bson.M{
			"token": token,
			"data": bson.M{
				"fullName": res.FullName,
				"id":       res.ID,
			},
		})
	})

}
