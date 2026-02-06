package ai_processor

import (
	"bufio"
	"bytes"
	"encoding/json"
	"feelgoodfeatures/internal/domain"
	"net/http"
	"strings"
)

func GenerateArticleSummary(prompt string) (string, error) {
	ollamaRequest := domain.OllamaRequest{
		//Model:  "llama3.2",
		Model:  "llama3.2:1b",
		Prompt: prompt,
	}

	ollamaRequestBody, _ := json.Marshal(ollamaRequest)
	ollamaResponseRaw, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(ollamaRequestBody))
	if err != nil {
		return "", err
	}
	defer ollamaResponseRaw.Body.Close()

	var articleSummary strings.Builder

	scanner := bufio.NewScanner(ollamaResponseRaw.Body)
	for scanner.Scan() {
		var ollamaResponse domain.OllamaResponse
		if err := json.Unmarshal(scanner.Bytes(), &ollamaResponse); err == nil {
			articleSummary.WriteString(ollamaResponse.Response)
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return articleSummary.String(), nil
}
