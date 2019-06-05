package Controllers

import (
	"../Helpers"
	"../Models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var categories []Models.User
	err := Models.GetUsers(&categories)
	if err != nil {
		Helpers.RespondJSON(c, 404, categories)
	} else {
		Helpers.RespondJSON(c, 200, categories)
	}
}

func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	if err != nil {
		Helpers.RespondJSON(c, 404, user)
	} else {
		Helpers.RespondJSON(c, 200, user)
	}
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUser(&user, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, user)
	} else {
		Helpers.RespondJSON(c, 200, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUser(&user, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, user)
	} else {
		Helpers.RespondJSON(c, 200, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, user)
	} else {
		Helpers.RespondJSON(c, 200, user)
	}
}

func AddOfferRequest(c *gin.Context) {
	var offerrequest Models.OfferRequest
	c.BindJSON(&offerrequest)

	id := c.Params.ByName("id")

	err := Models.AddOfferRequest(offerrequest, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, offerrequest)
	} else {
		Helpers.RespondJSON(c, 200, offerrequest)
	}
}

func AddOffer(c *gin.Context) {
	var offer Models.Offer
	c.BindJSON(&offer)

	id := c.Params.ByName("id")

	err := Models.AddOffer(offer, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, offer)
	} else {
		Helpers.RespondJSON(c, 200, offer)
	}
}

