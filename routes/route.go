package routes

import (
	"GoEvents/controllers"
	"github.com/gin-gonic/gin"
)

type Route interface {
	RegisterHandlers()
}

type route struct {
	engine     *gin.Engine
	controller controllers.Controller
}

func RegisterHandlers(engine *gin.Engine, controller controllers.Controller) Route {
	return &route{
		engine:     engine,
		controller: controller,
	}
}

func (r route) RegisterHandlers() {
	r.engine.GET("/employee", r.controller.GetAllEmployees())
	r.engine.GET("/employee/:id", r.controller.GetAccount())
	r.engine.POST("/employee", r.controller.CreateAccount())
	r.engine.PATCH("/employee/:id", r.controller.UpdateAccount())
	r.engine.DELETE("/employee/:id", r.controller.DeleteAccount())
}
