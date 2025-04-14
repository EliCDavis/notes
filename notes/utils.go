package notes

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"
)

func removeNonAlphanumeric(s string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(s, "")
}

func sanitizeTags(dirty []string) []string {
	out := make([]string, 0, len(dirty))

	for _, tag := range dirty {
		clean := removeNonAlphanumeric(tag)
		if clean == "" {
			continue
		}
		out = append(out, strings.ToLower(clean))
	}

	return out
}

func LoadJSON[T any](filename string) (T, error) {
	var data T
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return data, err
	}
	return data, json.Unmarshal(fileData, &data)
}

func SaveJSON(filename string, data any) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	content, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = file.Write(content)
	return err
}
