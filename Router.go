package main

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()                        // Getting Router object to set routes from gin
	router.GET("/users", GetAllUsers)              // Creating route for getting all users from system
	router.GET("/user/:ID", GetUserByID)           // Creating route for getting sepcific user from system using ID
	router.POST("/adduser", CreateUser)            // Creating Route for creating a user into system
	router.PUT("/updateuser", UpdateUser)          // Creating Route for Updating user Info
	router.DELETE("/deleteuser/:ID", DeleteAnUser) // Creating Route for deleting an user from db using id

	return router
}
