package openrouter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	defaultModel = "x-ai/grok-4.1-fast:free"
	apiURL       = "https://openrouter.ai/api/v1/chat/completions"
)

type Client struct {
	apiKey string
	model  string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		model:  defaultModel,
	}
}

type chatRequest struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatResponse struct {
	Choices []struct {
		Message message `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func (c *Client) GetSummary(cveID, description string) (string, error) {
	if c.apiKey == "" {
		return "", fmt.Errorf("OpenRouter API key is missing")
	}

	prompt := fmt.Sprintf("Please provide a concise summary and potential mitigation/solution for the following vulnerability:\n\nID: %s\nDescription: %s\n\nFormat the output as:\n**Summary:** [Summary here]\n\n**Solution:** [Solution here]", cveID, description)

	reqBody := chatRequest{
		Model: c.model,
		Messages: []message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	// Optional: Add HTTP-Referer and X-Title for OpenRouter rankings
	req.Header.Set("HTTP-Referer", "http://localhost:4200")
	req.Header.Set("X-Title", "CVE Search Tool")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned error status: %d, body: %s", resp.StatusCode, string(body))
	}

	var chatResp chatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if chatResp.Error != nil {
		return "", fmt.Errorf("API error: %s", chatResp.Error.Message)
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from API")
	}

	return chatResp.Choices[0].Message.Content, nil
}
