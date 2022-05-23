package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"art-museum/artist"
)

type contentHandler struct {
	artistService artist.Service
}

func NewArtistHandler(artistService artist.Service) *contentHandler {
	return &contentHandler{artistService}
}

func ResponCreator(a artist.Artist) artist.CreatorRespon {
	return artist.CreatorRespon{
		ID:          a.ID,
		Name:        a.Name,
		Nationality: a.Nationality,
		Description: a.Description,
		Birth_year:  a.Birth_year,
		Death_year:  a.Death_year,
	}
}

func (h *contentHandler) GetData(c *gin.Context) {
	content, err := h.artistService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var respon []artist.CreatorRespon

	for _, b := range content {
		creatorRespon := ResponCreator(b)

		respon = append(respon, creatorRespon)

	}

	c.JSON(http.StatusOK, gin.H{
		"data": content,
	})
}

func (h *contentHandler) GetContentByID(c *gin.Context) {
	idstring := c.Param("id")
	id, _ := strconv.Atoi(idstring)

	b, err := h.artistService.FindByID(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	content := ResponCreator(b)

	c.JSON(http.StatusOK, gin.H{
		"data": content,
	})
}

func (h *contentHandler) CreatorHandler(c *gin.Context) {
	var contentRequest artist.Creator

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

	content, err := h.artistService.Create(contentRequest)

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

func (h *contentHandler) UpdateHandler(c *gin.Context) {
	var contentRequest artist.Creator

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

	content, err := h.artistService.Update(id, contentRequest)

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

func (h *contentHandler) DeleteHandler(c *gin.Context) {
	idstring := c.Param("id")
	id, _ := strconv.Atoi(idstring)

	b, err := h.artistService.Delete(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	content := ResponCreator(b)

	c.JSON(http.StatusOK, gin.H{
		"data": content,
	})
}
