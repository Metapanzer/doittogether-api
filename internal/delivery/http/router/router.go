package router

import (
	"DoItTogether/internal/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

type router struct {
	App         *gin.Engine
	UserHandler *handler.UserHandler
}

type Router interface {
	Setup()
}

func NewRouter(app *gin.Engine, userHandler *handler.UserHandler) Router {

	return &router{
		App:         app,
		UserHandler: userHandler,
	}
}

func (r *router) Setup() {
	r.RegisterPublicEndPoints()
	r.RegisterPrivateEndPoints()
}

func (r *router) RegisterPublicEndPoints() {
	r.App.GET("/users", r.UserHandler.Register)
}

func (r *router) RegisterPrivateEndPoints() {
	r.App.GET("/private", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Private",
		})
	})
}
