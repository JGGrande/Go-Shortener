package controller

import (
	"log"
	"net/http"
	"shortener/src/model"
	"shortener/src/service"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	var shortener model.Shortener

	if err := ctx.ShouldBindJSON(&shortener); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdShortener, err := service.Create(shortener)

	if err != nil {
		log.Println("Error creating shortener:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shortener"})
		return
	}

	ctx.JSON(201, createdShortener)
}

func RedirectBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if slug == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Slug is required"})
		return
	}

	shortener, err := service.FindBySlug(slug)

	if err != nil {
		log.Println("Error finding shortener:", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Shortener not found"})
		return
	}

	if shortener.OriginalUrl == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Original URL not found"})
		return
	}

	ctx.Redirect(http.StatusFound, shortener.OriginalUrl)
}
