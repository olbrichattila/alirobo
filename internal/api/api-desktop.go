package api

import (
	"alibabarobotgame/internal/defaultconfig"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type desktopClient struct{}

// UserScore is a name, score and created at
type UserScore struct {
	Name      string
	Score     int
	CreatedAt string
}

// List of user scores
type UserScores []UserScore

func New() APIClient {
	return &desktopClient{}
}

func (c *desktopClient) AddScore(name string, score int) error {
	data := map[string]any{
		"name":  name,
		"score": score,
	}
	requestBody, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshaling error: %w", err)
	}

	req, err := http.NewRequest("POST", defaultconfig.ApiUrl+"add", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("request creation error: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return nil
}

func (c *desktopClient) Top10() (UserScores, error) {
	fmt.Println(defaultconfig.ApiUrl + "top")
	resp, err := http.Get(defaultconfig.ApiUrl + "top")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var scores UserScores
	err = json.NewDecoder(resp.Body).Decode(&scores)
	if err != nil {
		return nil, err
	}

	return scores, nil
}
