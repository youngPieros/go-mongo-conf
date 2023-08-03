package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DAO struct {
	databaseName   string
	collectionName string
	collection     *mongo.Collection
	ctx            context.Context
}

var GetDAO = func() func() *DAO {
	var instance *DAO
	return func() *DAO {
		if instance == nil {
			databaseName := MongoDataBase()
			collectionName := "panicmodes"
			collection := GetMongoClient().GetCollection(databaseName, collectionName)
			ctx := GetMongoClient().GetContext()
			instance = &DAO{databaseName, collectionName, collection, ctx}
		}
		return instance
	}
}()

func (dao *DAO) Load(name string) []PanicModeVariable {
	var table PanicTable
	filter := bson.M{"name": name}
	if err := dao.collection.FindOne(dao.ctx, filter).Decode(&table); err != nil {
		return []PanicModeVariable{}
	}
	return table.Variables
}

func (dao *DAO) DeleteVariable(table, variable string) error {
	filter := bson.M{"name": table}
	update := bson.M{"$pull": bson.M{"variables": bson.M{"name": variable}}}
	result, _ := dao.collection.UpdateMany(dao.ctx, filter, update)
	if result == nil || result.ModifiedCount != 1 {
		return errors.New("panic variable not found")
	}
	return nil
}

func (dao *DAO) DeleteTable(table string) error {
	filter := bson.M{"name": table}
	result, _ := dao.collection.DeleteOne(dao.ctx, filter)
	if result == nil || result.DeletedCount != 1 {
		return errors.New("panic table not found")
	}
	return nil
}

func (dao *DAO) SaveTable(table string, variables []PanicModeVariable) error {
	filter := bson.M{"name": table}
	update := bson.M{"$set": bson.M{"variables": variables}}
	option := options.Update().SetUpsert(true)
	_, _ = dao.collection.UpdateOne(dao.ctx, filter, update, option)
	return nil
}
