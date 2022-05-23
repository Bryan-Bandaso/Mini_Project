package handler

import (
	"art-museum/art"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ArtHandler struct {
	artService art.Service
}

func NewArtHandler(artService art.Service) *ArtHandler {
	return &ArtHandler{artService}
}

func ResponArt(a art.Art) art.ArtRespon {
	return art.ArtRespon{
		ID:            a.ID,
		Name_Art:      a.Name_Art,
		Name_Artist:   a.Name_Artist,
		Location:      a.Location,
		Type:          a.Type,
		Creation_Date: a.Creation_Date,
	}
}

func (h *ArtHandler) GetDataArt(c *gin.Context) {
	content, err := h.artService.FindAllArt()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var respon []art.ArtRespon

	for _, b := range content {
		artRespon := ResponArt(b)

		respon = append(respon, artRespon)

	}

	c.JSON(http.StatusOK, gin.H{
		"data": content,
	})
}

func (h *ArtHandler) GetContentArtByID(c *gin.Context) {
	idstring := c.Param("id")
	id, _ := strconv.Atoi(idstring)

	b, err := h.artService.FindArtByID(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	content := ResponArt(b)

	c.JSON(http.StatusOK, gin.H{
		"data": content,
	})
}

func (h *ArtHandler) CreateArt(c *gin.Context) {
	var contentRequest art.Arts

	err := c.ShouldBindJSON(&contentRequest)
	if err != nil {
		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	content, err := h.artService.CreateArt(contentRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": content,
	})
}

func (h *ArtHandler) UpdateArt(c *gin.Context) {
	var contentRequest art.Arts

	err := c.ShouldBindJSON(&contentRequest)
	if err != nil {
		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	idstring := c.Param("id")
	id, _ := strconv.Atoi(idstring)

	content, err := h.artService.UpdateArt(id, contentRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": content,
	})
}

func (h *ArtHandler) DeleteArt(c *gin.Context) {
	idstring := c.Param("id")
	id, _ := strconv.Atoi(idstring)

	b, err := h.artService.DeleteArt(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	content := ResponArt(b)

	c.JSON(http.StatusOK, gin.H{
		"data": content,
	})
}
