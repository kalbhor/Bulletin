package Controllers

import (
	"../Helpers"
	"../Models"
	"github.com/gin-gonic/gin"
)

func GetServices(c *gin.Context) {
	var services []Models.Service
	err := Models.GetServices(&services)
	if err != nil {
		Helpers.RespondJSON(c, 404, services)
	} else {
		Helpers.RespondJSON(c, 200, services)
	}
}

func CreateService(c *gin.Context) {
	var service Models.Service
	c.BindJSON(&service)
	err := Models.CreateService(&service)
	if err != nil {
		Helpers.RespondJSON(c, 404, service)
	} else {
		Helpers.RespondJSON(c, 200, service)
	}
}

func GetService(c *gin.Context) {
	id := c.Params.ByName("id")
	var service Models.Service
	err := Models.GetService(&service, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, service)
	} else {
		Helpers.RespondJSON(c, 200, service)
	}
}

func UpdateService(c *gin.Context) {
	var service Models.Service
	id := c.Params.ByName("id")
	err := Models.GetService(&service, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, service)
	}
	c.BindJSON(&service)
	err = Models.UpdateService(&service, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, service)
	} else {
		Helpers.RespondJSON(c, 200, service)
	}
}

func DeleteService(c *gin.Context) {
	var service Models.Service
	id := c.Params.ByName("id")
	err := Models.DeleteService(&service, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, service)
	} else {
		Helpers.RespondJSON(c, 200, service)
	}
}