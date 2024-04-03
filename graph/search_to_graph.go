package graph

import (
	"context"
	"fmt"
	"github.com/arangodb/go-driver"
)

type SearchToGraph struct {
	SearchViewName string
	GraphName      string
	SearchKey      string
	SearchValue    string
}

func (s SearchToGraph) SearchToGraph(ctx context.Context, db driver.Database) error {
	// Insert the SearchToGraph struct fields in the query text
	query := fmt.Sprintf(`
	LET matchingDocuments = (
	  FOR doc IN %s
		SEARCH ANALYZER(doc.%s == "%s", "identity")
		RETURN doc._id
	)
	
	FOR startVertexId IN matchingDocuments
	  FOR vertex, edge, path IN 1..1 ANY startVertexId GRAPH '%s'
		RETURN {path}
	`, s.SearchViewName, s.SearchKey, s.SearchValue, s.GraphName)

	// Run the query
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		return fmt.Errorf("failed to execute search to graph query: %w", err)
	}
	defer cursor.Close()

	// Process query results
	for {
		var data map[string]interface{}
		_, err = cursor.ReadDocument(ctx, &data)
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
