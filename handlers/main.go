package handlers

import (
	"example_web_site/common/zerrors"
	"example_web_site/db"
	"example_web_site/proto"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func InitRouter(client *mongo.Client, node *snowflake.Node) *gin.Engine {
	db.Init(client, node)
	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to my web site"})
	})
	api := router.Group("/api/v1")
	registerUserHandler(api)
	return router
}

func badRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, processError(err))
}

func internalError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, processError(err))
}

func processError(err error) interface{} {
	var result interface{}
	if e, ok := err.(*zerrors.BaseError); ok {
		result = proto.Response{
			Code:    e.Code,
			Message: e.Message,
		}
	} else if err != nil {
		result = gin.H{"message": err.Error(), "code": 8888}
	} else {
		result = gin.H{"message": "error", "code": 8888}
	}
	return result
}

func success(c *gin.Context, data interface{}) {
	response := proto.Response{
		Code:    0,
		Message: "Success",
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

func created(c *gin.Context, data interface{}) {
	response := proto.Response{
		Code:    0,
		Message: "Created",
		Data:    data,
	}
	c.JSON(http.StatusCreated, response)
}

func unauthorized(c *gin.Context, err *zerrors.BaseError) {
	c.JSON(http.StatusUnauthorized, processError(err))
}

func forbidden(c *gin.Context, err *zerrors.BaseError) {
	c.JSON(http.StatusForbidden, processError(err))
}

func errorResponse(c *gin.Context, err *zerrors.BaseError) {
	c.JSON(http.StatusForbidden, processError(err))
}
