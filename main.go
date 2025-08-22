package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

var Database []Student = []Student{}

func main() {
	r := gin.Default()
	//POST
	r.POST("/v1/students", func(ctx *gin.Context) {
		student := &Student{}
		err := ctx.ShouldBind(&student)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		Database = append(Database, *student)
		ctx.JSON(http.StatusOK, gin.H{"message": "Student created successfully", "data": student})
	})
	r.GET("/v1/students/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for _, student := range Database {
			if student.ID == id {
				ctx.JSON(http.StatusOK, gin.H{"message": "Student found", "data": student})
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
	})
	r.PUT("/v1/students/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")                   // retrieve student id
		studentToUpdate := &Student{}           //declare new student data
		err := ctx.ShouldBind(&studentToUpdate) // apply data to student object
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		for i, student := range Database {
			if student.ID == id {
				studentToUpdate.ID = id
				Database[i] = *studentToUpdate
				ctx.JSON(http.StatusOK, gin.H{"message": "Student updated successfully", "data": studentToUpdate})
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
	})

	r.GET("/v1/students", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,
			gin.H{"message": "Students fetched successfully", "data": Database})
	})
	r.DELETE("/v1/students/:id", func(ctx *gin.Context) {
		//5 elements 0 -> 4
		// numbers := []int{1, 2, 3, 4, 5}
		// numbers = append(numbers[:2], numbers[3:]...)
		// fmt.Printf("%v", numbers)
		id := ctx.Param("id")
		for i, student := range Database {
			if student.ID == id {
				Database = append(Database[:i], Database[i+1:]...)
				ctx.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
	})
	r.Run(":4000")
}
