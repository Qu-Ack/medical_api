package upload

import (
	"log"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

var sess, err = session.NewSession(&aws.Config{
	Region: aws.String("us-east-1"),
})

var svc = s3.New(sess)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateImageName(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)

}

func GetPresignedURL() (string, error) {
	imageName := GenerateImageName(8)

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String("kanteentest"),
		Key:    aws.String(imageName),
	})

	str, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Println("ERROR: GetPresignURL", err)
		return "", err
	}

	return str, nil
}

func HandleGetS3Url(ctx *gin.Context) {
	url, err := GetPresignedURL()

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{"status": "success", "url": url})
}
