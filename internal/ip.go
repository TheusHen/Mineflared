package internal

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"time"
)

var ipRegex = regexp.MustCompile(`^(\d{1,3}\.){3}\d{1,3}$`)

func GetIP() string {
	ip, err := fetchIPFromExternal()
	if err == nil && ipRegex.MatchString(ip) {
		return ip
	}
	return ""
}

// Busca IP via ipwho.is
func fetchIPFromExternal() (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("https://ipwho.is/")
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", errors.New("Error fetching IP from external service")
	}
	defer resp.Body.Close()

	var result struct {
		IP      string `json:"ip"`
		Success bool   `json:"success"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	if !result.Success || !ipRegex.MatchString(result.IP) {
		return "", errors.New("Could not get valid IP")
	}
	return result.IP, nil
}
