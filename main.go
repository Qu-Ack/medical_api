package main

import (
	"context"
	"time"

	"github.com/Qu-Ack/medical_api/internal/auth"
	"github.com/Qu-Ack/medical_api/internal/database"
	"github.com/Qu-Ack/medical_api/internal/upload"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type state struct {
	DB *mongo.Database
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error in loading environment variables")
	}
	client := database.ConnectDB()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	database := client.Database("voyage")

	state := state{
		DB: database,
	}

	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://example.com"}, // Replace with your frontend URLs
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth.Init()
	router.GET("health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.GET("/auth/google/start", auth.BeginGoogleAuth)
	router.GET("auth/google/callback", auth.OAuthCallback)

	router.POST("/doctor", state.handlePostDoctor)

	// so to upload an image, frontend would make a request to get a presigned url,
	// then it will make a PUT request on the url with the file and Image will be uploaded
	router.GET("/upload/s3url", upload.HandleGetS3Url)

	router.Run()
}
