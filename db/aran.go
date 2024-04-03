package db

import (
	"context"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"log"
)

var Database driver.Database

func ConnectArangoDB(dbURL, username, password, database string) error {
	ctx := context.Background()
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{dbURL},
	})
	if err != nil {
		log.Fatalf("Failed to create HTTP connection: %v", err)
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(username, password),
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	Database, err = client.Database(ctx, database)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	return nil
}
