package mongo

import (
	"context"
	"fmt"
	"github.com/youngPieros/go-mongo-conf/tools"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	ctx    context.Context
	client *mongo.Client
}

func createMongoClient() *MongoClient {
	client := MongoClient{}
	client.setClient()
	client.setContext()
	client.connect()
	client.ping()
	return &client
}

func (mongoClient *MongoClient) getMongoURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		MongoUser(), MongoPassword(), MongoHost(), MongoPort(), MongoDataBase(),
	)
}

func (mongoClient *MongoClient) setClient() {
	address := options.Client().ApplyURI(mongoClient.getMongoURI())
	client, err := mongo.NewClient(address)
	if err != nil {
		tools.Logger.Panicw("MONGOCONF_MONGO_CLIENT_CREATION_ERROR", "MONGO", mongoClient.getMongoURI())
	}
	mongoClient.client = client
}

func (mongoClient *MongoClient) setContext() {
	mongoClient.ctx = context.TODO()
}

func (mongoClient *MongoClient) disconnect() {
	if err := mongoClient.client.Disconnect(mongoClient.ctx); err != nil {
		tools.Logger.Panicw("MONGOCONF_MONGO_DISCONNECTION_ERROR", "MONGO", mongoClient.getMongoURI())
	}
}

func (mongoClient *MongoClient) connect() {
	if err := mongoClient.client.Connect(mongoClient.ctx); err != nil {
		tools.Logger.Panicw("MONGOCONF_MONGO_CONNECTION_ERROR", "MONGO", mongoClient.getMongoURI())
	}
}

func (mongoClient *MongoClient) ping() {
	if err := mongoClient.client.Ping(mongoClient.ctx, readpref.Primary()); err != nil {
		tools.Logger.Panicw("MONGOCONF_MONGO_PING_ERROR", "MONGO", mongoClient.getMongoURI())
	}
}

func (mongoClient *MongoClient) GetCollection(database, collection string) *mongo.Collection {
	return mongoClient.client.Database(database).Collection(collection)
}

func (mongoClient *MongoClient) GetContext() context.Context {
	return mongoClient.ctx
}

var GetMongoClient = func() func() *MongoClient {
	var instance *MongoClient
	return func() *MongoClient {
		if instance == nil {
			instance = createMongoClient()
		}
		return instance
	}
}()
