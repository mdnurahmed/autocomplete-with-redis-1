package controllers

import (
	"autocomplete/app/DTO"
	"autocomplete/app/services"
	"autocomplete/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type IAutocompleteController interface {
	Search(c *gin.Context)
	Insert(c *gin.Context)
	Delete(c *gin.Context)
}

type AutocompleteController struct {
	autocompleteService services.IAutocompleteService
}

func NewInstanceOfAutocompleteController(
	autocompleteService services.IAutocompleteService) AutocompleteController {
	return AutocompleteController{autocompleteService: autocompleteService}
}

func (a *AutocompleteController) Search(c *gin.Context) {
	m := DTO.Request{}
	m.Word = c.DefaultQuery("Word", "")
	if m.Word == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": DTO.NewErrorResponse("NoWordProvided", "Word has to be provided"),
		})
		return
	}
	result, err := a.autocompleteService.Search(m.Word)
	if err != nil {
		log.WithFields(log.Fields{
			"error_message": err.Error(),
		})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": DTO.NewErrorResponse("ServiceError", "Something Went Wrong"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})

}

func (a *AutocompleteController) Insert(c *gin.Context) {
	m := DTO.Request{}
	err := c.Bind(&m)
	if err != nil || m.Word == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": DTO.NewErrorResponse("NoWordProvided", "Word has to be provided"),
		})
		return
	}
	err = a.autocompleteService.Insert(m.Word)
	if err != nil {
		log.WithFields(log.Fields{
			"error_message": err.Error(),
		})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": DTO.NewErrorResponse("ServiceError", "Something Went Wrong"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Word inserted successfully",
	})
}

func (a *AutocompleteController) Delete(c *gin.Context) {

	err := a.autocompleteService.Delete(utils.Configuration.KeyName)
	if err != nil {
		log.WithFields(log.Fields{
			"erro_message": err.Error(),
		})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": DTO.NewErrorResponse("ServiceError", "Something Went Wrong"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Deleted successfully",
	})
}
