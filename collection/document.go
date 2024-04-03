package collection

import (
	"context"
	"fmt"
	"github.com/arangodb/go-driver"
	"strings"
)

func UpsertDocuments(ctx context.Context, db driver.Database, collectionName string, docs []map[string]any, keys []string) error {
	for _, doc := range docs {
		// Create key filter and search document for UPSERT query
		keyFilterParts := make([]string, 0, len(keys))
		for _, key := range keys {
			keyFilterParts = append(keyFilterParts, fmt.Sprintf(`"%s": @%s`, key, key))
		}
		keyFilter := "{" + strings.Join(keyFilterParts, ", ") + "}"

		// Prepare the UPSERT query
		query := fmt.Sprintf(`
			UPSERT %s
			INSERT @doc
			UPDATE @doc
			IN @@collection
		`, keyFilter)

		bindVars := map[string]interface{}{
			"@collection": collectionName,
			"doc":         doc,
		}
		for _, key := range keys {
			bindVars[key] = doc[key]
		}

		// Run the query
		cur, err := db.Query(ctx, query, bindVars)
		if err != nil {
			return fmt.Errorf("failed to execute upsert query for document %v: %w", doc, err)
		}
		cur.Close() // Close Cursor
	}

	return nil
}
