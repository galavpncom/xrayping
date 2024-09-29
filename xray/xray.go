package xray

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	"github.com/galavpncom/xrayping/utils"
)

type Manager struct {
	BinaryPath string
	Verbose    bool
}

func NewManager(binaryPath string, verbose bool) *Manager {
	return &Manager{BinaryPath: binaryPath, Verbose: verbose}
}

func (xm *Manager) Start(config map[string]interface{}) (*exec.Cmd, error) {
	configBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal in-memory config: %v", err)
	}

	cmd := exec.Command(xm.BinaryPath, "-c", "stdin:")
	cmd.Stdin = bytes.NewReader(configBytes)

	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("failed to start Xray: %v", err)
	}

	if xm.Verbose {
		fmt.Println(utils.WrapTextWithColor("Xray proxy started.", utils.ColorYellow))
	}

	time.Sleep(2 * time.Second) // Wait for Xray to establish the proxy
	return cmd, nil
}

func (xm *Manager) Stop(cmd *exec.Cmd) error {
	if err := cmd.Process.Kill(); err != nil {
		return fmt.Errorf("failed to stop Xray: %v", err)
	}
	if xm.Verbose {
		fmt.Println(utils.WrapTextWithColor("Xray proxy stopped.", utils.ColorYellow))
	}
	return nil
}
