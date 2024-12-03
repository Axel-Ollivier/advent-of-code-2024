package utils

import (
	"fmt"
	"io"
	"net/http"
)

func GetInput(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Cookie", "session=53616c7465645f5fe00d45fe687795c958f4b0bf922397aa79b8c9e19df45e0404e486a9a98ceb1ef273e7402088c1984c09b9d89972fbea1db5fb72678b9fff")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return string(body), nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
