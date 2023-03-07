package storage

import (
	"context"
	"fmt"
	"log"
	"nova/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	config *utils.Configuration
	client *mongo.Client
	link   string
}

func NewMongo(c *utils.Configuration) *Mongo {
	l := fmt.Sprintf(
		"mongodb://%s:%s@%s%s",
		c.Database.User, c.Database.Pass, c.Database.Host, c.Database.Port,
	)

	m := &Mongo{
		config: c,
		client: nil,
		link:   l,
	}
	return m
}

// public functions

func (m *Mongo) StoreObject(c string, a any) {
	m.connect()

	directory := m.client.Database(m.config.Database.Db).Collection(c)
	_, err := directory.InsertOne(context.TODO(), a)
	if err != nil {
		return
	}
}

func (m *Mongo) RefreshObject(c string, id any, a any) {
	m.connect()

	directory := m.client.Database(m.config.Database.Db).Collection(c)
	_, err := directory.ReplaceOne(context.TODO(), bson.M{"_id": id}, a)
	if err != nil {
		return
	}
}

func (m *Mongo) DeleteObject(c string, id any) {
	m.connect()

	directory := m.client.Database(m.config.Database.Db).Collection(c)
	_, err := directory.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return
	}
}

func (m *Mongo) RetrieveObject(c string, id any) *mongo.SingleResult {
	m.connect()

	directory := m.client.Database(m.config.Database.Db).Collection(c)
	result := directory.FindOne(context.TODO(), bson.M{"_id": id})
	return result
}

func (m *Mongo) SearchObject(c string, k string, v string) *mongo.SingleResult {
	m.connect()

	directory := m.client.Database(m.config.Database.Db).Collection(c)
	result := directory.FindOne(context.TODO(), bson.M{k: v})
	return result
}

func (m *Mongo) ListObjects(c string) *mongo.Cursor {
	m.connect()

	directory := m.client.Database(m.config.Database.Db).Collection(c)
	result, err := directory.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil
	}

	return result
}

// private functions

func (m *Mongo) connect() {
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.link))
	if err != nil {
		log.Fatal(err)
	}
	m.client = c
}
