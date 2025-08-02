package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mineflared-cli/internal"
	"net/http"
)

func UpdateDNS() {
	cfg := internal.GetConfig()
	internal.LoadConfig()

	ip := getCurrentIP()
	if ip == "" {
		fmt.Println(internal.GetTranslation("DNS_IP_ERROR"))
		return
	}
	if ip != cfg.IP {
		err := updateDNSRequest(cfg.Token, ip)
		if err != nil {
			fmt.Println(internal.GetTranslation("DNS_UPDATE_ERROR"), err)
			return
		}
		cfg.IP = ip
		internal.SaveConfig()
		fmt.Println(internal.GetTranslation("DNS_UPDATE_SUCCESS"))
	} else {
		fmt.Println(internal.GetTranslation("DNS_NO_UPDATE"))
	}
}

func getCurrentIP() string {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	var ip string
	fmt.Fscanf(resp.Body, "%s", &ip)
	return ip
}

func updateDNSRequest(token, ip string) error {
	body := map[string]string{"ip": ip}
	b, _ := json.Marshal(body)
	dnsCreateURL := internal.GetEnv("DNS_CREATE_URL", "http://localhost:3000/dns/create")
	req, _ := http.NewRequest("POST", dnsCreateURL, bytes.NewReader(b))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("status %d", resp.StatusCode)
	}
	return nil
}
