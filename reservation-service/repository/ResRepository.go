package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ResRepository struct {
	DatabaseConnection *mongo.Client
}
