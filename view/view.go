package view

import (
	"context"
	"fmt"
	"github.com/MelihEmreGuler/arango-methods/db"
	"github.com/arangodb/go-driver"
)

var ctx = context.Background()

func CreateSearchView(name string) error {
	//create options *ArangoSearchViewProperties
	viewOptions := &driver.ArangoSearchViewProperties{

		Links: map[string]driver.ArangoSearchElementProperties{},
	}

	_, err := db.Database.CreateArangoSearchView(ctx, name, viewOptions)
	if err != nil {
		return fmt.Errorf("failed to create ArangoSearch view: %w", err)
	}
	return nil
}

func CreateSearchAliasView(name string) error {
	//create options *ArangoSearchAliasViewProperties
	viewOptions := &driver.ArangoSearchAliasViewProperties{}

	_, err := db.Database.CreateArangoSearchAliasView(ctx, name, viewOptions)
	if err != nil {
		return fmt.Errorf("failed to create ArangoSearchAlias view: %w", err)
	}

	return nil
}
