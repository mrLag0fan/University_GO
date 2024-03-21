package scanner

import (
	"O1/internal/model"
	"encoding/json"
	"log"
	"os"
)

func Scan(path string) []model.Route {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var routes []model.Route
	err = json.Unmarshal(fileContent, &routes)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return routes
}
