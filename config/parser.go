package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Manager struct {
	Config map[string]interface{}
}

func NewManager(configPath string) (*Manager, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config map[string]interface{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &Manager{Config: config}, nil
}

func LoadIPs(ipListPath string) ([]string, error) {
	data, err := os.ReadFile(ipListPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read IP list file: %v", err)
	}

	ipList := strings.Split(strings.TrimSpace(string(data)), "\n")
	return ipList, nil
}

func (m *Manager) UpdateConfig(ip string) error {
	// Implementation to update in-memory config with new IP
	outbounds, ok := m.Config["outbounds"].([]interface{})
	if !ok || len(outbounds) == 0 {
		return fmt.Errorf("outbounds section is missing or empty")
	}
	firstOutbound := outbounds[0].(map[string]interface{})
	settings := firstOutbound["settings"].(map[string]interface{})
	vnext := settings["vnext"].([]interface{})
	vnext[0].(map[string]interface{})["address"] = ip
	return nil
}
