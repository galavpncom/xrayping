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
	configPath    string
	ipListPath    string
	xrayPath      string
	testURL       string
	socks5Address string
	retries       int
	verbose       bool
)

func Execute() error {
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "ipdelay",
	Short: "IP delay tester with Xray",
	Run: func(cmd *cobra.Command, args []string) {
		// Load the config and IP list
		configManager, err := config.NewManager(configPath)
		if err != nil {
			fmt.Println(utils.WrapTextWithColor(fmt.Sprintf("Error: %v", err), utils.ColorRed))
			return
		}

		ipList, err := config.LoadIPs(ipListPath)
		if err != nil {
			fmt.Println(utils.WrapTextWithColor(fmt.Sprintf("Error: %v", err), utils.ColorRed))
			return
		}

		// Initialize Xray and LatencyTester components
		xrayManager := xray.NewManager(xrayPath, verbose)
		latencyTester := latency.NewTester(testURL, socks5Address, verbose)

		concurrencyLimit := 5
		semaphore := make(chan struct{}, concurrencyLimit)

		var wg sync.WaitGroup
		var successIPs []latency.IPResult
		var mu sync.Mutex

		for _, ip := range ipList {
			if verbose {
				fmt.Println(utils.WrapTextWithColor(fmt.Sprintf("Starting test for IP: %s", ip),
					utils.ColorYellow))
			}
			wg.Add(1)
			go func(ip string) {
				defer wg.Done()
				semaphore <- struct{}{}
				latency.ProcessIP(ip,
					configManager,
					xrayManager,
					latencyTester,
					retries,
					semaphore,
					&wg,
					&successIPs,
					&mu)
			}(ip)
		}

		wg.Wait()

		if verbose {
			fmt.Println(utils.WrapTextWithColor("\n\nSuccessful IPs and their latencies:", utils.ColorGreen))
		}
		for _, result := range successIPs {
			fmt.Printf("%sIP: %s, Best Latency: %d ms%s\n",
				utils.ColorGreen,
				result.IP,
				result.Latency.Milliseconds(),
				utils.ColorReset)
		}
	},
}

func init() {
	rootCmd.Flags().StringVar(&configPath, "config", "", "Path to the Xray config file (required)")
	rootCmd.Flags().StringVar(&ipListPath, "ip-list", "", "Path to the IP list file (required)")
	rootCmd.Flags().StringVar(&xrayPath, "xray-path", "/usr/local/bin/xray/xray", "Path to the Xray binary")
	rootCmd.Flags().StringVar(&testURL, "url",
		"http://connectivitycheck.gstatic.com/generate_204",
		"URL to test for latency")
	rootCmd.Flags().StringVar(&socks5Address, "socks5", "127.0.0.1:10808", "SOCKS5 proxy address")
	rootCmd.Flags().IntVar(&retries, "retry", 3, "Number of retry attempts for latency tests")
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "Enable verbose output")
	rootCmd.MarkFlagRequired("config")
	rootCmd.MarkFlagRequired("ip-list")
}
