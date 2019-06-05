package Routers

import (
	"../Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/category/", Controllers.GetCategories)
	r.GET("/category/:id", Controllers.GetCategory)
	r.POST("/category", Controllers.CreateCategory)
	r.PUT("/category/:id", Controllers.UpdateCategory)
	r.DELETE("/category/:id", Controllers.DeleteCategory)

	r.GET("/service/", Controllers.GetServices)
	r.GET("/service/:id", Controllers.GetService)
	r.POST("/service", Controllers.CreateService)
	r.PUT("/service/:id", Controllers.UpdateService)
	r.DELETE("/service/:id", Controllers.DeleteService)

	r.GET("/offer/", Controllers.GetOffers)
	r.GET("/offer/:id", Controllers.GetOffer)
	r.POST("/offer", Controllers.CreateOffer)
	r.PUT("/offer/:id", Controllers.UpdateOffer)
	r.DELETE("/offer/:id", Controllers.DeleteOffer)

	r.GET("/user/", Controllers.GetUsers)
	r.GET("/user/:id", Controllers.GetUser)
	r.POST("/user", Controllers.CreateUser)
	r.PUT("/user/:id", Controllers.UpdateUser)
	r.DELETE("/user/:id", Controllers.DeleteUser)

	r.POST("/user/offer", Controllers.AddOffer)
	r.POST("/user/offer_request", Controllers.AddOfferRequest)

	return r

}