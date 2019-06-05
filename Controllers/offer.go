package Controllers

import (
	"../Helpers"
	"../Models"
	"github.com/gin-gonic/gin"
)

func GetOffers(c *gin.Context) {
	var offers []Models.Offer
	err := Models.GetOffers(&offers)
	if err != nil {
		Helpers.RespondJSON(c, 404, offers)
	} else {
		Helpers.RespondJSON(c, 200, offers)
	}
}

func CreateOffer(c *gin.Context) {
	var offer Models.Offer
	c.BindJSON(&offer)
	err := Models.CreateOffer(&offer)
	if err != nil {
		Helpers.RespondJSON(c, 404, offer)
	} else {
		Helpers.RespondJSON(c, 200, offer)
	}
}

func GetOffer(c *gin.Context) {
	id := c.Params.ByName("id")
	var offer Models.Offer
	err := Models.GetOffer(&offer, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, offer)
	} else {
		Helpers.RespondJSON(c, 200, offer)
	}
}

func UpdateOffer(c *gin.Context) {
	var offer Models.Offer
	id := c.Params.ByName("id")
	err := Models.GetOffer(&offer, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, offer)
	}
	c.BindJSON(&offer)
	err = Models.UpdateOffer(&offer, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, offer)
	} else {
		Helpers.RespondJSON(c, 200, offer)
	}
}

func DeleteOffer(c *gin.Context) {
	var offer Models.Offer
	id := c.Params.ByName("id")
	err := Models.DeleteOffer(&offer, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, offer)
	} else {
		Helpers.RespondJSON(c, 200, offer)
	}
}