package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Connection interface {
	Close()
	DB() *mongo.Database
}

type conn struct {
	session  *mongo.Client
	database *mongo.Database
}

func NewConnection(cfg Config) (Connection, error) {
	fmt.Println("database url:", cfg.GetDSN())
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	session, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.GetDSN()))
	if err != nil {
		return nil, err
	}
	err = session.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	databases, err := session.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	return &conn{session: session,
		database: session.Database(cfg.GetDBName())}, nil
}

func (c *conn) Close() {
	err := c.session.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
}

func (c *conn) DB() *mongo.Database {
	return c.database
}
