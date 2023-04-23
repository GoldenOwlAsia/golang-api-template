package utils

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func MakeFilenameUnique(filename string) string {
	ext := filepath.Ext(filename)
	name := filename[0 : len(filename)-len(ext)]
	name = CleanFileName(name)
	timestamp := time.Now().Unix()
	uniqueFilename := fmt.Sprintf("%s-%d%s", name, timestamp, ext)
	return uniqueFilename
}

func CleanFileName(fileName string) string {
	// Remove any characters that are not letters, numbers, or spaces
	reg := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)
	cleanedName := reg.ReplaceAllString(fileName, "_")

	// Replace any spaces with hyphens
	cleanedName = strings.ReplaceAll(cleanedName, " ", "-")

	// Convert the name to lower case
	cleanedName = strings.ToLower(cleanedName)

	return cleanedName
}
