package main

import (
	"fmt"

	"github.com/Qu-Ack/medical_api/internal/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	auth.Init()
	router.GET("health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.GET("/auth/google/start", auth.BeginGoogleAuth)
	router.GET("auth/google/callback", auth.OAuthCallback)

	router.Run()

	fmt.Println("Hello World")
}
