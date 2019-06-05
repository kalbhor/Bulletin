package Controllers

import (
	"../Helpers"
	"../Models"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []Models.Category
	err := Models.GetCategories(&categories)
	if err != nil {
		Helpers.RespondJSON(c, 404, categories)
	} else {
		Helpers.RespondJSON(c, 200, categories)
	}
}

func CreateCategory(c *gin.Context) {
	var category Models.Category
	c.BindJSON(&category)
	err := Models.CreateCategory(&category)
	if err != nil {
		Helpers.RespondJSON(c, 404, category)
	} else {
		Helpers.RespondJSON(c, 200, category)
	}
}

func GetCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	var category Models.Category
	err := Models.GetCategory(&category, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, category)
	} else {
		Helpers.RespondJSON(c, 200, category)
	}
}

func UpdateCategory(c *gin.Context) {
	var category Models.Category
	id := c.Params.ByName("id")
	err := Models.GetCategory(&category, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, category)
	}
	c.BindJSON(&category)
	err = Models.UpdateCategory(&category, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, category)
	} else {
		Helpers.RespondJSON(c, 200, category)
	}
}

func DeleteCategory(c *gin.Context) {
	var category Models.Category
	id := c.Params.ByName("id")
	err := Models.DeleteCategory(&category, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, category)
	} else {
		Helpers.RespondJSON(c, 200, category)
	}
}