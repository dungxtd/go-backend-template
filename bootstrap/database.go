package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sportgo-app/sportgo-go/mongo"
	"github.com/sportgo-app/sportgo-go/postgres"
)

func NewMongoDatabase(env *Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.MongoDBHost
	dbPort := env.MongoDBPort
	dbUser := env.MongoDBUser
	dbPass := env.MongoDBPass

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	client, err := mongo.NewClient(mongodbURI)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}

func NewPostgresDatabase(env *Env) postgres.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.PostgresDBHost
	dbPort := env.PostgresDBPort
	dbUser := env.PostgresDBUser
	dbPass := env.PostgresDBPass
	dbName := env.PostgresDBName
	sslMode := env.PostgresSSLMode

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslMode)

	if dbUser == "" || dbPass == "" {
		dsn = fmt.Sprintf("postgresql://%s:%s/%s?sslmode=%s", dbHost, dbPort, dbName, sslMode)
	}

	client, err := postgres.NewClient(dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
