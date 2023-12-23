package database

import (
	"strconv"
	"strings"
)

func (a *databaseAccessor) UploadFile(in NewMeme) (int, error) {
	// Build the insert query
	query, values, err := buildInsertQuery(in.Name, in.Content, in.Tags, in.Url, in.Description)
	if err != nil {
		return 0, err
	}

	// Execute the insert query
	var fileID int
	err = a.db.QueryRow(query, values...).Scan(&fileID)
	if err != nil {
		return 0, err
	}

	return fileID, nil
}

// buildInsertQuery dynamically builds the insert query based on content, URL, and tags presence
func buildInsertQuery(name string, content []byte, tags []string, url, description string) (string, []interface{}, error) {
	var queryBuilder strings.Builder
	var values []interface{}

	queryBuilder.WriteString("INSERT INTO files (name")
	values = append(values, name)

	if content != nil {
		queryBuilder.WriteString(", content")
		values = append(values, content)
	}

	if len(tags) > 0 {
		queryBuilder.WriteString(", tags")
		values = append(values, tags)
	}

	if url != "" {
		queryBuilder.WriteString(", url")
		values = append(values, url)
	}

	if description != "" {
		queryBuilder.WriteString(", description")
		values = append(values, description)
	}

	queryBuilder.WriteString(") VALUES (")

	for i := 1; i <= len(values); i++ {
		queryBuilder.WriteString("$" + strconv.Itoa(i))
		if i < len(values) {
			queryBuilder.WriteString(", ")
		}
	}

	queryBuilder.WriteString(") RETURNING id;")

	return queryBuilder.String(), values, nil
}
