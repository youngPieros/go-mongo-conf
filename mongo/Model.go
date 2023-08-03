package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PanicTable struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	Name      string              `bson:"name"`
	Variables []PanicModeVariable `bson:"variables"`
}

type PanicModeVariable struct {
	Name  string `bson:"name"`
	Type  string `bson:"type"`
	Value string `bson:"value"`
}
