package main

import (
	"fmt"
	"log"

	"github.com/Qu-Ack/medical_api/internal/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Qualification struct {
	Title      string `json:"title"`
	College    string `json:"college"`
	Start_year string `json:"start_year"`
	End_year   string `json:"end_year"`
}

func (s state) handlePostDoctor(c *gin.Context) {
	type body struct {
		MedicalRegistrationNumber      string          `json:"medical_registration_number"`
		Name                           string          `json:"name"`
		MedicalRegistrationDocumentUrl string          `json:"medical_registration_document_url"`
		Qualifications                 []Qualification `json:"qualifications"`
		Phone                          string          `json:"phone"`
		Specialization                 string          `json:"specialization"`
		JoinDate                       string          `json:"join_date"`
	}
	var doctor database.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	doctor.ID = primitive.NewObjectID()

	// Print out the body of the request
	fmt.Printf("Received Body: %+v\n", doctor)

	doctorsCollection := s.DB.Collection("doctors")

	result, err := doctorsCollection.InsertOne(c.Request.Context(), doctor)

	if err != nil {
		log.Println("Error inserting doctor into MongoDB:", err)
		c.JSON(500, gin.H{"error": "Failed to insert doctor"})
		return
	}

	// Respond with the received body
	c.JSON(200, gin.H{"message": "doctor added successfully", "result": result})

}
