package main

import (
	"art-museum/art"
	"art-museum/artist"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"art-museum/handler"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/art-museum?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection error")
	}
	fmt.Println("DB connection Succes")
	db.AutoMigrate(&artist.Artist{})
	db.AutoMigrate(&art.Art{})

	contentRepository := artist.NewRepository(db)
	contentService := artist.NewService(contentRepository)
	contentHandler := handler.NewArtistHandler(contentService)

	ArtRepository := art.NewRepository(db)
	ArtService := art.NewService(ArtRepository)
	ArtHandler := handler.NewArtHandler(ArtService)

	router := gin.Default()
	v1 := router.Group("/v1")

	//Artist
	v1.POST("/creator", contentHandler.CreatorHandler)
	v1.GET("/creator", contentHandler.GetData)
	v1.GET("/creator/:id", contentHandler.GetContentByID)
	v1.PUT("/creator/:id", contentHandler.UpdateHandler)
	v1.DELETE("/creator/:id", contentHandler.DeleteHandler)

	//Art
	v1.POST("/art", ArtHandler.CreateArt)
	v1.GET("/art", ArtHandler.GetDataArt)
	v1.GET("/art/:id", ArtHandler.GetContentArtByID)
	v1.PUT("/art/:id", ArtHandler.UpdateArt)
	v1.DELETE("/art/:id", ArtHandler.DeleteArt)

	router.Run()
}
