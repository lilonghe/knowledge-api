package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"lilonghe.net/knowledge/config"
	"lilonghe.net/knowledge/middleware"
	"lilonghe.net/knowledge/models"
)

var DB = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Get user value
	r.GET("/knowledge/all", func(c *gin.Context) {

		err, datas := models.GetKnowledgeList()
		err2, premises := models.GetAllKnowledgePremise()
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(200, gin.H{"error": err})
			return
		}
		if err2 != nil {
			fmt.Printf("%+v", err)
			c.JSON(200, gin.H{"error": err})
			return
		}

		c.JSON(200, gin.H{
			"knowledges":        datas,
			"knowledgePremises": premises,
		})
	})

	r.POST("/knowledge/add", func(c *gin.Context) {
		item := models.Knowledge{}
		err := c.ParseBody(&item)

		if err != nil {
			c.JSON(200, gin.H{"error": err.Error()})
		} else {
			err, id := models.AddKnowledge(item)
			if err != nil {
				c.JSON(200, gin.H{"error": err})
			} else {
				c.JSON(200, gin.H{"id": id})
			}
		}

	})

	return r
}

func main() {
	config.Init()
	migrateDB()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func migrateDB() {
	db := config.Store.Master()
	db.AutoMigrate(&models.Knowledge{})
	db.AutoMigrate(&models.KnowledgePremise{})
}
