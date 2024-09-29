package cmd

import (
	"fmt"
	"sync"

	"github.com/galavpncom/xrayping/config"
	"github.com/galavpncom/xrayping/latency"
	"github.com/galavpncom/xrayping/utils"
	"github.com/galavpncom/xrayping/xray"
	"github.com/spf13/cobra"
)

var (
	subnetListFile          string
	ipCount                 int
	randomTestConfigPath    string
	randomTestXrayPath      string
	randomTestSocks5Address string
	randomTestRetries       int
	randomTestTestURL       string
)

var randomTestCmd = &cobra.Command{
	Use:   "random-test",
	Short: "Randomly select IPs from a subnet list and test latency",
	Run: func(cmd *cobra.Command, args []string) {
		// Load the subnet list
		subnets, err := config.LoadIPs(subnetListFile)
		if err != nil {
			fmt.Println(utils.WrapTextWithColor(fmt.Sprintf("Error loading subnet list: %v", err), utils.ColorRed))
			return
		}

		// Generate random IPs from the subnets
		randomIPs, err := utils.GenerateRandomIPsFromSubnets(subnets, ipCount)
		if err != nil {
			fmt.Println(utils.WrapTextWithColor(fmt.Sprintf("Error generating random IPs: %v", err), utils.ColorRed))
			return
		}

		runLatencyTestForIPs(randomIPs)
	},
}

func runLatencyTestForIPs(ips []string) {
	configManager, err := config.NewManager(randomTestConfigPath)
	if err != nil {
		fmt.Println(utils.WrapTextWithColor(fmt.Sprintf("Error: %v", err), utils.ColorRed))
		return
	}

	xrayManager := xray.NewManager(randomTestXrayPath, false)
	latencyTester := latency.NewTester(randomTestTestURL, randomTestSocks5Address, false)

	concurrencyLimit := 5
	semaphore := make(chan struct{}, concurrencyLimit)

	var wg sync.WaitGroup
	var successIPs []latency.IPResult
	var mu sync.Mutex

	for _, ip := range ips {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			semaphore <- struct{}{}
			latency.ProcessIP(ip,
				configManager,
				xrayManager,
				latencyTester,
				randomTestRetries,
				semaphore, &wg,
				&successIPs,
				&mu)
		}(ip)
	}

	wg.Wait()

	fmt.Println("\n\nSuccessful IPs and their latencies:")
	for _, result := range successIPs {
		fmt.Printf("IP: %s, Best Latency: %d ms\n", result.IP, result.Latency.Milliseconds())
	}
}

func init() {
	rootCmd.AddCommand(randomTestCmd)

	randomTestCmd.Flags().StringVar(&subnetListFile, "subnet-list", "", "Path to the subnet list file")
	randomTestCmd.Flags().IntVar(&ipCount, "count", 10, "Number of random IPs to test from the subnet list")
	randomTestCmd.Flags().StringVar(&randomTestConfigPath, "config", "./config.json", "Path to the Xray config file")
	randomTestCmd.Flags().StringVar(&randomTestXrayPath, "xray-path", "/usr/local/bin/xray", "Path to the Xray binary")
	randomTestCmd.Flags().StringVar(&randomTestSocks5Address, "socks5", "127.0.0.1:10808", "SOCKS5 proxy address")
	randomTestCmd.Flags().IntVar(&randomTestRetries, "retry", 3, "Number of retry attempts for latency tests")
	randomTestCmd.Flags().StringVar(&randomTestTestURL,
		"url",
		"http://connectivitycheck.gstatic.com/generate_204",
		"URL to test latency against")
	randomTestCmd.MarkFlagRequired("subnet-list")
}
