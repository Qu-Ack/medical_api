package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Qualification struct {
	Title     string `bson:"title,omitempty" json:"title"`
	College   string `bson:"college,omitempty" json:"college"`
	StartYear string `bson:"start_year,omitempty" json:"start_year"`
	EndYear   string `bson:"end_year,omitempty" json:"end_year"`
}

type Doctor struct {
	ID                             primitive.ObjectID `bson:"_id,omitempty"`
	Name                           string             `bson:"name,omitempty" json:"name"`
	MedicalRegistrationNumber      string             `bson:"medical_registration_number,omitempty" json:"medical_registration_number"`
	Phone                          string             `bson:"phone,omitempty" json:"phone"`
	Specialization                 string             `bson:"specialization,omitempty" json:"specialization"`
	Qualifications                 []Qualification    `bson:"qualifications,omitempty" json:"qualifications"`
	MedicalRegistrationDocumentUrl string             `bson:"medical_registration_document_url,omitempty" json:"medical_registration_document_url"`
	JoinDate                       string             `bson:"join_date,omitempty" json:"join_date"`
}

func ConnectDB() *mongo.Client {

	mongoDbURI := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDbURI))

	if err != nil {
		panic(err)
	}

	return client

}
