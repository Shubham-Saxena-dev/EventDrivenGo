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
	r.engine.GET("/account", r.controller.GetAllEmployees())
	r.engine.GET("/account/:id", r.controller.GetAccount())
	r.engine.POST("/account", r.controller.CreateAccount())
	r.engine.PUT("/account/:id", r.controller.UpdateAccount())
	r.engine.DELETE("/account/:id", r.controller.DeleteAccount())
}
