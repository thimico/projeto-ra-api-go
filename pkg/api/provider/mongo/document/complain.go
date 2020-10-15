package document

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Complain struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Locale      Locale             `bson:"locale"`
	Company     Company            `bson:"company"`
}

type Locale struct {
	City  string `bson:"city"`
	State string `bson:"state"`
}

type Company struct {
	Title       string `bson:"title"`
	Description string `bson:"state"`
}
