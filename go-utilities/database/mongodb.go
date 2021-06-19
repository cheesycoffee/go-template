package database

import (
	"context"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectMongoClient = mongo.Connect

type (
	mongoInstance struct {
		read, write *mongo.Database
	}

	// MongoOption client connection
	MongoOption struct {
		ReadURI, WriteURI string
		DatabaseName      string
		ConnectTimeOut    time.Duration
	}

	// MongoDB clients
	MongoDB interface {
		GetReadDB() *mongo.Database
		GetWriteDB() *mongo.Database
	}
)

func newMongoDB(option MongoOption) MongoDB {
	mongoInst := new(mongoInstance)
	mongoInst.read = getMongoDB(option.ReadURI, option.DatabaseName, option.ConnectTimeOut)
	mongoInst.write = getMongoDB(option.WriteURI, option.DatabaseName, option.ConnectTimeOut)
	return mongoInst
}

func (m *mongoInstance) GetReadDB() *mongo.Database {
	return m.read
}

func (m *mongoInstance) GetWriteDB() *mongo.Database {
	return m.write
}

func getMongoDB(dsn, dbName string, connectTimeout time.Duration) *mongo.Database {
	if strings.TrimSpace(dsn) == "" || strings.TrimSpace(dbName) == "" {
		return nil
	}

	if connectTimeout <= 0*time.Second {
		connectTimeout = 30 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	client, err := connectMongoClient(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return client.Database(dbName)
}
