package handlers

import (
	"example_web_site/db"
	"github.com/gin-gonic/gin"
)

func registerUserHandler(api *gin.RouterGroup) {
	userGroup := api.Group("/user")
	userGroup.GET("", getUsers)
}

func getUsers(c *gin.Context) {
	users, err := db.UserRepo.GetUsers()
	if err != nil {
		internalError(c, err)
		return
	}
	success(c, users)

}
