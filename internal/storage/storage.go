package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func Save(data map[string]string) {
	storageType := viper.GetString("storage.type")

	switch storageType {

	case "file":
		saveToFile(data)
	case "stdout":
		fmt.Println(data)
	default:
		log.Fatalf("Unsupported storage type: %s", storageType)
	}
}

func saveToFile(data map[string]string) {
	filePath := viper.GetString("storage.file_path")
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	if _, err := file.Write(jsonData); err != nil {
		log.Fatalf("Failed to write data to file: %v", err)
	}

	if _, err := file.WriteString("\n"); err != nil {
		log.Fatalf("Failed to write newline to file: %v", err)
	}
}
