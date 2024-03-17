package queries

import (
	"context"
	"fmt"
	"os"

	"github.com/bit-wiz/data-store-a/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	col *mongo.Collection
}

var DB Mongo

func NewMongo(name string) error {
	opts := options.Client().ApplyURI(os.Getenv("MONGO_ME"))
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return err
	}

	collection := client.Database("somedb").Collection(name)

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB!")
	DB.col = collection

	return nil
}

func (m *Mongo) GetUsers(filter interface{}, findOptions options.FindOptions) ([]models.Citizen, error) {
	cur, err := m.col.Find(context.Background(), filter, &findOptions)
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())

	var data []models.Citizen

	for cur.Next(context.Background()) {
		var citizen models.Citizen
		if err := cur.Decode(&citizen); err != nil {
			return nil, err
		}
		data = append(data, citizen)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *Mongo) GetUser(id primitive.ObjectID) (models.Citizen, error) {
	var data models.Citizen
	err := m.col.FindOne(context.Background(), &bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&data)
	if err != nil {
		return models.Citizen{}, err
	}

	return data, nil
}

func (m *Mongo) CreateUser(user models.Citizen) error {
	_, err := m.col.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (m *Mongo) UpdateUser(id primitive.ObjectID, user models.Citizen) error {
	_, err := m.col.UpdateOne(context.Background(), &bson.D{primitive.E{Key: "_id", Value: id}}, &bson.D{primitive.E{Key: "$set", Value: user}})
	if err != nil {
		return err
	}
	return nil
}

func (m *Mongo) DeleteUser(id primitive.ObjectID) error {
	_, err := m.col.DeleteOne(context.Background(), &bson.D{primitive.E{Key: "_id", Value: id}})
	if err != nil {
		return err
	}
	return nil
}
