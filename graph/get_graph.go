package graph

import (
	"context"
	"fmt"
	"github.com/arangodb/go-driver"
)

func GetGraphData(ctx context.Context, db driver.Database, graphName string) error {
	// Insert the graph name directly in the query text
	query := fmt.Sprintf(`
        FOR v, e IN 1..2 ANY @startVertex GRAPH '%s'
        RETURN {vertex: v, edge: e}
    `, graphName)

	// Query parameters
	bindVars := map[string]interface{}{
		"startVertex": "test11/135788", // ID of the start node
	}

	// Run the query
	cursor, err := db.Query(ctx, query, bindVars)
	if err != nil {
		return fmt.Errorf("failed to execute graph query: %w", err)
	}
	defer cursor.Close()

	// Process query results
	for {
		var data map[string]interface{}
		_, err := cursor.ReadDocument(ctx, &data)
		if driver.IsNoMoreDocuments(err) {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read document: %w", err)
		}
		fmt.Printf("Data: %v\n", data)
	}

	return nil
}
