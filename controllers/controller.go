package controllers

import (
	"GoEvents/requests"
	"GoEvents/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Controller interface {
	GetAllEmployees() gin.HandlerFunc
	GetAccount() gin.HandlerFunc
	CreateAccount() gin.HandlerFunc
	UpdateAccount() gin.HandlerFunc
	DeleteAccount() gin.HandlerFunc
}

type controller struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) GetAllEmployees() gin.HandlerFunc {
	return func(context *gin.Context) {
		response, err := c.service.GetAllEmployees()
		if err != nil {
			handleError(context, err, http.StatusInternalServerError)
		}
		context.JSON(http.StatusOK, response)
	}
}

func (c *controller) GetAccount() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		response, err := c.service.GetAccount(id)

		if err != nil {
			handleError(context, err, http.StatusInternalServerError)
		}
		context.JSON(http.StatusOK, response)
	}
}

func (c *controller) CreateAccount() gin.HandlerFunc {
	return func(context *gin.Context) {
		var createRequest requests.AccountCreateRequest
		if err := context.ShouldBindJSON(&createRequest); err != nil {
			handleError(context, err, http.StatusBadRequest)
			return
		}
		response, err := c.service.CreateAccount(createRequest)
		if err != nil {
			handleError(context, err, http.StatusInternalServerError)
			return
		}
		context.JSON(http.StatusOK, response)
	}
}
func (c *controller) UpdateAccount() gin.HandlerFunc {
	return func(context *gin.Context) {

		var updateRequest requests.AccountUpdateRequest
		id := context.Param("id")

		if err := context.ShouldBindJSON(&updateRequest); err != nil {
			handleError(context, err, http.StatusBadRequest)
			return
		}
		err := c.service.UpdateAccount(id, updateRequest)
		if err != nil {
			handleError(context, err, http.StatusInternalServerError)
			return
		}
		context.JSON(http.StatusOK, "Account updated")
	}
}

func (c *controller) DeleteAccount() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		err := c.service.DeleteAccount(id)
		if err != nil {
			handleError(context, err, http.StatusInternalServerError)
			return
		}
		context.JSON(http.StatusOK, "Account deleted")
	}
}

func handleError(context *gin.Context, err error, statusCode int) {
	log.Error(err)
	context.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}
